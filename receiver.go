package main

import (
	"fmt"
	"net"
)

var message = ""

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection established.")

	// Read incoming data from the client
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	message = string(buf)
	fmt.Println("Received message:", message)

	// Check if the message equals a specified string
	if len(message) > 0 {
		conn.Close() // Close the connection to the client
		return
	}
}

// i suggest port to have the value of 3035
func startTCPServer(port string) {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// Handle error
		fmt.Println("Error:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// Handle error
			fmt.Println("Error:", err)
			continue
		}
		go handleConnection(conn)
	}
}
