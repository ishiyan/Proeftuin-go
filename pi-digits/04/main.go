package main

// https://www.angio.net/pi/pi-programs.html

// Compute Pi using Machin's formula, but with no
// special support for big numbers (shows how floating
// point precision limits how much you can compute)

import (
	"math"
	"fmt"
)

// Go doesn't provide a native arccot functon.
func arccot(x float64) float64 {
	return math.Atan(1.0/x)
}

func main() {
	pi := 4 * (4 * arccot(5) - arccot(239))
	fmt.Println("Pi is approximately: ", pi)
	fmt.Println("Pi should be:         3.1415926535897932")
	fmt.Println("                                       ^")
}
