package api

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func getAssets(c echo.Context) error {
	server := c.(*Context).GetServer()

	assets := server.GetAssets()
	if assets == nil {
		return c.String(http.StatusInternalServerError, "자산 정보가 없습니다.")
	}

	return c.JSON(http.StatusOK, assets)
}

func getAsset(c echo.Context) error {
	server := c.(*Context).GetServer()
	accountNumber := c.Param("account")

	assets := server.GetAssets()
	if assets == nil {
		return c.String(http.StatusInternalServerError, "자산 정보가 없습니다.")
	}

	asset, isExist := assets[accountNumber]
	if !isExist {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("%s 계좌에 해당하는 자산이 없습니다.", accountNumber))
	}

	return c.JSON(http.StatusOK, asset)
}
