package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/iksuddle/regex-rank/types"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>Login with <a href="/login">GitHub</a></h1>`)
}

func TestAuthRoute(c echo.Context) error {
	user := c.Get("user").(types.User)
    res := fmt.Sprintf("<h1>Welcome %s!</h1>", user.Username)
    res += fmt.Sprintf("\n<p>Your account was created on %s</p>", time.Unix(user.CreatedAt, 0).Format("Monday, January 2006"))
	return c.HTML(http.StatusOK, res)
}
