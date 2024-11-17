package routes

import "github.com/gofiber/fiber/v2"

func Web(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		if c.Path() != "/" && len(c.Path()) >= 4 && c.Path()[:4] == "/api" {
			return c.Next()
		}

		return c.SendFile("./resources/view/index.html")
	})
}
