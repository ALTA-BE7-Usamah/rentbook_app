package user

import (
	_entities "usamah/project-test-1/entities"
	_userRepository "usamah/project-test-1/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) CreatUser(user _entities.User) (_entities.User, error) {
	user, err := uuc.userRepository.CreatUser(user)
	return user, err
}

func (uuc *UserUseCase) GetUser(id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetUser(id)
	return user, rows, err
}

func (uuc *UserUseCase) UpdateUser(userUpdate _entities.User, id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetUser(id)
	if err != nil {
		return user, 0, err
	}
	if rows == 0 {
		return user, 0, nil
	}
	if userUpdate.Name != "" {
		user.Name = userUpdate.Name
	}
	if userUpdate.Email != "" {
		user.Email = userUpdate.Email
	}
	if userUpdate.Password != "" {
		user.Password = userUpdate.Password
	}

	updateUser, updateRows, updateErr := uuc.userRepository.UpdateUser(user)
	return updateUser, updateRows, updateErr
}

func (uuc *UserUseCase) DeleteUser(id int) (int, error) {
	rows, err := uuc.userRepository.DeleteUser(id)
	return rows, err
}
