package types

type BeginServiceDeliveryRequest struct {
	ClientID             string               `json:"clientID"`
	ServiceDeliveryToken ServiceDeliveryToken `json:"serviceDeliveryToken"`
	UnitsToSupply        int                  `json:"unitsToSupply"`
}
