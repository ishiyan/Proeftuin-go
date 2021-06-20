package sort

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWithTimeout(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "bubble sort took more than 500 ms")
		return
	}
	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestBubbleSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
