package domain

type ServicePriceResponse struct {

	ServerID string `json:"serverID"`
	Prices []Price `json:"prices"`
}
