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
	sessionStore.Options.MaxAge = 3600
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = true // some browsers consider http://localhost secure
}

func LoginHandler(c echo.Context) error {
	session, err := sessionStore.Get(c.Request(), sessionId)
	if err != nil {
		log.Printf("error when getting session: %v\n", err)
		return err
	}

	state := generateStateToken()
	session.Values["state"] = state

	session.Save(c.Request(), c.Response().Writer)

	// github doesn't support PKCE yet.
	url := authConfig.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func AuthCallbackHandler(c echo.Context) error {
	session, err := sessionStore.Get(c.Request(), sessionId)
	if err != nil {
		log.Printf("error when getting session: %v\n", err)
		return err
	}

	state := c.FormValue("state")
	if state != session.Values["state"] {
		log.Printf("state token mismatch: %v\n", err)
		return err
	}

	// exchange code for token
	code := c.FormValue("code")
	token, err := authConfig.Exchange(context.TODO(), code)
	if err != nil {
		log.Printf("error when exchanging code for token: %v\n", err)
		return err
	}

	client := authConfig.Client(context.TODO(), token)

	// create a request to retrieve user data
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

	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error when reading response data: %v\n", err)
		return err
	}

	var userData map[string]any
	err = json.Unmarshal(rawData, &userData)
	if err != nil {
		log.Printf("error while unmarshalling json: %v\n", err)
		return err
	}

	// github user id should always be int
	userGithubId := userData["id"].(float64)
	user := User{
		GitHubId:  int(userGithubId),
		Username:  userData["login"].(string),
		AvatarUrl: userData["avatar_url"].(string),
	}

	return c.JSON(http.StatusOK, user)
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

type User struct {
	Id        string `json:"id"`
	GitHubId  int    `json:"github_id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}
