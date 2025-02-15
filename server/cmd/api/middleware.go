package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (app *app) authRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get the jwt token
		cookie, err := c.Cookie("rgx_jwt")
		if err != nil {
			return newHTTPError(http.StatusForbidden, "could not retrieve jwt from cookie", err)
		}
		jwtString := cookie.Value

		// validate token
		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (any, error) {
			// validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return app.auth.jwtKey, nil
		})
		if err != nil {
			log.Printf("could not validate jwt token %v\n", err)
			return newHTTPError(http.StatusForbidden, "error validating token", err)
		}

		// validate token
		if !token.Valid {
			log.Println("invalid token", err)
			return newHTTPError(http.StatusForbidden, "invalid token", err)
		}

		// read claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// store user in context
			userId := int(claims["sub"].(float64))
			user, err := app.store.Users.GetUserById(userId)
			if err != nil {
				return newHTTPError(http.StatusInternalServerError, "user not found", err)
			}
			c.Set(contextUserKey, user)
		} else {
			return newHTTPError(http.StatusForbidden, "error reading claims", err)
		}

		return next(c)
	}
}
