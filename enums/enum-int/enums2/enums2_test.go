package enums2

import (
	"encoding/json"
	"testing"
)

// Run: go test -bench=.

func BenchmarkDuration(b *testing.B) {
	act := Enum2(999)

	for i := 0; i < b.N; i++ {
		act.Duration()
	}
}

func BenchmarkString(b *testing.B) {
	act := Enum2(999)

	for i := 0; i < b.N; i++ {
		act.String()
	}
}

func BenchmarkIsKnown(b *testing.B) {
	act := Enum2(999)

	for i := 0; i < b.N; i++ {
		act.IsKnown()
	}
}

type Test struct {
	Test Enum2 `json:"test,omitempty"`
}

func BenchmarkMarshalJSON(b *testing.B) {
	act := Test{Sec1}

	for i := 0; i < b.N; i++ {
		json.Marshal(act)
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	s := "{\"test\":\"sec1\"}"
	bs := []byte(s)
	var act Test

	for i := 0; i < b.N; i++ {
		json.Unmarshal(bs, &act)
	}
}

func BenchmarkMarshalJSONUnknown(b *testing.B) {
	act := Test{Enum2(999)}

	for i := 0; i < b.N; i++ {
		json.Marshal(act)
	}
}

func BenchmarkUnmarshalJSONUnknown(b *testing.B) {
	s := "{\"test\":\"foobar\"}"
	bs := []byte(s)
	var act Test

	for i := 0; i < b.N; i++ {
		json.Unmarshal(bs, &act)
	}
}
