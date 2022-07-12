package app

import (
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/impl"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
)

// InitOrder 실시간 주문 접수/체결/정정/취소/거부에 대한 작업 쓰레드 생성
func (es *EBestServer) InitOrder() error {
	resPath := es.config.AppSettings.ResPath

	// 종목별 메모리 할당
	for _, stockCode := range es.GetStockCodeList() {
		es.OrderBuyRequest[stockCode] = make(map[string]*model.OrderRequest)
		es.OrderSellRequest[stockCode] = make(map[string]*model.OrderRequest)
		es.OrderModifyRequest[stockCode] = make(map[string]*model.OrderRequest)
		es.OrderCancelRequest[stockCode] = make(map[string]*model.OrderRequest)
	}

	// 접수
	sc0RealTrade := impl.NewSC0()
	if sc0RealTrade == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewSC0 is nil")
	}
	sc0Real := ebest.NewReal(resPath, sc0RealTrade)
	if sc0Real == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewReal(SC0) is nil")
	}
	if e := sc0Real.SetInBlock(nil); e != nil {
		return e
	}
	sc0Real.Start()
	es.reals[ebest.SC0] = sc0Real

	// 체결
	sc1RealTrade := impl.NewSC1()
	if sc1RealTrade == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewSC1 is nil")
	}
	sc1Real := ebest.NewReal(resPath, sc1RealTrade)
	if sc1Real == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewReal(SC1) is nil")
	}
	if e := sc1Real.SetInBlock(nil); e != nil {
		return e
	}
	sc1Real.Start()
	es.reals[ebest.SC1] = sc1Real

	// 정정
	sc2RealTrade := impl.NewSC2()
	if sc2RealTrade == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewSC2 is nil")
	}
	sc2Real := ebest.NewReal(resPath, sc2RealTrade)
	if sc2Real == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewReal(SC2) is nil")
	}
	if e := sc2Real.SetInBlock(nil); e != nil {
		return e
	}
	sc2Real.Start()
	es.reals[ebest.SC2] = sc2Real

	// 취소
	sc3RealTrade := impl.NewSC3()
	if sc3RealTrade == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewSC3 is nil")
	}
	sc3Real := ebest.NewReal(resPath, sc3RealTrade)
	if sc3Real == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewReal(SC3) is nil")
	}
	if e := sc3Real.SetInBlock(nil); e != nil {
		return e
	}
	sc3Real.Start()
	es.reals[ebest.SC3] = sc3Real

	// 거부
	sc4RealTrade := impl.NewSC4()
	if sc4RealTrade == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewSC4 is nil")
	}
	sc4Real := ebest.NewReal(resPath, sc4RealTrade)
	if sc4Real == nil {
		return fmt.Errorf("App.TradeServer.InitStocks.NewReal(SC4) is nil")
	}
	if e := sc4Real.SetInBlock(nil); e != nil {
		return e
	}
	sc4Real.Start()
	es.reals[ebest.SC4] = sc4Real

	// 실시간 주문 상태
	doneChan := make(chan bool, 1)
	es.doneChans["stock"] = doneChan
	es.wg.Add(1)
	go es.startOrderReceiver()

	log.Info("order 초기화 완료")

	return nil
}

