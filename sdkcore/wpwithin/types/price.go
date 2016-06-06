package types

type Price struct {

	ServiceID int `json:"serviceID"`
	ID int `json:"priceID"`
	Description string `json:"priceDescription"`
	PricePerUnit int `json:"pricePerUnit"`
	UnitID int `json:"unitID"`
	UnitDescription string `json:"unitDescription"`
}

func NewPrice() (*Price, error) {

	result := &Price{}

	return result, nil
}