package controllers

import "github.com/gofiber/fiber/v2"

func ProductReadList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to get Product",
		"status":  200,
	})
}

func ProductReadDetail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to get Product detail",
		"status":  200,
	})
}

func ProductCreate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to create Product",
		"status":  200,
	})
}

func ProductUpdate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to update Product",
		"status":  200,
	})
}

func ProductDelete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to delete Product",
		"status":  200,
	})
}
