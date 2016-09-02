package types

type BeginServiceDeliveryResponse struct {
	ServerID             string               `json:"serverID"`
	ClientID             string               `json:"clientID"`
	ServiceDeliveryToken ServiceDeliveryToken `json:"serviceDeliveryToken"`
	UnitsToSupply        int                  `json:"unitsToSupply"`
}
