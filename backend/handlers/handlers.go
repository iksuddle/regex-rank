package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>Login with <a href="/login">GitHub</a></h1>`)
}

func TestAuth(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		log.Println(err)
		return err
	}
	jwtString := cookie.Value

	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		i := int64(claims["exp"].(float64))
		return c.String(http.StatusOK, fmt.Sprintf("%v %v", int(claims["sub"].(float64)), time.Unix(i, 0)))
	} else {
		fmt.Println(err)
	}

	return c.String(http.StatusForbidden, "there was a problemo")
}
