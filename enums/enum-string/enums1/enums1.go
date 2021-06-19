package enums1

import (
	"bytes"
	"fmt"
	"time"
)

type Enum1 string

// Aperiodic foo bar.
const Aperiodic Enum1 = "aperiodic" // Aperiodic does not have a period.
const Sec1 Enum1 = "1sec"           // Sec1 has a period of 1 second.

func (e Enum1) Duration() time.Duration {
	switch e {
	case Aperiodic:
		return time.Nanosecond
	case Sec1:
		return time.Second
	default:
		return time.Hour * 24
	}
}

func (e Enum1) IsKnown() bool {
	switch e {
	case Aperiodic, Sec1:
		return true
	default:
		return false
	}
}

func (e Enum1) MarshalJSON() ([]byte, error) {
	if !e.IsKnown() {
		return nil, fmt.Errorf("cannot marshal unknown enum1 value '%v'", e)
	}

	b := make([]byte, 0, len(e)+2)
	b = append(b, '"')
	b = append(b, e...)
	b = append(b, '"')
	return b, nil
}

func (e *Enum1) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	v := Enum1(d)
	if !v.IsKnown() {
		return fmt.Errorf("cannot unmarshal unknown enum1 value '%v'", v)
	}
	*e = v
	return nil
}
