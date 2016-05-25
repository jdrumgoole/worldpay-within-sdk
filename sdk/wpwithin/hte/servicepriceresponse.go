package hte
import "innovation.worldpay.com/worldpay-within-sdk/sdk/wpwithin/domain"

type ServicePriceResponse struct {

	ServerID string `json:"serverID"`
	Prices []domain.Price `json:"prices"`
}
