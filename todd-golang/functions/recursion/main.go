package main

import (
	"fmt"
)

// https://play.golang.org/p/t3WT_KXsA0E

func main() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	f := n
	for n > 1 {
		n--
		f *= n
	}
	return f
}
