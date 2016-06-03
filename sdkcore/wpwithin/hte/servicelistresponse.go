package hte

type ServiceListResponse struct {

	ServerID string `json:"serverID"`
	Services []ServiceDetails `json:"services"`
}
