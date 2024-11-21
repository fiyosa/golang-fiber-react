package policy

import (
	"go-fiber-react/app/helper"
	"go-fiber-react/app/http/middleware"
	"go-fiber-react/app/model"
	"go-fiber-react/lang"

	"github.com/gofiber/fiber/v2"
)

var User userPolicy

type userPolicy struct{}

func (*userPolicy) Update(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	roles := c.Locals("roles").([]string)
	user_id := c.Params("id")

	if middleware.Role.IsAdmin(roles) {
		return c.Next()
	}

	if user.Id == helper.Str2Int(user_id) {
		return c.Next()
	}

	return helper.Res.SendErrorMsg(c, lang.L.Convert(lang.L.Get().UNAUTHORIZED_ACCESS), fiber.StatusUnauthorized)
}
