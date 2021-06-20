package main

import "fmt"

// fizz_buzz is short program that prints each number from 1 to 100 on a new line
// For each multiple of 3, print "Fizz" instead of the number, for each multiple of 5, print "Buzz",
// and for numbers which are multiples of both 3 and 5, print "FizzBuzz"

func main() {
	for i := 1; i < 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizz-buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
