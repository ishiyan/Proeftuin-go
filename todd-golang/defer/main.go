package main

import (
	"fmt"
)

// https://play.golang.org/p/ssnC-JmXCFY

func main() {
	foo()
}

func foo() {
	defer func() {
		fmt.Println("i am deffered")
	}()

	fmt.Println("hello from foo")
}
