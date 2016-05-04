package servicediscovery
import (
	"net"
	log "github.com/Sirupsen/logrus"
	"time"
)

const (

	srvAddr = "127.0.0.1"
	port = 9999
	maxDatagramSize = 8192
)

type broadcasterImpl struct {


}

func (bcast *broadcasterImpl) StartBroadcast(timeoutMillis int32) error {

	log.Debug("Start svc broadcast..")

//	ServerAddr,err := net.ResolveUDPAddr(net.IPv4bcast.String(),"8980")
//
//	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
//
//		IP:ServerAddr.IP,
//		Port:ServerAddr.Port,
//	})
//
//	defer socket.Close()
//
//	if err != nil {
//
//		return err
//	}
//
//	i := 0
//	for {
//		msg := fmt.Sprintf("Hello %d", i)
//		socket.Write([]byte(msg))
//		fmt.Printf("Writing: %s\n", msg)
//
//		i++
//
//		time.Sleep(time.Second * 1)
//
//		if i == 5 {
//
//			break
//		}
//	}

	go ping(srvAddr)

	return nil
}

func (bcast *broadcasterImpl) StopBroadcast() error {

	log.Debug("Stop svc broadcast")

	return nil
}

func ping(str string) {

//	addr, err := net.ResolveUDPAddr("udp", str)
//
//	if err != nil {
//
//		log.Fatal(err)
//	}
//
//	c, err := net.DialUDP("udp", nil, addr)
//
//	for {
//
//		c.Write([]byte ("hello world\n"))
//		time.Sleep(1 * time.Second)
//	}

	BROADCAST_IPv4 := net.IPv4(255, 255, 255, 255)
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   BROADCAST_IPv4,
		Port: port,
	})

	if err != nil {

		log.Fatal(err)
		return
	}

	for {

		socket.Write([]byte("Hello UDP\n"))

		time.Sleep(1 * time.Second)
	}
}