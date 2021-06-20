package main

import (
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println(len(sha512.New().Sum(nil)) * 8)
}
