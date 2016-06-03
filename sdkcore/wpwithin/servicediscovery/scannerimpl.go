package servicediscovery
import (
	"fmt"
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
	"encoding/json"
)

type scannerImpl struct {

	run bool /* Used to stop scanning before timeout */
	stepSleep int /* Time to sleep before UDP reads */
	port int
	udpProtocol string
}

func (scanner *scannerImpl) ScanForServices(timeout int) ScanResult {

	/*
		This function works by setting up a UDP broadcast listener, returning a result object which includes a channel
		informing when scanning is finished.
		Inside the result is an error object and also a list of scanned services
		Error is != nil if there was a problem
	 */

	log.Debugf("Begin ScanForServices(timeout = %d)", timeout)

	result := ScanResult{}
	result.Complete = make(chan bool)
	result.Services = make(map[string]BroadcastMessage, 0)
	// Enable the scanner to run
	scanner.run = true
	// Calculate when the operation will expire based on the timeout duration
	timeoutTime := time.Now().Add(time.Duration(timeout) * time.Millisecond)
	timedOut := false

	// UDP Broadcast discovery
	srvAddr := &net.UDPAddr{
		IP: net.IPv4allrouter,
		Port:scanner.port,
	}

	go func() {

		// Reading incoming messages - 2kb buffer
		buf := make([]byte, 2048)

		for scanner.run && !timedOut {

			srvConn, err := net.ListenUDP(scanner.udpProtocol, srvAddr)
			if err != nil {

				result.Error = err

				scanner.run = false
			}

			// Defer closing connection in go routine instead of main routine as it will be closed before the go routine starts.
			defer srvConn.Close()

			// Wait for incoming UDP message
			srvConn.SetReadDeadline(time.Now().Add(time.Duration(scanner.stepSleep) * time.Millisecond))

			nRecv, addrRecv,err := srvConn.ReadFromUDP(buf)
			
			if err != nil {

				log.Error(err)
			}

			if nRecv > 0 { /* Did we actually receive any data? */

				log.Debugf("Did receive UDP message from %s: %s", addrRecv.String(), string(buf[0:nRecv]))

				var msg BroadcastMessage

				// Try to deserialize the message into a broadcast message
				// NB: Anybody can send a message here so not all messages are expected to be valid
				err = json.Unmarshal(buf[0:nRecv], &msg);

				if err != nil {

					// This is not neccessarily an error - could be a message from another source (ignore)
					log.WithFields(log.Fields{"Error: ": fmt.Sprintf("Err: %q", err.Error())}).Error("Did not decode UDP message")
				} else {

					log.Infof("Did decode broadcast message: %#v", msg)

					result.Services[msg.ServerID] = msg
				}
			}
			// Have we timed out? i.e. Is the current time greater or equal time out time
			timedOut = timeoutTime.Unix() <= time.Now().Unix()
		}

		log.WithFields(log.Fields{ "Timed out": timedOut, "Run": scanner.run, "Found": len(result.Services)}).Debug("Finish ScanForServices()")

		result.Complete <- true
	}()

	return result
}

func (scanner *scannerImpl) StopScanner() {

	scanner.run = false
}