package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("goroutine function")

	// goroutine with anonymous function
	go func() {
		fun("goroutine anonymous function")
	}()

	// goroutine with function value call
	fv := fun
	go fv("goroutine function value")

	// wait for goroutines to end
	fmt.Println("wait for goroutines ...")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("done..")
}
