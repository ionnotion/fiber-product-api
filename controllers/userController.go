package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/configs"
	"github.com/ionnotion/fiber-product-api/models"
)

func LoginHandler (c *fiber.Ctx) error {

	var input models.UserForm

	if err := c.BodyParser(&input); err!= nil{
		return err
	}

	newUser := models.User{Username: input.Username, Password: input.Password}
	configs.DB.Create(&newUser)

	return c.Status(200).JSON(newUser)
}

func RegisterHandler (c *fiber.Ctx) error {
	return nil
}