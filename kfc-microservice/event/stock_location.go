package event

import (
	"encoding/json"
	"fmt"
	"kfc-microservice/services"
)

//
type StockLocationRoot struct {
	StockLocation StockLocation `json:"stock_location"`
}

//
type StockLocation struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	VendorID string `json:"vendorId"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
}

//
func CreateStockLocation(payload string) (newStockLocation *StockLocationRoot, err error) {
	newStoresRoot := &services.StoresRoot{}
	err = json.Unmarshal([]byte(payload[1:len(payload)-1]), &newStoresRoot)

	newStore := newStoresRoot.Stores

	newStockLocation = &StockLocationRoot{
		StockLocation: StockLocation{
			Code:     newStore.StoreCode,
			Name:     newStore.Name,
			Phone:    newStore.ContactInformation[0].Phone,
			Address1: newStore.ContactInformation[0].Address,
			Address2: "blank",
			City:     newStore.LocationInformation[0].City,
			State:    "blank",
			VendorID: fmt.Sprint(newStore.VendorID),
			Lat:      newStore.LocationInformation[0].Latitud,
			Lng:      newStore.LocationInformation[0].Longitude,
		},
	}

	return newStockLocation, err
}
