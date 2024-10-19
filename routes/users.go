package routes

import (
	"github.com/enrichoalkalas01/learn-go-fiber.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterUsers(basePath string, api fiber.Router) {
	usersGroup := api.Group("/users")

	usersGroup.Get("/", controllers.UsersReadList)
	usersGroup.Get("/:id", controllers.UsersReadDetail)
	usersGroup.Post("/", controllers.UsersCreate)
	usersGroup.Put("/:id", controllers.UsersUpdate)
	usersGroup.Delete("/:id", controllers.UsersDelete)
}
