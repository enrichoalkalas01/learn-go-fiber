package main

import (
	"github.com/enrichoalkalas01/learn-go-fiber.git/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Use All Router From Routes
	routes.RouterApp("/api/v1", app)

	app.Listen(":5600")
}
