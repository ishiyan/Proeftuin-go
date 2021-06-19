package enums2

import (
	"bytes"
	"fmt"
	"time"
)

type Enum2 int

// Aperiodic foo bar.
const Aperiodic Enum2 = 1 // Aperiodic does not have a period.
const Sec1 Enum2 = 2      // Sec1 has a period of 1 second.

const aperiodic = "aperiodic"
const sec1 = "sec1"

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

func (e Enum2) String() string {
	switch e {
	case Aperiodic:
		return aperiodic
	case Sec1:
		return sec1
	default:
		return "unknown"
	}
}

func (e Enum2) IsKnown() bool {
	switch e {
	case Aperiodic, Sec1:
		return true
	default:
		return false
	}
}

func (e Enum2) MarshalJSON() ([]byte, error) {
	if !e.IsKnown() {
		return nil, fmt.Errorf("cannot marshal unknown enum2 value '%v'", e)
	}

	s := e.String()
	b := make([]byte, 0, len(s)+2)
	b = append(b, '"')
	b = append(b, s...)
	b = append(b, '"')
	// fmt.Printf("... Marshal JSON '%v' -> %v\n", s, b)
	return b, nil
}

func (e *Enum2) UnmarshalJSON(data []byte) error {
	d := bytes.Trim(data, "\"")
	s := string(d)
	switch s {
	case aperiodic:
		*e = Aperiodic
	case sec1:
		*e = Sec1
	default:
		return fmt.Errorf("cannot unmarshal unknown enum2 value '%v'", s)
	}
	return nil
}
