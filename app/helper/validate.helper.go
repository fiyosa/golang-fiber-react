package helper

import (
	"encoding/json"
	"go-fiber-react/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, input interface{}) error {
	if err := c.BodyParser(input); err != nil {
		return generateError(c, err)
	}

	if err := config.Validate.Struct(input); err != nil {
		return generateError(c, err)
	}

	return nil
}

func generateError(c *fiber.Ctx, err error) error {
	res := &Res{}
	newErrors := map[string]string{}
	msg := ""

	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		newMsg := "Json binding error: " + v.Field + " type error"
		newErrors[v.Field] = newMsg

	case validator.ValidationErrors:
		for _, e := range v {
			newMsg := e.Translate(config.Translator)
			newErrors[e.Field()] = newMsg
		}

	default:
		if v != nil {
			msg = v.Error()
		} else {
			msg = "Invalid data."
		}
	}

	return res.SendErrors(c, msg, newErrors)
}
