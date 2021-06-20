package main

import (
	"fmt"
	"math"
)

// both scale and size are with receiver type *Point, even though the size method does not
// modify any elements of the Point struct receiver
type Point struct {
	x, y float64
}

func (p *Point) scale(s float64) {
	p.x = p.x * s
	p.y = p.y * s
}

func (p *Point) size() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	// There are two reasons to use a pointer receiver.
	// 1. The first is so that the method can modify the value that its receiver points to.
	// 2. The second is to avoid copying the value on each method call. This can be more
	//    efficient if the receiver is a large struct, for example.

	// pointer receiver
	p := &Point{3, 4}
	fmt.Printf("before scaling: %+v size = %v\n", p, p.size())
	var sc float64 = 5
	p.scale(sc)
	fmt.Printf("after scaling: %+v size = %v\n", p, p.size())

}
