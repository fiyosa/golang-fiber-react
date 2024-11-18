package controller

import (
	"go-fiber-react/app/helper"
	"go-fiber-react/app/http/middleware"
	"go-fiber-react/app/model"
	"go-fiber-react/app/repository"
	"go-fiber-react/app/request"
	"go-fiber-react/config"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

var AuthController authController

type authController struct{}

func (*authController) Login(c *fiber.Ctx) error {
	jwt := &middleware.Jwt{}

	validated := &request.LoginAuthReq{}
	if err, isOk := helper.Validate(c, validated); !isOk {
		return err
	}

	user := &model.User{}
	if err := config.G.Where(&model.User{Username: validated.Username}).First(&user).Error; err != nil {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().AUTH_FAILED))
	}

	if !helper.Hash.Verify(validated.Password, user.Password) {
		helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().AUTH_FAILED))
	}

	token, err := jwt.Create(helper.Int2Str(user.Id))
	if err != nil {
		return helper.Res.SendErrorMsg(c, err.Error(), fiber.StatusInternalServerError)
	}

	auth := &model.Auth{
		UserId: user.Id,
		Token:  token,
	}

	if err := repository.Auth.Create(auth); err != nil {
		return helper.Res.SendErrorMsg(c, err.Error(), fiber.StatusInternalServerError)
	}

	return helper.Res.SendCustom(c, request.LoginAuthRes{
		Token: token,
	}, fiber.StatusOK)
}

func (*authController) Register(c *fiber.Ctx) error {
	validated := &request.RegisterReq{}
	if err, isOk := helper.Validate(c, validated); !isOk {
		return err
	}

	user := &model.User{}
	config.G.Where(&model.User{Username: validated.Username}).First(&user)
	if user.Id != 0 {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().ALREADY_EXIST, fiber.Map{"operator": lang.L.Get().USER}))
	}

	newPassword, err := helper.Hash.Create(validated.Password)
	if err != nil {
		return helper.Res.SendErrorMsg(c, err.Error())
	}

	user.Username = validated.Username
	user.Name = validated.Name
	user.Password = newPassword
	repository.User.Create(user)

	id, _ := helper.Hash.EncodeId(user.Id)
	result := request.RegisterRes{
		Data: request.UserShowRes{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
		Message: lang.L.Convert(lang.L.Get().SAVED_SUCCESSFULLY, fiber.Map{"operator": lang.L.Get().USER}),
	}

	return helper.Res.SendCustom(c, result, fiber.StatusOK)
}
