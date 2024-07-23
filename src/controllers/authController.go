package controllers

import (
	"fmt"

	models "godoc/src/model/sql"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	user := new(models.User)

	body := c.Body()
	fmt.Println("Body: ", string(body))

	if err := c.BodyParser(user); err != nil {
		return err
	}
	fmt.Printf("Parsed user: %+v\n", user)
	return nil
}

func Logout(c *fiber.Ctx) error {
	user := new(models.User)

	body := c.Body()
	fmt.Println("Body: ", string(body))

	if err := c.BodyParser(user); err != nil {
		return err
	}
	fmt.Printf("Parsed user: %+v\n", user)
	return nil
}

func Passcode(c *fiber.Ctx) error {
	return nil
}
