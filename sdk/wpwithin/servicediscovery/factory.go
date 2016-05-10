package servicediscovery
import "net"

func NewScanner() (Scanner, error) {

	result := &scannerImpl{}

	return result, nil
}

func NewBroadcaster(host string, port int, stepSleep int) (Broadcaster, error) {

	result := &broadcasterImpl{

		stepSleep: stepSleep,
		run: false,
		host: net.IPv4(255, 255, 255, 255),
		port: port,
		udpProtocol: "udp4",
	}

	return result, nil
}