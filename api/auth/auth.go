package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"skcloud.io/cloudzcp/zcpctl-backend/util"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		c.Logger().Debug("aaaaaa")

		user := new(User)
		if err = c.Bind(user); err != nil {
			return
		}

		token := createToken(*user)

		cookie := new(http.Cookie)
		cookie.Name = "zcp-cli-token"
		cookie.Value = token
		cookie.Expires = time.Now().Add(1 * time.Hour)
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, util.SetSuccessTrue(token))

	}

}

func createToken(user User) string {

	token := ""

	token = CreateToken(user)

	return token

}
