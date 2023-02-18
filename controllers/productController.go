package controllers

import "github.com/gofiber/fiber/v2"

func GetProducts (c *fiber.Ctx) error {
	return c.Status(200).SendString("test")
}

func PostProducts (c *fiber.Ctx) error {
	return nil
}