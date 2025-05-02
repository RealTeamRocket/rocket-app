package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AuthService struct {
	JwtSecret string
}

func NewAuthService(jwtSecret string) *AuthService {
	return &AuthService{JwtSecret: jwtSecret}
}

func (a *AuthService) GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(a.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *AuthService) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.JwtSecret), nil
	})

	return token, err
}

func (a *AuthService) ValidateToken(token *jwt.Token) (uuid.UUID, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["user_id"].(string)
		userUUID, err := uuid.Parse(userID)
		if err != nil {
			return uuid.Nil, fmt.Errorf("invalid user ID format")
		}
		return userUUID, nil
	}
	return uuid.Nil, fmt.Errorf("invalid token claims")
}