func (es *EBestServer) startOrderReceiver() {
	defer es.wg.Done()

	doneChan := make(chan bool, 1)
	es.doneChans["orderReceiver"] = doneChan

	sc0Real := es.reals[ebest.SC0]
	sc1Real := es.reals[ebest.SC1]
	sc2Real := es.reals[ebest.SC2]
	sc3Real := es.reals[ebest.SC3]
	sc4Real := es.reals[ebest.SC4]

	for {
		select {
		case <-sc0Real.GetReceivedRealDataChan():
			sc0OutBlock := sc0Real.GetOutBlock().(res.SC0OutBlock)
			log.Debug("App.TradeServer.SC0.GetOutBlock", log.Any("SC0OutBlock", sc0OutBlock))

			order := model.NewOrder()
			order.SC0OutBlock = &sc0OutBlock
			es.orderAcceptMutex.Lock()
			es.OrderAccept[sc0OutBlock.Ordno] = order
			es.orderAcceptMutex.Unlock()
		case <-sc1Real.GetReceivedRealDataChan():
			sc1OutBlock := sc1Real.GetOutBlock().(res.SC1OutBlock)
			log.Debug("App.TradeServer.SC1.GetOutBlock", log.Any("SC1OutBlock", sc1OutBlock))

			order := model.NewOrder()
			order.SC1OutBlock = &sc1OutBlock
			es.orderAgreeMutex.Lock()
			es.OrderAgree[sc1OutBlock.Ordno] = order
			es.orderAgreeMutex.Unlock()

			// 체결 후 자산 및 잔고 업데이트
			es.UpdateAssets()
			es.UpdateBalances(sc1OutBlock.Accno)
		case <-sc2Real.GetReceivedRealDataChan():
			sc2OutBlock := sc2Real.GetOutBlock().(res.SC2OutBlock)
			log.Debug("App.TradeServer.SC2.GetOutBlock", log.Any("SC2OutBlock", sc2OutBlock))

			order := model.NewOrder()
			order.SC2OutBlock = &sc2OutBlock
			es.orderModifyMutex.Lock()
			es.OrderModify[sc2OutBlock.Ordno] = order
			es.orderModifyMutex.Unlock()
		case <-sc3Real.GetReceivedRealDataChan():
			sc3OutBlock := sc3Real.GetOutBlock().(res.SC3OutBlock)
			log.Debug("App.TradeServer.SC3.GetOutBlock", log.Any("SC3OutBlock", sc3OutBlock))

			order := model.NewOrder()
			order.SC3OutBlock = &sc3OutBlock
			es.orderCancelMutex.Lock()
			es.OrderCancel[sc3OutBlock.Ordno] = order
			es.orderCancelMutex.Unlock()
		case <-sc4Real.GetReceivedRealDataChan():
			sc4OutBlock := sc4Real.GetOutBlock().(res.SC4OutBlock)
			log.Debug("App.TradeServer.SC4.GetOutBlock", log.Any("SC4OutBlock", sc4OutBlock))

			order := model.NewOrder()
			order.SC4OutBlock = &sc4OutBlock
			es.orderRejectMutex.Lock()
			es.OrderReject[sc4OutBlock.Ordno] = order
			es.orderRejectMutex.Unlock()
		case <-doneChan:
			return
		}
	}
}

func (es *EBestServer) GetOrderRequestNumbers(orderRequestType string, stockCode string) ([]string, error) {
	var orders map[string]*model.OrderRequest
	var orderNumbers []string
	var e error

	switch orderRequestType {
	case model.OrderRequestTypeBuy:
		es.orderBuyRequestMutex.RLock()
		buyRequests, ok := es.OrderBuyRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 유효하지 않음: 단축코드: %s", stockCode)
		} else if len(buyRequests) == 0 {
			e = fmt.Errorf("매수요청 목록이 없음: 단축코드: %s", stockCode)
		} else {
			orders = buyRequests
		}
		es.orderBuyRequestMutex.RUnlock()
	case model.OrderRequestTypeSell:
		es.orderSellRequestMutex.RLock()
		sellRequests, ok := es.OrderSellRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 유효하지 않음: 단축코드: %s", stockCode)
		} else if len(sellRequests) == 0 {
			e = fmt.Errorf("매도요청 목록이 없음: 단축코드: %s", stockCode)
		} else {
			orders = sellRequests
		}
		es.orderSellRequestMutex.RUnlock()
	case model.OrderRequestTypeModify:
		es.orderModifyRequestMutex.RLock()
		orderModifies, ok := es.OrderModifyRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 유효하지 않음: 단축코드: %s", stockCode)
		} else if len(orderModifies) == 0 {
			e = fmt.Errorf("정정요청 목록이 없음: 단축코드: %s", stockCode)
		} else {
			orders = orderModifies
		}
		es.orderModifyRequestMutex.RUnlock()
	case model.OrderRequestTypeCancel:
		es.orderCancelRequestMutex.RLock()
		orderCancels, ok := es.OrderCancelRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 유효하지 않음: 단축코드: %s", stockCode)
		} else if len(orderCancels) == 0 {
			e = fmt.Errorf("취소요청 목록이 없음: 단축코드: %s", stockCode)
		} else {
			orders = orderCancels
		}
		es.orderCancelRequestMutex.RUnlock()
	default:
		return nil, fmt.Errorf("주문종류가 유효하지 않음: %s", orderRequestType)
	}

	if e == nil {
		for ordNumber := range orders {
			orderNumbers = append(orderNumbers, ordNumber)
		}
	}

	return orderNumbers, e
}

