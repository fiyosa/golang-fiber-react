package routes

import (
	"go-fiber-react/app/http/controller"
	"go-fiber-react/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	r := app.Group("/api")

	Auth := middleware.Auth
	UserController := &controller.UserController{}
	AuthController := &controller.AuthController{}

	r.Post("/auth/login", AuthController.Login)
	r.Post("/auth/register", AuthController.Register)

	r.Get("/user", Auth(), UserController.Index)

	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "API not found"})
	})
}
