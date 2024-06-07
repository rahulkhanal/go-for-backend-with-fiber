package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items = []Item{}

func registerItemRoutes(app *fiber.App) {
	itemGroup := app.Group("/items")

	itemGroup.Get("/", getAllItems)
	itemGroup.Post("/", createItem)
	itemGroup.Patch("/:id", updateItem)
	itemGroup.Delete("/:id", deleteItem)
}

func getAllItems(c *fiber.Ctx) error {
	return c.JSON(items)
}

func createItem(c *fiber.Ctx) error {
	item := new(Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	item.ID = len(items) + 1
	items = append(items, *item)
	return c.Status(fiber.StatusCreated).JSON(item)
}

func updateItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	item := new(Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, existingItem := range items {
		if existingItem.ID == id {
			items[i] = *item
			items[i].ID = id
			return c.JSON(items[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Item not found")
}

func deleteItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return c.SendString("Item deleted")
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Item not found")
}
