package main

import (
	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.NewConfig()

	db := database.NewDB(database.NewMySQLConfig(config))

	handlers.InitAuth(config, db)

    e := echo.New()

	e.GET("/", handlers.IndexHandler)
	e.GET("/login", handlers.LoginHandler)
	e.GET("/login/callback", handlers.LoginCallbackHandler)
	e.GET("/test", handlers.TestAuthRoute, handlers.AuthRoute)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
