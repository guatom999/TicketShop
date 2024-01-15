package users

import "gorm.io/gorm"

type (
	Users struct {
		gorm.Model
		Email    string `gorm:"unique"`
		Username string `gorm:"unique"`
		Password string
	}

	UserRegisterReq struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserLoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserPassPort struct {
		Username     string `json:"username"`
		Email        string `json:"email"`
		AccessToken  string `json:"access_token" `
		ReFreshToken string `json:"refresh_token"`
	}
)
