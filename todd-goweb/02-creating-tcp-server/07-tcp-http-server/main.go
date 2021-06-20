package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	uri := request(conn)

	// write response
	respond(conn, uri)
}

func request(conn net.Conn) string {
	i := 0
	uri := ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			fmt.Println("***METHOD", strings.Fields(ln)[0])
			uri = strings.Fields(ln)[1]
			fmt.Println("***URI", uri)
			fmt.Println("***HTTP version", strings.Fields(ln)[2])
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return uri
}

func respond(conn net.Conn, uri string) {
	withURI := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Your URI is</strong>&nbsp;%s</body></html>`
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Your have no URI</strong></body></html>`
	if len(uri) > 0 {
		body = fmt.Sprintf(withURI, uri)
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
