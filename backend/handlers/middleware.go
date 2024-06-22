package handlers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// todo use Authorization header instead of cookie
func AuthRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("jwt")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		jwtString := cookie.Value

		token, err := jwt.Parse(jwtString, func(token *jwt.Token) (any, error) {
			// validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// read claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// store user in context
			userId := int(claims["sub"].(float64))
			c.Set("user_id", userId)
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, "error reading claims")
		}

		return next(c)
	}
}
