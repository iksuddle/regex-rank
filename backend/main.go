package main

import (
	"net/http"

	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/labstack/echo/v4"
)

func main() {
	config := config.NewConfig()

	db := database.NewDB(database.NewMySQLConfig(config))

	InitAuth(config, db)

	e := echo.New()

	e.GET("/", indexHandler)
	e.GET("/login", LoginHandler)
	e.GET("/login/callback", LoginCallbackHandler)
	// e.GET("/logout", LogoutHandler)

	e.Logger.Fatal(e.Start(":" + config.Port))
}

func indexHandler(c echo.Context) error {
	indexHTML := `<h1>login with <a href="/login">Github</a></h1>`
	return c.HTML(http.StatusOK, indexHTML)
}
