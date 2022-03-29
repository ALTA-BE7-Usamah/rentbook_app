package routes

import (
	_authHandler "usamah/project-test-1/delivery/handler/auth"
	_bookHandler "usamah/project-test-1/delivery/handler/book"
	_rentHandler "usamah/project-test-1/delivery/handler/rent"
	_userHandler "usamah/project-test-1/delivery/handler/user"
	_middlewares "usamah/project-test-1/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterBookPath(e *echo.Echo, bh *_bookHandler.BookHandler) {
	e.GET("/books", bh.GetAllBookHandler())
	e.GET("/books/:id", bh.GetBookHandler())
	e.POST("/books", bh.CreateBookHandler(), _middlewares.JWTMiddleware())
	e.PUT("/books/:id", bh.UpdateBookHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/books/:id", bh.DeleteBookHandler(), _middlewares.JWTMiddleware())
}

func RegisterRentPath(e *echo.Echo, rh *_rentHandler.RentHandler) {
	e.POST("/rent", rh.RentBookHandler(), _middlewares.JWTMiddleware())
	e.GET("/rent", rh.GetListRentHandler(), _middlewares.JWTMiddleware())
	e.GET("/rent/:id", rh.GetRentByIDHandler(), _middlewares.JWTMiddleware())
	e.POST("/rent/:id", rh.ReturnBookHandler(), _middlewares.JWTMiddleware())
}
