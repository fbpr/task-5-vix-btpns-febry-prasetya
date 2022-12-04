package helpers

import (
	"time"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

var secretKey = []byte("jwt_secret_key")

func GenerateJWT(username string, email string) (accessToken string, err error) {
	accessTokenExp := time.Now().Add(1 * time.Hour).Unix()
	claims := &JwtCustomClaims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTokenExp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString(secretKey)
	return
}

func ValidateJWT(signedToken string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtCustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return token.Claims.(*JwtCustomClaims), nil
}
