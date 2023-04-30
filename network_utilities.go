package main

import (
	"fmt"
	"net"
)

func GetLocalIPAddresses() []string {
	var result []string
	// Get a list of network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	// Iterate over the network interfaces
	for _, iface := range ifaces {
		// Only consider interfaces that are up and not loopback
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			// Get the addresses for this interface
			addrs, err := iface.Addrs()
			if err != nil {
				panic(err)
			}

			// Iterate over the addresses for this interface
			for _, addr := range addrs {
				// Only consider IPv4 addresses
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
					fmt.Println("IPv4 Address:", ipnet.IP.String())
					result = append(result, ipnet.IP.String())
				}
			}
		}
	}

	return result
}
