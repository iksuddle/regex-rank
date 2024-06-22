package types

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func NewJWT(userId int, jwtKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"sub": userId,
	})

	jwtString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
