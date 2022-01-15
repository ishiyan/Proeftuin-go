package assignment

import (
	"fmt"
	"testing"
)

func TestExecution(t *testing.T) {
	t.Parallel()

	t.Run("plain build", func(t *testing.T) {
		t.Parallel()
		a := NewAccountBuilder("foo").Build()

		b := AccountClient(a)
		if b != nil {
			fmt.Print("good")
		} else {
			fmt.Print("bad")
		}
	})
}
