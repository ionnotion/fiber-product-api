package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/helpers"
	"github.com/ionnotion/fiber-product-api/models"
)

func Authentication (c *fiber.Ctx) error {

	err := helpers.VerifyToken(c)

	if err != nil {
		return c.Status(401).JSON(models.Response{Message: "Invalid Authorization"})
	}


	
	if c.Locals("id") == nil {
		return c.Status(401).JSON(models.Response{Message: "Unauthorized"})
	}

	return c.Next()
}