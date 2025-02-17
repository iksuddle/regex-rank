package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"suddle.dev/regex-rank/config"
	"suddle.dev/regex-rank/database"
	"suddle.dev/regex-rank/types"
)

const stateSessionName = "rgx-state"

var authConfig *oauth2.Config
var userStore database.UserStore
var sessionStore *sessions.CookieStore

func InitAuth(c config.Config, db *sqlx.DB) {
	authConfig = &oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		RedirectURL:  "http://localhost:" + c.Port + "/login/callback",
		Endpoint:     github.Endpoint,
		Scopes:       []string{"read:user"},
	}

	userStore = database.UserStore{
		Db: db,
	}

	sessionStore = sessions.NewCookieStore([]byte(c.SessionKey))
	sessionStore.Options.Path = "/"
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.SameSite = http.SameSiteLaxMode
	sessionStore.Options.Secure = true // some browsers consider localhost secure
}

func LoginHandler(c echo.Context) error {
	s, err := sessionStore.Get(c.Request(), stateSessionName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not decode session")
	}

	stateToken := generateStateToken()
	s.Values["state"] = stateToken

	s.Save(c.Request(), c.Response().Writer)

	url := authConfig.AuthCodeURL(stateToken)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func LoginCallbackHandler(c echo.Context) error {
	s, err := sessionStore.Get(c.Request(), stateSessionName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not decode session")
	}

	// verify state token
	stateToken := c.FormValue("state")
	if stateToken != s.Values["state"] {
		return echo.NewHTTPError(http.StatusForbidden, "state token mismatch")
	}

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
	userData := types.UserData{}
	if err := json.NewDecoder(res.Body).Decode(&userData); err != nil {
		return err
	}

	// check if user exists and create them if needed
	user, err := userStore.GetUserById(userData.Id)
	if err != nil {
		user = types.NewUser(userData)
		err = userStore.CreateUser(user)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, user)
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
