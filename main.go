package main

import (
	"fmt"
	"log"
	"usamah/project-test-1/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authHandler "usamah/project-test-1/delivery/handler/auth"
	_authRepository "usamah/project-test-1/repository/auth"
	_authUseCase "usamah/project-test-1/usecase/auth"

	_userHandler "usamah/project-test-1/delivery/handler/user"
	_userRepository "usamah/project-test-1/repository/user"
	_userUseCase "usamah/project-test-1/usecase/user"

	_bookHandler "usamah/project-test-1/delivery/handler/book"
	_bookRepository "usamah/project-test-1/repository/book"
	_bookUseCase "usamah/project-test-1/usecase/book"

	_rentHandler "usamah/project-test-1/delivery/handler/rent"
	_rentRepository "usamah/project-test-1/repository/rent"
	_rentUseCase "usamah/project-test-1/usecase/rent"

	_middlewares "usamah/project-test-1/delivery/middlewares"
	_routes "usamah/project-test-1/delivery/routes"
	_utils "usamah/project-test-1/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	bookRepo := _bookRepository.NewBookRepository(db)
	bookUseCase := _bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := _bookHandler.NewBookHandler(bookUseCase)

	rentRepo := _rentRepository.NewRentRepository(db)
	rentUseCase := _rentUseCase.NewRentUseCase(rentRepo)
	rentHandler := _rentHandler.NewRentHandler(rentUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())

	_routes.RegisterUserPath(e, userHandler)
	_routes.RegisterBookPath(e, bookHandler)
	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterRentPath(e, rentHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
