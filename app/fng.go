package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
	"time"
)

// InitFNGs 기업 정보 초기화
// - 주식 정보 초기화(InitStocks)가 요구됨
func (es *EBestServer) InitFNGs() error {
	queryFNGs := <-es.store.FNG().GetAll()
	if queryFNGs.Err != nil {
		return fmt.Errorf("InitFNGs: %s", queryFNGs.Err)
	}

	var FNGs map[string]*model.FNG
	if queryFNGs.Data != nil {
		FNGs = queryFNGs.Data.(map[string]*model.FNG)
		es.FNGs = FNGs
	}

	es.wg.Add(1)
	go es.updateFNGs()

	log.Info("기업정보 초기화 완료")

	return nil
}

func (es *EBestServer) FinalizeFNGs() error {
	es.fngsMutex.Lock()
	fngs := es.FNGs
	es.fngsMutex.Unlock()

	query := <-es.store.FNG().SaveAll(fngs)
	if query.Err != nil {
		return fmt.Errorf("FinalizeFNGs: %s", query.Err)
	}
	log.Info("FNG 정보 저장 완료")

	return nil
}

// updateFNGs Goroutine
func (es *EBestServer) updateFNGs() {
	defer es.wg.Done()

	doneChan := make(chan bool, 1)
	es.doneChans["updateFNGs"] = doneChan

	es.stocksMutex.Lock()
	stocks := es.Stocks
	es.stocksMutex.Unlock()

	for code, stock := range stocks {
		if isETF(stock) || isETN(stock) || isBlueChip(stock) {
			continue
		}

		req := model.NewQueryRequest(ebest.T3320, false, res.T3320InBlock{Gicode: code})
		if e := es.requestServer.Request(ebest.T3320, req); e != nil {
			log.Error(e.Error())
			return
		}

		select {
		case resp := <-req.RespChan:
			if resp.Error != nil {
				log.Error(resp.Error.Error())
			} else {
				t3320OutBlock := resp.OutBlocks[0].(res.T3320OutBlock)
				t3320OutBlock1 := resp.OutBlocks[1].(res.T3320OutBlock1)

				es.fngsMutex.Lock()
				fng := es.FNGs[code]

				if fng == nil {
					es.FNGs[code] = model.NewFNG(stock.Hname, code, t3320OutBlock, t3320OutBlock1)
				} else if fng.T3320OutBlock.Gsym != t3320OutBlock.Gsym {
					fng.UpdateAt = time.Now()
					fng.T3320OutBlock = t3320OutBlock
				} else if fng.T3320OutBlock1.Gsym != t3320OutBlock1.Gsym {
					fng.UpdateAt = time.Now()
					fng.T3320OutBlock1 = t3320OutBlock1
				}

				es.fngsMutex.Unlock()
			}
		case <-doneChan:
			return
		}
	}
}

func (es *EBestServer) GetFNG(stockCode string) *model.FNG {
	es.fngsMutex.RLock()
	fng := es.FNGs[stockCode]
	es.fngsMutex.RUnlock()

	return fng
}
