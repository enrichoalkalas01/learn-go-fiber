package controllers

import (
	"fmt"

	schemasql "github.com/enrichoalkalas01/learn-go-fiber.git/models/schema-sql"
	repo "github.com/enrichoalkalas01/learn-go-fiber.git/repositories/repo-postgresql"
	"github.com/enrichoalkalas01/learn-go-fiber.git/utils"
	"github.com/gofiber/fiber/v2"
)

func ProductsReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams((c))
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
		})
	}

	data, total, err := repo.GetListProduct(repo.GetListProductParams{
		Search: &searchQuery,
		Page:   &page,
		Size:   &size,
		Order:  &order,
		SortBy: &sortBy,
	})

	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	totalArray := len(data)

	return utils.SuccessHandler(c, utils.SuccessHandlerProps{
		Data: data,
		Pagination: &utils.PaginationProps{
			Query:     &searchQuery,
			Page:      &page,
			Size:      &size,
			Total:     &totalArray,
			TotalData: &total,
		},
	})
}

func ProductsReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    fiber.ErrBadRequest.Error(),
		})
	}

	data, err := repo.GetDetailProduct(id)
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return utils.SuccessHandler(c, utils.SuccessHandlerProps{
		Data: data,
	})
}

type ProductRequest struct {
	ProductName string  `json:"product_name" validate:"required,min=1,max=200"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Stock       int     `json:"stock" validate:"required,min=1"`
	CategoryID  *uint   `json:"category_id"` // optional for relation catergory
}

func ProductsCreate(c *fiber.Ctx) error {
	var body ProductRequest

	// Validation Body Request
	errorsMessage, err := utils.ValidateStruct(c, &body)
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Errors:     errorsMessage,
		})
	}

	// Setup Data Passing
	productData := &schemasql.Product{
		ProductName: body.ProductName,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
	}

	// Create Data
	if err := repo.CreateProduct(*productData); err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return utils.SuccessHandler(c)
}

func ProductsUpdate(c *fiber.Ctx) error {
	var body ProductRequest

	// Parsing Validation Params
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// Validation Body Request
	errorsMessage, err := utils.ValidateStruct(c, &body)
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Errors:     errorsMessage,
		})
	}

	// Setup Data Passing
	productData := &schemasql.Product{
		ProductName: body.ProductName,
		Description: body.Description,
		Price:       body.Price,
		Stock:       body.Stock,
	}

	// Update Data
	if err := repo.UpdateProduct(*productData, id); err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	return utils.SuccessHandler(c)
}

func ProductsDelete(c *fiber.Ctx) error {
	// Parsing Validation Params
	id, err := utils.ValidationIdParams(c)
	if err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	// Delete Data
	if err := repo.SoftDeleteProduct(id); err != nil {
		return utils.ErrorHandler(c, utils.FailedHandlerProps{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
		})
	}

	fmt.Println(id)

	return utils.SuccessHandler(c)
}
