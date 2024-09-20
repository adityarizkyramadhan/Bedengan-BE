package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetExpiredToken(tokenString string) (time.Duration, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
	// Load secret key from environment variable
	var secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return 0, errors.New("SECRET_KEY environment variable not set")
	}
	var jwtSecret = []byte(secretKey)

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	// Get the claims from the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the 'exp' (expiration) field exists
		if exp, ok := claims["exp"].(float64); ok {
			// Convert the expiration time to time.Duration
			expirationTime := time.Unix(int64(exp), 0)
			now := time.Now()

			// Calculate the remaining duration before the token expires
			remainingTime := expirationTime.Sub(now)

			// If the token is already expired
			if remainingTime < 0 {
				return 0, errors.New("token is already expired")
			}

			return remainingTime, nil
		}
		return 0, errors.New("expiration (exp) not found in token")
	}

	return 0, errors.New("invalid token")
}
