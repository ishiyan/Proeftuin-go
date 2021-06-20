package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
		go serve(c)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var i int
	var method, uri string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if len(ln) == 0 {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		fmt.Println(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			method = xs[0]
			uri = xs[1]
			fmt.Println("Method:", method)
			fmt.Println("URI:", uri)
		}
		i++
	}

	var body string
	switch {
	case method == "GET" && uri == "/":
		body = indexBody
	case method == "GET" && uri == "/apply":
		body = applyBody
	case method == "POST" && uri == "/apply":
		body = applyPostBody
	default:
		body = defaultBody
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

const defaultBody = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Code Gangsta</title>
</head>
<body>
	<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
</body>
</html>
`

const indexBody = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>GET INDEX</title>
</head>
<body>
	<h1>"GET INDEX"</h1>
	<a href="/">index</a><br>
	<a href="/apply">apply</a><br>
</body>
</html>
`

const applyBody = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>GET DOG</title>
</head>
<body>
	<h1>"GET APPLY"</h1>
	<a href="/">index</a><br>
	<a href="/apply">apply</a><br>
	<form action="/apply" method="POST">
	<input type="hidden" value="In my good death">
	<input type="submit" value="submit">
	</form>
</body>
</html>
`

const applyPostBody = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>POST APPLY</title>
</head>
<body>
	<h1>"POST APPLY"</h1>
	<a href="/">index</a><br>
	<a href="/apply">apply</a><br>
</body>
</html>
`
