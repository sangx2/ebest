package api

import (
	"github.com/labstack/echo"
	"net/http"
)

func getAccounts(c echo.Context) error {
	server := c.(*Context).GetServer()

	accounts := server.GetAccounts()
	if accounts == nil {
		return c.String(http.StatusInternalServerError, "계좌 정보가 없습니다.")
	}

	return c.JSON(http.StatusOK, accounts)
}
