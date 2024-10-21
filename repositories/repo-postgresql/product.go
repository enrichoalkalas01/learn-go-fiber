package repopostgresql

import (
	"fmt"

	"github.com/enrichoalkalas01/learn-go-fiber.git/models"
	schemasql "github.com/enrichoalkalas01/learn-go-fiber.git/models/schema-sql"
	"github.com/enrichoalkalas01/learn-go-fiber.git/utils"
)

type GetListProductParams struct {
	Search *string
	Page   *int
	Size   *int
	Order  *string
	SortBy *string
}

func GetListProduct(params GetListProductParams) ([]schemasql.Product, int64, error) {
	var products []schemasql.Product
	var totalData int64

	// Filter Query List
	paramsValid := utils.ValidationParamsListMethod(utils.ParamsListMethod{
		Search: params.Search,
		Page:   params.Page,
		Size:   params.Size,
		Order:  params.Order,
		SortBy: params.SortBy,
	})

	queryData := models.PGDatabase().Model(&schemasql.Product{})

	if paramsValid.Search != nil && *paramsValid.Search != "" {
		queryData = queryData.Where("product_name ILIKE ?", fmt.Sprintf("%%%s%%", *paramsValid.Search))
	}

	queryData = queryData.Order(fmt.Sprintf("%s %s", *paramsValid.SortBy, *paramsValid.Order)).
		Limit(*paramsValid.Size).Offset((*paramsValid.Page - 1) * *paramsValid.Size)

	err := queryData.Find(&products).Error
	if err != nil {
		fmt.Println("Failed to get data list : ", err)
		return nil, 0, err
	}

	if err := queryData.Model(&products).Count(&totalData).Error; err != nil {
		fmt.Println("Failed to get total data")
	}

	fmt.Println("Success to get data products")

	return products, totalData, err
}

func GetDetailProduct(id int) (interface{}, error) {
	var product schemasql.Product
	var productId int = id

	err := models.PGDatabase().First(&product, productId).Error
	if err != nil {
		fmt.Println("product not found", err)
	}

	fmt.Println("Success Get Product Detail", productId)

	return product, err
}

func CreateProduct(params schemasql.Product) error {
	var CategoryID *uint

	if params.CategoryID != nil {
		CategoryID = params.CategoryID
	} else {
		CategoryID = nil
	}

	// Menggunakan schema.Product
	product := &schemasql.Product{
		ProductName: params.ProductName,
		Description: params.Description,
		Price:       params.Price,
		Stock:       params.Stock,
		CategoryID:  CategoryID, // Bisa bernilai nil karena optional
	}

	err := models.PGDatabase().Create(&product).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Product Saved : ", product)

	return err
}

func UpdateProduct(params schemasql.Product, id int) error {
	// Buat variable penyimpan data dari hasil query gorm, dan harus berdasarkan schema table yang sudah dibuat.
	var product schemasql.Product
	var productId uint = uint(id) // rubah id biasa ke uniq int / uint

	var CategoryID *uint

	if params.CategoryID != nil {
		CategoryID = params.CategoryID
	} else {
		CategoryID = nil
	}

	// buat error handlers, dan ketika berhasil, data yang sudah di query akan dimasukan ke var product yang diatas. err hanya untuk menyimpan error dan menjadi error handler
	err := models.PGDatabase().First(&product, productId).Error
	if err != nil {
		fmt.Println("Product not found or error occurred:", err)
		return err
	}

	updateData := schemasql.Product{
		ProductName: params.ProductName,
		Description: params.Description,
		Price:       params.Price,
		Stock:       params.Stock,
		CategoryID:  CategoryID, // Bisa bernilai nil karena optional
	}

	if err := models.PGDatabase().Model(product).Updates(updateData).Error; err != nil {
		fmt.Println("Failed to update product :", err)
		return err
	}

	fmt.Println("Update Existing Product:", product)

	return err
}

func SoftDeleteProduct(id int) error {
	var product schemasql.Product
	var productId uint = uint(id)

	err := models.PGDatabase().First(&product, productId).Error
	if err != nil {
		fmt.Println("Product not found or error occurred:", err)
		return err
	}

	errDel := models.PGDatabase().Delete(&product, productId)
	if errDel != nil {
		fmt.Println("Failed to delete product :", err)
		return err
	}

	// fmt.Println(product)
	// fmt.Println(productId)

	return err
}

func HardDeleteProduct() {
	var product schemasql.Product
	var productId int = 2

	// Cari produk yang sudah di soft delete berdasarkan ID
	if err := models.PGDatabase().Unscoped().Where("id = ? AND deleted_at IS NOT NULL", productId).First(&product).Error; err != nil {
		fmt.Println("soft deleted product not found: %v", err)
	}

	// Hapus produk secara permanen
	if err := models.PGDatabase().Unscoped().Delete(&product).Error; err != nil {
		fmt.Println("failed to permanently delete soft deleted product: %v", err)
	}

	fmt.Println("Soft deleted product permanently deleted:", productId)
}
