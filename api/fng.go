package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func getFNG(c echo.Context) error {
	server := c.(*Context).GetServer()

	stockCode := c.Param("stockCode")

	fng := server.GetFNG(stockCode)
	if fng == nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s 기업 정보가 없습니다.", stockCode))
	}

	return c.JSON(http.StatusOK, fng)
}
