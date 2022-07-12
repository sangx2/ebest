package model

import (
	"github.com/sangx2/ebest-sdk/res"
	"time"
)

const (
	OrderTypeAccept = "접수"
	OrderTypeAgree  = "체결"
	OrderTypeModify = "정정"
	OrderTypeCancel = "취소"
	OrderTypeReject = "거부"
)

type Order struct {
	Date string `json:"날짜"`

	SC0OutBlock *res.SC0OutBlock `json:"접수,omitempty"`
	SC1OutBlock *res.SC1OutBlock `json:"체결,omitempty"`
	SC2OutBlock *res.SC2OutBlock `json:"정정,omitempty"`
	SC3OutBlock *res.SC3OutBlock `json:"취소,omitempty"`
	SC4OutBlock *res.SC4OutBlock `json:"거부,omitempty"`
}

func NewOrder() *Order {
	return &Order{
		Date: time.Now().Format("2006-01-02"),
	}
}
