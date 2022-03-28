package user

import (
	_entities "usamah/project-test-1/entities"
)

type UserUseCaseInterface interface {
	CreatUser(user _entities.User) (_entities.User, error)
	GetUser(id int) (_entities.User, int, error)
	UpdateUser(userUpdate _entities.User, id int) (_entities.User, int, error)
	DeleteUser(id int) (int, error)
}
