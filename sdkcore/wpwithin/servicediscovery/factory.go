package servicediscovery
import "net"

func NewScanner(port, stepSleep int) (Scanner, error) {

	result := &scannerImpl{
		run: false,
		stepSleep: stepSleep,
	}

	if comm, err := NewUDPComm(port, "udp4"); err != nil {

		return nil, err
	} else {

		result.comm = comm
	}

	return result, nil
}

func NewBroadcaster(host string, port int, stepSleep int) (Broadcaster, error) {

	result := &broadcasterImpl{

		stepSleep: stepSleep,
		run: false,
		host: net.IPv4bcast.To4(),
		port: port,
		udpProtocol: "udp4",
	}

	return result, nil
}