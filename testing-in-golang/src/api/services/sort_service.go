package services

import (
	"proeftuin/testing-in-golang/src/api/utils/sort"
	)

func Sort(elements []int) {
	if len(elements) <= 20000 {
		sort.BubbleSort(elements)
		return
	}
	sort.Sort(elements)
}
