package main

import (
	"net/http"

	"github.com/iksuddle/regex-rank/types"
	"github.com/labstack/echo/v4"
)

const contextUserKey string = "rgx-user"

func (app *app) getUserHandler(c echo.Context) error {
	user, ok := c.Get(contextUserKey).(*types.User)
	if !ok {
		return newHTTPError(http.StatusInternalServerError, "could not retrieve user from context", nil)
	}

	return c.JSON(http.StatusOK, user)
}

func (app *app) deleteUserHandler(c echo.Context) error {
	user, ok := c.Get(contextUserKey).(*types.User)
	if !ok {
		return newHTTPError(http.StatusInternalServerError, "could not retrieve user from context", nil)
	}

	err := app.store.Users.DeleteUser(user.Id)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error deleting user", err)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "user deleted successfully",
	})
}
