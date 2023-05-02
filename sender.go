package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// the first client send the data to the second client using this
// the port value must be the same as for the receiver and i suggest both to be 3035
func SendDataToServer(message string, serverIPAddress, port string) {
	// Set up the request data
	data := []byte(message)

	// Create a new request with POST method and request body
	// serverIPAddress = "http://192.168.1.4"
	req, err := http.NewRequest("POST", serverIPAddress+":"+port+"/path/to/endpoint", bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the content type header to specify form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()
}
