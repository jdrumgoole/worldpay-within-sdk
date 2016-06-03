package servicediscovery

type Broadcaster interface {

	StartBroadcast(msg BroadcastMessage, timeoutMillis int) (chan bool, error)
	StopBroadcast() error
}