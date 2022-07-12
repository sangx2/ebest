package api

import (
	"github.com/labstack/echo"
	"github.com/sangx2/ebest/model"
	"net/http"
)

func getBuyOrderRequestNumbers(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	orderNumbers, e := server.GetOrderRequestNumbers(model.OrderRequestTypeBuy, stockCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, orderNumbers)
}

func getSellOrderRequestNumbers(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	orderNumbers, e := server.GetOrderRequestNumbers(model.OrderRequestTypeSell, stockCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, orderNumbers)
}

func getModifyOrderRequestNumbers(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	orderNumbers, e := server.GetOrderRequestNumbers(model.OrderRequestTypeModify, stockCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, orderNumbers)
}

func getCancelOrderRequestNumbers(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	orderNumbers, e := server.GetOrderRequestNumbers(model.OrderRequestTypeCancel, stockCode)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, orderNumbers)
}

func getBuyOrderRequest(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")
	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrderRequest(stockCode, model.OrderRequestTypeBuy, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getSellOrderRequest(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")
	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrderRequest(stockCode, model.OrderRequestTypeSell, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getModifyOrderRequest(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")
	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrderRequest(stockCode, model.OrderRequestTypeModify, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getCancelOrderRequest(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")
	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrderRequest(stockCode, model.OrderRequestTypeCancel, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getAcceptOrder(c echo.Context) error {
	server := c.(*Context).GetServer()

	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrder(model.OrderTypeAccept, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getAgreeOrder(c echo.Context) error {
	server := c.(*Context).GetServer()

	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrder(model.OrderTypeAgree, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getModifyOrder(c echo.Context) error {
	server := c.(*Context).GetServer()

	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrder(model.OrderTypeModify, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getCancelOrder(c echo.Context) error {
	server := c.(*Context).GetServer()

	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrder(model.OrderTypeCancel, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func getRejectOrder(c echo.Context) error {
	server := c.(*Context).GetServer()

	orderNumber := c.Param("orderNumber")

	order, e := server.GetOrder(model.OrderTypeReject, orderNumber)
	if e != nil {
		return c.String(http.StatusInternalServerError, e.Error())
	}

	return c.JSON(http.StatusOK, order)
}
