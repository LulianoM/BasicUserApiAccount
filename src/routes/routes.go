package routes

import (
	"basicuserapiaccount/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Get("ping", controllers.Ping)
}
