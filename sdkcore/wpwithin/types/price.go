package types

type Price struct {

	ID int `json:"priceID"`
	Description string `json:"priceDescription"`
	PricePerUnit *PricePerUnit `json:"pricePerUnit"`
	UnitID int `json:"unitID"`
	UnitDescription string `json:"unitDescription"`
}

func NewPrice() (*Price, error) {

	result := &Price{}

	return result, nil
}
