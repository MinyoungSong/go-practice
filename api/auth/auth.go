package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		// user := User{Name: name, Email: email}
		// userbyte, _ := json.Marshal(user)

		c.Logger().Debug("aaaaaa")

		user := new(User)
		if err = c.Bind(user); err != nil {
			return
		}

		// user.Department = new(type department struct { Code string, Name string})

		// e.Logger.Debug(json.Marshal(user))
		// e.Logger.Debug(string(userbyte))

		token := createToken(*user)

		type ResponseBody struct {
			Success bool        `json:"success"`
			Message string      `json:"message"`
			Errors  string      `json:"errors"`
			Data    interface{} `json:"data"`
		}

		res := new(ResponseBody)
		res.Success = true
		res.Message = ""
		res.Errors = ""
		res.Data = token

		return c.JSON(http.StatusOK, res)

	}

}

func createToken(user User) string {

	token := ""

	token = CreateToken(user)

	return token

}
