package repository

import (
	"go-fiber-react/app/model"
	"go-fiber-react/config"
)

var Auth auth

type auth struct{}

func (*auth) Create(u *model.Auth) error {
	return config.G.Create(&u).Error
}
