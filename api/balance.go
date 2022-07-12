package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func getBalances(c echo.Context) error {
	server := c.(*Context).GetServer()

	accountNumber := c.Param("account")

	balances := server.GetBalances(accountNumber)
	if balances == nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s 계좌의 잔고가 없습니다.", accountNumber))
	}

	return c.JSON(http.StatusOK, balances)
}
