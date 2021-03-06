package rent

import (
	_entities "usamah/project-test-1/entities"
)

type RentUseCaseInterface interface {
	RentBook(rent _entities.Rent, bookID uint) (_entities.Rent, error)
	GetListRent(userID uint) ([]_entities.Rent, error)
	GetRentByID(id uint, idToken uint) (_entities.Rent, int, error)
	ReturnBook(id uint, idToken uint) (_entities.Rent, int, error)
}
