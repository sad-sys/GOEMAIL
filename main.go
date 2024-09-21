package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("SMTP Listener functioning")

	//Listner for incoming TCP connections on port 25
	listener, err := net.Listen("tcp", ":25")
	if err != nil {
		log.Fatalf("Error starting TCP Listener: %v", err)
	}
	defer listener.Close()

	fmt.Println("Listening on port 25...")
}
