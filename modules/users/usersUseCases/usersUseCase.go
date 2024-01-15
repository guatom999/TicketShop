package usersUseCases

import (
	"context"
	"errors"
	"log"

	"github.com/guatom999/TicketShop/config"
	"github.com/guatom999/TicketShop/modules/users"
	"github.com/guatom999/TicketShop/modules/users/usersRepositories"
	"github.com/guatom999/TicketShop/pkg/authen"
	"golang.org/x/crypto/bcrypt"
)

type UsersUseCaseService interface {
	Register(pctx context.Context, req *users.UserRegisterReq) error
	Login(pctx context.Context, req *users.UserLoginReq) (*users.UserPassPort, error)
}

type userUseCase struct {
	userRepo usersRepositories.UsersRepositoryService
	cfg      *config.Config
}

func NewUserUseCase(userRepo usersRepositories.UsersRepositoryService, cfg *config.Config) UsersUseCaseService {
	return &userUseCase{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (u *userUseCase) Register(pctx context.Context, req *users.UserRegisterReq) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return errors.New("error: failed to hash password")
	}

	if err := u.userRepo.CreateUser(pctx, &users.Users{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}); err != nil {
		return err
	}

	return nil

}

func (u *userUseCase) Login(pctx context.Context, req *users.UserLoginReq) (*users.UserPassPort, error) {
	result, err := u.userRepo.CreadentialSearch(pctx, req)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(req.Password)); err != nil {
		log.Printf("Error: User Creadential doesn't match:%s", err.Error())
		return nil, errors.New("error: email or password invalid")
	}

	accessToken := u.userRepo.AccessToken(u.cfg, &authen.Claims{
		PlayerId: result.Email,
	})

	refreshToken := u.userRepo.RefreshToken(u.cfg, &authen.Claims{
		PlayerId: result.Email,
	})

	return &users.UserPassPort{
		Username:     result.Username,
		Email:        result.Email,
		AccessToken:  accessToken,
		ReFreshToken: refreshToken,
	}, nil
}
