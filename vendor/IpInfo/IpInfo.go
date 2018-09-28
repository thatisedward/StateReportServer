package IpInfo

import (
	"fmt"
	"net"
	"os"
)

func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ip := ""

	for _, address := range addrs {

		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			//eliminate loop back address start with 127

			if ipNet.IP.To4() != nil {
				if ipNet.IP.String()[:3] != "121" {
					ip = ipNet.IP.String()
				}
			}
		}
	}

	return ip
}
