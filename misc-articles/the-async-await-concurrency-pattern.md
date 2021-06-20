# The async/await concurrency pattern in Golang

[source](https://madeddu.xyz/posts/go-async-await/)

The simplest `async\await` example simulates a workload of 2 seconds and asynchronously waits for it to be completed. Also, since the run of a script from the shell is synchronous, you have to await for the execution of `myAsyncFunction` from inside an async context, otherwise the `Node.js` runtime will complaint.

```js
const sleep = require('util').promisify(setTimeout)
async function myAsyncFunction() {
    await sleep(2000)
    return 2
};

(async function() {
    const result = await myAsyncFunction();
    // outputs `2` after two seconds
    console.log(result);
})();
```

How can we achieve the same behavior with a Golang?

```go
package main

import (
  "fmt"
  "time"
)

func myAsyncFunction() <-chan int32 {
  r := make(chan int32, 1)
  fmt.Println("a")
  go func() {
    defer close(r)
    // func() core (meaning, the operation to be completed)
    fmt.Println("c")
    time.Sleep(time.Second * 2)
    fmt.Println("d")
    r <- 2
    fmt.Println("e")
  }()
  fmt.Println("b")
  return r
}

func main() {
  r := <-myAsyncFunction()
  // outputs `2` after two seconds
  fmt.Println(r)
}
```

the async function explicitly returns a `<-chan [yourType]` where `yourType` could be whatever you want. In this case, it's a simple `int32` number. Within the function you want to run asynchronously, create a channel by using the make(chan [your_type]) and return the created channel at the end of the function. Finally, start an anonymous goroutine by the `go myAsyncFunction() {...}` and implement the function's logic inside that anonymous function. Return the result by sending the value to the channel. At the beginning of the anonymous function, add ``defer close(r)` to close the channel once done.

To `await` behavior is implemented by simply read the value from channel, with `r := <-myAsyncFunction()`.

Another common scenario is when you start multiple async tasks then wait for all of them to finish and gather their results. We can achieve it is by using the `Promise.all()` primitive.

```js
const myAsyncFunction = (s) => {
    return new Promise((resolve) => {
        setTimeout(() => resolve(s), 2000);
    })
};

(async function() {
    const result = await Promise.all([
        myAsyncFunction(2),
        myAsyncFunction(3)
    ]);
    // outputs `[2, 3]` after three seconds
    console.log(result);
})();
```

The `await` this time is done across a list of `Promises`: the `.all()` signature takes an array as input. The `.all()` resolves all promises passed as an iterable object, short-circuits when an input value is rejected, is resolved successfully when all the promises in the array are resolved and rejected at first rejected of them.

The same behavior with a Golang.

```go
package main

import (
  "fmt"
  "time"
)

func myAsyncFunction(s int32) <-chan int32 {
  r := make(chan int32)
  fmt.Println("c")
  go func() {
    defer close(r)
    // func() core (meaning, the operation to be completed)
    fmt.Println("e")
    time.Sleep(time.Second * 2)
    fmt.Println("f")
    r <- s
    fmt.Println("g")
  }()
  fmt.Println("d")
  return r
}

func main() {
  fmt.Println("a")
  firstChannel, secondChannel := myAsyncFunction(2), myAsyncFunction(3)
  fmt.Println("b")
  first, second := <-firstChannel, <-secondChannel
  // outputs `2, 3` after three seconds
  fmt.Println(first, second)
}
```

In both snippets of code we just packaged a function taking as parameter the number of seconds to simulate a workload. The `await` is implemented using the channels receive operation, nothing more than the `<-` operator.

Sometimes, a piece of data can be received from several sources to avoid high latencies, or there're cases that multiple results are generated but they're equivalent and the only first response is consumed. This first-response-win pattern is quite popular.

```js
const myAsyncFunction = (s) => {
    return new Promise((resolve) => {
        setTimeout(() => resolve(s), 2000);
    })
};

(async function() {
    const result = await Promise.race([
        myAsyncFunction(2),
        myAsyncFunction(3)
    ]);
    // outputs `2` after three seconds
    console.log(result);
})();
```

The expected behavior is that `2` is always returned before the second `Promise` returned by `myAsyncFunction(3)` got resolved. This is natural due to the nature of `.race()` that implements the first-win pattern mentioned above.

In Golang, this can be obtained similarly by using the `select` statement.

```go
package main

import (
  "fmt"
  "time"
)

func myAsyncFunction(s int32) <-chan int32 {
  r := make(chan int32)
  fmt.Println("a")
  go func() {
    defer close(r)
    // func() core (meaning, the operation to be completed)
    fmt.Println("c")
    time.Sleep(time.Second * 2)
    fmt.Println("d")
    r <- s
    fmt.Println("e")
  }()
  fmt.Println("b")
  return r
}

func main() {
  var r int32
  select {
  case r = <-myAsyncFunction(2):
  case r = <-myAsyncFunction(3):
  }
  // outputs `2` after three seconds
  fmt.Println(r)
}
```

You can use Go's select statement to implement concurrency patterns and wait on multiple channel operations. In the snippet above, we use select to await both of the values simultaneously, choosing, in this case, the first one that arrives: once again, `2` is always returned before a value appear is retrieved from the channel populated by the `myAsyncFunction(3)`.

However, we've seen that basic sends and receives on channels are blocking. We can use select with a `default` clause to implement non-blocking sends, receives, and even non-blocking multi-way selects. Let's take the example exposed by the [gobyexample](https://gobyexample.com/non-blocking-channel-operations) site.

```go
package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  default:
    fmt.Println("no message received")
  }

  msg := "hi"
  select {
  case messages <- msg:
    fmt.Println("sent message", msg)
  default:
    fmt.Println("no message sent")
  }

  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  case sig := <-signals:
    fmt.Println("received signal", sig)
  default:
    fmt.Println("no activity")
  }
}
```

The code above implements a non-blocking receive. If a value is available on messages then select will take the `<-messages` case with that value. If not it will immediately take the default case. A non-blocking send works similarly. Here `msg` cannot be sent to the `messages` channel, because the channel has no buffer and there is no receiver. Therefore the `default` case is selected. We can use multiple cases above the `default` clause to implement a multi-way non-blocking select. Here we attempt non-blocking receives on both `messages` and `signals`.
