package main

import "fmt"

// this is a recursive function
func fibo_r(n int) int {
	if n <= 1 {
		return n
	}
	return fibo_r(n-1) + fibo_r(n-2)
}

// as an another sample for the closures, the fiboi() is a function that returns a function that returns an int
// this is an iterative function
func fibo_i() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y
		return r
	}
}

func main() {
	n := 10
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibo_r(i))
	}
	fmt.Println()

	next_fibo := fibo_i()
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", next_fibo())
	}
	fmt.Println()

}
