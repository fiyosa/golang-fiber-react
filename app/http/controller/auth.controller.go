package controller

import (
	"go-fiber-react/app/helper"
	"go-fiber-react/app/http/middleware"
	"go-fiber-react/app/http/request/request_auth"
	"go-fiber-react/app/http/resource/resource_auth"
	"go-fiber-react/app/http/resource/resource_user"
	"go-fiber-react/app/model"
	"go-fiber-react/app/repository"
	"go-fiber-react/config"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

var Auth authController

type authController struct{}

func (*authController) Login(c *fiber.Ctx) error {
	validated := &request_auth.Login{}
	if err, isOk := helper.Validate(c, validated); !isOk {
		return err
	}

	user := &model.User{}
	if err := config.G.Where(&model.User{Username: validated.Username}).First(&user).Error; err != nil {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().AUTH_FAILED))
	}

	if !helper.Hash.Verify(validated.Password, user.Password) {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().AUTH_FAILED))
	}

	hashId, err := helper.Hash.EncodeId(user.Id)
	if err != nil {
		return helper.Res.SendException(c, err)
	}

	token, err := middleware.Jwt.Create(hashId)
	if err != nil {
		return helper.Res.SendException(c, err)
	}

	auth := &model.Auth{
		UserId: user.Id,
		Token:  token,
	}

	if err := repository.Auth.Create(auth); err != nil {
		return helper.Res.SendErrorMsg(c, err.Error(), fiber.StatusInternalServerError)
	}

	return helper.Res.SendCustom(c, resource_auth.Login{
		Token: token,
	}, fiber.StatusOK)
}

func (*authController) Register(c *fiber.Ctx) error {
	validated := &request_auth.Register{}
	if err, isOk := helper.Validate(c, validated); !isOk {
		return err
	}

	tx := config.G.Begin()

	user := &model.User{}
	config.G.Where(&model.User{Username: validated.Username}).First(&user)
	if user.Id != 0 {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().ALREADY_EXIST, fiber.Map{"operator": lang.L.Get().USER}))
	}

	newPassword, err := helper.Hash.Create(validated.Password)
	if err != nil {
		return helper.Res.SendException(c, err)
	}

	user.Username = validated.Username
	user.Name = validated.Name
	user.Password = newPassword
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return helper.Res.SendException(c, err)
	}

	roleUser := &model.Role{}
	if err := tx.Model(&model.Role{}).Where("name = ?", "user").First(roleUser).Error; err != nil {
		tx.Rollback()
		return helper.Res.SendException(c, err)
	}

	uhr := &model.UserHasRole{
		UserId: user.Id,
		RoleId: roleUser.Id,
	}
	if err := tx.Create(uhr).Error; err != nil {
		tx.Rollback()
		return helper.Res.SendException(c, err)
	}

	tx.Commit()

	id, _ := helper.Hash.EncodeId(user.Id)
	result := resource_auth.Register{
		Data: resource_user.Show{
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
