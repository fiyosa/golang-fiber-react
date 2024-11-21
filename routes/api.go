package routes

import (
	"go-fiber-react/app/http/controller"
	"go-fiber-react/app/http/middleware"
	"go-fiber-react/app/policy"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App) {
	r := app.Group("/api")
	Auth := middleware.Auth

	r.Post("/auth/login", controller.Auth.Login)
	r.Post("/auth/register", controller.Auth.Register)
	r.Get("/auth/user", Auth(), controller.User.Auth)

	r.Get("/user", Auth(), policy.User.Update, controller.User.Index)
	r.Get("/user/:id", Auth(), controller.User.Show)
	r.Put("/user/:id", Auth(), controller.User.Update)

	r.Get("/test/job", controller.Test.Job)
	r.Get("/test/event", controller.Test.Event)

	r.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "API not found"})
	})
}
