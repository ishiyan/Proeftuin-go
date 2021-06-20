package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var name string
	if len(os.Args) > 1 {
		fmt.Println(os.Args[0])
		fmt.Println(os.Args[1])
		name = os.Args[1]
	} else {
		name = "Miss Moneypenny"
	}
	str := fmt.Sprint(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hello World!</title>
</head>
<body>
<h1>` +
		name + `</h1>
</body>
</html>
`)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}

// at the terminal:
// go run main.go Todd
