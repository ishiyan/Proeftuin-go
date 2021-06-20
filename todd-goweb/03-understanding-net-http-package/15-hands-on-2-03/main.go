package main

import (
	"bufio"
	"fmt"
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

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(c)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		fmt.Println(ln)
	}

	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected.")
}
