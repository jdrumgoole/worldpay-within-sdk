package servicediscovery

type ScanResult struct {

	Complete chan bool
	Services map[string]BroadcastMessage
	Error error
}
