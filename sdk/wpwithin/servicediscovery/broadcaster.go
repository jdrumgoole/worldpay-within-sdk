package servicediscovery

type Broadcaster interface {

	StartBroadcast(timeoutMillis int32) error
	StopBroadcast() error
}