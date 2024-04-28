package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const sessionName = "auth-session"

var authConfig *oauth2.Config
var sessionStore *sessions.CookieStore

// port needed for auth redirect url
func InitAuth(port string) {
	authConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  "http://localhost:" + port + "/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionStore.Options.Path = "/"
	sessionStore.Options.MaxAge = 3600
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = true // some browsers consider http://localhost secure
}

func LoginHandler(c echo.Context) error {
	verifier := oauth2.GenerateVerifier()

	// todo: handle error
	session, _ := sessionStore.Get(c.Request(), sessionName)
	session.Values["verifier"] = verifier

	state := generateStateToken()
	session.Values["state"] = state

	session.Save(c.Request(), c.Response().Writer)

	url := authConfig.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthCallbackHandler(c echo.Context) error {
	// todo: handle error
	session, _ := sessionStore.Get(c.Request(), sessionName)

	state := c.FormValue("state")
	if state != session.Values["state"] {
		log.Fatal("state mismatch")
	}

	verifier := session.Values["verifier"].(string)
	// exchange code for token
	code := c.FormValue("code")
	token, err := authConfig.Exchange(context.Background(), code, oauth2.VerifierOption(verifier))
	if err != nil {
		log.Fatal("error when exchanging auth code")
	}

	// todo: handle error
	// get user data
	req, _ := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("error when retrieving  user info")
	}
	defer res.Body.Close()

	// todo: handle error
	userInfoJSON, _ := io.ReadAll(res.Body)

	var userInfo map[string]any
	json.Unmarshal(userInfoJSON, &userInfo)

	return c.JSON(http.StatusOK, userInfo)
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
