package servicediscovery
import (
	"net"
	log "github.com/Sirupsen/logrus"
	"time"
	"encoding/json"
	"innovation.worldpay.com/worldpay-within-sdk/sdkcore/wpwithin/domain"
)

type broadcasterImpl struct {

	run bool
	stepSleep int
	host net.IP
	port int
	udpProtocol string
}

func (bcast *broadcasterImpl) StartBroadcast(msg domain.ServiceMessage, timeoutMillis int) (chan bool, error) {

	log.Debug("Start svc broadcast")

	chDone := make(chan bool)

	var udpConn *net.UDPConn

	// Enable broadcaster to run
	bcast.run = true

	// Determine when the broadcast operation will expire
	timeoutTime := time.Now().Add(time.Duration(timeoutMillis) * time.Millisecond)
	timedOut := false

	go func(){

		for bcast.run && !timedOut {

			log.Debug("Broadcasting..")

			BROADCAST_IPv4 := bcast.host
			conn, err := net.DialUDP(bcast.udpProtocol, nil, &net.UDPAddr{
				IP:   BROADCAST_IPv4,
				Port: int(bcast.port),
			})

			if err != nil {

				log.Error(err)

			} else {

				log.Debug("Did open UDP broadcast socket")

				jsonBytes, err := json.Marshal(msg)

				if err != nil {

					log.Error(err)
					continue
				}

				_, err = conn.Write(jsonBytes)

				if err != nil {

					log.Error(err)
					continue
				}

				log.Debug("Did successfully write broadcast message")
			}

			// Sleep before broadcasting another message - don't want to flood network
			time.Sleep(time.Duration(bcast.stepSleep) * time.Millisecond)

			// Determine if operation has timed out yet
			timedOut = timeoutTime.Unix() <= time.Now().Unix()
		}

		// Clean up connection if still active
		if udpConn != nil {

			err := udpConn.Close()

			if err != nil {

				log.Error(err)
			} else {

				log.Debug("Did close UDP broadcast socket")
			}
		}

		log.WithFields(log.Fields{ "Timed out": timedOut, "Run": bcast.run }).Debug("Finished broadcasting")

		chDone <- true
	}()

	return chDone, nil
}

func (bcast *broadcasterImpl) StopBroadcast() error {

	log.Debug("Stop svc broadcast")

	bcast.run = false

	return nil
}