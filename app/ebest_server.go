package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/impl"
	"github.com/sangx2/ebest-sdk/interfaces"
	"github.com/sangx2/ebest/model"
	"github.com/sangx2/ebest/store"
	"github.com/sangx2/ebest/store/filestore"
	"github.com/sangx2/go-servers/request"
	log "github.com/sangx2/golog"
	"sync"
	"time"
)

var (
	ResNames = []string{
		ebest.CSPAQ12200,
		ebest.CSPAT00600,
		ebest.CSPAT00700,
		ebest.CSPAT00800,
		ebest.T0424,
		ebest.T1101,
		ebest.T1305,
		ebest.T1511,
		ebest.T3320,
		ebest.T8424,
		ebest.T8436,
	}
)

type EBestServer struct {
	// 계좌
	Accounts []*model.Account
	// map[계좌번호]자산
	Assets map[string]*model.Asset
	// map[계좌번호]잔고
	Balances map[string][]*model.Balance
	// map[종목코드]종목
	Stocks map[string]*model.Stock
	// map[종목코드]기업정보
	FNGs map[string]*model.FNG

	// map[종목코드]map[주문번호]매매요청
	OrderBuyRequest    map[string]map[string]*model.OrderRequest `json:"매수 요청"`
	OrderSellRequest   map[string]map[string]*model.OrderRequest `json:"매도 요청"`
	OrderModifyRequest map[string]map[string]*model.OrderRequest `json:"수정 요청"`
	OrderCancelRequest map[string]map[string]*model.OrderRequest `json:"취소 요청"`

	// map[주문번호]주문
	OrderAccept map[string]*model.Order `json:"접수,omitempty"`
	OrderAgree  map[string]*model.Order `json:"체결,omitempty"`
	OrderModify map[string]*model.Order `json:"정정,omitempty"`
	OrderCancel map[string]*model.Order `json:"취소,omitempty"`
	OrderReject map[string]*model.Order `json:"거부,omitempty"`

	// ebest
	*ebest.EBest
	queries       map[string]*ebest.Query
	reals         map[string]*ebest.Real
	realDoneChans map[string]chan bool

	// server
	requestServer *request.Server

	// store
	store store.Store

	// mutex
	accountsMutex sync.RWMutex
	assetsMutex   sync.RWMutex
	balancesMutex sync.RWMutex
	stocksMutex   sync.RWMutex
	fngsMutex     sync.RWMutex

	orderBuyRequestMutex    sync.RWMutex
	orderSellRequestMutex   sync.RWMutex
	orderModifyRequestMutex sync.RWMutex
	orderCancelRequestMutex sync.RWMutex

	orderAcceptMutex sync.RWMutex
	orderAgreeMutex  sync.RWMutex
	orderModifyMutex sync.RWMutex
	orderCancelMutex sync.RWMutex
	orderRejectMutex sync.RWMutex

	doneChans map[string]chan bool

	wg sync.WaitGroup

	config *model.Config
}

func NewEBestServer(config *model.Config) *EBestServer {
	eBest := ebest.NewEBest(config.AppSettings.ID, config.AppSettings.Passwd, config.AppSettings.CertPasswd,
		config.AppSettings.Server, ebest.Port, config.AppSettings.ResPath)
	if eBest == nil {
		log.Error("NewEBest is nil")
		return nil
	}
	es := &EBestServer{
		Assets:   make(map[string]*model.Asset),
		Balances: make(map[string][]*model.Balance),
		Stocks:   make(map[string]*model.Stock),
		FNGs:     make(map[string]*model.FNG),

		OrderBuyRequest:    make(map[string]map[string]*model.OrderRequest),
		OrderSellRequest:   make(map[string]map[string]*model.OrderRequest),
		OrderModifyRequest: make(map[string]map[string]*model.OrderRequest),
		OrderCancelRequest: make(map[string]map[string]*model.OrderRequest),

		OrderAccept: make(map[string]*model.Order),
		OrderAgree:  make(map[string]*model.Order),
		OrderModify: make(map[string]*model.Order),
		OrderCancel: make(map[string]*model.Order),
		OrderReject: make(map[string]*model.Order),

		EBest: eBest,

		queries:       make(map[string]*ebest.Query),
		reals:         make(map[string]*ebest.Real),
		realDoneChans: make(map[string]chan bool),

		requestServer: request.NewServer(config.AppSettings.QueueSize),

		doneChans: make(map[string]chan bool),

		config: config,
	}
	return es
}

