package model

import (
	"encoding/json"
	"github.com/sangx2/ebest-sdk/res"
	"strings"
)

type Stock struct {
	res.T8436OutBlock
}

func NewStock(t8436OutBlock res.T8436OutBlock) *Stock {
	s := &Stock{
		T8436OutBlock: t8436OutBlock,
	}

	return s
}

func (s *Stock) IsKOSPI() bool {
	if strings.Compare(s.T8436OutBlock.Gubun, "1") != 0 {
		return false
	}
	return true
}

func (s *Stock) ToJson() []byte {
	b, e := json.Marshal(s)
	if e != nil {
		return nil
	}

	return b
}
