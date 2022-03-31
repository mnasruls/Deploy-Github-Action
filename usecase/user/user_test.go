package user

import (
	_entities "be7/layered/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "alta", data[0].Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.GetUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "alta", data.Name)
		assert.Equal(t, "alta@mail.com", data.Email)
		assert.Equal(t, "12345", data.Password)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.GetUser(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.DeleteUser(1)
		assert.Nil(t, err)
		assert.Equal(t, "deleted", data.Name)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.DeleteUser(1)
		assert.NotNil(t, err)
		assert.Equal(t, "alta", data.Name)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.CreateUser(_entities.User{
			Name: "alta", Email: "alta@mail.com", Password: "12345"})
		assert.Nil(t, err)
		assert.Equal(t, "alta", data.Name)
		assert.Equal(t, "alta@mail.com", data.Email)
		assert.Equal(t, "12345", data.Password)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.CreateUser(_entities.User{
			Name: "alta", Email: "alta@mail.com", Password: "12345"})
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{}, data)
	})
}

func TestUpdatedUser(t *testing.T) {
	t.Run("TestGetUserSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.UpdatedUser(_entities.User{
			Name: "alta", Email: "alta@mail.com", Password: "12345",
		}, 1)
		assert.Nil(t, err)
		assert.Equal(t, "alta", data.Name)
		assert.Equal(t, "alta@mail.com", data.Email)
		assert.Equal(t, "12345", data.Password)
	})

	t.Run("TestGetUserError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.UpdatedUser(_entities.User{
			Name: "alta", Email: "alta@mail.com", Password: "12345",
		}, 1)
		assert.NotNil(t, err)
		assert.Equal(t, _entities.User{
			Name: "nasrul", Email: "nasrul@mail.com", Password: "12345",
		}, data)
	})
}

// mock succes
type mockUserRepository struct{}

func (m mockUserRepository) GetAll() ([]_entities.User, error) {
	return []_entities.User{
		{Name: "alta", Email: "alta@mail.com", Password: "12345"},
	}, nil
}

func (m mockUserRepository) GetUser(id int) (_entities.User, int, error) {
	return _entities.User{
		Name: "alta", Email: "alta@mail.com", Password: "12345",
	}, 1, nil
}

func (m mockUserRepository) DeleteUser(id int) (_entities.User, error) {
	return _entities.User{
		Name: "deleted", Email: "deleted", Password: "deleted",
	}, nil
}

func (m mockUserRepository) UpdatedUser(user _entities.User) (_entities.User, error) {
	return _entities.User{
		Name: "alta", Email: "alta@mail.com", Password: "12345",
	}, nil
}
func (m mockUserRepository) CreateUser(user _entities.User) (_entities.User, error) {
	return _entities.User{
		Name: "alta", Email: "alta@mail.com", Password: "12345",
	}, nil
}

//  mock error

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) GetAll() ([]_entities.User, error) {
	return nil, fmt.Errorf("error")
}
func (m mockUserRepositoryError) GetUser(id int) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error get data user")
}
func (m mockUserRepositoryError) DeleteUser(id int) (_entities.User, error) {
	return _entities.User{
		Name: "alta", Email: "alta@mail.com", Password: "12345",
	}, fmt.Errorf("error delete data user")
}
func (m mockUserRepositoryError) CreateUser(_entities.User) (_entities.User, error) {
	return _entities.User{}, fmt.Errorf("error create data user")
}
func (m mockUserRepositoryError) UpdatedUser(_entities.User) (_entities.User, error) {
	return _entities.User{
		Name: "nasrul", Email: "nasrul@mail.com", Password: "12345",
	}, fmt.Errorf("error updated data user")
}
