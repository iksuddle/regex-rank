package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/iksuddle/regex-rank/types"
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `<h1>Login with <a href="/login">GitHub</a></h1>`)
}

func ProtectedTest(c echo.Context) error {
    log.Println("ERROR B")
	user := c.Get("user").(types.User)
	return c.HTML(http.StatusOK, fmt.Sprintf("<h1>Welcome %s!</h1>", user.Username))
}

func AuthRoute(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
        log.Println("entering auth route")
		cookie, err := c.Cookie("jwt")
		if err != nil {
			log.Println(err)
			return err
		}
        log.Println("auth route 1")
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
        log.Println("auth route 2")
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
            log.Println("ERROR A")
			// store user in context
			user := types.User{
				Id:        int(claims["sub"].(float64)),
				Username:  claims["name"].(string),
				AvatarUrl: claims["picture"].(string),
			}
			c.Set("user", user)
		} else {
			fmt.Println(err)
            log.Println("auth route 3")
		}

		return next(c)
	}
}

// func TestAuth(c echo.Context) error {
// 	cookie, err := c.Cookie("jwt")
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	jwtString := cookie.Value
//
// 	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (any, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return jwtKey, nil
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok {
// 		i := int64(claims["exp"].(float64))
// 		return c.String(http.StatusOK, fmt.Sprintf("%v %v", int(claims["sub"].(float64)), time.Unix(i, 0)))
// 	} else {
// 		fmt.Println(err)
// 	}
//
// 	return c.String(http.StatusForbidden, "there was a problemo")
// }
