package helper

import (
	"encoding/json"
	"go-fiber-react/config"
	"go-fiber-react/lang"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, input interface{}) (err error, isOk bool) {
	if err := c.BodyParser(input); err != nil {
		return generateError(c, err), false
	}

	if err := config.Validate.Struct(input); err != nil {
		return generateError(c, err), false
	}

	return nil, true
}

func generateError(c *fiber.Ctx, err error) error {
	newErrors := map[string]string{}
	msg := "Invalid data"

	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		field := strings.ToLower(v.Field)
		newErrors[field] = "Json binding error: " + field + " type error"

	case validator.ValidationErrors:
		for _, e := range v {
			field := strings.ToLower(e.Field())
			newErrors[field] = strings.ToLower(e.Translate(config.Translator))
		}

	default:
		if v != nil {
			msg = v.Error()
		} else {
			msg = lang.L.Get().SOMETHING_WENT_WRONG
		}
	}

	return Res.SendErrors(c, msg, newErrors)
}
