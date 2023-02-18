package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/configs"
	"github.com/ionnotion/fiber-product-api/models"
)

func GetProducts (c *fiber.Ctx) error {
	var id = c.Locals("id").(int64)
	var products []models.Product

	configs.DB.Where("user_id=?",id).Find(&products)

	return c.Status(200).JSON(products)
}

func PostProducts (c *fiber.Ctx) error {
	var input models.ProductForm

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	var id = c.Locals("id").(int64)
		
	var newProduct models.Product
	if err := c.BodyParser(&newProduct); err != nil {
		return err
	}
	
	newProduct = models.Product{ProductName: newProduct.ProductName,ProductAmount: newProduct.ProductAmount,ProductPrice: input.ProductPrice,UserId: id}

	err := configs.DB.Create(&newProduct)

	if err.Error != nil {
		return c.Status(400).JSON(models.Response{Message: err.Error.Error()})
	}

	return c.Status(200).JSON(models.Response{Message: "Post New Product Success"})
}