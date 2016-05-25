package domain
import "fmt"

type Service struct {

	Uid string `json:"uid"`
	Name string `json:"name"`
	Description string `json:"description`
	prices map[string]Price `json:"prices"`
}

func NewService() (*Service, error) {

	result := &Service{}

	result.prices = make(map[string]Price, 0)

	return result, nil
}

func (service *Service) AddPrice(price Price) error {

	fmt.Printf("Add price. Price UID = %s\n", price.Uid)

	service.prices[price.Uid] = price

	return nil
}

func (service *Service) RemovePrice(price Price) error {

	fmt.Println("Remove price..")

	delete(service.prices, price.Uid)

	return nil
}

func (service *Service) Prices() map[string]Price {

	return service.prices
}