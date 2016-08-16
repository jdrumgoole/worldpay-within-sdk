package types

type PaymentResponse struct {

	ServerID string `json:"serverID"`
	ClientID string `json:"clientID"`
	TotalPaid int `json:"totalPaid"`
	ServiceDeliveryToken *ServiceDeliveryToken `json:"serviceDeliveryToken"`
	ClientUUID string `json:"client-uuid"`
}
