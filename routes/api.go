package routes

import (
	"go-fiber-react/app/http/controller"
	"go-fiber-react/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	r := app.Group("/api")

	Auth := middleware.Auth

	r.Post("/auth/login", controller.AuthController.Login)
	r.Post("/auth/register", controller.AuthController.Register)
	r.Get("/auth/user", Auth(), controller.UserController.Auth)

	r.Get("/test/job", controller.TestController.Job)
	r.Get("/test/event", controller.TestController.Event)

	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "API not found"})
	})
}
