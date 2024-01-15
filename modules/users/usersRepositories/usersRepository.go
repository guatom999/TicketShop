package usersRepositories

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/guatom999/TicketShop/config"
	"github.com/guatom999/TicketShop/modules/users"
	"github.com/guatom999/TicketShop/pkg/authen"
	"gorm.io/gorm"
)

type UsersRepositoryService interface {
	CreateUser(pctx context.Context, req *users.Users) error
	CreadentialSearch(pctx context.Context, req *users.UserLoginReq) (*users.Users, error)
	AccessToken(cfg *config.Config, claims *authen.Claims) string
	RefreshToken(cfg *config.Config, claims *authen.Claims) string
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepositoryService {
	return &usersRepository{db: db}
}

func (r *usersRepository) CreateUser(pctx context.Context, req *users.Users) error {

	tx := r.db.Create(req)

	if tx.Error != nil {
		log.Println("Error: Failed to Register:", tx.Error)
		return errors.New("error: failed to register")
	}

	return nil
}

func (r *usersRepository) CreadentialSearch(pctx context.Context, req *users.UserLoginReq) (*users.Users, error) {
	//Login

	user := new(users.Users)

	tx := r.db.Where("email = ? ", strings.ToLower(req.Email)).First(&user)
	if tx.Error != nil {
		log.Println("Error: Failed To Login")
		return nil, errors.New("error: failed to login")
	}

	return user, nil
}

func (r *usersRepository) AccessToken(cfg *config.Config, claims *authen.Claims) string {

	return authen.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &authen.Claims{
		PlayerId: claims.PlayerId,
	}).SignToken()

}

func (r *usersRepository) RefreshToken(cfg *config.Config, claims *authen.Claims) string {

	return authen.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &authen.Claims{
		PlayerId: claims.PlayerId,
	}).SignToken()

}
