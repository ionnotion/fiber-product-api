package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/configs"
	"github.com/ionnotion/fiber-product-api/controllers"
	"github.com/ionnotion/fiber-product-api/middlewares"
	"github.com/ionnotion/fiber-product-api/models"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(models.Response{Message: "Hello World"})
	})

	configs.GormConnect()

	user := app.Group("/user")
	user.Post("/login",controllers.LoginHandler)
	user.Post("/register",controllers.RegisterHandler)
	
	product := app.Group("/product",middlewares.Authentication)
	product.Get("/",controllers.GetProducts)
	product.Post("/",controllers.PostProducts)

	app.Listen(":8080")
}
