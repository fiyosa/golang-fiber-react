package controller

import (
	"go-fiber-react/app/helper"
	"go-fiber-react/app/model"
	"go-fiber-react/app/repository"
	"go-fiber-react/app/request"
	"go-fiber-react/config"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (*UserController) Auth(c *fiber.Ctx) error {
	res := &helper.Res{}
	l := &lang.L{}
	ru := &repository.User{}
	hash := &helper.Hash{}

	user := &model.User{}
	if err := ru.First(c, user); err != nil {
		return err
	}

	id, _ := hash.EncodeId(user.Id)
	return res.SendData(
		c,
		l.Convert(l.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": l.Get().USER}),
		&request.UserShowRes{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
	)
}

func (*UserController) Index(c *fiber.Ctx) error {
	res := &helper.Res{}
	l := &lang.L{}

	config.Log("this is log")
	config.Logf("this is log: %v", "format")

	return res.SendSuccess(c, l.Convert(l.Get().SAVED_SUCCESSFULLY, fiber.Map{"operator": l.Get().USER}))
}
