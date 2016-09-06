package types

import "fmt"

type Service struct {
	Id          int           `json:"serviceID"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	prices      map[int]Price `json:"prices"`
}

func NewService() (*Service, error) {

	result := &Service{}

	result.prices = make(map[int]Price, 0)

	return result, nil
}

func (service *Service) AddPrice(price Price) error {

	if _, exists := service.prices[price.ID]; exists {

		return fmt.Errorf("A price with that ID (%d) already exists for that service.", price.ID)
	}

	service.prices[price.ID] = price

	return nil
}

func (service *Service) RemovePrice(price Price) error {

	delete(service.prices, price.ID)

	return nil
}

func (service *Service) Prices() map[int]Price {

	return service.prices
}
