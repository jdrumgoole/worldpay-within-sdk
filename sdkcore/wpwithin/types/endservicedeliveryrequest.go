package types

type EndServiceDeliveryRequest struct {
	ClientID             string               `json:"clientID"`
	ServiceDeliveryToken ServiceDeliveryToken `json:"serviceDeliveryToken"`
	UnitsReceived        int                  `json:"unitsReceived"`
}
