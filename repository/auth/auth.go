package auth

import (
	_middlewares "be7/layered/delivery/middleware"
	_entities "be7/layered/entities"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository struct {
	database *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		database: db,
	}
}

func (ar *AuthRepository) Login(email string, password string) (string, error) {
	var user _entities.User
	tx := ar.database.Where("email = ?", email).Find(&user)
	if tx.Error != nil {
		return "failed", tx.Error
	}
	if tx.RowsAffected == 0 {
		return "user not found", errors.New("user not found")
	}
	if user.Password != password {
		return "password incorrect", errors.New("password incorrect")
	}
	token, err := _middlewares.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return "create token failed", err
	}
	return token, nil
}
