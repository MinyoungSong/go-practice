package cluster

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"skcloud.io/cloudzcp/zcpctl-backend/db"
	"skcloud.io/cloudzcp/zcpctl-backend/util"
)

func GetClsuterList() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		db.Select()

		return c.JSON(http.StatusOK, util.SetSuccessTrue(""))

	}

}
