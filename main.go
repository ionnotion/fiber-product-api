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

	err := configs.GormConnect()

	if err != nil {
		panic(err)
	}

	user := app.Group("/user")
	user.Post("/login",controllers.LoginHandler)
	user.Post("/register",controllers.RegisterHandler)
	
	product := app.Group("/product")
	product.Use(middlewares.Authentication)
	product.Get("/",controllers.GetProducts)
	product.Post("/",controllers.PostProducts)

	app.Listen(":8080")
}
