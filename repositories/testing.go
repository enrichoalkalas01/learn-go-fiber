package repositories

import (
	"fmt"
	"log"

	"github.com/enrichoalkalas01/learn-go-fiber.git/models"
	"github.com/enrichoalkalas01/learn-go-fiber.git/utils"
)

func Testing() {
	models.PGConnection()
	GetListProduct()

	// fmt.Println("here1")
}

func GetListProduct() {
	params := utils.ValidationParamsListMethod(utils.ParamsListMethod{
		Search: nil,
		Page:   nil,
		Size:   nil,
		Order:  nil,
		SortBy: nil,
	})

	fmt.Println(*params.Search)
	fmt.Println("Success to get data products")
}

func GetDetailProduct() {
	var product models.Product
	var productId int = 4

	err := models.PGDatabase().First(&product, productId).Error
	if err != nil {
		log.Fatal("product not found", err)
	}

	fmt.Println("Success Get Product Detail", productId)
}

func CreateProduct() {
	// Menggunakan schema.Product
	product := &models.Product{
		ProductName: "product 1",
		Description: "product 1 description",
		Price:       10000,
		Stock:       10,
		CategoryID:  nil, // Bisa bernilai nil karena optional
	}

	err := models.PGDatabase().Create(&product).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product Saved : ", product)
}

func UpdateProduct() {
	// Buat variable penyimpan data dari hasil query gorm, dan harus berdasarkan schema table yang sudah dibuat.
	var product models.Product
	var productId int = 1

	// buat error handlers, dan ketika berhasil, data yang sudah di query akan dimasukan ke var product yang diatas. err hanya untuk menyimpan error dan menjadi error handler
	err := models.PGDatabase().First(&product, productId).Error
	if err != nil {
		log.Fatal("Product not found or error occurred:", err)
	}

	updateData := models.Product{
		ProductName: "product 2",
		Description: "product 21 description",
		Price:       12000,
		Stock:       12,
		CategoryID:  nil, // Bisa bernilai nil karena optional
	}

	if err := models.PGDatabase().Model(product).Updates(updateData).Error; err != nil {
		log.Fatal("Failed to update product :", err)
	}

	fmt.Println("Update Existing Product:", product)
}

func SoftDeleteProduct() {
	var product models.Product
	var productId int = 2

	err := models.PGDatabase().First(&product, productId).Error
	if err != nil {
		log.Fatal("Product not found or error occurred:", err)
	}

	errDel := models.PGDatabase().Delete(&product, productId)
	if errDel != nil {
		log.Fatal("Failed to delete product :", err)
	}

	fmt.Println(product)
	fmt.Println(productId)
}

func HardDeleteProduct() {
	var product models.Product
	var productId int = 2

	// Cari produk yang sudah di soft delete berdasarkan ID
	if err := models.PGDatabase().Unscoped().Where("id = ? AND deleted_at IS NOT NULL", productId).First(&product).Error; err != nil {
		log.Fatal("soft deleted product not found: %v", err)
	}

	// Hapus produk secara permanen
	if err := models.PGDatabase().Unscoped().Delete(&product).Error; err != nil {
		log.Fatal("failed to permanently delete soft deleted product: %v", err)
	}

	fmt.Println("Soft deleted product permanently deleted:", productId)
}
