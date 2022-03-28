package rent

import (
	"errors"
	_entities "usamah/project-test-1/entities"

	"gorm.io/gorm"
)

type RentRepository struct {
	database *gorm.DB
}

func NewRentRepository(db *gorm.DB) *RentRepository {
	return &RentRepository{
		database: db,
	}
}

func (rr *RentRepository) RentBook(rent _entities.Rent, bookID uint) (_entities.Rent, error) {
	var rentFind _entities.Rent
	txFind := rr.database.Where("book_id = ?", bookID).Find(&rentFind)
	if txFind.Error != nil {
		return rent, txFind.Error
	}
	if rentFind.BookID == bookID {
		return rent, errors.New("book is not available")
	}

	tx := rr.database.Save(&rent)
	if tx.Error != nil {
		return rent, tx.Error
	}
	return rent, nil
}

func (rr *RentRepository) GetListRent(userID uint) ([]_entities.Rent, error) {
	var rents []_entities.Rent
	tx := rr.database.Where("user_id = ?", userID).Find(&rents)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rents, nil
}

func (rr *RentRepository) GetRentByID(id uint, idToken uint) (_entities.Rent, int, error) {
	var rent _entities.Rent
	tx := rr.database.Find(&rent, id)
	if tx.Error != nil {
		return rent, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return rent, 0, nil
	}
	if rent.UserID != idToken {
		return rent, 0, errors.New("id not recognise")
	}
	return rent, int(tx.RowsAffected), nil
}
