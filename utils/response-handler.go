package utils

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

type SuccessHandlerProps struct {
	Status     bool             `json:"status"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	Data       interface{}      `json:"data,omitempty"`
	Pagination *PaginationProps `json:"pagination,omitempty"`
}

type FailedHandlerProps struct {
	Status     bool        `json:"status"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors"`
}

type PaginationProps struct {
	Query     *string `json:"query"`
	Page      *int    `json:"page"`
	Size      *int    `json:"size"`
	Total     *int    `json:"total"`
	TotalData *int64  `json:"totalData"`
	SortBy    *string `json:"sortBy"`
	OrderBy   *string `json:"orderBy"`
}

func SuccessHandler(c *fiber.Ctx, params ...SuccessHandlerProps) error {
	response := SuccessHandlerProps{
		Status:     true,
		StatusCode: fiber.StatusOK,
		Message:    "Successfull to run process",
		Data:       nil,
		Pagination: nil,
	}

	if len(params) > 0 {
		p := params[0]

		if !isZero(p.Message) {
			response.Message = p.Message
		}

		if !isZero(p.StatusCode) {
			response.StatusCode = p.StatusCode
		}

		if !isZero(p.Status) {
			response.Status = p.Status
		}

		if !isZero(p.Data) {
			response.Data = p.Data
		}

		if !isZero(p.Pagination) {
			pagination := PaginationProps{
				Query:     nil,
				Page:      nil,
				Size:      nil,
				Total:     nil,
				TotalData: nil,
				OrderBy:   nil,
				SortBy:    nil,
			}

			if p.Pagination.Query != nil {
				pagination.Query = p.Pagination.Query
			}

			if p.Pagination.Page != nil {
				pagination.Page = p.Pagination.Page
			}

			if p.Pagination.Size != nil {
				pagination.Size = p.Pagination.Size
			}

			if p.Pagination.Total != nil {
				pagination.Total = p.Pagination.Total
			}

			if p.Pagination.TotalData != nil {
				pagination.TotalData = p.Pagination.TotalData
			}

			if p.Pagination.OrderBy != nil {
				pagination.OrderBy = p.Pagination.OrderBy
			}

			if p.Pagination.SortBy != nil {
				pagination.SortBy = p.Pagination.SortBy
			}

			response.Pagination = &pagination
		}
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func ErrorHandler(c *fiber.Ctx, params ...FailedHandlerProps) error {
	response := FailedHandlerProps{
		Status:     false,
		StatusCode: fiber.StatusBadGateway,
		Message:    "Failed to run process",
		Errors:     nil,
	}

	if len(params) > 0 {
		p := params[0]

		if !isZero(p.Message) {
			response.Message = p.Message
		}

		if !isZero(p.StatusCode) {
			response.StatusCode = p.StatusCode
		}

		if !isZero(p.Status) {
			response.Status = p.Status
		}

		if !isZero(p.Errors) {
			response.Errors = p.Errors
		}
	}

	return c.Status(response.StatusCode).JSON(response)
}

func isZero(value interface{}) bool {
	return value == nil || reflect.ValueOf(value).IsZero()
}
