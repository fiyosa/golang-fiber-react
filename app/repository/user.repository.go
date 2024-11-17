package repository

import (
	"go-fiber-react/app/helper"
	"go-fiber-react/app/model"
	"go-fiber-react/config"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

type User struct{}

func (*User) Auth(c *fiber.Ctx, u *model.User) error {
	res := &helper.Res{}
	l := &lang.L{}
	user := c.Locals("user")

	if user == "" {
		return res.SendErrorMsg(c, l.Convert(l.Get().UNAUTHORIZED_ACCESS))
	}
	userObj, ok := user.(model.User)
	if !ok {
		return res.SendErrorMsg(c, l.Convert(l.Get().UNAUTHORIZED_ACCESS))
	}
	*u = userObj
	return nil
}

func (*User) First(c *fiber.Ctx, u *model.User) error {
	res := &helper.Res{}
	l := &lang.L{}

	user := c.Locals("user")
	if user == "" {
		return res.SendErrorMsg(c, l.Convert(l.Get().UNAUTHORIZED_ACCESS))
	}
	userObj, ok := user.(model.User)
	if !ok {
		return res.SendErrorMsg(c, l.Convert(l.Get().UNAUTHORIZED_ACCESS))
	}

	*u = userObj
	return nil
}

func (*User) Create(u *model.User) error {
	return config.G.Create(&u).Error
}
