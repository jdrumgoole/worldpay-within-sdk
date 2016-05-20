package utils
import (
	"errors"
	"net"
)

// Return the IPv4 external address of this device.
// Note external does not necessarily mean WAN IP. On most networks it will be the LAN IP of device as opposed
// to internal localhost address (127.0.0.1)
func ExternalIPv4() (string, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}

			return ip.String(), nil
		}
	}
	return "", errors.New("Device does not appear to be network connected.")
}