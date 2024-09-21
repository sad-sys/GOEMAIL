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

	for {
		// Accept an incoming connection
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Error accepting connection %v", err)
			continue
		}
		//Handle connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Send SMTP greeting
	_, err := conn.Write([]byte("220 Welcome to Go SMTP Listener Sadiq.com \r\n"))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
		return
	}

	// Read data from the client
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Error reading from connection: %v", err)
			return
		}

		// Echo the data back to the client for now (just for testing)
		conn.Write(buffer[:n])
	}
}
