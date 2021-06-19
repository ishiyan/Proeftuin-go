// Preventing reimplementation of interfaces, or better sum types for go
package main

import (
	"fmt"
)

// People often (somewhat rightfully) complain that go doesn't have decent
// sum/enum types.  The usual wisdom is, to use interfaces for that. If you
// make the interface include an unexported method, in theory, there can't be
// any implementations of that interface. The method is unexported, so no type
// of a third-party package can ever provide an implementation. This is sometimes
// called a "marker-method":

type SumType interface {
	sumType()
}

type A int

func (a A) sumType() {}

type B float64

func (b B) sumType() {}

type C string

func (c C) sumType() {}

func String(s SumType) string {
	switch v := s.(type) {
	case A:
		return fmt.Sprintf("A(%d)", v)
	case B:
		return fmt.Sprintf("B(%f)", v)
	case C:
		return fmt.Sprintf("C(%q)", v)
	default:
		panic("invalid SumType")
	}
}

func DemoSumType() {
	a, b, c := A(42), B(13.37), C("fnord")
	fmt.Println(String(a), String(b), String(c))
}

// As far as I am aware, until now, the usual response has been, that this
// still doesn't prevent reimplementations (and I have argued in the past that
// it is theoretically impossible to prevent), due to type embeddings. So
// someone could do, in another package:

type Embedded struct {
	SumType
}

func DemoSumType2() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered:", v)
		}
	}()

	E := Embedded{A(42)}
	fmt.Println(String(E))
}

// Whenever you define an interface type I, anyone can trivially re-implement
// it, by defining type T struct { I }. However, with a slight adjustment, this
// stays true, but you still get safe(r) sum types. Instead of making the
// method a pure marker, we make it return a value, using recursive interfaces:

type BetterSumType interface {
	betterSumType() BetterSumType
}

type BetterA int

func (a BetterA) betterSumType() BetterSumType {
	return a
}

type BetterB float64

func (b BetterB) betterSumType() BetterSumType {
	return b
}

type BetterC string

func (c BetterC) betterSumType() BetterSumType {
	return c
}

func BetterString(s BetterSumType) string {
	switch v := s.betterSumType().(type) {
	case BetterA:
		return fmt.Sprintf("BetterA(%d)", v)
	case BetterB:
		return fmt.Sprintf("BetterB(%f)", v)
	case BetterC:
		return fmt.Sprintf("BetterC(%q)", v)
	default:
		panic("invalid BetterSumType")
	}
}

// The effect is, that while re-implementations are possible, you only ever
// have to deal with well-knows implementations. Because the only valid
// implementations of the *method* betterSumType happen in this package, and we
// can ensure that all of them return a well-known implementation. The reason
// this works is, that embedding isn't inheritance. If you embed a type T into
// a struct S, the receiver of the promoted method will still be T, not S, so
// embedded values will be "unpacked" first:

type BetterEmbedded struct {
	BetterSumType
}

func DemoBetterSumType() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered:", v)
		}
	}()

	a, b, c, e := BetterA(42), BetterB(13.37), BetterC("fnord"), BetterEmbedded{BetterC("works")}
	fmt.Println(BetterString(a), BetterString(b), BetterString(c), BetterString(e))
}

// A downside compared to real sum types is that it's still possible for the
// embedded value (or the value passed into BetterString) to be nil, in which case
// s.betterSumType() in line 99 will panic. This is equivalent to the same
// situation before or to someone passing in a nil-value as the parameter.
//
// All of this might or might not be news for you. It is news to me, so I
// wanted to write it down and publish it. It gives me warm fuzzy feelings, that
// even though go is simple and I consider myself pretty familiar with the
// language, I still discover new corners every once in a while. Let me know,
// what you think. :)
//
// Merovius
//

func main() {
	DemoSumType()
	DemoSumType2()
	fmt.Println()
	DemoBetterSumType()
}

// © 2016 Axel Wagner <mero@merovius.de>, published under CC-BY
// https://creativecommons.org/licenses/by/2.0/
