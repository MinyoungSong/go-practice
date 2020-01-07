package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	// "github.com/labstack/echo/engine/fasthttp"

	"skcloud.io/cloudzcp/zcpctl-backend/api/auth"
	"skcloud.io/cloudzcp/zcpctl-backend/api/cluster"
	"skcloud.io/cloudzcp/zcpctl-backend/db"
)

type responseDataFormat struct {
	success bool
	message string
	errors  interface{}
	data    interface{}
}

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	// e.Use(echo.Gzip())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	// e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	authGroup := e.Group("/api/auth")
	{
		authGroup.POST("/login", auth.Login())
	}

	clusterGroup := e.Group("/api/cluster")
	{
		clusterGroup.GET("", cluster.GetClsuterList())
		clusterGroup.GET("/", cluster.GetClsuterListMD())
		clusterGroup.GET("/:cluster_name", cluster.GetClsuterList())
	}

	db.InitDB()

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

	// logger.SetLevel(log.DEBUG)

	// // 첫 화면
	// e.GET("/", func(c echo.Context) error {
	// 	// echo.Logger.Debug("ddddd")
	// 	return c.String(http.StatusOK, "Hello World!")
	// })

	// e.GET("/users/:id", getUser)
	// e.GET("/show", show)
	// e.POST("/saveForm", save)

	// logger.Fatal(e.Start(":1323")) // localhost:1323
}

// func getUser(c echo.Context) error {

// 	id := c.Param("ids")
// 	return c.String(http.StatusOK, id)
// }

// func show(c echo.Context) error {

// 	team := c.QueryParam("team")
// 	member := c.QueryParam("member")

// 	return c.String(http.StatusOK, "team : "+team+"member : "+member)
// }

// func save(c echo.Context) (err error) {

// 	// e.Logger.Debug("Save Form API!!!!")
// 	// e.Logger.Info("Save Form API!!!!")
// 	// e.Logger.Warn("Save Form API!!!!")
// 	// e.Logger.Error("Save Form API!!!!")
// 	// e.Logger.Fatal("Save Form API!!!!")

// 	// name := c.FormValue("name")
// 	// email := c.FormValue("email")

// 	type User struct {
// 		Name       string `json:"name"`
// 		Email      string `json:"email"`
// 		Department struct {
// 			Code string `json:"code"`
// 			Name string `json:"name"`
// 		} `json:"department"`
// 	}

// 	// user := User{Name: name, Email: email}
// 	// userbyte, _ := json.Marshal(user)

// 	user := new(User)
// 	if err = c.Bind(user); err != nil {
// 		return
// 	}

// 	// user.Department = new(type department struct { Code string, Name string})

// 	// e.Logger.Debug(json.Marshal(user))
// 	// e.Logger.Debug(string(userbyte))

// 	return c.JSON(http.StatusOK, user)
// 	// return c.String(http.StatusOK, "name:"+name+", email:"+email)
// }
