package types

type ServiceListResponse struct {

	ServerID string `json:"serverID"`
	Services []ServiceDetails `json:"services"`
}
