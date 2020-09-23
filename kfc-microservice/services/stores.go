package services

//StoresRoot ...
type StoresRoot struct {
	Stores Stores `json:"stores"`
}

//
type Stores struct {
	Active              bool             `json:"active"`
	ContactInformation  []ContactInfo    `json:"contactInfo"`
	DeliveryInformation []DeliveryInfo   `json:"deliveryInfo"`
	LocationInformation []LocationInfo   `json:"locationInfo"`
	Name                string           `json:"name"`
	Schedules           []StoreSchedules `json:"schedules"`
	Services            []StoreServices  `json:"services"`
	StoreChannels       []Channels       `json:"storeChannels"`
	StoreCode           string           `json:"storeCode"`
	StoreID             int              `json:"storeId"`
	TaxesInformation    []TaxesInfo      `json:"taxesInfo"`
	VendorID            int              `json:"vendorId"`
}

//
type ContactInfo struct {
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

//
type DeliveryInfo struct {
	CookTime     string `json:"cookTime"`
	MinimumOrder string `json:"minimumOrder"`
	ShippingCost string `json:"shippingCost"`
}

//
type LocationInfo struct {
	City      string `json:"city"`
	Latitud   string `json:"latitud"`
	Longitude string `json:"longitude"`
}

//
type StoreSchedules struct {
	Day       string `json:"day"`
	EndTime   string `json:"endTime"`
	StartTime string `json:"startTime"`
}

//
type StoreServices struct {
	Active bool   `json:"active"`
	Name   string `json:"name"`
}

//
type Channels struct {
	ChannelID string `json:"channelId"`
}

//
type TaxesInfo struct {
	VatRatePercentage int `json:"vatRatePercentage"`
}
