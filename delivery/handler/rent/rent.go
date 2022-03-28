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
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success rent book", rent))
	}
}

func (rh *RentHandler) GetListRent() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		rents, err := rh.rentUseCase.GetListRent(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all rent", rents))
	}
}

func (rh *RentHandler) GetRentByID() echo.HandlerFunc {
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
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get rent", rent))
	}
}
