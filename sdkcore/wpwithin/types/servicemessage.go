package types

type ServiceMessage struct {
	DeviceDescription string `json:"deviceDescription"`
	Hostname          string `json:"hostname"`
	PortNumber        int    `json:"portNumber"`
	ServerID          string `json:"serverID"`
	UrlPrefix         string `json:"urlPrefix"`
	Scheme            string `json:"scheme"`
}
