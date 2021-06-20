package main

import "fmt"

// return a function which is returning int
func int_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// go supports anonymous functions which can form closures
	next_int := int_seq()
	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())

	// as we can see from the output, it printed 1, 2, and then 3, and so on
	// this means the int_seq() is keeping track of the values of i even though it's gone
	// out of scope (outside its function body)
	// that's the key point of a closure

	// a closure is a function value (next_int()) that references variables (i) from outside
	// its body (int_seq()), but still remembers the value

	// the function may access and assign to the referenced variables; in this sense the
	// function is "bound" to the variables

	// let's create another separate instance of the function, int_seq()
	next_int2 := int_seq()
	fmt.Println(next_int2())
	fmt.Println(next_int2())
	fmt.Println(next_int2())

	// we see the two instances are independent of each other, they are separate instances!
}
