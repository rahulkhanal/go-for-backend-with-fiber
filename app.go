package main

import (
	db "godoc/src/config"
	route "godoc/src/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	app := fiber.New()
	db.Connect()          //database
	godotenv.Load(".env") //load env
	route.Setup(app)      //routing
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
