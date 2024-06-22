package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>Login with <a href="/login">GitHub</a></h1>`)
}

func TestAuthRoute(c echo.Context) error {
	userId, ok := c.Get("user_id").(int)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user not found (probably didn't login)")
	}
	user, err := userStore.GetUserById(userId)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "user not found", err)
	}
	res := fmt.Sprintf("<h1>Welcome %s!</h1>", user.Username)
	res += fmt.Sprintf("\n<p>Your account was created on %s</p>", user.CreatedAt.Local().Format("Monday, 02 January 2006 at 3:04PM"))
	res += fmt.Sprintf("\n"+`<img src="%s" width="200" height="200">`, user.AvatarUrl)
	return c.HTML(http.StatusOK, res)
}
