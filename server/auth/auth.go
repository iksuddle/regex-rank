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

type MsgResponse struct {
	Msg string `json:"msg"`
}

type Auth struct {
	authConfig   *oauth2.Config
	userStore    database.UserStore
	sessionStore *sessions.CookieStore
	redirectURL  string
}

const stateSessionName = "rgx-state"
const authSessionName = "auth-session"

func InitAuth(c config.Config, db *sqlx.DB) Auth {
	auth := Auth{}

	auth.authConfig = &oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		RedirectURL:  "http://localhost:" + c.Port + "/login/callback",
		Endpoint:     github.Endpoint,
		Scopes:       []string{"read:user"},
	}

	auth.userStore = database.UserStore{
		Db: db,
	}

	auth.sessionStore = sessions.NewCookieStore([]byte(c.SessionKey))
	auth.sessionStore.Options.Path = "/"
	auth.sessionStore.Options.HttpOnly = true
	auth.sessionStore.Options.SameSite = http.SameSiteLaxMode
	auth.sessionStore.Options.Secure = true // some browsers consider localhost secure

	auth.redirectURL = c.ClientUrl + "login"

	return auth
}

func (auth Auth) LoginHandler(c echo.Context) error {
	s, err := auth.sessionStore.Get(c.Request(), stateSessionName)
	if err != nil {
		return httpError(err)
	}

	stateToken := generateStateToken()
	s.Values["state"] = stateToken

	s.Save(c.Request(), c.Response().Writer)

	url := auth.authConfig.AuthCodeURL(stateToken)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (auth Auth) LoginCallbackHandler(c echo.Context) error {
	s, err := auth.sessionStore.Get(c.Request(), stateSessionName)
	if err != nil {
		return httpError(err)
	}

	// verify state token
	stateToken := c.FormValue("state")
	if stateToken != s.Values["state"] {
		return httpError(err)
	}

	// exchange code for token
	authCode := c.FormValue("code")
	token, err := auth.authConfig.Exchange(context.TODO(), authCode)
	if err != nil {
		return httpError(err)
	}

	// creates http client with the token
	client := auth.authConfig.Client(context.TODO(), token)

	// get info from github using token
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return httpError(err)
	}

	res, err := client.Do(req)
	if err != nil {
		return httpError(err)
	}
	defer res.Body.Close()

	// decode user data
	userData := types.UserData{}
	if err := json.NewDecoder(res.Body).Decode(&userData); err != nil {
		return httpError(err)
	}

	// check if user exists and create them if needed
	user, err := auth.userStore.GetUserById(userData.Id)
	if err != nil {
		user = types.NewUser(userData)
		err = auth.userStore.CreateUser(user)
		if err != nil {
			return httpError(err)
		}
	}

	// store user in session
	authSession, err := auth.sessionStore.Get(c.Request(), authSessionName)
	if err != nil {
		return httpError(err)
	}

	authSession.Values["user_id"] = user.Id
	authSession.Values["user_name"] = user.Username
	authSession.Values["avatar_url"] = user.AvatarUrl

	err = authSession.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return httpError(err)
	}

	return c.Redirect(http.StatusPermanentRedirect, auth.redirectURL)
}

func (auth Auth) GetCurrentUser(c echo.Context) error {
	authSession, err := auth.sessionStore.Get(c.Request(), authSessionName)
	if err != nil {
		return httpError(err)
	}

	if authSession.IsNew {
		return c.JSON(http.StatusUnauthorized, MsgResponse{Msg: "session not found"})
	}

	// construct user
	// no need to check errors since we know user is stored in the session
	user := types.UserData{
		Id:        getIntFromSession(authSession, "user_id"),
		Username:  getStringFromSession(authSession, "user_name"),
		AvatarUrl: getStringFromSession(authSession, "avatar_url"),
	}

	return c.JSON(http.StatusOK, user)
}

func getStringFromSession(session *sessions.Session, key string) string {
	if val, ok := session.Values[key].(string); ok {
		return val
	}
	return ""
}

func getIntFromSession(session *sessions.Session, key string) int64 {
	if val, ok := session.Values[key].(int64); ok {
		return val
	}
	return -1
}

func generateStateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func httpError(err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
