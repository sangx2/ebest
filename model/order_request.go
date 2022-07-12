package model

import (
	"github.com/sangx2/ebest-sdk/res"
	"time"
)

const (
	OrderRequestTypeBuy    = "매수요청"
	OrderRequestTypeSell   = "매도요청"
	OrderRequestTypeModify = "정정요청"
	OrderRequestTypeCancel = "취소요청"
)

type OrderRequest struct {
	State string

	Date string `json:"날짜"`

	CSPAT00600OutBlock1 *res.CSPAT00600OutBlock1 `json:"현물정상주문1,omitempty"`
	CSPAT00600OutBlock2 *res.CSPAT00600OutBlock2 `json:"현물정상주문2,omitempty"`
	CSPAT00700OutBlock1 *res.CSPAT00700OutBlock1 `json:"현물정정주문1,omitempty"`
	CSPAT00700OutBlock2 *res.CSPAT00700OutBlock2 `json:"현물정정주문2,omitempty"`
	CSPAT00800OutBlock1 *res.CSPAT00800OutBlock1 `json:"현물취소주문1,omitempty"`
	CSPAT00800OutBlock2 *res.CSPAT00800OutBlock2 `json:"현물취소주문2,omitempty"`
}

func NewOrderRequest(state string) *OrderRequest {
	return &OrderRequest{
		State: state,
		Date:  time.Now().Format("2006-01-02"),
	}
}
