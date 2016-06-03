package types

type TotalPriceRequest struct {

	ClientID string `json:"clientID"`
	SelectedNumberOfUnits int `json:"selectedNumberOfUnits"`
	SelectedPriceId int `json:"selectedPriceID"`
	ClientUUID string `json:"client-uuid"`
}
