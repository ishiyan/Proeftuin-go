# Implementing Async/Await

[source](https://hackernoon.com/asyncawait-in-golang-an-introductory-guide-ol1e34sg)

```go
package async

import "context"

// Future interface has the method signature for await
type Future interface {
  Await() interface{}
}

type future struct {
  await func(ctx context.Context) interface{}
}

func (f future) Await() interface{} {
  return f.await(context.Background())
}

// Exec executes the async function
func Exec(f func() interface{}) Future {
  var result interface{}
  c := make(chan struct{})
  go func() {
    defer close(c)
    result = f()
  }()
  return future{
    await: func(ctx context.Context) interface{} {
      select {
      case <-ctx.Done():
        return ctx.Err()
      case <-c:
        return result
      }
    },
  }
}
```

We add a `Future` interface that has the `Await` method signature.
Next, we add a `future` struct that holds one value, a function signature of the `await` function.
Now futute struct implements `Future` interface's `Await` method by invoking its own await function.

Next in the `Exec` function, we execute the passed function asynchronously in goroutine.
And we return the `await` function.
It waits for the channel to close or context to read from.
Based on whichever happens first, it either returns the error or the result which is an interface.

Usage.

```go
func DoneAsync() int {
  fmt.Println("Warming up ...")
  time.Sleep(3 * time.Second)
  fmt.Println("Done ...")
  return 1
}

func main() {
  fmt.Println("Let's start ...")
  future := async.Exec(func() interface{} {
    return DoneAsync()
  })
  fmt.Println("Done is running ...")
  val := future.Await()
  fmt.Println(val)
}
```

It looks much cleaner, we are not explicitly working with goroutine or channels here.
Our `DoneAsync` function has been changed to a completely synchronous nature.
In the `main` function, we use the async package's `Exec` method to handle `DoneAsync`.
Which starts the execution of `DoneAsync`.
The control flow is returned back to `main` function which can execute other pieces of code.
Finally, we make blocking call to `Await` and read back data.
