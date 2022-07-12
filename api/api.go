package api

import "github.com/labstack/echo"

const (
	BaseURI = "/api/v1"
)

type API struct {
	Accounts *echo.Route // `{BaseURI}/accounts`

	Assets *echo.Route // `{BaseURI}/assets`

	Balances *echo.Route // `{BaseURI}/accounts/:account/balances`

	Stocks *echo.Route // `{BaseURI}/stocks`

	Orders *echo.Route // `{BaseURI}/orders`
}

func NewAPI(echo *echo.Echo) *API {
	a := &API{}

	root := echo.Group(BaseURI)

	a.Accounts = root.GET("/accounts", getAccounts)

	a.Assets = root.GET("/assets", getAssets)
	a.Assets = root.GET("/assets/:account", getAsset)

	a.Balances = root.GET("/accounts/:account/balances", getBalances)

	a.Stocks = root.GET("/stocks", getStockCodeList)
	a.Stocks = root.GET("/stocks/:stockCode", getStock)

	a.Stocks = root.POST("/stocks/:stockCode/buy", postBuy)
	a.Stocks = root.POST("/stocks/:stockCode/sell", postSell)
	a.Stocks = root.POST("/stocks/:stockCode/modify", postModify)
	a.Stocks = root.POST("/stocks/:stockCode/cancel", postCancel)

	a.Orders = root.GET("/stocks/:stockCode/order/buy/requests", getBuyOrderRequestNumbers)
	a.Orders = root.GET("/stocks/:stockCode/order/sell/requests", getSellOrderRequestNumbers)
	a.Orders = root.GET("/stocks/:stockCode/order/modify/requests", getModifyOrderRequestNumbers)
	a.Orders = root.GET("/stocks/:stockCode/order/cancel/requests", getCancelOrderRequestNumbers)

	a.Orders = root.GET("/stocks/:stockCode/order/buy/requests/:orderNumber", getBuyOrderRequest)
	a.Orders = root.GET("/stocks/:stockCode/order/sell/requests/:orderNumber", getSellOrderRequest)
	a.Orders = root.GET("/stocks/:stockCode/order/modify/requests/:orderNumber", getModifyOrderRequest)
	a.Orders = root.GET("/stocks/:stockCode/order/cancel/requests/:orderNumber", getCancelOrderRequest)

	a.Orders = root.GET("/order/accept/:orderNumber", getAcceptOrder)
	a.Orders = root.GET("/order/agree/:orderNumber", getAgreeOrder)
	a.Orders = root.GET("/order/modify/:orderNumber", getModifyOrder)
	a.Orders = root.GET("/order/cancel/:orderNumber", getCancelOrder)
	a.Orders = root.GET("/order/reject/:orderNumber", getRejectOrder)

	a.Orders = root.GET("/stocks/:stockCode/", nil)

	return a
}
