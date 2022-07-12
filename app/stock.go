package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
)

// InitStocks 전체 종목 초기화
func (es *EBestServer) InitStocks() error {
	t8436Req := model.NewQueryRequest(ebest.T8436, false, res.T8436InBlock{Gubun: "0"})
	if e := es.requestServer.Request(ebest.T8436, t8436Req); e != nil {
		return e
	} else {
		if resp := <-t8436Req.RespChan; resp.Error != nil {
			return resp.Error
		} else {
			t8436OutBlocks := resp.OutBlocks[0].([]res.T8436OutBlock)
			for _, t8436OutBlock := range t8436OutBlocks {
				es.Stocks[t8436OutBlock.Shcode] = model.NewStock(t8436OutBlock)
			}
		}
	}
	log.Info("stock 초기화 완료")

	return nil
}

func (es *EBestServer) Buy(acntNo, code, amount, price,
	ordprcPtnCode, mgntrnCode, loanDt, ordCndiTpCode string) (*model.OrderRequest, error) {
	req := model.NewQueryRequest(ebest.CSPAT00600, false, res.CSPAT00600InBlock1{
		AcntNo:        acntNo,
		InptPwd:       es.config.AccountSettings.Accounts[acntNo],
		IsuNo:         code,
		OrdQty:        amount,
		OrdPrc:        price,
		BnsTpCode:     "2",
		OrdprcPtnCode: ordprcPtnCode,
		MgntrnCode:    mgntrnCode,
		LoanDt:        loanDt,
		OrdCndiTpCode: ordCndiTpCode,
	})
	if e := es.requestServer.Request(ebest.CSPAT00600, req); e != nil {
		return nil, fmt.Errorf("요청 서버 에러: %v", e)
	}
	resp := <-req.RespChan
	if resp.Error != nil {
		return nil, fmt.Errorf("매수 요청 에러: %v", resp.Error)
	}
	if (resp.OutBlocks == nil) || (len(resp.OutBlocks) != 2) {
		return nil, fmt.Errorf("매수 요청 에러: 유효하지 않는 OutBlocks")
	}

	cspat00600OutBlock1 := resp.OutBlocks[0].(res.CSPAT00600OutBlock1)
	cspat00600OutBlock2 := resp.OutBlocks[1].(res.CSPAT00600OutBlock2)

	order := model.NewOrderRequest(model.OrderRequestTypeBuy)
	order.CSPAT00600OutBlock1 = &cspat00600OutBlock1
	order.CSPAT00600OutBlock2 = &cspat00600OutBlock2

	if order.CSPAT00600OutBlock1.IsuNo == "" || order.CSPAT00600OutBlock2.OrdNo == "" {
		return nil, fmt.Errorf("매수 요청 에러: 유효하지 않는 응답 데이터: IsuNo(%s): OrdNo(%s)",
			order.CSPAT00600OutBlock1.IsuNo, order.CSPAT00600OutBlock2.OrdNo)
	}

	es.orderBuyRequestMutex.Lock()
	es.OrderBuyRequest[order.CSPAT00600OutBlock1.IsuNo][order.CSPAT00600OutBlock2.OrdNo] = order
	es.orderBuyRequestMutex.Unlock()

	return order, nil
}

func (es *EBestServer) Sell(acntNo, code, amount, price,
	ordprcPtnCode, mgntrnCode, loanDt, ordCndiTpCode string) (*model.OrderRequest, error) {
	req := model.NewQueryRequest(ebest.CSPAT00600, false, res.CSPAT00600InBlock1{
		AcntNo:        acntNo,
		InptPwd:       es.config.AccountSettings.Accounts[acntNo],
		IsuNo:         code,
		OrdQty:        amount,
		OrdPrc:        price,
		BnsTpCode:     "1",
		OrdprcPtnCode: ordprcPtnCode,
		MgntrnCode:    mgntrnCode,
		LoanDt:        loanDt,
		OrdCndiTpCode: ordCndiTpCode,
	})
	if e := es.requestServer.Request(ebest.CSPAT00600, req); e != nil {
		return nil, fmt.Errorf("요청 서버 에러: %v", e)
	}
	resp := <-req.RespChan
	if resp.Error != nil {
		return nil, fmt.Errorf("매도 요청 에러: %v", resp.Error)
	}
	if (resp.OutBlocks == nil) || (len(resp.OutBlocks) != 2) {
		return nil, fmt.Errorf("매도 요청 에러: 유효하지 않는 OutBlocks")
	}

	cspat00600OutBlock1, _ := resp.OutBlocks[0].(res.CSPAT00600OutBlock1)
	cspat00600OutBlock2, _ := resp.OutBlocks[1].(res.CSPAT00600OutBlock2)

	order := model.NewOrderRequest(model.OrderRequestTypeSell)
	order.CSPAT00600OutBlock1 = &cspat00600OutBlock1
	order.CSPAT00600OutBlock2 = &cspat00600OutBlock2

	if order.CSPAT00600OutBlock1.IsuNo == "" || order.CSPAT00600OutBlock2.OrdNo == "" {
		return nil, fmt.Errorf("매도 요청 에러: 유효하지 않는 응답 데이터: IsuNo(%s): OrdNo(%s)",
			order.CSPAT00600OutBlock1.IsuNo, order.CSPAT00600OutBlock2.OrdNo)
	}
	es.orderSellRequestMutex.Lock()
	es.OrderSellRequest[order.CSPAT00600OutBlock1.IsuNo][order.CSPAT00600OutBlock2.OrdNo] = order
	es.orderSellRequestMutex.Unlock()

	return order, nil
}

