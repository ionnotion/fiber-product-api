package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/configs"
	"github.com/ionnotion/fiber-product-api/models"
)

func LoginHandler(c *fiber.Ctx) error {

	var input models.UserForm

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	if err := configs.DB.Where("username=?", input.Username).First(&input).Error; err != nil {
		return c.Status(404).JSON(models.Response{Message: "Invalid username / password"})
	}

	return c.Status(200).JSON(input)
}

func RegisterHandler(c *fiber.Ctx) error {

	var input models.UserForm

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	newUser := models.User{Username: input.Username, Password: input.Password}
	newUser.BeforeSave()
	configs.DB.Create(&newUser)

	return nil
}
