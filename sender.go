package main

import (
	"fmt"
	"net"
)

// the first client send the data to the second client using this
// the port value must be the same as for the receiver and i suggest both to be 3035
func SendDataToServer(data string, port string) {
	// Establish a connection to the server
	conn, err := net.Dial("tcp", "192.168.1.2:"+port) // TODO: making the IP address a dynamic value
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// Send a message to the server
	message := data
	// fmt.Println("Sending message to server:", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	// fmt.Println("Message sent successfully!")
}
