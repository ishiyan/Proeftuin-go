package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	io.WriteString(c, "I see you connected.")

}
