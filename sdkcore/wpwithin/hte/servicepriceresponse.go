package hte
import "innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"

type ServicePriceResponse struct {

	ServerID string `json:"serverID"`
	Prices []domain.Price `json:"prices"`
}
