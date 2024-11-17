package bootstrap

import (
	"go-fiber-react/config"
	"go-fiber-react/routes"

	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	config.Env()
	config.DB()
	routes.Console()
	config.Logger()
	config.I18n()

	app := config.App()
	app.Use(config.Cors())

	routes.Api(app)
	routes.Web(app)

	return app
}
