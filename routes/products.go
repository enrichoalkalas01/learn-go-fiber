package routes

import (
	"github.com/enrichoalkalas01/learn-go-fiber.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouterProducts(basePath string, api fiber.Router) {
	productsGroup := api.Group(basePath)

	productsGroup.Get("/", controllers.ProductsReadList)
	productsGroup.Get("/:id", controllers.ProductsReadDetail)
	productsGroup.Post("/", controllers.ProductsCreate)
	productsGroup.Put("/:id", controllers.ProductsUpdate)
	productsGroup.Delete("/:id", controllers.ProductsDelete)
}
