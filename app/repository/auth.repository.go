package repository

import (
	"go-fiber-react/app/model"
	"go-fiber-react/config"
)

type Auth struct{}

func (*Auth) Create(u *model.Auth) error {
	return config.G.Create(&u).Error
}
