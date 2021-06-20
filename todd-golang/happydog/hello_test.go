package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "hello"
	if actual := Hello(); actual != expected {
		t.Errorf("Expected %q, actual %s", expected, actual)
	}
}
