# Unlimited Buffering with Low Overhead

[source](https://rogpeppe.wordpress.com/2010/02/10/unlimited-buffering-with-low-overhead/)

In Go, channels have a fixed length buffer.
Sometimes it is useful to add a buffer of unlimited length to a channel.
The first question is what the interface should look like.
I can think of three immediate possibilities (assume T is an arbitrary type – if Go had generics, this would be a generic function):

Given a channel, make sure that no writes to that channel will block, and return a channel from which the buffered values can be read:

```go
func Buffer(in <-chan T) <-chan T
```

Given a channel, return a channel that will buffer writes to that channel:

```go
func Buffer(out chan<- T) chan <-T
```

Given two channels, connect them via a buffering process:

```go
func Buffer(in <-chan T, out chan<- T)
```

Of these possibilities, on balance I think I prefer the second, as no operations will be performed on the original channel except when a value is written on the returned channel.

## Double-linked list implementation

Here is one simple, and relatively slow implementation. It uses the doubly-linked list implementation from the Go library. I timed it at 2076ns per item transferred on my machine. Note the code that runs before the select statement each time through the loop, which works out whether we want to be sending a value, and when it is time to finish. This relies on the fact that in a Go select statement, operations on nil channels are ignored.

```go
import "container/list"
func BufferList(out chan<- T) chan<- T {
  in := make(chan T, 100)
  go func() {
    var buf = list.New()
    for {
      outc := out
      var v T
      n := buf.Len()
      if n == 0 {
        // buffer empty: don't try to send on output
        if in == nil {
          close(out)
          return
        }
        outc = nil
      }else{
        v = buf.Front().Value.(T)
      }
      select {
      case e := <-in:
        if closed(in) {
          in = nil
        } else {
          buf.PushBack(e)
        }
      case outc <- v:
        buf.Remove(buf.Front())
      }
    }
  }()
  return in
}
```

The above implementation allocates a new linked list item for every value transferred.

## Circular buffer implementation

Here’s an alternative implementation that uses an array as a circular buffer, amortising allocations over time by doubling the size of the buffer when it overflows, and shrinking it when there is too much space.
Although the basic structure is similar, the code is more complex, and the time saving is modest – I timed it at 1729ns per item transferred, an improvement of 17%.
Removing the code to shrink the buffer does not make it significantly faster.

```go
func BufferRingOrig(out chan<- T) chan<- T {
  in := make(chan T, 100)
  go func() {
    var zero T
    var buf = make([]T, 10)
    var i = 0 // location of first value in buffer.
    var n = 0 // number of items in buffer.
    for {
      outc := out
      switch {
      case n == 0:
        // buffer empty: don't try to send on output
        if in == nil {
          close(out)
          return
        }
        outc = nil

      case n == len(buf):
        // buffer full: expand it
        b := make([]T, n*2)
        copy(b, buf[i:])
        copy(b[n-i:], buf[0:i])
        i = 0
        buf = b

      case len(buf) > 128 && n*3 < len(buf):
        // buffer too big, shrink it
        b := make([]T, len(buf) / 2)
        j := i + n
        if j > len(buf) {
          // wrap around
          k := j - len(buf)
          j = len(buf)
          copy(b, buf[i:j])
          copy(b[j - i:], buf[0:k])
        }else{
          // contiguous
          copy(b, buf[i:j])
        }
        i = 0
        buf = b
      }
      select {
      case e := <-in:
        if closed(in) {
          in = nil
        } else {
          j := i + n
          if j >= len(buf) {
            j -= len(buf)
          }
          buf[j] = e
          n++
        }
      case outc <- buf[i]:
        buf[i] = zero
        if i++; i == len(buf) {
          i = 0
        }
        n--
      }
    }
  }()
  return in
}
```

I wondered if the unnecessary tests before the select statement were making any significant difference to the time taken.
Although it makes it easy to preserve the invariants, there is no need to test whether the buffer is empty when a value has just been placed in it, for example.

Here is a version that only does the tests when necessary.
Interestingly, this change actually made the code run marginally slower (1704ns per item)

```go
func BufferRing(out chan<- T) chan<- T {
  in := make(chan T, 100)
  go func() {
    var zero T
    var buf = make([]T, 10)
    var i = 0 // location of first value in buffer.
    var n = 0 // number of items in buffer.
    var outc chan<- T
    for {
      select {
      case e := <-in:
        if closed(in) {
          in = nil
          if n == 0 {
            close(out)
            return
          }
        } else {
          j := i + n
          if j >= len(buf) {
            j -= len(buf)
          }
          buf[j] = e
          n++
          if n == len(buf) {
            // buffer full: expand it
            b := make([]T, n*2)
            copy(b, buf[i:])
            copy(b[n-i:], buf[0:i])
            i = 0
            buf = b
          }
          outc = out
        }
      case outc <- buf[i]:
        buf[i] = zero
        if i++; i == len(buf) {
          i = 0
        }
        n--
        if n == 0 {
          // buffer empty: don't try to send on output
          if in == nil {
            close(out)
            return
          }
          outc = nil
        }
        if len(buf) > 128 && n*3 < len(buf) {
          // buffer too big, shrink it
          b := make([]T, len(buf) / 2)
          j := i + n
          if j > len(buf) {
            // wrap around
            k := j - len(buf)
            j = len(buf)
            copy(b, buf[i:j])
            copy(b[j - i:], buf[0:k])
          }else{
            // contiguous
            copy(b, buf[i:j])
          }
          i = 0
          buf = b
        }
      }
    }
  }()
  return in
}
```

## Loop improvements

Although the speed improvement from the above piece of code was disappointing, the change paves the way for a change that really does make a difference.
A select statement in Go is significantly more costly than a regular channel operation.
In the code below, we loop receiving or sending values as long as we can do so without blocking.
Here’s a version of the list-based code that does this.
I measured it at 752ns per item, an improvement of 63% over the original, or 2.7x faster.

```go
func BufferListCont(out chan<- T) chan<- T {
  in := make(chan T, 100)
  go func() {
    var buf = list.New()
    var outc chan<- T
    var v T
    for {
      select {
      case e := <-in:
        if buf.Len() == 0 && !closed(in) {
          outc = out
          v = e
        }
        for {
          if closed(in) {
            in = nil
            if buf.Len() == 0 {
              close(out)
              return
            }
            break
          }
          buf.PushBack(e)
          var ok bool
          if e, ok = <-in; !ok {
            break
          }
        }
  
      case outc <- v:
        for {
          buf.Remove(buf.Front())
          if buf.Len() == 0 {
            // buffer empty: don't try to send on output
            if in == nil {
              close(out)
              return
            }
            outc = nil
            break
          }
          v = buf.Front().Value.(T)
          if ok := outc <- v; !ok {
            break
          }
        }
      }
    }
  }()
  return in
}
```

One objection to the above code is that in theory if there was a fast enough producer on another processor, the buffer process could spend forever feeding values into the buffer, without ever trying to write them out.
Although I believe that in practice the risk is negligible, it’s easy to guard against anyway, by only adding a fixed maximum number of values before returning to the select statement.

Here’s my final implementation, using the looping technique and with the guard added in.

I timed it at 427ns per item transferred, an improvement of 79% over the original version, or almost 5x faster.
Using a buffered channel directly is only 2.4x faster than this.

```go
func BufferRingContCheck(out chan<- T) chan<- T {
  in := make(chan T, 100)
  go func() {
    var zero T
    var buf = make([]T, 10)
    var i = 0 // location of first value in buffer.
    var n = 0 // number of items in buffer.
    var outc chan<- T
    for {
      select {
      case e := <-in:
        for added := 0; added < 1000; added++ {
          if closed(in) {
            in = nil
            if n == 0 {
              close(out)
              return
            }
            break
          }
          j := i + n
          if j >= len(buf) {
            j -= len(buf)
          }
          buf[j] = e
          n++
          outc = out    // enable output
          if n == len(buf) {
            // buffer full: expand it
            b := make([]T, n*2)
            copy(b, buf[i:])
            copy(b[n-i:], buf[0:i])
            i = 0
            buf = b
          }
          var ok bool
          if e, ok = <-in; !ok {
            break
          }
        }
      case outc <- buf[i]:
        for {
          buf[i] = zero
          if i++; i == len(buf) {
            i = 0
          }
          n--
          if n == 0 {
            // buffer empty: don't try to send on output
            if in == nil {
              close(out)
              return
            }
            outc = nil
            break
          }
          if len(buf) > 128 && n*3 < len(buf) {
            // buffer too big, shrink it
            b := make([]T, len(buf) / 2)
            j := i + n
            if j > len(buf) {
              // wrap around
              k := j - len(buf)
              j = len(buf)
              copy(b, buf[i:j])
              copy(b[j - i:], buf[0:k])
            }else{
              // contiguous
              copy(b, buf[i:j])
            }
            i = 0
            buf = b
          }
          if ok := outc <- buf[i]; !ok {
            break
          }
        }
      }
    }
  }()
  return in
}
```

Obviously the final code is significantly bigger and more complex than the original.
Which implementation should we choose?
Lacking generics, this code cannot usefully be put into a library, as most channels are not of type chan interface{}.

Given this, in most instances, perhaps the first version is to be preferred, as it’s smaller to cut and paste, and easier to understand.
In cases where performance is crucial, the final version can easily be substituted.

## Discussion

### 1

You could also use a circular linked list for the buffer, which is faster and use less memory than a doubly-linked list, and still quite simple.

For my prime sieve I settled on an array-based buffer that expands but never shrinks, throttling the read branch of the “select” statement.
Having a low threshold seems to be marginally faster, but using “added < 100" gives incorrect results, which is strange …

Obviously, the array buffer a lot less memory than all other approaches.

```go
import "container/ring"

// Use a goroutine to receive values from `out` and store them
// in an auto-expanding buffer, so that sending to `out` never blocks.
// Return a channel which serves as a sending proxy to to `out`.
func sendproxy(out chan<- int) chan<- int {
  in := make(chan int, 100)
  go func() {
    n := 1000 // the allocated length of the circular queue
    first := ring.New(n)
    last := first
    var c chan<- int
    var e int
    var ok bool
    for {
      select {
      case e = <-in:
        for added := 0; added < 1000; added++ {
          if closed(in) {
            in = nil
            break
          }
          last.Value = e
          if last.Next() == first {
            // buffer full: expand it
            last.Link(ring.New(n))
            n *= 2
          }
          last = last.Next()
          if e, ok = <-in; !ok {
            break
          }
        }
        c = out // enable output
        e = first.Value.(int)
      case c <- e:
        for {
          first = first.Next()
          if first == last {
            // buffer empty: disable output
            if in == nil {
              close(out)
              return
            }
            c = nil
            break
          }
          e = first.Value.(int)
          if ok = c <- e; !ok {
            break
          }
        }
      }
    }
  }()
  return in
}
```

### 2

I noticed a small thing with your input buffering.
Specifically, you add a value to the buffer, and then if it is full, you expand the buffer.

However, what if you changed it to

- (1) Accept the value
- (2) Check if the buffer is currently full, if so, expand it
- (3) insert the value
- (4) if the buffer is now full, break, allowing a chance for the consumer side to run
- (5) loop to (1)

That should decrease the number of times you need to resize if you have a fast producer.

Imagine both a producer and consumer that can handle channel updates as fast as the buffering code.
So if you loop on the producer you’ll never block, and if you loop on the consumer you’ll only stop when the buffer is empty.

With the current code, you’ll buffer 1000 units from the producer, and then yield them to the consumer.
That will cause it to size up to 1000, then size down to 128 on each pass.

If you instead ‘break’ at the point the buffer fills up, that would let the buffer stay constant sized.
Since the producer will fill it, the consumer will empty it.
And I think it would stay at the 128 entry level.

#### Answer

it’s an interesting point.
I think though that there’s no particular need to adjust the code – just some of the numbers pertaining to the buffer hysteresis.

The key is that when there’s a roughly constant flow between producer and consumer, the buffer should not need resizing.

The numbers “128” and “1000” are arbitrary.
If the minimum buffer size was changed from 128 to 1000, I don’t think you’d see any buffer resizing, as any fluctuation of up to 1000 values will be absorbed by the hysteresis.

You could try both ways and see how the performance is impacted, if you like.

### 3

Based on some of the ideas here (and elsewhere) I have implemented a package containing an infinitely-buffered channel among other types.
I ended up using go’s built-in slices and append function which is quite fast (though I haven’t benchmarked it against the version presented here).

[Code](https://github.com/eapache/channels)
[Queue](https://github.com/eapache/queue)
[Documentation](https://godoc.org/github.com/eapache/channels)
