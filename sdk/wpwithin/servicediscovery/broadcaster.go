package servicediscovery

type Broadcaster interface {

	StartBroadcast(msg BroadcastMessage, timeoutMillis int) error
	StopBroadcast() error
}