package main

import (
	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.NewConfig()

	db := database.NewDB(database.NewMySQLConfig(config))

	handlers.InitAuth(config, db)

	e := echo.New()

	// todo: move url to .env
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	e.GET("/login", handlers.LoginHandler)
	e.GET("/login/callback", handlers.LoginCallbackHandler)
	e.GET("/logout", handlers.LogoutHandler)

	e.GET("/test", handlers.TestAuthRoute, handlers.AuthRoute)
	e.GET("/user", handlers.GetUser, handlers.AuthRoute)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
