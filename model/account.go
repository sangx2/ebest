package model

import (
	"encoding/json"
)

// Account : 계좌 정보
type Account struct {
	Number     string `json:"계좌번호"`
	Name       string `json:"이름"`
	DetailName string `json:"상세"`
	NickName   string `json:"별명"`
}

func NewAccount(number, name, detailName, nickName string) *Account {
	return &Account{
		Number:     number,
		Name:       name,
		DetailName: detailName,
		NickName:   nickName,
	}
}

func (a *Account) ToJson() []byte {
	b, e := json.Marshal(a)
	if e != nil {
		return nil
	}

	return b
}
