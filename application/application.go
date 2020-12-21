package application

import (
	"fmt"
	"net"
	"strings"
)

// GetMacAddress - return macaddress of machine
func GetMacAddress() string {
	var currentNetworkHardwareName string
	var currentIP string

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				currentIP = ipnet.IP.String()
			}
		}
	}

	interfaces, _ := net.Interfaces()
	for _, interf := range interfaces {

		if addrs, err := interf.Addrs(); err == nil {
			for _, addr := range addrs {
				if strings.Contains(addr.String(), currentIP) {
					currentNetworkHardwareName = interf.Name
				}
			}
		}
	}

	netInterface, err := net.InterfaceByName(currentNetworkHardwareName)
	return netInterface.HardwareAddr.String()
}
