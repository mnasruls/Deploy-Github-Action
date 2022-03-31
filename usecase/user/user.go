package user

import (
	_entities "be7/layered/entities"
	_userRepository "be7/layered/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.userRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetUser(id int) (_entities.User, int, error) {
	users, rows, err := uuc.userRepository.GetUser(id)
	return users, rows, err
}

func (uuc *UserUseCase) DeleteUser(id int) (_entities.User, error) {
	users, err := uuc.userRepository.DeleteUser(id)
	return users, err
}
func (uuc *UserUseCase) CreateUser(user _entities.User) (_entities.User, error) {
	users, err := uuc.userRepository.CreateUser(user)
	return users, err
}
func (uuc *UserUseCase) UpdatedUser(users _entities.User, id int) (_entities.User, error) {
	users.ID = uint(id)
	users, err := uuc.userRepository.UpdatedUser(users)
	return users, err
}
