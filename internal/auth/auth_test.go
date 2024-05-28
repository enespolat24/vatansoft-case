package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestVerifyPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)

	err = VerifyPassword(hashedPassword, password)
	assert.NoError(t, err)

	err = VerifyPassword(hashedPassword, "wrongpassword")
	assert.Error(t, err)
}

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestVerifyToken(t *testing.T) {
	tokenString, err := GenerateToken(1)
	assert.NoError(t, err)

	token, err := VerifyToken(tokenString)
	assert.NoError(t, err)
	assert.NotNil(t, token)

	// Test with invalid token
	invalidToken := "invalidtoken"
	_, err = VerifyToken(invalidToken)
	assert.Error(t, err)
}
