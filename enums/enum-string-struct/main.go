package main

import (
	"encoding/json"
	"enum-string-struct/enums2"
	"fmt"
)

type Foo struct {
	First  enums2.Enum2 `json:"first,omitempty"`
	Second enums2.Enum2 `json:"second,omitempty"`
}

func NewEnumValue() {
	fmt.Println("NewEnumValue()")
	empty := enums2.Enum2{}
	fmt.Println("new empty enum2 value:", empty)

	var null enums2.Enum2
	fmt.Println("new null enum2 value:", null)

	// normal := enums2.Enum2{"normal"} // implicit assignment to unexported field value
	// fmt.Println("new normal enum2 value:", normal)
}

func Good() {
	fmt.Println("Good()")
	f := Foo{enums2.Aperiodic, enums2.Sec1}
	fmt.Println("f:", f)
	fmt.Println("f.First:", f.First.Duration(), "f.Second:", f.Second.Duration())
	fmt.Println("marshal f")
	fbs, err := json.Marshal(f)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f marshaled:", string(fbs))
	fmt.Println("unmarshal f")
	b := Foo{}
	if err := json.Unmarshal(fbs, &b); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f unmarshaled:", b)
	if f != b {
		fmt.Println("error: f not equal b, f", f, "b:", b)
	}
}

func BadUnmarshaling() {
	fmt.Println("BadUnmarshaling()")
	s := "{\"first\":\"aperiodic\",\"second\":\"z1sec\"}"
	bs := []byte(s)
	fmt.Println("string:", s)
	fmt.Println("unmarshal")
	b := Foo{}
	if err := json.Unmarshal(bs, &b); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f unmarshaled:", b)
}

func Empty() {
	fmt.Println("Empty()")
	f := Foo{}
	fmt.Println("f:", f)
	fmt.Println("f.First:", f.First.Duration(), "f.Second:", f.Second.Duration())
	fmt.Println("marshal f")
	fbs, err := json.Marshal(f)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f marshaled:", string(fbs))
	fmt.Println("unmarshal f")
	b := Foo{}
	if err := json.Unmarshal(fbs, &b); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f unmarshaled:", b)
}

func Nil() {
	fmt.Println("Nil()")
	var f Foo
	fmt.Println("f:", f)
	fmt.Println("f.First:", f.First.Duration(), "f.Second:", f.Second.Duration())
	fmt.Println("marshal f")
	fbs, err := json.Marshal(f)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f marshaled:", string(fbs))
	fmt.Println("unmarshal f")
	b := Foo{}
	if err := json.Unmarshal(fbs, &b); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f unmarshaled:", b)
}

func main() {
	NewEnumValue()
	Good()
	BadUnmarshaling()
	Empty()
	Nil()
	fmt.Println("done")
}
