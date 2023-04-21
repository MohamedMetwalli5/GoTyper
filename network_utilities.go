package main

import (
	"fmt"
	"net"
)

func GetLocalIPAddresses() []string {
	var result []string
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error:", err)
		return result
	}

	for _, intf := range interfaces {
		addrs, err := intf.Addrs()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				fmt.Println("IP address:", ipnet.IP.String())
				result = append(result, ipnet.IP.String())
			}
		}
	}

	return result
}
