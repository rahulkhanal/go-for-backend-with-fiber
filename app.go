package main

import (
	"godoc/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":9000"))
}
