package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/fiber-product-api/configs"
	"github.com/ionnotion/fiber-product-api/helpers"
	"github.com/ionnotion/fiber-product-api/models"
)

func LoginHandler(c *fiber.Ctx) error {

	var input models.UserForm

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	// fmt.Println(input)

	var foundUser models.User

	if err := configs.DB.Where("username=?", input.Username).First(&foundUser).Error; err != nil {
		return c.Status(404).JSON(models.Response{Message: "Invalid username / password"})
	}

	// fmt.Println(foundUser)

	err := helpers.VerifyPassword(foundUser.Password,input.Password)

	// fmt.Println(err)

	if err != nil {
		return c.Status(404).JSON(models.Response{Message: "Invalid username / password"})
	}

	token, err := helpers.GenerateToken(foundUser.Id)

	if err != nil {
		return c.Status(404).JSON(models.Response{Message: "Invalid username / password"})
	}

	return c.Status(200).JSON(models.Response{Message: "Login Success", LoggedUser: foundUser.Username, Access_token: token})
}

func RegisterHandler(c *fiber.Ctx) error {

	var input models.UserForm

	if err := c.BodyParser(&input); err != nil {
		return err
	}

	newUser := new(models.User)
	if err := c.BodyParser(newUser); err != nil {
		return err
	}

	err := configs.DB.Create(&newUser)

	if err.Error != nil {
		return c.Status(400).JSON(models.Response{Message: err.Error.Error()})
	}

	return c.Status(201).JSON(models.Response{Message: "Register Success"})
}
