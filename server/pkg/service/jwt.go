package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zerodoctor/shawarma/pkg/model"
)

// JWT is the structure to hold the payload data
type JWT struct {
	model.User
	jwt.StandardClaims
}

// CreateJWTToken creates a jwt token and returns the string
func CreateJWTToken(user model.User, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &JWT{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(exp).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "shawarma server",
		},
	})

	return token.SignedString(getJWTSecret())
}

// VerifyJWTToken verifies the jwt token and returns the user payload
func VerifyJWTToken(tokenString string) (model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return getJWTSecret(), nil
	})

	if err != nil {
		return model.User{}, err
	}

	if claims, ok := token.Claims.(*JWT); ok && token.Valid {
		return claims.User, nil
	}

	return model.User{}, errors.New("invalid token")
}

func getJWTSecret() []byte {
	jwtSecret := os.Getenv("SHAW_JWT_SECRET")
	if jwtSecret == "" {
		log.Panic("SHAW_JWT_SECRET is not set")
	}

	return []byte(jwtSecret)
}
