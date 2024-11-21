package repository

import (
	"errors"
	"go-fiber-react/app/model"
	"go-fiber-react/config"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

var Auth authRepository

type authRepository struct{}

func (*authRepository) User(c *fiber.Ctx, u *model.User) error {
	user := c.Locals("user")
	if user == nil {
		return errors.New(lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS))
	}

	userObj, ok := user.(model.User)
	if !ok {
		return errors.New(lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS))
	}

	*u = userObj
	return nil
}

func (*authRepository) Create(u *model.Auth) error {
	return config.G.Create(&u).Error
}
