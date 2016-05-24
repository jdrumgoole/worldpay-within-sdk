package servicediscovery

type Scanner interface {

	ScanForServices(timeout int) ScanResult

	StopScanner()
}