package main

import (
	// "github.com/enrichoalkalas01/learn-go-fiber.git/routes"
	"github.com/enrichoalkalas01/learn-go-fiber.git/repositories"
	"github.com/gofiber/fiber/v2"
	// "fmt"
)

func main() {
	app := fiber.New()

	repositories.Testing()
	
	// // Use All Router From Routes
	// routes.RouterApp("/api/v1", app)

	app.Listen(":5600")
}
