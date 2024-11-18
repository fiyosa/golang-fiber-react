package repository

import (
	"errors"
	"go-fiber-react/app/model"
	"go-fiber-react/config"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

var User user

type user struct{}

func (*user) Auth(c *fiber.Ctx, u *model.User) error {
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

func (*user) First(c *fiber.Ctx, u *model.User) error {
	return nil
}

func (*user) Create(u *model.User) error {
	return config.G.Create(&u).Error
}
