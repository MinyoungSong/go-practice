package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"skcloud.io/cloudzcp/zcpctl-backend/util"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type VerifyAuthConfig struct {
	Skipper middleware.Skipper
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
		cookie.Path = "/" // 반드시 입력 필요!!!
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, util.SetSuccessTrue(token))

	}

}

func createToken(user User) string {

	token := ""

	token = CreateToken(user)

	return token

}

func VerifyAuth(config VerifyAuthConfig) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			if config.Skipper(c) {
				return next(c)
			}

			var token string = ""
			cookie, err := c.Cookie("zcp-cli-token")

			if err != nil {
				token = c.Request().Header.Get("zcp-cli-token")
			} else {
				token = cookie.Value
			}

			if token == "" {
				return c.JSON(http.StatusUnauthorized, util.SetSuccessFalse("token is not verified", nil))
			}

			result := VerifyToken(token)

			if result["active"] == true {
				// fmt.Println(fmt.Sprint(c.Request().Header.Get("username"))) //req header에서 username 가져오기
				c.Request().Header.Add("username", fmt.Sprintf("%v", result["username"]))

			} else {
				return c.JSON(http.StatusUnauthorized, util.SetSuccessFalse("token is not verified", nil))

			}

			return next(c)
		}

	}

}
