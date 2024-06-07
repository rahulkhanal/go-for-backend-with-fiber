package main

import (
	"godoc/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load(".env")      // load env
	routes.RegisterRoutes(app) //routing
	println("khai" + os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
