package servicediscovery

type Scanner interface {

	ScanForServices(timeout int) ScanResult
	SetServerGuidFilter(filter string) error
	StopScanner()
}