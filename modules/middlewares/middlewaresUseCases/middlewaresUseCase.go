package middlewaresUseCases

import (
	"log"

	"github.com/guatom999/TicketShop/config"
	"github.com/guatom999/TicketShop/pkg/authen"
	"github.com/labstack/echo/v4"
)

type MiddlwareUsecaseService interface {
	JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error)
}

type middlewareUsecase struct {
	// middlewareRepo middlewareRepositories.MiddlwareRepositoryService
}

func NewMiddlewareUsecase(
// middlewareRepo middlewareRepositories.MiddlwareRepositoryService
) MiddlwareUsecaseService {
	return &middlewareUsecase{
		// middlewareRepo: middlewareRepo,
	}
}

func (u *middlewareUsecase) JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error) {

	claims, err := authen.ParseToken(cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		return nil, err
	}

	log.Println("claims is", claims)

	return c, nil

}
