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

func Auth(permission ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &model.User{}
		if err := authentication(c, user); err != nil {
			return err
		}
		if err := authorization(c, user, permission...); err != nil {
			return err
		}
		return c.Next()
	}
}

func authorization(c *fiber.Ctx, user *model.User, permissions ...string) error {
	if len(permissions) == 0 {
		return nil
	}

	roles := []string{}
	repository.Role.GetMany(user.Id, &roles)

	getPermissions := []string{}
	repository.Permission.GetPermissions(roles, &getPermissions)

	check := false
	for _, v := range getPermissions {
		if v == permissions[0] {
			check = true
		}
	}

	if !check {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().PERMISSION_FAILED))
	}

	return nil
}

func authentication(c *fiber.Ctx, user *model.User) error {
	getToken := c.Get("Authorization")
	jwt := &Jwt{}

	if getToken == "" {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS))
	}

	tokenParts := strings.Split(getToken, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS))
	}

	token := tokenParts[1]
	if _, err := jwt.Verify(token); err != nil {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS))
	}

	auth := &model.Auth{}
	config.G.Preload("User").Where(&model.Auth{Token: token}).First(&auth)
	if auth.Id == 0 {
		return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS))
	}

	*user = auth.User
	config.Log(auth)
	c.Locals("user", auth.User)
	return nil
}
