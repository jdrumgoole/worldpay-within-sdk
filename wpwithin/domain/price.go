package domain

type Price struct {

	Uid string
	PricePerUnit int
	UnitId int
	Description string
	UnitDescription string
}

func NewPrice() (*Price, error) {

	result := &Price{}

	return result, nil
}
