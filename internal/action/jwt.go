package action

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var KEY = []byte("somekey")
var NOTOKEN = []byte("You're Unauthorized due to No token in the header")
var UNAUTHORIZE = []byte("You're Unauthorized")
var CANTPROVE = []byte("You're Unauthorized due to invalid token")
var TOKENERROR = []byte("You're Unauthorized due to error parsing the JWT")

type CustomClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func GenerateJWT(login string) (string, error) {
	claims := CustomClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "test",
		},
		login,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

	// })
	tokenString, err := token.SignedString(KEY)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
