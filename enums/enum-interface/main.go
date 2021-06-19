package main

import (
	"enum-interfaces/searchrequests"
	"fmt"
)

// Let's test the properties.

func typeSwitch(e searchrequests.SearchRequest) string {
	switch e.(type) {
	case searchrequests.Universal, searchrequests.Web, searchrequests.Images, searchrequests.Local, searchrequests.News, searchrequests.Products, searchrequests.Video:
		return "matches valid cases (if you remember to list all of them)"
	case nil:
		return "an enum value can be nil ☹"
	default:
		return "embedding does not circumvent type switch"
	}
}

func main() {
	// Iterable and can be used as strings or ints.
	for _, e := range searchrequests.Iter() {
		fmt.Printf("%d: %s\n", e.Int(), e)
	}

	fmt.Println()

	fmt.Println(searchrequests.FromInt(4))
	fmt.Println(searchrequests.FromInt(-4))

	fmt.Println(searchrequests.FromString("LOCAL"))
	fmt.Println(searchrequests.FromString("local"))

	fmt.Println()

	// Comparable.
	var v, n searchrequests.SearchRequest = searchrequests.Video{}, searchrequests.News{}
	if v != n /* && n == n */ {
		fmt.Println("comparison works")
	}

	// Type switches work, sorta.
	fmt.Println(typeSwitch(v))

	// An Enum value can be nil.
	var test searchrequests.SearchRequest
	fmt.Println(typeSwitch(test))

	// A false Enum value can be created, but Merovius' trick prevents it from messing up the type switch.
	test = struct{ searchrequests.SearchRequest }{searchrequests.Local{}}
	fmt.Println(typeSwitch(test))

	//and they don't mess with comparison
	if test == searchrequests.SearchRequest(searchrequests.Local{}) {
		fmt.Println("comparison broken ☹")
	}
}
