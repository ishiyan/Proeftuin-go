package main

// notice that we imported packages enclosed with '()', no comma(',') between them
import (
	"bytes"
	"fmt" // fmt is the name of a package that includes a variety of functions related to formatting and output to the screen
	"math"

	"bogo.to.bogo/string_util" // notice that when we import our string_util library we use its full name
)

// outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available
// k := 3

// note that we specified the return type, in this case, a string
func greetings(name string) string {
	return "Hello " + name
}

// we use a short form instead of 'x int, y int'
func add(x, y int) int {
	return x + y
}

type Cities struct {
	name     string
	location [2]int
}

func main() {
	fmt.Println("Hello, world!")

	/////////////////////////////////////////////////////////////////////////////////////////////
	// data types and variables
	/////////////////////////////////////////////////////////////////////////////////////////////

	// variables in Go are created by first using the var keyword, then specifying the variable name (x), the type (string)
	// and finally assigning a value
	// the Go can infer the type, we do not have to explicitly specify the type. So, we can write the code without string
	// var x string = "Hello, world!";
	var x = "Hello, world!"
	var year int
	year = 2019
	var next_year int64 = 2020
	fmt.Println(x, year, next_year)

	// we can print out the types of variables using Printf() and %T
	fmt.Printf("type of x: %T\n", x)
	fmt.Printf("type of year: %T\n", year)
	fmt.Printf("type of next_year: %T\n", next_year)

	var i, j int = 1, 2
	// inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
	k := 3
	// := introduces "a new variable", using it twice does not redeclare a second variable, so it's illegal
	// k := 4
	// k = 5 is legal, because, it just reassigns a new value to k
	k = 5
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("%T %T %T %T %T %T\n", i, j, k, c, python, java)

	/////////////////////////////////////////////////////////////////////////////////////////////
	// byte and rune
	/////////////////////////////////////////////////////////////////////////////////////////////

	// Go has integer types called byte and rune that are aliases for uint8 and int32 data types, respectively

	// the byte and rune data types are used to distinguish characters from integer values

	// in Go, there is no char data type. It uses byte and rune to represent character values

	// the byte data type represents ASCII characters while the rune data type represents a more broader set of
	// Unicode characters that are encoded in UTF-8 format

	// the default type for character values is rune, which means, if we don't declare a type explicitly when
	// declaring a variable with a character value, then Go will infer the type as rune

	// characters are expressed by enclosing them in single quotes like this: 'a'
	var myLetter = 'R' // type inferred as rune which is the default type for character values
	fmt.Printf("rune = %T\n", myLetter)

	// we can create a byte variable by explicitly specifying the type
	var anotherLetter byte = 'B'
	fmt.Printf("byte = %T\n", anotherLetter)

	// for example, a byte variable with value 'a' is converted to the integer 97 while a rune variable with
	// a unicode value '~' is converted to the corresponding unicode codepoint U+007E, where U+ means unicode
	// and the numbers are hexadecimal, which is essentially an integer
	my_byte := byte('a')
	my_rune := '~'
	fmt.Printf("%c = %d, %c = %U\n", my_byte, my_byte, my_rune, my_rune)

	// bytes.Join() concatenates the elements of input slice to create a new byte slice. In other words,
	// Join concatenates the elements of s to create a new byte slice.
	// The separator sep is placed between elements in the resulting slice:
	// Join(s [][]byte, sep []byte) []byte
	b1 := byte('a')
	b2 := []byte("A")
	b3 := []byte{'a', 'b', 'c'}
	fmt.Printf("b1 = %c, b2 = %c, b3 = %s\n", b1, b2, b3)
	s1 := []byte("Hello")
	s2 := []byte("world")
	s3 := [][]byte{s1, s2}
	s4 := bytes.Join(s3, []byte(", "))
	s5 := []byte{}
	s5 = bytes.Join(s3, []byte("--"))
	s6 := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("s1 = %s\n", s1)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("s3 = %s\n", s3)
	fmt.Printf("s4 = %s\n", s4)
	fmt.Printf("s5 = %s\n", s5)
	fmt.Printf("%s\n", bytes.Join(s6, []byte(", ")))

	/////////////////////////////////////////////////////////////////////////////////////////////
	// packages
	/////////////////////////////////////////////////////////////////////////////////////////////

	fmt.Println(math.Floor(1.7), math.Ceil(9.2))

	// see string_util/string_util.go
	// inside of the hello.go file we only use the last part of the name (package string_util)
	fmt.Println(string_util.Reverse("reverse me"))

	/////////////////////////////////////////////////////////////////////////////////////////////
	// functions
	/////////////////////////////////////////////////////////////////////////////////////////////

	// the Go syntax for a function looks like this:
	// func name(param type) returntype {
	//	 // body
	// }

	fmt.Println(greetings("world"))
	fmt.Println(add(3, 4))

	/////////////////////////////////////////////////////////////////////////////////////////////
	// arrays
	/////////////////////////////////////////////////////////////////////////////////////////////

	// The slice type is an abstraction built on top of Go's array type, and so to understand slices
	// we must first understand arrays.

	// An array variable denotes the entire array; unlike C, array name is not a pointer to the first
	// array element. This means that when we assign or pass around an array value we will make a copy
	// of its contents. (To avoid the copy you could pass a pointer to the array, but then that's a
	// pointer to an array, not an array.

	var fruits [2]string
	fruits[0] = "pomegranate"
	fruits[1] = "rambutan"
	fmt.Println(fruits)

	// note that we can declare and assign at the same time
	fruits2 := [2]string{"pomegranate", "rambutan"}
	fmt.Println(fruits2)

	// or like this
	ids := []int{1, 3, 5, 7, 9}
	fmt.Println(ids)

	/////////////////////////////////////////////////////////////////////////////////////////////
	// slices
	/////////////////////////////////////////////////////////////////////////////////////////////

	// We don't see arrays too often in Go code. Slices, though, are everywhere. They build on arrays
	// to provide great power and convenience.

	// The type specification for a slice is []T, where T is the type of the elements of the slice.
	// Unlike an array type, a slice type has no specified length.

	// A slice is a data structure describing a contiguous section of an array stored separately from
	// the slice variable itself. A slice is not an array but it describes a piece of an array.

	// a slice literal is declared just like an array literal, except we leave out the element count
	alphabet := []string{"a", "b", "c", "d", "e"}
	fmt.Println(alphabet)

	// A slice has both a length and a capacity. The length of a slice is the number of elements it
	// contains. The capacity of a slice is the number of elements in the underlying array, counting
	// from the first element in the slice. The length and capacity of a slice s can be obtained using
	// the expressions len(s) and cap(s)

	// make (type, len, cap)
	sli := make([]int, 5, 10)
	fmt.Printf("len=%d, cap=%d, %v\n", len(sli), cap(sli), sli)

	// When the capacity argument is omitted, it defaults to the specified length. Here's a more
	// succinct version of the same code
	sli2 := make([]int, 5)
	fmt.Printf("len=%d, cap=%d, %v\n", len(sli2), cap(sli2), sli2)

	// The length and capacity of a slice can be inspected using the built-in len and cap functions
	fmt.Println(len(sli2) == 5, cap(sli2) == 5, sli2)

	fruits_slice := []string{"pomegranate", "rambutan", "mangosteen", "jackfruit"}
	fmt.Println(fruits_slice)
	fmt.Println(len(fruits_slice))
	fmt.Println(fruits_slice[1:3])
	fmt.Println(fruits_slice[1:4])
	fmt.Println(fruits_slice[3:])
	fmt.Println(fruits_slice[:])

	/////////////////////////////////////////////////////////////////////////////////////////////
	// byte slices
	/////////////////////////////////////////////////////////////////////////////////////////////

	// We'll create a byte slice from a string literal "abc" and append a byte to the byte slice.
	// Then, we convert the byte slice into a string with the string() built-in method. A byte slice
	// has a length, which we retrieve with len. Also, we can access individual bytes

	sli3 := []byte("abc")
	// print initial byte slice
	fmt.Println(sli3)
	// append a byte
	sli3 = append(sli3, byte('d'))
	// print string reprsentation of the byte slice
	fmt.Println(string(sli3))
	// length of the byte slice
	fmt.Println(len(sli3))
	// 1st byte
	fmt.Println(sli3[0])

	/////////////////////////////////////////////////////////////////////////////////////////////
	// a function taking and returning a slice
	/////////////////////////////////////////////////////////////////////////////////////////////

	// see prime_numbers/prime_numbers.go and shuffling_deck/shuffling_deck.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// conditionals
	/////////////////////////////////////////////////////////////////////////////////////////////

	// if / else
	a4, b4 := 1, 10
	if a4 < b4 {
		fmt.Printf("%d is less than %d\n", a4, b4)
	} else if a4 == b4 {
		fmt.Printf("%d equals %d\n", a4, b4)
	} else {
		fmt.Printf("%d is greater than %d\n", a4, b4)
	}

	// switch
	galaxy := "M87"
	switch galaxy {
	case "Milky Way":
		fmt.Printf("Galaxy name is a 'Milky Way'\n")
	case "Andromeda":
		fmt.Printf("Galaxy name is 'Andromeda'\n")
	case "M87":
		fmt.Printf("Galaxy name is 'M87'\n")
	}

	/////////////////////////////////////////////////////////////////////////////////////////////
	// loops
	/////////////////////////////////////////////////////////////////////////////////////////////

	// see fizz_buzz/fizz_buzz.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// maps
	/////////////////////////////////////////////////////////////////////////////////////////////

	// One of the most useful data structures in computer science is the hash table. It offers fast
	// lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.
	// A map maps keys to values.

	// to initialize a map, we can use the built-in make() function

	// define map
	moons := make(map[string]string)
	// assign
	moons["Earth"] = "Moon"
	moons["Jupiter"] = "Europa"
	moons["Saturn"] = "Titan"
	fmt.Println(moons)
	// delete
	delete(moons, "Saturn")
	fmt.Println(moons)

	// declare and define a map
	moons2 := map[string]string{"Earth": "Moon", "Jupiter": "Europa", "Saturn": "Titan"}
	fmt.Println(moons2)

	// see balanced_parentheses/balanced_parentheses.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// range
	/////////////////////////////////////////////////////////////////////////////////////////////

	// the range form of the for loop iterates over a slice or map

	// one of the simplest range examples
	for _, i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i)
	}

	// range with Arrays - []int{}
	ids2 := []int{0, 10, 20, 30, 40, 50, 60}
	for i, id := range ids2 {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// if do not want to use index
	for _, id := range ids2 {
		fmt.Printf("ID: %d\n", id)
	}
	// sum of ids
	sum := 0
	for _, id := range ids2 {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// range with string indexes and runes
	for i, c := range "ㄱㄴㄷㄹㅁㅂㅅㅇㅈㅊㅋㅌㅍㅎ" {
		fmt.Printf("%#U starts at byte position %d\n", c, i)
	}

	// range with maps - map[string]string{}
	moons3 := map[string]string{"Earth": "Moon", "Jupiter": "Europa", "Saturn": "Titan"}
	for k, v := range moons3 {
		fmt.Printf("%s: %s\n", k, v)
	}

	// range with maps - map[string]int{}
	numbers := map[string]int{"Uno": 1, "Dos": 2, "Tres": 3, "Cuatro": 4, "Cinco": 5}
	for k, v := range numbers {
		fmt.Println(k, v)
	}

	// range with channel
	// we'll iterate over 5 values in the 'queue' channel
	queue := make(chan string, 5)
	queue <- "Enceladus"
	queue <- "Titan"
	queue <- "Europa"
	queue <- "Ganemede"
	queue <- "Io"
	close(queue)
	// this 'range' iterates over each element as it's received from 'queue'
	// because we closed the channel above, the iteration terminates after receiving the 5 queues
	for q := range queue {
		fmt.Println(q)
	}

	// range with struct
	// create empty slice of struct pointers
	cities := []*Cities{}
	// create struct and append it to the slice
	ct := new(Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 0
	cities = append(cities, ct)
	// create another struct
	ct = new(Cities)
	ct.name = "London"
	ct.location[0] = 34
	ct.location[1] = 51
	cities = append(cities, ct)
	// print
	for i := range cities {
		c := cities[i]
		fmt.Println("City: ", *c)
	}

	/////////////////////////////////////////////////////////////////////////////////////////////
	// pointers
	/////////////////////////////////////////////////////////////////////////////////////////////

	// a pointer holds the memory address of a value
	a9 := 10
	b9 := &a9 // the & operator generates a pointer to its operand
	fmt.Println(a9, b9)
	fmt.Printf("%T %T\n", a9, b9)

	// the * operator denotes the pointer's underlying value
	i9, j9 := 42, 2701
	p9 := &i9        // points to i9
	fmt.Println(*p9) // read i9 through the pointer
	*p9 = 21         // set i9 through the pointer
	fmt.Println(i9)  // see the new value of i9
	p9 = &j9         // point to j9
	*p9 = *p9 / 37   // divide j9 through the pointer
	fmt.Println(j9)  // see the new value of j9

	/////////////////////////////////////////////////////////////////////////////////////////////
	// anonymous functions
	/////////////////////////////////////////////////////////////////////////////////////////////

	// see anonymous_functions/anonymous_functions.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// closures
	/////////////////////////////////////////////////////////////////////////////////////////////

	// see closures/closures.go
	// see fibonacci_sequence_using_closures/fibonacci_sequence_using_closures.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// structs and receiver methods
	/////////////////////////////////////////////////////////////////////////////////////////////

	// see structs_and_receiver_methods/structs_and_receiver_methods.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// value or pointer receiver
	/////////////////////////////////////////////////////////////////////////////////////////////

	// see value_or_pointer_receiver/value_or_pointer_receiver.go

	/////////////////////////////////////////////////////////////////////////////////////////////
	// xxxxx
	/////////////////////////////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////////////////////////////
	// xxxxx
	/////////////////////////////////////////////////////////////////////////////////////////////

	/////////////////////////////////////////////////////////////////////////////////////////////
	// xxxxx
	/////////////////////////////////////////////////////////////////////////////////////////////

}
