package types

type ServicePriceResponse struct {

	ServerID string `json:"serverID"`
	Prices []Price `json:"prices"`
}
