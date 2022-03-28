package book

import (
	_entities "usamah/project-test-1/entities"

	"gorm.io/gorm"
)

type BookRepository struct {
	database *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		database: db,
	}
}

func (br *BookRepository) GetAllBook() ([]_entities.Book, error) {
	var books []_entities.Book
	tx := br.database.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return books, nil
}

func (br *BookRepository) GetBook(id int) (_entities.Book, int, error) {
	var book _entities.Book
	tx := br.database.Find(&book, id)
	if tx.Error != nil {
		return book, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return book, 0, nil
	}
	return book, int(tx.RowsAffected), nil
}

func (br *BookRepository) CreatBook(book _entities.Book) (_entities.Book, error) {
	tx := br.database.Save(&book)
	if tx.Error != nil {
		return book, tx.Error
	}
	return book, nil
}

func (br *BookRepository) UpdateBook(bookUpdate _entities.Book) (_entities.Book, int, error) {
	tx := br.database.Save(&bookUpdate)
	if tx.Error != nil {
		return bookUpdate, 0, tx.Error
	}
	return bookUpdate, int(tx.RowsAffected), nil
}

func (br *BookRepository) DeleteBook(id int) (int, error) {
	var book _entities.Book
	tx := br.database.Delete(&book, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
