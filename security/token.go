package security

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

const AuthTokenValidTime = time.Hour * 3

func GenerateToken(email string) (string, error) {
	authToken, err := CreateAuthToken(email)

	if err != nil {
		return "", err
	}

	return authToken, nil
}

func CreateAuthToken(email string) (string, error) {
	authTokenExp := time.Now().Add(AuthTokenValidTime)

	claims := jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(authTokenExp),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(viper.GetString("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
