package rent

import (
	"net/http"
	"strconv"
	"usamah/project-test-1/delivery/helper"
	_middlewares "usamah/project-test-1/delivery/middlewares"
	"usamah/project-test-1/entities"
	_rentUseCase "usamah/project-test-1/usecase/rent"

	"github.com/labstack/echo/v4"
)

type RentHandler struct {
	rentUseCase _rentUseCase.RentUseCaseInterface
}

func NewRentHandler(rentUseCase _rentUseCase.RentUseCaseInterface) *RentHandler {
	return &RentHandler{
		rentUseCase: rentUseCase,
	}
}

func (rh *RentHandler) RentBookHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var newRent entities.Rent
		err := c.Bind(&newRent)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		newRent.UserID = uint(idToken)
		bookID := newRent.BookID

		rent, error := rh.rentUseCase.RentBook(newRent, bookID)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("not available"))
		}
		responseRent := map[string]interface{}{
			"ID":      rent.ID,
			"user_id": rent.UserID,
			"book_id": rent.BookID,
			"address": rent.Address,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success rent book", responseRent))
	}
}

func (rh *RentHandler) GetListRentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		rents, err := rh.rentUseCase.GetListRent(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseRent := []map[string]interface{}{}
		for i := 0; i < len(rents); i++ {
			response := map[string]interface{}{
				"ID":      rents[i].ID,
				"user_id": rents[i].UserID,
				"book_id": rents[i].BookID,
				"user": map[string]interface{}{
					"ID":    rents[i].User.ID,
					"name":  rents[i].User.Name,
					"email": rents[i].User.Email},
				"book": map[string]interface{}{
					"ID":        rents[i].Book.ID,
					"title":     rents[i].Book.Title,
					"author":    rents[i].Book.Author,
					"publisher": rents[i].Book.Publisher},
			}
			responseRent = append(responseRent, response)
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all rent", responseRent))
	}
}

func (rh *RentHandler) GetRentByIDHandler() echo.HandlerFunc {
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
		rent, rows, err := rh.rentUseCase.GetRentByID(uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		responseRent := map[string]interface{}{
			"ID":      rent.ID,
			"user_id": rent.UserID,
			"book_id": rent.BookID,
			"user": map[string]interface{}{
				"ID":    rent.User.ID,
				"name":  rent.User.Name,
				"email": rent.User.Email},
			"book": map[string]interface{}{
				"ID":        rent.Book.ID,
				"title":     rent.Book.Title,
				"author":    rent.Book.Author,
				"publisher": rent.Book.Publisher},
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get rent", responseRent))
	}
}

func (rh *RentHandler) ReturnBookHandler() echo.HandlerFunc {
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

		_, rows, error := rh.rentUseCase.ReturnBook(uint(id), uint(idToken))
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success return book"))
	}
}
