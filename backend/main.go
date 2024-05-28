package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	InitDB(NewMySQLConfig())

	InitAuth()

	e := echo.New()

	e.GET("/", indexHandler)
	e.GET("/login", LoginHandler)
	e.GET("/auth/callback", AuthCallbackHandler)
	e.GET("/logout", LogoutHandler)

	e.Logger.Fatal(e.Start(":" + Envs.Port))
}

func indexHandler(c echo.Context) error {
	indexHTML := `<h1>login with <a href="/login">GitHub</a></h1>`
	return c.HTML(http.StatusOK, indexHTML)
}
