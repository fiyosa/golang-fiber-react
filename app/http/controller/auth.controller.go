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

type AuthController struct{}

func (*AuthController) Login(c *fiber.Ctx) error {
	res := &helper.Res{}
	l := &lang.L{}
	hash := &helper.Hash{}
	jwt := &middleware.Jwt{}
	ra := &repository.Auth{}

	validated := &request.LoginAuthReq{}
	if err := helper.Validate(c, validated); err != nil {
		return err
	}

	user := &model.User{}
	if err := config.G.Where(&model.User{Username: validated.Username}).First(&user).Error; err != nil {
		return res.SendErrorMsg(c, l.Convert(l.Get().AUTH_FAILED))
	}

	if !hash.Verify(validated.Password, user.Password) {
		res.SendErrorMsg(c, l.Convert(l.Get().AUTH_FAILED))
	}

	token, err := jwt.Create(helper.Int2Str(user.Id))
	if err != nil {
		return res.SendErrorMsg(c, err.Error(), fiber.StatusInternalServerError)
	}

	auth := &model.Auth{
		UserId: user.Id,
		Token:  token,
	}

	if err := ra.Create(auth); err != nil {
		return res.SendErrorMsg(c, err.Error(), fiber.StatusInternalServerError)
	}

	return res.SendCustom(c, request.LoginAuthRes{
		Token: token,
	}, fiber.StatusOK)
}

func (*AuthController) Register(c *fiber.Ctx) error {
	res := &helper.Res{}
	l := &lang.L{}
	hash := &helper.Hash{}
	ru := &repository.User{}

	validated := &request.RegisterReq{}
	if err := helper.Validate(c, validated); err != nil {
		return err
	}

	user := &model.User{}
	config.G.Where(&model.User{Username: validated.Username}).First(&user)
	if user.Id != 0 {
		return res.SendErrorMsg(c, l.Convert(l.Get().ALREADY_EXIST, fiber.Map{"operator": l.Get().USER}))
	}

	newPassword, err := hash.Create(validated.Password)
	if err != nil {
		return res.SendErrorMsg(c, err.Error())
	}

	user.Username = validated.Username
	user.Name = validated.Name
	user.Password = newPassword
	ru.Create(user)

	id, _ := hash.EncodeId(user.Id)
	result := request.RegisterRes{
		Data: request.UserShowRes{
			Id:        id,
			Username:  user.Username,
			Name:      user.Name,
			CreatedAt: helper.Time2Str(user.CreatedAt),
			UpdatedAt: helper.Time2Str(user.UpdatedAt),
		},
		Message: l.Convert(l.Get().SAVED_SUCCESSFULLY, fiber.Map{"operator": l.Get().USER}),
	}

	return res.SendCustom(c, result, fiber.StatusOK)
}
