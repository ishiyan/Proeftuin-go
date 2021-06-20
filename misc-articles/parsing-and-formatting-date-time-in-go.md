# Parsing and formatting date/time in Go

Go takes an interesting approach to [parsing strings to time objects, and formatting time objects as strings](http://golang.org/pkg/time/).
Instead of using codes like most languages to represent component parts of a date/time string representation—like %Y for a 4-digit year like “2011” or %b for an abbreviated month name like “Feb”—Go uses a mnemonic device: there is a standard time, which is:

```go
Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
```

Or put another way:

```go
01/02 03:04:05PM '06 -0700
```

Instead of having to remember or lookup the traditional formatting codes for functions like strftime, you just count one-two-three-four and each place in the standard time corresponds to a component of a date/time object (the Time type in Go): one for day of the month, two for the month, three for the hour (in 12-hour time), four for the minutes, etc.

The way you put this into action is by putting together the parts of the standard time in a layout string that matches the format of either the string representation you want to parse into a Time object or the opposite direction, when you want to generate a string representation from an Time object.

## Parsing

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  // layout, value
  t, err := time.Parse("2006-01-02 15:04:05.000000", "2011-05-19 03:34:11.123456")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(t)

  // layout, value
  t, err = time.Parse("2006-01-02 15:04:05", "2011-05-19 08:34:11")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(t)

  // layout, value
  t, err = time.Parse("2006-01-02 15:04", "2011-05-19 08:16")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(t)

  // layout, value
  t, err = time.Parse("2006-01-02", "2011-05-19")
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(t)
}
```

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    value  := "Thu, 05/19/11, 10:47PM"
    // Writing down the way the standard time would look like formatted our way
    layout := "Mon, 01/02/06, 03:04PM"
    t, _ := time.Parse(layout, value)
    fmt.Println(t)
}

// => "Thu May 19 22:47:00 +0000 2011"
```

## Formatting

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.SecondsToLocalTime(1305861602)
    t.ZoneOffset = -4*60*60
    fmt.Println(t.Format("2006-01-02 15:04:05 -0700"))
}

// => "2011-05-20 03:20:02 -0400"
```
