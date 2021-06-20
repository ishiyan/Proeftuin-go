package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// write server program to handle concurrent client connections.
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		handleConn(conn)
	}
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
