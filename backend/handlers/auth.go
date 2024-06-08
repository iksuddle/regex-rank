package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/sessions"
	"github.com/iksuddle/regex-rank/config"
	"github.com/iksuddle/regex-rank/database"
	"github.com/iksuddle/regex-rank/types"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

const stateSessionName = "state-token"

var userStore *database.UserStore

var authConfig *oauth2.Config
var sessionStore *sessions.CookieStore

var jwtKey []byte

func InitAuth(config *config.Config, db *sqlx.DB) {
	authConfig = &oauth2.Config{
		ClientID:     config.ClientId,
		ClientSecret: config.ClientSecret,
		RedirectURL:  "http://localhost:" + config.Port + "/login/callback",
		Scopes: []string{
			"read:user",
			"user:email",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	sessionStore = sessions.NewCookieStore([]byte(config.SessionKey))
	sessionStore.Options.Path = "/"
	sessionStore.Options.MaxAge = 86400 // one day
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.SameSite = http.SameSiteLaxMode
	sessionStore.Options.Secure = true // some browsers consider http://localhost secure

	userStore = database.NewUserStore(db)

	jwtKey = []byte(config.JWTKey)
}

func LoginHandler(c echo.Context) error {
	stateSession, err := sessionStore.Get(c.Request(), stateSessionName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when getting session: %v\n", err))
	}

	state := generateStateToken()
	stateSession.Values["state"] = state

	stateSession.Save(c.Request(), c.Response().Writer)

	url := authConfig.AuthCodeURL(state) // github doesn't support PKCE yet
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func LoginCallbackHandler(c echo.Context) error {
	stateSession, err := sessionStore.Get(c.Request(), stateSessionName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when getting session: %v\n", err))
	}

	// verify that the states match
	state := c.FormValue("state")
	if state != stateSession.Values["state"] {
		return echo.NewHTTPError(http.StatusInternalServerError, "state token mismatch")
	}
	// delete the state token
	stateSession.Options.MaxAge = -1
	stateSession.Save(c.Request(), c.Response().Writer)

	// exchange code for token
	code := c.FormValue("code")
	token, err := authConfig.Exchange(context.TODO(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when exchanging code for token: %v\n", err))
	}

	client := authConfig.Client(context.TODO(), token)

	// make a request to retrieve user data
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when creating request to retrieve user data: %v\n", err))
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when retrieving user data from github: %v\n", err))
	}
	defer res.Body.Close()

	// get user data from response
	var userData map[string]any
	if err = json.NewDecoder(res.Body).Decode(&userData); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when decoding user data: %v\n", err))
	}

	// check if user exists
	userGithubId := int(userData["id"].(float64))
	user, err := userStore.GetUserById(userGithubId)
	if err != nil {
		// user does not exist
		user, err = types.NewUserFromData(userData)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when creating user: %v\n", err))
		}
		// create new user in db
		err = userStore.CreateUser(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when inserting user: %v\n", err))
		}
	}

	// create jwt
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"sub":     user.Id,
		"name":    user.Username,
		"picture": user.AvatarUrl,
		"created": user.CreatedAt,
	})

	jwtString, err := jwt.SignedString(jwtKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("error when signing jwt: %v\n", err))
	}

	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    jwtString,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
	})

	return c.JSON(http.StatusOK, types.NewJWTResponse(jwtString))
}

func LogoutHandler(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	return c.HTML(http.StatusOK, "<h1>Logged Out.</h1>")
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

var loggedInView = `
<h1>Welcome %s</h1>
<p>Your <code>id</code> is <code>%d</code></p>
<img src="%s" width="200" height="200" "border-radius:50%%;">
`
