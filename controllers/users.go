package controllers

import "github.com/gofiber/fiber/v2"

func UserReadList(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to get User",
		"status":  200,
	})
}

func UserReadDetail(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to get User detail",
		"status":  200,
	})
}

func UserCreate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to create User",
		"status":  200,
	})
}

func UserUpdate(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to update User",
		"status":  200,
	})
}

func UserDelete(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Successfull to delete User",
		"status":  200,
	})
}
