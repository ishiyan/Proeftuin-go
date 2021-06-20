package main

import "fmt"

type sized interface {
	getWidth() int
	setWidth(width int)
	getHeight() int
	setHeight(height int)
}

type rectangle struct {
	width, height int
}

// vvv !! POINTER

func (r *rectangle) getWidth() int {
	return r.width
}

func (r *rectangle) setWidth(width int) {
	r.width = width
}

func (r *rectangle) getHeight() int {
	return r.height
}

func (r *rectangle) setHeight(height int) {
	r.height = height
}

// Modified LSP
// If a function takes an interface and
// works with a type T that implements this
// interface, any structure that aggregates T
// should also be usable in that function.

type square struct {
	rectangle
}

func newSquare(size int) *square {
	sq := square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *square) setWidth(width int) {
	s.width = width
	s.height = width
}

func (s *square) setHeight(height int) {
	s.width = height
	s.height = height
}

// one of ways to fix it

type square2 struct {
	size int
}

func (s *square2) rectangle() rectangle {
	return rectangle{s.size, s.size}
}

func useIt(sized sized) {
	width := sized.getWidth()
	sized.setHeight(10)
	expectedArea := 10 * width
	actualArea := sized.getWidth() * sized.getHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rc := &rectangle{2, 3}
	useIt(rc)

	sq := newSquare(5)
	useIt(sq)
}
