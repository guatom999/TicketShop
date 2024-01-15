package middlwareHandlers

import (
	"net/http"
	"strings"

	"github.com/guatom999/TicketShop/config"
	"github.com/guatom999/TicketShop/modules/middlewares/middlewaresUseCases"
	"github.com/labstack/echo/v4"
)

type MiddlwareHandlerService interface {
	JwtAuthorize(next echo.HandlerFunc) echo.HandlerFunc
}

type middlewareHandler struct {
	cfg               *config.Config
	middlewareUseCase middlewaresUseCases.MiddlwareUsecaseService
}

func NewMiddlewareHandler(cfg *config.Config, middlewareUseCase middlewaresUseCases.MiddlwareUsecaseService) MiddlwareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUseCase: middlewareUseCase,
	}
}

func (h *middlewareHandler) JwtAuthorize(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		clamis, err := h.middlewareUseCase.JwtAuthorization(c, h.cfg, accessToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
			// return response.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}

		return next(clamis)
	}
}
