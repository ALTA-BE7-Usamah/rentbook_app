package book

import (
	_entities "usamah/project-test-1/entities"
)

type BookRepositoryInterface interface {
	GetAllBook() ([]_entities.Book, error)
	GetBook(id int) (_entities.Book, int, error)
	CreatBook(book _entities.Book) (_entities.Book, error)
	UpdateBook(bookUpdate _entities.Book) (_entities.Book, int, error)
	DeleteBook(id int) (int, error)
}
