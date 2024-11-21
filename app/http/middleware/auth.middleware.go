package middleware

import (
	"go-fiber-react/app/helper"
	"go-fiber-react/app/model"
	"go-fiber-react/app/repository"
	"go-fiber-react/config"
	"go-fiber-react/lang"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	getPermissions *[]string
)

func Auth(permission ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &model.User{}
		if err, isOk := authentication(c, user); !isOk {
			return err
		}
		if err, isOk := authorization(c, permission...); !isOk {
			return err
		}
		return c.Next()
	}
}

func authentication(c *fiber.Ctx, user *model.User) (error, bool) {
	getToken := c.Get("Authorization")

	if getToken == "" {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS)), false
	}

	tokenParts := strings.Split(getToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS)), false
	}

	token := tokenParts[1]
	if _, err := Jwt.Verify(token); err != nil {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS)), false
	}

	auth := &model.Auth{}
	if err := config.G.Preload("User").Where(&model.Auth{Token: token}).First(&auth).Error; err != nil {
		return helper.Res.SendErrorMsg(c, err.Error()), false
	}
	if auth.Id == 0 {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS)), false
	}

	roles := &[]string{}
	if err := repository.Role.GetMany(auth.UserId, roles); err != nil {
		return helper.Res.SendErrorMsg(c, err.Error()), false
	}

	permissions := &[]string{}
	if err := repository.Permission.GetManyByUserId(auth.UserId, permissions); err != nil {
		return helper.Res.SendErrorMsg(c, err.Error()), false
	}
	getPermissions = permissions

	*user = auth.User
	c.Locals("user", auth.User)
	c.Locals("roles", *roles)
	c.Locals("permissions", *permissions)
	return nil, true
}

func authorization(c *fiber.Ctx, permission ...string) (error, bool) {
	if len(permission) == 0 {
		return nil, true
	}

	check := false
	for _, v := range *getPermissions {
		if v == permission[0] {
			check = true
		}
	}

	if !check {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().PERMISSION_FAILED)), false
	}

	return nil, true
}
