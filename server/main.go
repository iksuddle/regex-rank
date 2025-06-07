package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"suddle.dev/regex-rank/auth"
	"suddle.dev/regex-rank/config"
	"suddle.dev/regex-rank/database"
)

func main() {
	godotenv.Load()

	config := config.NewConfig()

	db := database.NewDB(config)
	auth := auth.InitAuth(config, db)

	e := echo.New()

	e.GET("/login", auth.LoginHandler)
	e.GET("/login/callback", auth.LoginCallbackHandler)
	e.GET("/me", auth.GetCurrentUser)

	e.Start(":" + config.Port)
}
