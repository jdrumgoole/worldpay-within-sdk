package servicediscovery

type BroadcastMessage struct {

	DeviceDescription string `json:"deviceDescription"`
	Hostname string `json:"hostname"`
	PortNumber int `json:"portNumber"`
	ServerID string `json:"serverID"`
	UrlPrefix string `json:"urlPrefix"`
}