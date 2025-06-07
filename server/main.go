package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	e.GET("/login", auth.LoginHandler)
	e.GET("/login/callback", auth.LoginCallbackHandler)
	e.GET("/me", auth.GetCurrentUser)

	e.Start(":" + config.Port)
}
