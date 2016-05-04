package servicediscovery

type Scanner interface {

	ScanForServices(timeout int) error
	SetServerGuidFilter(filter string) error
	StopScanner()
}