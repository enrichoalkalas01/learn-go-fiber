package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/enrichoalkalas01/learn-go-fiber.git/models"
	"github.com/enrichoalkalas01/learn-go-fiber.git/routes"
	"github.com/gofiber/fiber/v2"
	// "fmt"
)

func main() {
	// Open Connection DB Postgresql
	models.PGConnection()

	// Handle signal os for  closing the database connection if the app has been stopped
	ctx := make(chan os.Signal, 1)
	signal.Notify(ctx, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ctx
		models.PGCloseConnection()
		fmt.Println("Database connection closed")
		os.Exit(0)
	}()

	// Init Fiber
	app := fiber.New()

	// Use All Router From Routes
	routes.RouterApp("/api/v1", app)

	log.Fatal(app.Listen(":5600"))
}
