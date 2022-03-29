package rent

import (
	_entities "usamah/project-test-1/entities"
	_rentRepository "usamah/project-test-1/repository/rent"
)

type RentUseCase struct {
	rentRepository _rentRepository.RentRepositoryInterface
}

func NewRentUseCase(rentRepo _rentRepository.RentRepositoryInterface) RentUseCaseInterface {
	return &RentUseCase{
		rentRepository: rentRepo,
	}
}

func (ruc *RentUseCase) RentBook(rent _entities.Rent, bookID uint) (_entities.Rent, error) {
	rent, err := ruc.rentRepository.RentBook(rent, bookID)
	return rent, err
}

func (ruc *RentUseCase) GetListRent(userID uint) ([]_entities.Rent, error) {
	rents, err := ruc.rentRepository.GetListRent(userID)
	return rents, err
}

func (ruc *RentUseCase) GetRentByID(id uint, idToken uint) (_entities.Rent, int, error) {
	rent, rows, err := ruc.rentRepository.GetRentByID(id, idToken)
	return rent, rows, err
}

func (ruc *RentUseCase) ReturnBook(rent _entities.Rent, id uint, idToken uint) (_entities.Rent, int, error) {
	rentReturn, rows, err := ruc.rentRepository.GetRentByID(uint(id), uint(idToken))
	if err != nil {
		return rent, 0, err
	}
	if rows == 0 {
		return rent, 0, nil
	}
	if rent.ReturnStatus != "" {
		rentReturn.ReturnStatus = rent.ReturnStatus
	}

	updateRent, updateRows, updateErr := ruc.rentRepository.ReturnBook(rentReturn)
	return updateRent, updateRows, updateErr
}
