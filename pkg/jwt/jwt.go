package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int64, username, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       id,
			"username": username,
			"exp":      time.Now().Add(60 * time.Minute).Unix(),
		},
	)

	key := []byte(secretKey)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateToken(tokenStr, secretKey string, withClaimValidation bool) (int64, string, error) {
	var (
		key    = []byte(secretKey)
		claims = jwt.MapClaims{}
		token  *jwt.Token
		err    error
	)

	if withClaimValidation {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithoutClaimsValidation())
	}

	if err != nil {
		return 0, "", err
	}

	if !token.Valid {
		return 0, "", errors.New("Invalid token")
	}

	userID := int64(claims["id"].(float64))
	username := claims["username"].(string)

	return userID, username, nil
}
