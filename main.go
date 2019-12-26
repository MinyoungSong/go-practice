package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	// "github.com/labstack/echo/v4/middleware"
)

var e = echo.New()
var logger = e.Logger

type responseDataFormat struct {
	success bool
	message string
	errors  interface{}
	data    interface{}
}

func main() {

	logger.SetLevel(log.DEBUG)

	// 첫 화면
	e.GET("/", func(c echo.Context) error {
		// echo.Logger.Debug("ddddd")
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/saveForm", save)

	logger.Fatal(e.Start(":1323")) // localhost:1323
}

func getUser(c echo.Context) error {

	id := c.Param("ids")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {

	team := c.QueryParam("team")
	member := c.QueryParam("member")

	return c.String(http.StatusOK, "team : "+team+"member : "+member)
}

func save(c echo.Context) (err error) {

	e.Logger.Debug("Save Form API!!!!")
	e.Logger.Info("Save Form API!!!!")
	e.Logger.Warn("Save Form API!!!!")
	e.Logger.Error("Save Form API!!!!")
	// e.Logger.Fatal("Save Form API!!!!")

	// name := c.FormValue("name")
	// email := c.FormValue("email")

	type User struct {
		Name       string `json:"name"`
		Email      string `json:"email"`
		Department struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"department"`
	}

	// user := User{Name: name, Email: email}
	// userbyte, _ := json.Marshal(user)

	user := new(User)
	if err = c.Bind(user); err != nil {
		return
	}

	// user.Department = new(type department struct { Code string, Name string})

	// e.Logger.Debug(json.Marshal(user))
	// e.Logger.Debug(string(userbyte))

	return c.JSON(http.StatusOK, user)
	// return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