func (es *EBestServer) GetOrderRequest(stockCode, orderRequestType, orderNumber string) (*model.OrderRequest, error) {
	var orderRequestMap map[string]*model.OrderRequest
	var orderRequest *model.OrderRequest
	var e error

	switch orderRequestType {
	case model.OrderRequestTypeBuy:
		es.orderBuyRequestMutex.RLock()
		orderRequests, ok := es.OrderBuyRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 존재하지 않음: 단축코드: %s", stockCode)
		} else {
			orderRequestMap = orderRequests
		}
		es.orderBuyRequestMutex.RUnlock()
	case model.OrderRequestTypeSell:
		es.orderSellRequestMutex.RLock()
		orderRequests, ok := es.OrderSellRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 존재하지 않음: 단축코드: %s", stockCode)
		} else {
			orderRequestMap = orderRequests
		}
		es.orderSellRequestMutex.RUnlock()
	case model.OrderRequestTypeModify:
		es.orderModifyRequestMutex.RLock()
		orderModifies, ok := es.OrderModifyRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 존재하지 않음: 단축코드: %s", stockCode)
		} else {
			orderRequestMap = orderModifies
		}
		es.orderModifyRequestMutex.RUnlock()
	case model.OrderRequestTypeCancel:
		es.orderCancelRequestMutex.RLock()
		orderCancels, ok := es.OrderCancelRequest[stockCode]
		if !ok {
			e = fmt.Errorf("단축주문코드가 존재하지 않음: 단축코드: %s", stockCode)
		} else {
			orderRequestMap = orderCancels
		}
		es.orderCancelRequestMutex.RUnlock()
	default:
		return nil, fmt.Errorf("주문종류가 유효하지 않음: %s", orderRequestType)
	}

	if e == nil {
		o, ok := orderRequestMap[orderNumber]
		if !ok {
			e = fmt.Errorf("주문번호가 존재하지 않음: 주문번호: %s", orderNumber)
		} else {
			orderRequest = o
		}
	}

	return orderRequest, e
}

func (es *EBestServer) GetOrder(orderType, orderNumber string) (*model.Order, error) {
	var order *model.Order
	var e error

	switch orderType {
	case model.OrderTypeAccept:
		es.orderAcceptMutex.RLock()
		if orderAccept, ok := es.OrderAccept[orderNumber]; !ok {
			e = fmt.Errorf("주문번호가 존재하지 않음: %s", orderNumber)
		} else {
			order = orderAccept
		}
		es.orderAcceptMutex.RUnlock()
	case model.OrderTypeAgree:
		es.orderAgreeMutex.RLock()
		if orderAgree, ok := es.OrderAgree[orderNumber]; !ok {
			e = fmt.Errorf("주문번호가 존재하지 않음: %s", orderNumber)
		} else {
			order = orderAgree
		}
		es.orderAgreeMutex.RUnlock()
	case model.OrderTypeModify:
		es.orderModifyMutex.RLock()
		if orderModify, ok := es.OrderModify[orderNumber]; !ok {
			e = fmt.Errorf("주문번호가 존재하지 않음: %s", orderNumber)
		} else {
			order = orderModify
		}
		es.orderModifyMutex.RUnlock()
	case model.OrderTypeCancel:
		es.orderCancelMutex.RLock()
		if orderCancel, ok := es.OrderCancel[orderNumber]; !ok {
			e = fmt.Errorf("주분번호가 존재하지 않음: %s", orderNumber)
		} else {
			order = orderCancel
		}
		es.orderCancelMutex.RUnlock()
	case model.OrderTypeReject:
		es.orderRejectMutex.RLock()
		if orderReject, ok := es.OrderReject[orderNumber]; !ok {
			e = fmt.Errorf("주분번호가 존재하지 않음: %s", orderNumber)
		} else {
			order = orderReject
		}
		es.orderRejectMutex.RUnlock()
	default:
		return nil, fmt.Errorf("주문종류가 유효하지 않음: %s", orderType)
	}

	return order, e
}
