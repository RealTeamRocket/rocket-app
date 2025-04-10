package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	jwtSecret := "testsecret"
	authService := NewAuthService(jwtSecret)
	userID := uuid.New()

	tokenString, err := authService.GenerateToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestParseToken(t *testing.T) {
	jwtSecret := "testsecret"
	authService := NewAuthService(jwtSecret)
	userID := uuid.New()

	tokenString, err := authService.GenerateToken(userID)
	assert.NoError(t, err)

	token, err := authService.ParseToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.True(t, token.Valid)
}

func TestValidateToken(t *testing.T) {
	jwtSecret := "testsecret"
	authService := NewAuthService(jwtSecret)
	userID := uuid.New()

	tokenString, err := authService.GenerateToken(userID)
	assert.NoError(t, err)

	token, err := authService.ParseToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.True(t, token.Valid)

	parsedUserID, err := authService.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, parsedUserID)
}

func TestInvalidToken(t *testing.T) {
	jwtSecret := "testsecret"
	authService := NewAuthService(jwtSecret)

	// Create an invalid token string
	invalidTokenString := "invalid.token.string"

	token, err := authService.ParseToken(invalidTokenString)
	assert.Error(t, err)
	assert.Nil(t, token.Claims)
	assert.False(t, token.Valid)
}

func TestExpiredToken(t *testing.T) {
	jwtSecret := "testsecret"
	authService := NewAuthService(jwtSecret)
	userID := uuid.New()

	// Create a token with a short expiration time
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(time.Second * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	assert.NoError(t, err)

	// Wait for the token to expire
	time.Sleep(time.Second * 2)

	parsedToken, err := authService.ParseToken(tokenString)
	assert.Error(t, err)
	assert.NotNil(t, parsedToken)
	assert.False(t, parsedToken.Valid)
}
