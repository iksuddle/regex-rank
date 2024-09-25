package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/types"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const sessionName = "rgx-session"

type auth struct {
	authConfig   *oauth2.Config
	sessionStore *sessions.CookieStore
	userStore    *database.UserStore
	jwtKey       []byte
}

func initAuth(config *config) *auth {
	a := &auth{
		authConfig: &oauth2.Config{
			ClientID:     config.clientId,
			ClientSecret: config.clientSecret,
			RedirectURL:  "http://localhost:" + config.port + "/login/callback",
			Endpoint:     github.Endpoint,
			Scopes: []string{
				"read:user",
				"user:email",
			},
		},
		sessionStore: sessions.NewCookieStore([]byte(config.sessionKey)),
		jwtKey:       []byte(config.jwtKey),
	}

	a.sessionStore.Options.Path = "/"
	a.sessionStore.Options.HttpOnly = true
	a.sessionStore.Options.SameSite = http.SameSiteLaxMode
	a.sessionStore.Options.Secure = true // some browsers consider http://localhost secure

	return a
}

func (app *app) loginHandler(c echo.Context) error {
	session, err := app.auth.sessionStore.Get(c.Request(), sessionName)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error when getting session", err)
	}

	state := generateStateToken()
	session.Values["state"] = state

	session.Save(c.Request(), c.Response().Writer)

	url := app.auth.authConfig.AuthCodeURL(state) // github doesn't support PKCE yet
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (app *app) loginCallbackHandler(c echo.Context) error {
	session, err := app.auth.sessionStore.Get(c.Request(), sessionName)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error when getting session", err)
	}

	// verify that the states match
	state := c.FormValue("state")
	if state != session.Values["state"] {
		return newHTTPError(http.StatusInternalServerError, "state token mismatch", nil)
	}

	// exchange code for token
	code := c.FormValue("code")
	token, err := app.auth.authConfig.Exchange(context.TODO(), code)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error when exchanging code for token", err)
	}

	client := app.auth.authConfig.Client(context.TODO(), token)

	// make a request to retrieve user data
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error when creating request to retrieve user data", err)
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error when retrieving user data from github", err)
	}
	defer res.Body.Close()

	// get user data from response
	var userData map[string]any
	if err = json.NewDecoder(res.Body).Decode(&userData); err != nil {
		return newHTTPError(http.StatusInternalServerError, "error when decoding user data", err)
	}

	// check if user exists
	userGithubId := int(userData["id"].(float64))
	user, err := app.store.Users.GetUserById(userGithubId)
	if err != nil {
		// user does not exist
		user, err = types.NewUserFromData(userData)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, "error when creating user", err)
		}
		// create new user in db
		err = app.store.Users.CreateUser(user)
		if err != nil {
			return newHTTPError(http.StatusInternalServerError, "error when inserting user", err)
		}
	}

	jwt, err := types.NewJWT(user.Id, app.auth.jwtKey)
	if err != nil {
		return newHTTPError(http.StatusInternalServerError, "error creating jwt", err)
	}

	c.SetCookie(&http.Cookie{
		Name:     "rgx_jwt",
		Value:    jwt,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400,
	})

	// todo: move url to .env
	url := "http://localhost:5173/login"
	return c.Redirect(http.StatusPermanentRedirect, url)
}

func (app *app) logoutHandler(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:   "rgx_jwt",
		MaxAge: -1,
	})

	return c.JSON(http.StatusOK, map[string]string{"message": "successfully logged out"})
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
