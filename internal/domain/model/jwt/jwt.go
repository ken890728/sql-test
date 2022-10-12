package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	UserId string
}

var jwtKey = []byte("ggWQTJSyKu&NcKrAEYrB&yT23Uuy9V@vIIGtH!&8x*Im8y*@90")

func NewToken(userId string) (string, error) {
	authClaims := AuthClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		UserId: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func VerifyToken(tokenStr string) (AuthClaims, error) {
	var authClaims AuthClaims
	token, err := jwt.ParseWithClaims(tokenStr, &authClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method error")
		}
		return jwtKey, nil
	})

	if err != nil {
		return authClaims, err
	}

	if !token.Valid {
		return authClaims, errors.New("invalid token")
	}

	return authClaims, nil
}
