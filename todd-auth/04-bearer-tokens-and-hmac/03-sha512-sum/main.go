package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println(sha512.New().Sum(nil))
}
