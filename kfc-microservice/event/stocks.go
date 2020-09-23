package event

import (
	"encoding/json"
	"fmt"
	"kfc-microservice/services"
)

//StockRoot ...
type StockRoot struct {
	Stock Stock `json:"product"`
}

//Stock ...
type Stock struct {
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	ExternalID       string          `json:"external_id"`
	Price            float32         `json:"price"`
	ShippingCategory string          `json:"shipping_category"`
	VendorID         int             `json:"vendor_id"`
	PriceLocations   []PriceLocation `json:"price_locations"`
}

//PriceLocation ...
type PriceLocation struct {
	LocationCode string  `json:"location_code"`
	Price        float32 `json:"price"`
}

//
func CreateStock(payload string) (newStock *StockRoot, err error) {
	newProductRoot := &services.ProductsRoot{}
	err = json.Unmarshal([]byte(payload[1:len(payload)-1]), &newProductRoot)

	newProduct := newProductRoot.Products
	newList := newProductRoot.List

	newStock = &StockRoot{
		Stock: Stock{
			Name:             newProduct.Name,
			Description:      newProduct.Description,
			ExternalID:       fmt.Sprint(newProduct.ProductID),
			Price:            newProduct.PriceInformation[0].Price,
			ShippingCategory: "Default",
			VendorID:         newList.VendorID,
			PriceLocations: []PriceLocation{
				{
					LocationCode: "2",
					Price:        90,
				},
			},
		},
	}

	return newStock, err
}