func (es *EBestServer) Init() error {
	// set EBest
	if e := es.EBest.Connect(); e != nil {
		return fmt.Errorf("Init: ebest Connect: %v", e)
	}
	log.Info("eBest 접속 성공")

	if e := es.EBest.Login(); e != nil {
		es.EBest.Disconnect()
		return fmt.Errorf("Init: ebest Login: %v", e)
	}
	log.Info("eBest 로그인 성공")

	// 1. create query
	for _, resName := range ResNames {
		var duration = time.Second

		var trade interfaces.QueryTrade
		switch resName {
		case ebest.CSPAQ12200:
			trade = impl.NewCSPAQ12200()
			duration = time.Second * 6 // 이놈만 특별함
		case ebest.CSPAT00600:
			trade = impl.NewCSPAT00600()
		case ebest.CSPAT00700:
			trade = impl.NewCSPAT00700()
		case ebest.CSPAT00800:
			trade = impl.NewCSPAT00800()
		case ebest.T1101:
			trade = impl.NewT1101()
		case ebest.T1305:
			trade = impl.NewT1305()
		case ebest.T1511:
			trade = impl.NewT1511()
		case ebest.T3320:
			trade = impl.NewT3320()
		case ebest.T0424:
			trade = impl.NewT0424()
		case ebest.T8424:
			trade = impl.NewT8424()
		case ebest.T8436:
			trade = impl.NewT8436()
		default:
			return fmt.Errorf("Init: invalid ResName: %s", resName)
		}

		query := ebest.NewQuery(es.config.AppSettings.ResPath, trade)
		es.queries[resName] = query

		if e := es.requestServer.AddLimitersWithFunc(resName,
			[]*request.Limiter{request.NewLimiter(duration, int32(query.TPS)),
				request.NewLimiter(time.Minute*10, int32(query.LPP))}, es.QueryCallback); e != nil {
			log.Error("RequestServer.AddLimitersWithFunc", log.Err(e))
			return fmt.Errorf("Init: AddLimitersWithFunc: %v", e)
		}
	}
	es.requestServer.Start()

	// store
	// TODO: database supplier
	if !es.config.SQLSettings.Enable {
		es.store = filestore.NewFileSupplier(es.config.AppSettings.DataPath)
	}
	if es.store == nil {
		return fmt.Errorf("Init: store is nil")
	}

	if e := es.InitAccount(); e != nil {
		return fmt.Errorf("Init: %w", e)
	}

	if e := es.InitAssets(); e != nil {
		return fmt.Errorf("Init: %w", e)
	}

	if e := es.InitBalance(); e != nil {
		return fmt.Errorf("Init: %w", e)
	}

	if e := es.InitStocks(); e != nil {
		return fmt.Errorf("Init: %w", e)
	}

	if e := es.InitOrder(); e != nil {
		return fmt.Errorf("Init: %w", e)
	}

	if e := es.InitFNGs(); e != nil {
		return fmt.Errorf("Init: %w", e)
	}

	return nil
}

func (es *EBestServer) Start() error {

	return nil
}

func (es *EBestServer) Shutdown() {
	if e := es.FinalizeFNGs(); e != nil {
		log.Error("finalize FNGs error", log.Err(e))
	}

	if e := es.FinalizeBalance(); e != nil {
		log.Error("finalize balance error", log.Err(e))
	}

	if e := es.FinalizeAssets(); e != nil {
		log.Error("finalize assets error", log.Err(e))
	}

	// eBest - reals
	for _, recvDoneChan := range es.realDoneChans {
		recvDoneChan <- true
		close(recvDoneChan)
	}
	for _, r := range es.reals {
		r.Close()
	}

	// eBest - queries
	if es.requestServer != nil {
		es.requestServer.Shutdown()
	}
	for _, q := range es.queries {
		q.Close()
	}

	es.EBest.Disconnect()
	log.Info("eBest 접속 종료")
}
