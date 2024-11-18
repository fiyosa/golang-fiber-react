package helper

import (
	"encoding/json"
	"go-fiber-react/config"
	"go-fiber-react/lang"

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
	msg := ""

	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		newMsg := "Json binding error: " + v.Field + " type error"
		newErrors[v.Field] = newMsg
		msg = "Invalid data"

	case validator.ValidationErrors:
		for _, e := range v {
			newMsg := e.Translate(config.Translator)
			newErrors[e.Field()] = newMsg
		}
		msg = "Invalid data"

	default:
		if v != nil {
			msg = v.Error()
		} else {
			msg = lang.L.Get().SOMETHING_WENT_WRONG
		}
	}

	return Res.SendErrors(c, msg, newErrors)
}
