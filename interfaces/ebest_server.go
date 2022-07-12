package interfaces

import "github.com/sangx2/ebest/model"

type EBestServer interface {
	GetAccounts() []*model.Account

	GetAssets() map[string]*model.Asset

	GetBalances(accountNumber string) []*model.Balance

	GetStockCodeList() map[string]string
	GetStock(stockCode string) *model.Stock

	Buy(acntNo, stockCode, amount, price, ordprcPtnCode, mgntrnCode, loanDt, ordCndiTpCode string) (*model.OrderRequest, error)
	Sell(acntNo, stockCode, amount, price, ordprcPtnCode, mgntrnCode, loanDt, ordCndiTpCode string) (*model.OrderRequest, error)
	Modify(orgOrdNo, acntNo, stockCode, amount, price, ordPrcPtnCode, ordCndiTpCode string) (*model.OrderRequest, error)
	Cancel(orgOrdNo, acntNo, stockCode, amount string) (*model.OrderRequest, error)

	GetOrderRequestNumbers(orderType string, stockCode string) ([]string, error)
	GetOrderRequest(stockCode, orderType, orderNumber string) (*model.OrderRequest, error)
	GetOrder(orderType, orderNumber string) (*model.Order, error)
}
