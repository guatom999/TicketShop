package server

import (
	"net/http"

	"github.com/guatom999/TicketShop/modules/users/userHandlers"
	"github.com/guatom999/TicketShop/modules/users/usersRepositories"
	"github.com/guatom999/TicketShop/modules/users/usersUseCases"
	"github.com/labstack/echo/v4"
)

func tesT(c echo.Context) error {
	return c.JSON(http.StatusOK, "test")
}

func (s *server) userModule() {
	userRepo := usersRepositories.NewUsersRepository(s.db)
	userUseCase := usersUseCases.NewUserUseCase(userRepo, s.cfg)
	userHandler := userHandlers.NewUserHandler(userUseCase)

	userRouter := s.app.Group("/user")

	userRouter.GET("/test", tesT, s.middleware.JwtAuthorize) // func(c echo.Context) error {
	// 	// Access request information, perform actions, and return a response
	// 	return c.String(http.StatusOK, "Hello, World!")
	// },
	// 	s.middleware.JwtAuthorize,

	userRouter.POST("/register", userHandler.Register)
	userRouter.POST("/login", userHandler.Login)
}
