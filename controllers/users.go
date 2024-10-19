package controllers

import (
	"fmt"

	"github.com/enrichoalkalas01/learn-go-fiber.git/utils"
	"github.com/gofiber/fiber/v2"
)

// Request Body Field & Validator ( tag )
type UsersRequest struct {
	Username    *string `json:"username" validate:"required,min=1,max=200"`
	Password    *string `json:"password" validate:"required,min=1,max=200"`
	Firstname   *string `json:"firstname" validate:"required,min=1,max=200"`
	Lastname    *string `json:"lastname" validate:"required,min=1,max=200"`
	Phonenumber *string `json:"phonenumber" validate:"required,min=1,max=14"`
	Email       *string `json:"email" validate:"required,email"`
}

func UsersReadList(c *fiber.Ctx) error {
	searchQuery, page, size, order, sortBy, err := utils.ValidationQueryParams((c)) // Parsing Automate Query Params
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfull to get Users",
		"status":  200,
		"pagination": fiber.Map{
			"search":  searchQuery,
			"page":    page,
			"size":    size,
			"order":   order,
			"sort_by": sortBy,
		},
	})
}

func UsersReadDetail(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	return c.JSON(fiber.Map{
		"message": "Successfull to get Users detail",
		"status":  200,
	})
}

func UsersCreate(c *fiber.Ctx) error {
	var body UsersRequest

	// Sebelum kita menggunakan function, kita harus bisa membuat setup params dahulu sebelum masuk kedalam function nya
	// jika tidak ada setup params, go akan menganggap bahwa code akan error karena params wajib di isi / di setup
	test := utils.CreateTokenJWT(&utils.CreateTokenParams{
		ExpiredNumber: nil, // Memicu nilai default
		ExpiredType:   nil, // Memicu nilai default
		SourceFrom:    nil, // Memicu nilai default
		TokenType:     nil, // Memicu nilai default
		UserData: &map[string]interface{}{
			"id":       1, // Mengambil data dari request body, contoh
			"username": "",
			"email":    "",
			"type":     "user", // Atau sesuaikan dengan tipe user
		},
	})

	errorsMap, err := utils.ValidateStruct(c, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"errors":  errorsMap,
			"status":  fiber.StatusBadRequest,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfull to create Users",
		"status":  200,
		"data":    body,
		"test":    test,
	})
}

func UsersUpdate(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	var body UsersRequest

	// Memvalidasi request body
	errorsMap, err := utils.ValidateStruct(c, &body)
	if err != nil {
		// Jika validasi gagal, kembalikan respons yang lebih rinci
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "validation failed",
			"errors":  errorsMap, // Mengembalikan detail kesalahan validasi
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(body)

	return c.JSON(fiber.Map{
		"message": "Successfull to update Users",
		"status":  200,
		"data":    body,
	})
}

func UsersDelete(c *fiber.Ctx) error {
	id, err := utils.ValidationIdParams((c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	fmt.Println(id)

	return c.JSON(fiber.Map{
		"message": "Successfull to delete Users",
		"status":  200,
	})
}
