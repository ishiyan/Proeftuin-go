package main

import (
	"fmt"
)

// https://play.golang.org/p/g1Z7ZyjUOwS

func main() {
	func() {
		fmt.Println("hello from anonymous func")
	}()
}
