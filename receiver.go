package main

import (
	"fmt"
	"net"
	"os"
)

var message = ""

// i suggest port to have the value of 3035
func startTCPServer(port string) string {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	// Wait for a connection
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
		os.Exit(1)
	}

	// Read message from client
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		os.Exit(1)
	}

	// Save message to a string variable
	message = string(buf[:n])

	// Close the connection
	conn.Close()
	return message
}
