package routes

import (
	"basicuserapiaccount/src/controllers"
	"basicuserapiaccount/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)

	user := api.Group("user")
	userAuthenticated := user.Use(middlewares.IsAuthenticated)
	userAuthenticated.Post("logout", controllers.Logout)
	userAuthenticated.Put("users/infosensitive", controllers.UpdateSensitiveInfo)
	userAuthenticated.Put("users/password", controllers.UpdatePassword)

	api.Get("ping", controllers.Ping)
}
