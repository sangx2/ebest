package model

import (
	"encoding/json"
	"github.com/sangx2/ebest-sdk/res"
	"io"
	"time"
)

type Balance struct {
	Date string `json:"날짜"`

	res.T0424OutBlock1
}

func NewBalance(t0424OutBlock1s res.T0424OutBlock1) *Balance {
	now := time.Now()

	return &Balance{
		Date: now.Format("2006-01-02"),

		T0424OutBlock1: t0424OutBlock1s,
	}
}

func BalanceFromJson(data io.Reader) *Balance {
	var balance *Balance

	json.NewDecoder(data).Decode(&balance)

	return balance
}

func (bal *Balance) ToJson() []byte {
	b, e := json.Marshal(bal)
	if e != nil {
		return nil
	}

	return b
}
