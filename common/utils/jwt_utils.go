package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	"time"
)

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().UTC().After(expirationTime) {
			return nil, fmt.Errorf("token has expired")
		}
	}

	return token, nil
}

func PreAuthorize(w http.ResponseWriter, r *http.Request) (error, bool) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, true
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return nil, true
	}
	return err, false
}
