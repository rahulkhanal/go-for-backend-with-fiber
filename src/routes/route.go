package routes

import (
	"godoc/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	//authentication routes
	app.Post("/user/:id/login", controllers.Login)
	app.Post("/user", controllers.Logout)
	app.Get("/user/:id/passcode", controllers.Passcode)
}
