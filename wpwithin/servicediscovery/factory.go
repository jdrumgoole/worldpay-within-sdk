package servicediscovery

func NewScanner() (Scanner, error) {

	result := &scannerImpl{}

	return result, nil
}

func NewBroadcaster(description, host, svcUid, urlPrefix string, port int32) (Broadcaster, error) {

	result := &broadcasterImpl{}

	return result, nil
}