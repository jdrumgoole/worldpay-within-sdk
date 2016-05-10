package servicediscovery
import (
	"net"
	log "github.com/Sirupsen/logrus"
	"time"
	"encoding/json"
)

type broadcasterImpl struct {

	run bool
	stepSleep int
	host net.IP
	port int
	udpProtocol string
}

func (bcast *broadcasterImpl) StartBroadcast(msg BroadcastMessage, timeoutMillis int) error {

	log.Debug("Start svc broadcast")

	var udpConn *net.UDPConn

	bcast.run = true

	timeoutTime := time.Now().Add(time.Duration(timeoutMillis) * time.Millisecond)

	timedOut := false

	for bcast.run && !timedOut {

		log.Debug("Broadcasting now")

		BROADCAST_IPv4 := bcast.host
		conn, err := net.DialUDP(bcast.udpProtocol, nil, &net.UDPAddr{
			IP:   BROADCAST_IPv4,
			Port: int(bcast.port),
		})

		if err != nil {

			log.Error(err)
			break

		} else {

			log.Debug("Did open UDP broadcast socket")

			jsonBytes, err := json.Marshal(msg)

			if err != nil {

				log.Error(err)
				break
			}

			_, err = conn.Write(jsonBytes)

			if err != nil {

				log.Error(err)
				break;
			}

			log.Debug("Did successfully write broadcast message")
		}

		time.Sleep(time.Duration(bcast.stepSleep) * time.Millisecond)

		timedOut = timeoutTime.Second() >= time.Now().Second()
	}

	if udpConn != nil {

		err := udpConn.Close()

		if err != nil {

			log.Error(err)
		} else {

			log.Debug("Did close UDP broadcast socket")
		}
	}

	log.WithFields(log.Fields{ "Timed out": timedOut, "Run": bcast.run }).Debug("Finished broadcasting")

	return nil
}

func (bcast *broadcasterImpl) StopBroadcast() error {

	log.Debug("Stop svc broadcast")

	bcast.run = false

	return nil
}