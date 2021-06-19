package enums2

import (
	"bytes"
	"fmt"
	"time"
)

type Enum2 struct {
	value string
}

var Aperiodic = Enum2{"aperiodic"}
var Sec1 = Enum2{"1sec"}

func (e Enum2) Duration() time.Duration {
	switch e {
	case Aperiodic:
		return time.Nanosecond
	case Sec1:
		return time.Second
	default:
		return time.Hour * 24
	}
}

func (e Enum2) IsPredefined() bool {
	switch e {
	case Aperiodic, Sec1:
		return true
	default:
		return false
	}
}

func (e Enum2) MarshalJSON() ([]byte, error) {
	if !e.IsPredefined() {
		return nil, fmt.Errorf("cannot marshal unknown enum2 value '%v'", e.value)
	}

	b := make([]byte, 0, len(e.value)+2)
	b = append(b, '"')
	b = append(b, e.value...)
	b = append(b, '"')
	return b, nil
}

func (e *Enum2) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	v := Enum2{string(d)}
	if !v.IsPredefined() {
		return fmt.Errorf("cannot unmarshal unknown enum2 value '%v'", v.value)
	}
	e.value = v.value
	return nil
}

func (e Enum2) MarshalText() ([]byte, error) {
	if !e.IsPredefined() {
		return nil, fmt.Errorf("cannot marshal unknown enum2 value '%v'", e)
	}

	b := make([]byte, 0, len(e.value))
	b = append(b, e.value...)
	return b, nil
}

func (e *Enum2) UnmarshalText(data []byte) error {
	v := Enum2{string(data)}
	if !v.IsPredefined() {
		return fmt.Errorf("cannot unmarshal unknown enum2 value '%v'", v)
	}
	e.value = v.value
	return nil
}
