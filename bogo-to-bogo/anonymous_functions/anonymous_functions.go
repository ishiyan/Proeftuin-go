package main

import "fmt"

// a regular function (with its name)
func say_hello(msg string) {
	fmt.Println(msg)
}

// a regular function returning anonymous function
// return type func(string) -- anonymous function taking a string
func regular_f_returning_anonymous_f() func(string) {
	// returns an anonymous function which is an inner function
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {
	// a regular function (with its name)
	say_hello("Hello from a regular function")

	// an anonymous function (without any function name)
	func(msg string) {
		fmt.Println(msg)
	}("Hello from an anonymous function")
	// note that the () means call and execute the anonymous function
	// if we write only func(){} without () in func(){}(), we only declare the anonymous function without calling it

	print_fnc := regular_f_returning_anonymous_f()
	print_fnc("Hello from returned anonymous function")
}
