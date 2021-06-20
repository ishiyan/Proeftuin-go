package main

import (
	"fmt"
)

// https://play.golang.org/p/dLt69Y14OA9

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Odd sum:", odd(sum, ii...))
}

func sum(xi ...int) int {
	tot := 0
	for _, v := range xi {
		tot += v
	}
	return tot
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
