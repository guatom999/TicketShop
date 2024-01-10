package usersRepositories

import "gorm.io/gorm"

type UsersRepositoryService interface {
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepositoryService {
	return &usersRepository{db: db}
}
