# Process large csv file and limit goroutines

[source](https://stackoverflow.com/questions/56325466/process-large-csv-file-and-limit-goroutines)

Read a csv file (~1M row).
Each row contain a HTTP link to an image which we need to download.
Check [this](https://godoc.org/golang.org/x/sync/errgroup#Group.Go).

```go
package main

import (
    "context"
    "encoding/csv"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "os/signal"
    "sync"
    "time"
)

func worker(ctx context.Context, dst chan string, src chan []string) {
    for {
        select {
        case url, ok := <-src: // you must check for readable state of the channel.
            if !ok {
                return
            }
            dst <- fmt.Sprintf("out of %v", url) // do somethingg useful.
        case <-ctx.Done(): // if the context is cancelled, quit.
            return
        }
    }
}

func main() {

    // create a context
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    // that cancels at ctrl+C
    go onSignal(os.Interrupt, cancel)

    // parse command line arguments
    var filename string
    var numberOfWorkers int
    flag.StringVar(&filename, "filename", "", "src file")
    flag.IntVar(&numberOfWorkers, "c", 2, "concurrent workers")
    flag.Parse()

    // check arguments
    if filename == "" {
        log.Fatal("filename required")
    }

    start := time.Now()

    csvfile, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer csvfile.Close()

    reader := csv.NewReader(csvfile)

    // create the pair of input/output channels for the controller=>workers com.
    src := make(chan []string)
    out := make(chan string)

    // use a waitgroup to manage synchronization
    var wg sync.WaitGroup

    // declare the workers
    for i := 0; i < numberOfWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            worker(ctx, out, src)
        }()
    }

    // read the csv and write it to src
    go func() {
        for {
            record, err := reader.Read()
            if err == io.EOF {
                break
            } else if err != nil {
                log.Fatal(err)
            }
            src <- record // you might select on ctx.Done().
        }
        close(src) // close src to signal workers that no more job are incoming.
    }()

    // wait for worker group to finish and close out
    go func() {
        wg.Wait() // wait for writers to quit.
        close(out) // when you close(out) it breaks the below loop.
    }()

    // drain the output
    for res := range out {
        fmt.Println(res)
    }

    fmt.Printf("\n%2fs", time.Since(start).Seconds())
}

func onSignal(s os.Signal, h func()) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, s)
    <-c
    h()
}
```
