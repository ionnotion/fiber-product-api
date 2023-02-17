package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/models"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(models.Response{Message: "Hello World"})
	})

	app.Listen(":8080")
}
