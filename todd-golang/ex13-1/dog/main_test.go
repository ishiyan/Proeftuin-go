package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{input: 0, want: 0},
		{input: 1, want: 7},
		{input: 2, want: 14},
		{input: 10, want: 70},
	}
	for _, tt := range tests {
		if got := Years(tt.input); got != tt.want {
			t.Errorf("Years(%v) = %v, want %v", tt.input, got, tt.want)
		}
		if got := YearsTwo(tt.input); got != tt.want {
			t.Errorf("YearsTwo(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(8))
	// Output:
	// 56
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(9))
	// Output:
	// 63
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(i)
	}
}
