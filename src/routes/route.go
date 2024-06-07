package routes

import (
	"godoc/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//authentication routes
	app.Post("/cashiers/:cashierId/login", controllers.Login)
	app.Post("/cashiers/:cashierId/logout", controllers.Logout)
	app.Get("/cashiers/:cashierId/passcode", controllers.Passcode)
}
