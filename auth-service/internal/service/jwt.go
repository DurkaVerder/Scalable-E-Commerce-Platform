package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func (s *ServiceManager) generateJWT(userId int) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	standartClaims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	claims := &Claims{
		StandardClaims: standartClaims,
		UserID:         userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil

}

func (s *ServiceManager) validateJWT(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return fmt.Errorf("token is expired")
		}
	} else {
		return fmt.Errorf("invalid token")
	}

	return nil
}
