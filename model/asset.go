package model

import (
	"encoding/json"
	"github.com/sangx2/ebest-sdk/res"
	"io"
	"time"
)

type Asset struct {
	AccountNum string `json:"계좌번호"`
	Date       string `json:"날짜"`

	res.CSPAQ12200OutBlock2
}

func NewAsset(accountNum string, cspaq12200 res.CSPAQ12200OutBlock2) *Asset {
	now := time.Now()

	return &Asset{
		AccountNum: accountNum,
		Date:       now.Format("2006-01-02"),

		CSPAQ12200OutBlock2: cspaq12200,
	}
}

func AssetFromJson(data io.Reader) *Asset {
	var asset *Asset

	json.NewDecoder(data).Decode(&asset)

	return asset
}

func (a *Asset) ToJson() []byte {
	b, e := json.Marshal(a)
	if e != nil {
		return nil
	}

	return b
}
