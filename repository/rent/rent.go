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
	var rentFind []_entities.Rent
	// mendapatkan data Book menggunakan ID Book
	txFind := rr.database.Where("book_id = ?", bookID).Find(&rentFind)
	if txFind.Error != nil {
		return rent, txFind.Error
	}
	// memastikan apakah buku tersebut available atau not available
	for i := 0; i < len(rentFind); i++ {
		if rentFind[i].BookID == bookID && rentFind[i].ReturnStatus == "" {
			return rent, errors.New("book is not available")
		}
	}
	// mengubah status buku
	var book _entities.Book
	rr.database.Find(&book, bookID)
	book.Status = "not available"
	rr.database.Save(&book)

	// jika book yang ingin disewa available
	tx := rr.database.Save(&rent)
	if tx.Error != nil {
		return rent, tx.Error
	}
	return rent, nil
}

// mencari daftar buku yang disewa oleh user yang login
func (rr *RentRepository) GetListRent(userID uint) ([]_entities.Rent, error) {
	var rents []_entities.Rent
	tx := rr.database.Where("user_id = ?", userID).Find(&rents)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return rents, nil
}

// mencari data penyewaan buku sesuai id penyewaan
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

func (rr *RentRepository) ReturnBook(rent _entities.Rent) (_entities.Rent, int, error) {
	//mengubah status buku menjadi available
	var book _entities.Book
	rr.database.Find(&book, rent.BookID)
	book.Status = "available"
	rr.database.Save(&book)

	//mengupdate status pengembalian buku
	tx := rr.database.Save(&rent)
	if tx.Error != nil {
		return rent, 0, tx.Error
	}
	return rent, int(tx.RowsAffected), nil
}
