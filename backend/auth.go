package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

const stateSessionId = "state-token"
const authSessionId = "rgx-auth"
const userIdKey = "user-id"

var authConfig *oauth2.Config
var sessionStore *sessions.CookieStore

// ensure .env is loaded before calling InitAuth
func InitAuth() {
	authConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:" + Envs.Port + "/login/callback",
		Scopes: []string{
			"read:user",
			"user:email",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionStore.Options.Path = "/"
	sessionStore.Options.MaxAge = 86400 // one day
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.SameSite = http.SameSiteLaxMode
	sessionStore.Options.Secure = true // some browsers consider http://localhost secure
}

// func LogoutHandler(c echo.Context) error {
// 	session, err := sessionStore.Get(c.Request(), sessionId)
// 	if err != nil {
// 		return err
// 	}
//
// 	if session.IsNew {
// 		return c.Redirect(http.StatusTemporaryRedirect, "/")
// 	}
//
// 	session.Options.MaxAge = -1
// 	if err = session.Save(c.Request(), c.Response().Writer); err != nil {
// 		return err
// 	}
//
// 	return c.HTML(http.StatusOK, "<h1>Logged out.</h1>")
// }

func LoginHandler(c echo.Context) error {
	stateSession, err := sessionStore.Get(c.Request(), stateSessionId)
	if err != nil {
		log.Printf("error when getting session: %v\n", err)
		return err
	}

	state := generateStateToken()
	stateSession.Values["state"] = state

	stateSession.Save(c.Request(), c.Response().Writer)

	url := authConfig.AuthCodeURL(state) // github doesn't support PKCE yet
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthCallbackHandler(c echo.Context) error {
	stateSession, err := sessionStore.Get(c.Request(), stateSessionId)
	if err != nil {
		log.Printf("error when getting session: %v\n", err)
		return err
	}

	// verify that the states match
	state := c.FormValue("state")
	if state != stateSession.Values["state"] {
		log.Printf("state token mismatch: %v\n", err)
		return err
	}
	// delete the state token
	stateSession.Options.MaxAge = -1
	defer stateSession.Save(c.Request(), c.Response().Writer)

	// exchange code for token
	code := c.FormValue("code")
	token, err := authConfig.Exchange(context.TODO(), code)
	if err != nil {
		log.Printf("error when exchanging code for token: %v\n", err)
		return err
	}

	client := authConfig.Client(context.TODO(), token)

	// make a request to retrieve user data
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		log.Printf("error when creating request to retrieve user data: %v\n", err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		log.Printf("error when retrieving user data from github: %v\n", err)
		return err
	}
	defer res.Body.Close()

	// get user data from response
	var userData map[string]any
	if err = json.NewDecoder(res.Body).Decode(&userData); err != nil {
		log.Println("error 1")
		return err
	}

	user, err := NewUserFromData(userData)
	if err != nil {
		log.Println("error 2")
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf(loggedInView, user.Username, user.GitHubId, user.AvatarUrl))
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

var loggedInView = `
<h1>Welcome %s</h1>
<p>Your <code>github_id</code> is <code>%d</code></p>
<img src="%s" width="200" height="200" style="border-radius:50%%;border:0.5rem solid #DDDDDD">
`