func (es *EBestServer) Modify(orgOrdNo, acntNo, code, amount, price, ordPrcPtnCode, ordCndiTpCode string) (*model.OrderRequest, error) {
	req := model.NewQueryRequest(ebest.CSPAT00700, false, res.CSPAT00700InBlock1{
		OrgOrdNo:      orgOrdNo,
		AcntNo:        acntNo,
		InptPwd:       es.config.AccountSettings.Accounts[acntNo],
		IsuNo:         code,
		OrdQty:        amount,
		OrdprcPtnCode: ordPrcPtnCode,
		OrdCndiTpCode: ordCndiTpCode,
		OrdPrc:        price,
	})
	if err := es.requestServer.Request(ebest.CSPAT00700, req); err != nil {
		return nil, fmt.Errorf("요청 서버 에러: %v", err)
	}

	resp := <-req.RespChan
	if resp.Error != nil {
		return nil, fmt.Errorf("정정 요청 에러: %v", resp.Error)
	}
	if (resp.OutBlocks == nil) || (len(resp.OutBlocks) != 2) {
		return nil, fmt.Errorf("정정 요청 에러: 유효하지 않는 OutBlocks")
	}

	cspat00700OutBlock1, _ := resp.OutBlocks[0].(res.CSPAT00700OutBlock1)
	cspat00700OutBlock2, _ := resp.OutBlocks[1].(res.CSPAT00700OutBlock2)

	order := model.NewOrderRequest(model.OrderRequestTypeModify)
	order.CSPAT00700OutBlock1 = &cspat00700OutBlock1
	order.CSPAT00700OutBlock2 = &cspat00700OutBlock2

	if order.CSPAT00700OutBlock1.IsuNo == "" || order.CSPAT00700OutBlock2.OrdNo == "" {
		return nil, fmt.Errorf("정정 요청 에러: 유효하지 않는 응답 데이터: IsuNo(%s): OrdNo(%s)",
			order.CSPAT00700OutBlock1.IsuNo, order.CSPAT00700OutBlock2.OrdNo)
	}
	es.orderModifyRequestMutex.Lock()
	es.OrderModifyRequest[order.CSPAT00700OutBlock1.IsuNo][order.CSPAT00700OutBlock2.OrdNo] = order
	es.orderModifyRequestMutex.Unlock()

	return order, nil
}

func (es *EBestServer) Cancel(orgOrdNo, acntNo, code, amount string) (*model.OrderRequest, error) {
	req := model.NewQueryRequest(ebest.CSPAT00800, false, res.CSPAT00800InBlock1{
		OrgOrdNo: orgOrdNo,
		AcntNo:   acntNo,
		InptPwd:  es.config.AccountSettings.Accounts[acntNo],
		IsuNo:    code,
		OrdQty:   amount,
	})
	if err := es.requestServer.Request(ebest.CSPAT00800, req); err != nil {
		return nil, fmt.Errorf("요청 서버 에러: %v", err)
	}
	resp := <-req.RespChan
	if resp.Error != nil {
		return nil, fmt.Errorf("취소 요청 에러: %v", resp.Error)
	}
	if (resp.OutBlocks == nil) || (len(resp.OutBlocks) != 2) {
		return nil, fmt.Errorf("취소 요청 에러: 유효하지 않는 OutBlocks")
	}

	cspat00800OutBlock1, _ := resp.OutBlocks[0].(res.CSPAT00800OutBlock1)
	cspat00800OutBlock2, _ := resp.OutBlocks[1].(res.CSPAT00800OutBlock2)

	order := model.NewOrderRequest(model.OrderRequestTypeCancel)
	order.CSPAT00800OutBlock1 = &cspat00800OutBlock1
	order.CSPAT00800OutBlock2 = &cspat00800OutBlock2

	if order.CSPAT00800OutBlock1.IsuNo == "" || order.CSPAT00800OutBlock2.OrdNo == "" {
		return nil, fmt.Errorf("취소 요청 에러: 유효하지 않는 응답 데이터: IsuNo(%s): OrdNo(%s)",
			order.CSPAT00800OutBlock1.IsuNo, order.CSPAT00800OutBlock2.OrdNo)
	}
	es.orderCancelRequestMutex.Lock()
	es.OrderCancelRequest[order.CSPAT00800OutBlock1.IsuNo][order.CSPAT00800OutBlock2.OrdNo] = order
	es.orderCancelRequestMutex.Unlock()

	return order, nil
}

func (es *EBestServer) GetStockCodeList() map[string]string {
	stockCodeList := make(map[string]string)

	es.stocksMutex.RLock()
	for stockCode, stock := range es.Stocks {
		stockCodeList[stock.Hname] = stockCode
	}
	es.stocksMutex.RUnlock()

	return stockCodeList
}

func (es *EBestServer) GetStock(stockCode string) *model.Stock {
	es.stocksMutex.RLock()
	stock := es.Stocks[stockCode]
	es.stocksMutex.RUnlock()

	return stock
}
