package routes

import (
	"basicuserapiaccount/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Get("ping", controllers.Ping)
	api.Post("createuser", controllers.Register)
	api.Get("users", controllers.GetAllUser)
}
