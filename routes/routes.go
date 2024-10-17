package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouterApp(basePath string, app *fiber.App) {
	api := app.Group(basePath)

	RouterUsers("/users", api)
	// RouterProducts("/users", api)
	RouterTemplate("/template", api)
}
