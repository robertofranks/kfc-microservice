package services

//
type ProductsRoot struct {
	List     List     `json:"list"`
	Products Products `json:"products"`
}

//
type Products struct {
	Active           bool        `json:"active"`
	Description      string      `json:"description"`
	Name             string      `json:"name"`
	PriceInformation []PriceInfo `json:"priceInfo"`
	ProductID        int         `json:"productId"`
	ProductModifiers []Modifiers `json:"productModifiers"`
	Tags             []string    `json:"tags"`
	TaxInformation   []TaxInfo   `json:"taxInfo"`
	Type             string      `json:"type"`
}

//
type PriceInfo struct {
	Price float32 `json:"price"`
}

//
type Modifiers struct {
	ModifierID string `json:"modifierId"`
	Position   int    `json:"position"`
}

//
type TaxInfo struct {
	VatRatePercentage float32 `json:"vatRatePercentage"`
}

//
type List struct {
	ChannelID string `json:"channelId"`
	ListID    string `json:"listId"`
	ListName  string `json:"listName"`
	StoreID   int    `json:"storeId"`
	VendorID  int    `json:"vendorId"`
}
