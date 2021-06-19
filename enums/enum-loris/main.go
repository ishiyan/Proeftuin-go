package main

import (
	"enum-loris/days"
	"fmt"
)

func main() {
	//january := days.Day{"january"} // implicit assignment of unexported field 'value' in day.Day literal
	//getTask(january)

	//var march struct {
	//	value string
	//}
	//march.value = "march"
	//getTask(march) // cannot use march (type struct { value string }) as type day.Day in argument to getTask

	getTask(days.Monday)
	getTask(days.Tuesday)
	iterateDays()
}

func getTask(d days.Day) string {
	if d == days.Monday {
		fmt.Println("today is ", d, "!") // today is Monday !
		return "running"
	}
	if d == days.Tuesday {
		fmt.Println("today is ", d, "!") // today is Tuesday !
		return "running"
	}

	return "nothing to do"
}

func iterateDays() {
	fmt.Println("iterateDays:")
	for _, d := range days.Days {
		fmt.Println(d)
	}
}
