# Concurrent Idioms: Broadcasting values in Go with linked channels

[source](https://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/)

Channels work very well if lots of writers are funneling values to a single reader, but it’s not immediately clear how multiple readers can all wait for values from a single writer.

Here’s what a Go API for doing this might look like.

```go
type Broadcaster ...

func NewBroadcaster() Broadcaster
func (b Broadcaster) Write(v interface{})
func (b Broadcaster) Listen() chan interface{}
```

The broadcast channel is created with NewBroadcaster, and values can be written to it with Write.
To listen on the channel, we call Listen, which gives us a new channel from which we can receive the values written.

This post is about an implementation where the writer never blocks; a slow reader with a fast writer can fill all of memory if it goes on for long enough. And it’s not particularly efficient.

Here’s the heart of it:

```go
type broadcast struct {
  c  chan broadcast;
  v  interface{};
}
```

This is what I call a “linked channel” (analagously to a linked list).
But even more than a linked list, it’s an [Ouroboros data structure](http://wadler.blogspot.com/2009/11/list-is-odd-creature.html).
That is, an instance of the structure can be sent down the channel that is inside itself.

Or the other way around. If I have a value of type chan broadcast around, then I can read a broadcast value b from it, giving me the arbitrary value b.v, and another value of the original type, b.c, allowing me to repeat the process.

The other part of the puzzle comes from the way that a buffered channel can used as a one-use one-to-many broadcast object. If I’ve got a buffered channel of some type T:

```go
var c = make(chan T, 1)
```

then any process reading from it will block until a value is written.
When we want to broadcast a value, we simply write it to the channel. This value will only go to a single reader, however we observe the convention that if you read a value from the channel, you always put it back immediately.

```go
func wait(c chan T) T {
  v := <-c
  c <- v;
  return v;
}
```

Putting the two pieces together, we can see that if the channel inside the broadcast struct is buffered in the above way, then we can get an endless stream of one-use one-to-many broadcast channels, each with an associated value.

Here’s the code:

```go
package broadcast

type broadcast struct {
  c  chan broadcast;
  v  interface{};
}

type Broadcaster struct {
  // private fields:
  Listenc  chan chan (chan broadcast);
  Sendc  chan<- interface{};
}

type Receiver struct {
  // private fields:
  C chan broadcast;
}

// create a new broadcaster object.
func NewBroadcaster() Broadcaster {
  listenc := make(chan (chan (chan broadcast)));
  sendc := make(chan interface{});
  go func() {
    currc := make(chan broadcast, 1);
    for {
      select {
      case v := <-sendc:
        if v == nil {
          currc <- broadcast{};
          return;
        }
        c := make(chan broadcast, 1);
        b := broadcast{c: c, v: v};
        currc <- b;
        currc = c;
      case r := <-listenc:
        r <- currc
      }
    }
  }();
  return Broadcaster{
    Listenc: listenc,
    Sendc: sendc,
  };
}

// start listening to the broadcasts.
func (b Broadcaster) Listen() Receiver {
  c := make(chan chan broadcast, 0);
  b.Listenc <- c;
  return Receiver{<-c};
}

// broadcast a value to all listeners.
func (b Broadcaster) Write(v interface{})  { b.Sendc <- v }

// read a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
  b := <-r.C;
  v := b.v;
  r.C <- b;
  r.C = b.c;
  return v;
}
```

This implementastion has the nice property that there’s no longer any need for a central registry of listeners. A Receiver value encapsulates a place in the stream of values and can be copied at will – each copy will receive an identical copy of the stream. There’s no need for an Unregister function either. Of course, if the readers don’t keep up with the writers, memory can be used indefinitely, but… isn’t this quite neat?

Here’s some example code using it:

```go
package main

import (
  "fmt";
  "broadcast";
  "time";
)

var b = broadcast.NewBroadcaster();

func listen(r broadcast.Receiver) {
  for v := r.Read(); v != nil; v = r.Read() {
    go listen(r);
    fmt.Println(v);
  }
}

func main() {
  r := b.Listen();
  go listen(r);
  for i := 0; i  < 10; i++ {
    b.Write(i);
  }
  b.Write(nil);

  time.Sleep(3 * 1e9);
}
```

## Discussion

### 1

to complicated: chan chan chan broadcast
These types make it simpler:

```go
type Msg struct {
inp chan Msg
msg string
}
type Broadcast struct {
inp chan Msg
waitForAll chan bool
}
type Receiver struct {
inp chan Msg
waitForAll chan Msg
}
```

### 2

Here’s the code in the Go playground, changed slightly for Go 1:

```go
// This package holds the code from
// http://rogpeppe.wordpress.com/2009/12/01/concurrent-idioms-1-broadcasting-values-in-go-with-linked-channels/
// updated to Go 1 standard. In particular, it's now OK to pass around
// by-value objects containing private fields, and we don't need to use
// semicolons.
package main

import (
  "fmt"
  "time"
)

var b = NewBroadcaster()

func main() {
  r := b.Listen()
  go listen(r)
  for i := 0; i < 10; i++ {
    b.Write(i)
  }
  b.Write(nil)

  time.Sleep(3 * 1e9)
}

func listen(r Receiver) {
  for v := r.Read(); v != nil; v = r.Read() {
    go listen(r)
    fmt.Println(v)
  }
}

type broadcast struct {
  c chan broadcast
  v interface{}
}

// Broadcaster allows
type Broadcaster struct {
  listenc chan chan (chan broadcast)
  sendc   chan<- interface{}
}

// Receiver can be used to wait for a broadcast value.
type Receiver struct {
  c chan broadcast
}

// NewBroadcaster returns a new broadcaster object.
func NewBroadcaster() Broadcaster {
  listenc := make(chan (chan (chan broadcast)))
  sendc := make(chan interface{})
  go func() {
    currc := make(chan broadcast, 1)
    for {
      select {
      case v := <-sendc:
        if v == nil {
          currc <- broadcast{}
          return
        }
        c := make(chan broadcast, 1)
        b := broadcast{c: c, v: v}
        currc <- b
        currc = c
      case r := <-listenc:
        r <- currc
      }
    }
  }()
  return Broadcaster{
    listenc: listenc,
    sendc:   sendc,
  }
}

// Listen starts returns a Receiver that
// listens to all broadcast values.
func (b Broadcaster) Listen() Receiver {
  c := make(chan chan broadcast, 0)
  b.listenc <- c
  return Receiver{<-c}
}

// Write broadcasts a a value to all listeners.
func (b Broadcaster) Write(v interface{}) {
  b.sendc <- v
}

// Read reads a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
  b := <-r.c
  v := b.v
  r.c <- b
  r.c = b.c
  return v
}
```

### 3

I thought it’s unnecessary to use `Broadcast.listenc` to listen to register process.
Add a field `Broadcast.cc` refers the newest broadcast, and then the method `Broadcast.Listen` could only just return `b.cc`

```go
type broadcast struct {
  c chan broadcast
  v interface{}
}

// Broadcaster allows
type Broadcaster struct {
  cc    chan broadcast
  sendc chan<- interface{}
}

// Receiver can be used to wait for a broadcast value.
type Receiver struct {
  c chan broadcast
}

// NewBroadcaster returns a new broadcaster object.
func NewBroadcaster() Broadcaster {
  cc := make(chan broadcast, 1)
  sendc := make(chan interface{})
  b := Broadcaster{
    sendc: sendc,
    cc:    cc,
  }

  go func() {
    for {
      select {
      case v := <-sendc:
        if v == nil {
          b.cc <- broadcast{}
          return
        }
        c := make(chan broadcast, 1)
        newb := broadcast{c: c, v: v}
        b.cc <- newb
        b.cc = c
      }
    }
  }()

  return b
}

// Listen starts returns a Receiver that
// listens to all broadcast values.
func (b Broadcaster) Listen() Receiver {
  return Receiver{b.cc}
}

// Write broadcasts a a value to all listeners.
func (b Broadcaster) Write(v interface{}) {
  b.sendc <- v
}

// Read reads a value that has been broadcast,
// waiting until one is available if necessary.
func (r *Receiver) Read() interface{} {
  b := <-r.c
  v := b.v
  r.c <- b
  r.c = b.c
  return v
}
```
