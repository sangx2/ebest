package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type StockOrderRequest struct {
	AccountNo string `json:"계좌번호,omitempty"`

	// for modify, cancel
	OrderNo string `json:"주문번호,omitempty"`

	Price  string `json:"가격,omitempty"`
	Amount string `json:"수량,omitempty"`

	// 00@지정가,03@시장가,05@조건부지정가,06@최유리지정가,07@최우선지정가,61@장개시전시간외종가,81@시간외종가,82@시간외단일가}
	OrdPrcPtnCode string `json:"호가유형코드,omitempty"`

	// {000:보통,003:유통/자기융자신규,005:유통대주신규,007:자기대주신규,101:유통융자상환,103:자기융자상환,105:유통대주상환,
	//             107:자기대주상환,180:예탁담보대출상환(신용)}
	MgntrnCode string `json:"신용거래코드,omitempty"`

	LoanDt string `json:"대출일,omitempty"`

	OrdCndiTpCode string `json:"주문조건구분,omitempty"` // {0:없음,1:IOC,2:FOK}
}

func postBuy(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	req := new(StockOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	order, e := server.Buy(req.AccountNo, stockCode, req.Amount, req.Price,
		req.OrdPrcPtnCode, req.MgntrnCode, req.LoanDt, req.OrdCndiTpCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func postSell(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	req := new(StockOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	order, e := server.Sell(req.AccountNo, stockCode, req.Amount, req.Price,
		req.OrdPrcPtnCode, req.MgntrnCode, req.LoanDt, req.OrdCndiTpCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func postModify(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	req := new(StockOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	order, e := server.Modify(req.OrderNo, req.AccountNo, stockCode, req.Amount, req.Price,
		req.OrdPrcPtnCode, req.OrdCndiTpCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func postCancel(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	req := new(StockOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	order, e := server.Cancel(req.OrderNo, req.AccountNo, stockCode, req.Amount)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getStockCodeList(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockList := server.GetStockCodeList()
	if stockList == nil {
		return c.String(http.StatusInternalServerError, "종목 코드 목록이 없습니다.")
	}

	return c.JSON(http.StatusOK, stockList)
}

func getStock(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	stock := server.GetStock(stockCode)
	if stock == nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s에 해당하는 종목 정보가 없습니다.", stockCode))
	}

	return c.JSON(http.StatusOK, stock)
}
