package model

import (
	"encoding/json"
	"github.com/sangx2/ebest-sdk/res"
	"io"
	"time"
)

// FNG : 기업정보
type FNG struct {
	CreateAt time.Time
	UpdateAt time.Time

	Name string `json:"종목명"`
	Code string `json:"종목코드"`

	res.T3320OutBlock  `json:"기업기본정보"`
	res.T3320OutBlock1 `json:"기업재무정보"`
}

func NewFNG(name, code string, t3320OutBlock res.T3320OutBlock, t3320OutBLock1 res.T3320OutBlock1) *FNG {
	now := time.Now()

	return &FNG{
		CreateAt: now,
		UpdateAt: now,

		Name: name,
		Code: code,

		T3320OutBlock:  t3320OutBlock,
		T3320OutBlock1: t3320OutBLock1,
	}
}

func FNGFromJson(data io.Reader) *FNG {
	var fng *FNG

	json.NewDecoder(data).Decode(&fng)

	return fng
}

func (p *FNG) ToJson() []byte {
	b, e := json.Marshal(p)
	if e != nil {
		return nil
	}

	return b
}
