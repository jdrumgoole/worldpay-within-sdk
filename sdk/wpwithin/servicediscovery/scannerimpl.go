package servicediscovery
import (
	"fmt"
	"net"
)

type scannerImpl struct {

}

func (scanner *scannerImpl) ScanForServices(timeout int) error {

	fmt.Println("Start scan for services..")

	/* Lets prepare a address at any address at port 10001*/
	ServerAddr,err := net.ResolveUDPAddr(net.IPv4bcast.String(),"10001")

	if err != nil {

		fmt.Println("Error: ",err)

		return err
	}

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)

	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)

		if err != nil {

			fmt.Println("Error: ",err)

			return err
		}
	}

	return nil
}

func (scanner *scannerImpl) SetServerGuidFilter(filter string) error {

	return nil
}

func (scanner *scannerImpl) StopScanner() {


}