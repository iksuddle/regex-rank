package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>Login with <a href="/login">GitHub</a></h1>`)
}
