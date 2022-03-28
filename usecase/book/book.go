package book

import (
	"errors"
	_entities "usamah/project-test-1/entities"
	_bookRepository "usamah/project-test-1/repository/book"
)

type BookUseCase struct {
	bookRepository _bookRepository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo _bookRepository.BookRepositoryInterface) BookUseCaseInterface {
	return &BookUseCase{
		bookRepository: bookRepo,
	}
}

func (buc *BookUseCase) GetAllBook() ([]_entities.Book, error) {
	books, err := buc.bookRepository.GetAllBook()
	return books, err
}

func (buc *BookUseCase) GetBook(id int) (_entities.Book, int, error) {
	book, rows, err := buc.bookRepository.GetBook(id)
	return book, rows, err
}

func (buc *BookUseCase) CreatBook(book _entities.Book) (_entities.Book, error) {
	book, err := buc.bookRepository.CreatBook(book)
	return book, err
}

func (buc *BookUseCase) UpdateBook(bookUpdate _entities.Book, id int, idToken int) (_entities.Book, int, error) {
	book, rows, err := buc.bookRepository.GetBook(id)
	if err != nil {
		return book, 0, err
	}
	if rows == 0 {
		return book, 0, nil
	}
	if book.UserID != uint(idToken) {
		return book, 1, errors.New("unauthorized")
	}
	if bookUpdate.Title != "" {
		book.Title = bookUpdate.Title
	}
	if bookUpdate.Author != "" {
		book.Author = bookUpdate.Author
	}
	if bookUpdate.Publisher != "" {
		book.Publisher = bookUpdate.Publisher
	}

	updateBook, updateRows, updateErr := buc.bookRepository.UpdateBook(book)
	return updateBook, updateRows, updateErr
}

func (buc *BookUseCase) DeleteBook(id int, idToken int) (int, error) {
	book, rows, err := buc.bookRepository.GetBook(id)
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, nil
	}
	if book.UserID != uint(idToken) {
		return 1, errors.New("unauthorized")
	}

	rowsDelete, errDelete := buc.bookRepository.DeleteBook(id)
	return rowsDelete, errDelete
}
