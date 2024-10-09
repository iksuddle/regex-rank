package main

import (
	"fmt"
	"log"

	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type app struct {
	config *config.Config
	db     *sqlx.DB
	auth   *auth
	store  *database.Store
}

func newHTTPError(code int, message string, err error) *echo.HTTPError {
	msg := fmt.Sprintf("%s: %s", message, err.Error())
	log.Println(msg)
	return echo.NewHTTPError(code, msg)
}
