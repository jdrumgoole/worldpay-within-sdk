package types

type TotalPriceRequest struct {
	ClientID              string `json:"clientID"`
	SelectedNumberOfUnits int    `json:"selectedNumberOfUnits"`
	SelectedPriceID       int    `json:"selectedPriceID"`
	ClientUUID            string `json:"client-uuid"`
}
