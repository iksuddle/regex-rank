package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/types"
	"github.com/labstack/echo/v4"
)

var userStore *database.UserStore

func DeleteUser(c echo.Context) error {
	user, ok := c.Get(contextUserKey).(*types.User)
	if !ok {
		return newHTTPError(http.StatusInternalServerError, "could not retrieve user from context", nil)
	}

	err := userStore.DeleteUser(user.Id)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error deleting user", err)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "user deleted successfully",
	})
}

func GetUser(c echo.Context) error {
	user, ok := c.Get(contextUserKey).(*types.User)
	if !ok {
		return newHTTPError(http.StatusInternalServerError, "could not retrieve user from context", nil)
	}

	return c.JSON(http.StatusOK, user)
}

func TestAuthRoute(c echo.Context) error {
	// get user from the context
	user, ok := c.Get(contextUserKey).(*types.User)
	if !ok {
		return newHTTPError(http.StatusInternalServerError, "could not retrieve user from context", nil)
	}

	res := fmt.Sprintf("<h1>Welcome %s!</h1>", user.Username)
	res += fmt.Sprintf("\n<p>Your account was created on %s</p>", user.CreatedAt.Local().Format("Monday, 02 January 2006 at 3:04PM"))
	res += fmt.Sprintf("\n"+`<img src="%s" width="200" height="200">`, user.AvatarUrl)
	return c.HTML(http.StatusOK, res)
}

func newHTTPError(code int, message string, err error) *echo.HTTPError {
	msg := fmt.Sprintf("%s: %s", message, err.Error())
	log.Println(msg)
	return echo.NewHTTPError(code, msg)

}
