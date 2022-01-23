package routes

import (
	"basicuserapiaccount/src/controllers"
	"basicuserapiaccount/src/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	user := api.Group("user")
	user.Post("register", controllers.Register)
	user.Post("login", controllers.Login)

	userAuthenticated := user.Use(middlewares.IsAuthenticated)
	userAuthenticated.Post("logout", controllers.Logout)
	userAuthenticated.Put("users/infosensitive", controllers.UpdateSensitiveInfo)
	userAuthenticated.Put("users/password", controllers.UpdatePassword)

}
