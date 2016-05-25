package domain

type Price struct {

	Uid string `json:"uid"`
	Description string `json:"description"`
	PricePerUnit int `json:"pricePerUnit"`
	UnitId int `json:"unitId"`
	UnitDescription string `json:"unitDescription"`
}

func NewPrice() (*Price, error) {

	result := &Price{}

	return result, nil
}
