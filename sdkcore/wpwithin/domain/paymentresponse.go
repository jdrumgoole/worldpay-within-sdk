package domain

type PaymentResponse struct {

	ServerID string `json:"serverID"`
	ClientID string `json:"clientID"`
	TotalPaid int `json:"totalPaid"`
	ServiceDeliveryToken string `json:"serviceDeliveryToken"`
	ClientUUID string `json:"client-uuid"`
}
