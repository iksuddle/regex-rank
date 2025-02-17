package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"suddle.dev/regex-rank/config"
	"suddle.dev/regex-rank/database"
	"suddle.dev/regex-rank/types"
)

var authConfig *oauth2.Config
var userStore *database.UserStore

func InitAuth(c *config.Config, db *sqlx.DB) {
	authConfig = &oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		RedirectURL:  "http://localhost:" + c.Port + "/login/callback",
		Endpoint:     github.Endpoint,
		Scopes:       []string{"read:user"},
	}
	userStore = &database.UserStore{
		Db: db,
	}
}

func LoginHandler(c echo.Context) error {
	// get url for github login
	url := authConfig.AuthCodeURL("state")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func LoginCallbackHandler(c echo.Context) error {
	// exchange code for token
	authCode := c.FormValue("code")
	token, err := authConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		return err
	}

	// creates http client with the token
	client := authConfig.Client(context.TODO(), token)

	// get info from github using token
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// decode user data
	userData := &types.UserData{}
	if err := json.NewDecoder(res.Body).Decode(userData); err != nil {
		return err
	}

	// create user
	user := types.NewUser(userData)
	err = userStore.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
