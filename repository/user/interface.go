package user

import (
	_entities "be7/layered/entities"
)

type UserRepositoryInterface interface {
	GetAll() ([]_entities.User, error)
	GetUser(id int) (_entities.User, int, error)
	DeleteUser(id int) (_entities.User, error)
	CreateUser(user _entities.User) (_entities.User, error)
	UpdatedUser(users _entities.User) (_entities.User, error)
}
