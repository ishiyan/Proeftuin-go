package main

import (
	"encoding/json"
	"enum-string/enums1"
	"fmt"
)

type Foo struct {
	First  enums1.Enum1 `json:"first,omitempty"`
	Second enums1.Enum1 `json:"second,omitempty"`
}

func Good() {
	fmt.Println("Good()")
	f := Foo{enums1.Aperiodic, enums1.Sec1}
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
	fmt.Println("f unmarshalled:", b)
}

func BadUnmarshalling() {
	fmt.Println("BadUnmarshalling()")
	s := "{\"first\":\"aperiodic\",\"second\":\"z1sec\"}"
	bs := []byte(s)
	fmt.Println("string:", s)
	fmt.Println("unmarshal")
	b := Foo{}
	if err := json.Unmarshal(bs, &b); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("f unmarshalled:", b)
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
	fmt.Println("f unmarshalled:", b)
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
	fmt.Println("f unmarshalled:", b)
}

func main() {
	Good()
	BadUnmarshalling()
	Empty()
	Nil()
	fmt.Println("done")
}
