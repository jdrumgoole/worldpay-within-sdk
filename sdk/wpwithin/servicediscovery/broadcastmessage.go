package servicediscovery

type BroadcastMessage struct {

	Description string `json:"description"`
	Host string `json:"host"`
	SvcUid string `json:"svcUid"`
	UrlPrefix string `json:"urlPrefix"`
}