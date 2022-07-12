package app

import (
	"errors"
	"fmt"
	"github.com/sangx2/ebest-sdk/ebest"
	"github.com/sangx2/ebest-sdk/res"
	"github.com/sangx2/ebest/model"
	log "github.com/sangx2/golog"
	"strings"
)

func (es *EBestServer) QueryCallback(req interface{}) {
	var err error
	var inBlock interface{}

	request := req.(*model.Request)
	defer close(request.RespChan)
	log.Debug("QueryCallback", log.Any("Request", request))

	// TODO: 기능 추가
	switch request.ResName {
	case ebest.CSPAQ12200:
		inBlock = request.InBlocks[0].(res.CSPAQ12200InBlock1)
	case ebest.CSPAT00600:
		inBlock = request.InBlocks[0].(res.CSPAT00600InBlock1)
	case ebest.CSPAT00700:
		inBlock = request.InBlocks[0].(res.CSPAT00700InBlock1)
	case ebest.CSPAT00800:
		inBlock = request.InBlocks[0].(res.CSPAT00800InBlock1)
	case ebest.T1101:
		inBlock = request.InBlocks[0].(res.T1101InBlock)
	case ebest.T1305:
		inBlock = request.InBlocks[0].(res.T1305InBlock)
	case ebest.T1511:
		inBlock = request.InBlocks[0].(res.T1511InBlock)
	case ebest.T3320:
		inBlock = request.InBlocks[0].(res.T3320InBlock)
	case ebest.T0424:
		inBlock = request.InBlocks[0].(res.T0424InBlock)
	case ebest.T8424:
		inBlock = request.InBlocks[0].(res.T8424InBlock)
	case ebest.T8436:
		inBlock = request.InBlocks[0].(res.T8436InBlock)
	default:
		request.RespChan <- model.NewResponse(nil,
			fmt.Errorf("%s 가 구현되어 있지 않음", request.ResName))
		return
	}

	query := es.queries[request.ResName]

	err = query.SetInBlock(inBlock)
	if err != nil {
		var errQuery *ebest.ErrQuery
		if errors.As(err, &errQuery) {
			request.RespChan <- model.NewResponse(nil, fmt.Errorf("SetInBlock: %s", errQuery.Error()))
		} else {
			request.RespChan <- model.NewResponse(nil, err)
		}
		return
	}

	ret := query.Request(request.IsOccurs)
	log.Debug("QueryCallback.Request", log.Int("ret", ret))
	if ret < 0 {
		request.RespChan <- model.NewResponse(nil, fmt.Errorf("query 요청: %d: %s", ret, es.EBest.GetErrorMessage(ret)))
		return
	}

	receivedMsg, e := query.GetReceiveMessage()
	log.Debug("QueryCallback.GetReceiveMessage", log.String("receivedMsg", receivedMsg), log.Err(e))
	if e != nil {
		request.RespChan <- model.NewResponse(nil, fmt.Errorf("%s", e.Error()))
		return
	}

	log.Debug("QueryCallback.GetReceiveDataChan", log.String("ReceiveData", <-query.GetReceiveDataChan()))
	// TODO: receivedMsg 리스트업
	switch {
	case
		strings.Contains(receivedMsg, "00000"), // 00000: 조회완료
		strings.Contains(receivedMsg, "00040"), // 00040: 모의투자 매수주문이 완료 되었습니다.
		strings.Contains(receivedMsg, "00039"), // 00039: 모의투자 매도주문이 완료 되었습니다.
		strings.Contains(receivedMsg, "00462"), // 00462: 모의투자 정정주문이 완료 되었습니다.
		strings.Contains(receivedMsg, "00463"), // 00463: 모의투자 취소주문이 완료 되었습니다.
		strings.Contains(receivedMsg, "00136"): // 00136: 모의투자 조회가 완료되었습니다.
		request.RespChan <- model.NewResponse(query.GetOutBlocks(), nil)
	default:
		request.RespChan <- model.NewResponse(nil, fmt.Errorf("%s", receivedMsg))
	}
}
