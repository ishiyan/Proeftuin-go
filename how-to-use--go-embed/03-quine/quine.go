package main

import (
	_ "embed"
	"fmt"
)

//go:embed quine.go
var src string

func main() {
	fmt.Print(src)
}
