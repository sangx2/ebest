package app

import (
	"github.com/sangx2/ebest/model"
	"strings"
)

func isETF(stock *model.Stock) bool {
	switch stock.Etfgubun {
	case "1": // ETF
		return true
	}
	return false
}

func isETN(stock *model.Stock) bool {
	switch stock.Etfgubun {
	case "2": // ETN
		return true
	}
	return false
}

// isBlueChip: 우선주
func isBlueChip(stock *model.Stock) bool {
	if strings.HasSuffix(stock.Hname, "우") || strings.HasSuffix(stock.Hname, "우B") ||
		strings.HasSuffix(stock.Hname, "우C") || strings.HasSuffix(stock.Hname, "우(전환)") {
		return true
	}
	return false
}
