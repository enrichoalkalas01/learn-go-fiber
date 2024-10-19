package repositories

import (
	"github.com/enrichoalkalas01/learn-go-fiber.git/models"
)

func Testing() {
	models.PGConnection()

	// search := ""
	// page := 2
	// size := 5
	// order := ""
	// sortBy := "id"

	// rGetListProduct, err := repopgprd.GetListProduct(repopgprd.GetListProductParams{
	// 	Search: &search,
	// 	Page:   &page,
	// 	Size:   &size,
	// 	Order:  &order,
	// 	SortBy: &sortBy,
	// })

	// if err != nil {
	// 	fmt.Println("Failed to get product list:", err)
	// } else {
	// 	fmt.Printf("Filtered product list: %+v\n", rGetListProduct)
	// }
}
