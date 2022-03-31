package auth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("TestLoginSucces", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepository{})
		data, err := authUseCase.Login("nasrul@mail.com", "12345")
		assert.Nil(t, err)
		assert.Equal(t, "Token", data)
	})
	t.Run("TesLoginError", func(t *testing.T) {
		authUseCase := NewAuthUseCase(mockAuthRepositoryError{})
		data, err := authUseCase.Login("nas@mail.com", "123")
		assert.NotNil(t, err)
		assert.Equal(t, "Wrong email/password", data)
	})
}

type mockAuthRepository struct{}

func (m mockAuthRepository) Login(email string, password string) (string, error) {
	return "Token", nil
}

type mockAuthRepositoryError struct{}

func (m mockAuthRepositoryError) Login(email string, password string) (string, error) {
	return "Wrong email/password", fmt.Errorf("wrong email/password")
}
