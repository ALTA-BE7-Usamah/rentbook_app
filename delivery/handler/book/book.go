package book

import (
	"net/http"
	"strconv"
	"usamah/project-test-1/delivery/helper"
	_middlewares "usamah/project-test-1/delivery/middlewares"
	"usamah/project-test-1/entities"
	_bookUseCase "usamah/project-test-1/usecase/book"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase _bookUseCase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase _bookUseCase.BookUseCaseInterface) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (bh *BookHandler) GetAllBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.bookUseCase.GetAllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		responseBooks := []map[string]interface{}{}
		for i := 0; i < len(books); i++ {
			response := map[string]interface{}{
				"ID":        books[i].ID,
				"Title":     books[i].Title,
				"Catagory":  books[i].Catagory,
				"Author":    books[i].Author,
				"Publisher": books[i].Publisher,
				"Status":    books[i].Status,
			}
			responseBooks = append(responseBooks, response)
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all books", responseBooks))
	}
}

func (bh *BookHandler) GetBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		book, rows, err := bh.bookUseCase.GetBook(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		responseBook := map[string]interface{}{
			"ID":        book.ID,
			"Title":     book.Title,
			"Catagory":  book.Catagory,
			"Author":    book.Author,
			"Publisher": book.Publisher,
			"Status":    book.Status,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get user", responseBook))
	}
}

func (bh *BookHandler) CreateBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var newBook entities.Book
		err := c.Bind(&newBook)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}
		newBook.UserID = uint(idToken)
		_, error := bh.bookUseCase.CreatBook(newBook)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create user"))
	}
}

func (bh *BookHandler) UpdateBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		var updateBook entities.Book
		errBind := c.Bind(&updateBook)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		book, rows, error := bh.bookUseCase.UpdateBook(updateBook, id, idToken)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		responseBook := map[string]interface{}{
			"ID":        book.ID,
			"Title":     book.Title,
			"Catagory":  book.Catagory,
			"Author":    book.Author,
			"Publisher": book.Publisher,
			"Status":    book.Status,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update book", responseBook))
	}
}

func (bh *BookHandler) DeleteBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		rows, err := bh.bookUseCase.DeleteBook(id, idToken)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("unauthorized"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success deleted book"))
	}
}
