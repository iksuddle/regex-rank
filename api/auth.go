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

const sessionId = "auth-session"

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

	session, err := sessionStore.Get(c.Request(), sessionId)
	if err != nil {
		return err
	}
	session.Values["verifier"] = verifier

	state := generateStateToken()
	session.Values["state"] = state

	session.Save(c.Request(), c.Response().Writer)

	url := authConfig.AuthCodeURL(state, oauth2.S256ChallengeOption(verifier))
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthCallbackHandler(c echo.Context) error {
	session, err := sessionStore.Get(c.Request(), sessionId)
	if err != nil {
		return err
	}

	state := c.FormValue("state")
	if state != session.Values["state"] {
		log.Fatal("state mismatch")
	}

	// exchange code for token
	code := c.FormValue("code")
	verifier := session.Values["verifier"].(string)
	token, err := authConfig.Exchange(context.TODO(), code, oauth2.VerifierOption(verifier))
	if err != nil {
		return err
	}

	client := authConfig.Client(context.TODO(), token)

	// create a request to retrieve user data
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	userInfoJSON, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var userInfo map[string]any
	err = json.Unmarshal(userInfoJSON, &userInfo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, userInfo)
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
