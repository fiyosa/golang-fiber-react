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

var UserController userController

type userController struct{}

func (*userController) Auth(c *fiber.Ctx) error {
	user := &model.User{}
	if err := repository.User.Auth(c, user); err != nil {
		return helper.Res.SendErrorMsg(c, err.Error())
	}

	id, _ := helper.Hash.EncodeId(user.Id)
	return helper.Res.SendData(
		c,
		lang.L.Convert(lang.L.Get().RETRIEVED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}),
		&request.UserShowRes{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
	)
}

func (*userController) Index(c *fiber.Ctx) error {
	config.Log("this is log")
	config.Logf("this is log: %v", "format")

	return helper.Res.SendSuccess(c, lang.L.Convert(lang.L.Get().SAVED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}))
}
