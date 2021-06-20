# Use Go Channels as Promises and Async/Await

[source](https://levelup.gitconnected.com/use-go-channels-as-promises-and-async-await-ee62d93078ec)

Let’s experiment with a simple use case: `await` a result from an `async` function.

```js
const longRunningTask = async () => {
    // Simulate a workload.
    sleep(3000)
    return Math.floor(Math.random() * Math.floor(100))
}

const r = await longRunningTask()
console.log(r)
```

In Go it will be.

```go
package main

import (
  "fmt"
        "math/rand"
  "time"
)

func longRunningTask() <-chan int32 {
  r := make(chan int32)

  go func() {
    defer close(r)
    
    // Simulate a workload.
    time.Sleep(time.Second * 3)
    r <- rand.Int31n(100)
  }()

  return r
}

func main() {
  r := <-longRunningTask()
  fmt.Println(r)
}
```

## Promise.all()

It’s very common that we start multiple async tasks then wait for all of them to finish and gather their results.

```js
const longRunningTask = async () => {
    // Simulate a workload.
    sleep(3000)
    return Math.floor(Math.random() * Math.floor(100))
}

const [a, b, c] = await Promise.all(longRunningTask(), longRunningTask(), longRunningTask())
console.log(a, b, c)
```

```go
package main

import (
  "fmt"
        "math/rand"
  "time"
)

func longRunningTask() <-chan int32 {
  r := make(chan int32)

  go func() {
    defer close(r)
    
    // Simulate a workload.
    time.Sleep(time.Second * 3)
    r <- rand.Int31n(100)
  }()

  return r
}

func main() {
  aCh, bCh, cCh := longRunningTask(), longRunningTask(), longRunningTask()
  a, b, c := <-aCh, <-bCh, <-cCh
  
  fmt.Println(a, b, c)
}
```

Note we can not do `<-longRun(), <-longRun(), <-longRun()`, which will `longRun()` one by one instead all in once.

## Promise.race()

Sometimes, a piece of data can be received from several sources to avoid high latencies, or there’re cases that multiple results are generated but they’re equivalent and the only first response is consumed. This first-response-win pattern, therefore, is quite popular.

```js
const one = async () => {
    // Simulate a workload.
    sleep(Math.floor(Math.random() * Math.floor(2000)))
    return 1
}

const two = async () => {
    // Simulate a workload.
    sleep(Math.floor(Math.random() * Math.floor(1000)))
    sleep(Math.floor(Math.random() * Math.floor(1000)))
    return 2
}

const r = await Promise.race(one(), two())
console.log(r)
```

```go
package main

import (
  "fmt"
  "math/rand"
  "time"
)

func one() <-chan int32 {
  r := make(chan int32)

  go func() {
    defer close(r)

    // Simulate a workload.
    time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2000)))
    r <- 1
  }()

  return r
}

func two() <-chan int32 {
  r := make(chan int32)

  go func() {
    defer close(r)

    // Simulate a workload.
    time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
    time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
    r <- 2
  }()

  return r
}

func main() {
  var r int32
  select {
  case r = <-one():
  case r = <-two():
  }

  fmt.Println(r)
}
```

`select-case` is the pattern that Go designed specifically for racing channel operations.
We can even do more stuff within each case, but we’re focusing only on the result so we just leave them all empty.

## Promise.then() and Promise.catch()

Because Go’s error propagation model is very different from Javascript, there’s any clean way to replicate `Promise.then()` and `Promise.catch()`. In Go, error is returned along with return values instead of being thrown as exception. Therefore, if your function can fail, you can consider changing your return `<-chan ReturnType` into `<-chan ReturnAndErrorType`, which is a struct holding both the result and error.
