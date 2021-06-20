package main

import "fmt"

// check if integer is a prime or not
func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// takes a slice as its argument and gathers primes. Then, returning it to the caller:
func getPrimes(n int, a []int) []int {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			a = append(a, i)
		}
	}
	return a
}

func main() {
	n := 100
	ans := []int{}
	fmt.Printf("%+v\n", getPrimes(n, ans[0:]))
}
