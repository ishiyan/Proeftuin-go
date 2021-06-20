package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Bond"] = 42
	if age, ok := m["Bond"]; ok {
		// exists
		fmt.Println(age, ok)
	}

	// m["Moneypenny"] = 21
	if age, ok := m["Moneypenny"]; !ok {
		// does not exist
		fmt.Println(age, ok)
	}
}