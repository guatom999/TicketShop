package userHandlers

import (
	"context"

	"github.com/guatom999/TicketShop/modules/users"
	"github.com/guatom999/TicketShop/modules/users/usersUseCases"
	"github.com/labstack/echo/v4"
)

type UserHandlerService interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
}

type userHandler struct {
	userUseCase usersUseCases.UsersUseCaseService
}

func NewUserHandler(userUseCase usersUseCases.UsersUseCaseService) UserHandlerService {
	return &userHandler{userUseCase: userUseCase}
}

func (h *userHandler) Register(c echo.Context) error {

	ctx := context.Background()

	req := new(users.UserRegisterReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, echo.ErrBadRequest.Error())
	}

	if err := h.userUseCase.Register(ctx, req); err != nil {
		return c.JSON(404, err.Error())
	}

	return c.JSON(200, "Resgietr Success")

}

func (h *userHandler) Login(c echo.Context) error {
	ctx := context.Background()

	req := new(users.UserLoginReq)

	if err := c.Bind(req); err != nil {
		return c.JSON(400, echo.ErrBadRequest.Error())
	}

	result, err := h.userUseCase.Login(ctx, req)
	if err != nil {
		return c.JSON(404, err.Error())
	}

	return c.JSON(200, result)
}
