package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

	reader := bufio.NewReader(conn)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("Error reading from the connection: %v", err)
			return
		}

		if readLine(conn, line) {
			break
		}
	}

}

func readLine(conn net.Conn, line string) bool {
	line = strings.TrimSpace(line)
	log.Printf("Received %s", line)

	if strings.HasPrefix(strings.ToUpper(line), "HELO") {
		returnHello(conn, line)
	} else if strings.HasPrefix(strings.ToUpper(line), "QUIT") {
		returnQuit(conn, line)
		// Returning true signals that we should close the connection
		return true
	} else {
		conn.Write([]byte("500 Unrecognized command\r\n"))
	}
	return false
}

func returnHello(conn net.Conn, line string) {
	domain := strings.TrimSpace(line[4:])
	response := fmt.Sprintf("250 Hello %s, pleased to meet you\r\n", domain)
	conn.Write([]byte(response))
}

func returnQuit(conn net.Conn, line string) {
	conn.Write([]byte("221 Bye\r\n"))
	log.Printf("Connection closed by client")
}
