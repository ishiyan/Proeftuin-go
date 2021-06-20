package main

// dotGo 2017 - Bill Kennedy - Behavior Of Channels
// https://www.youtube.com/watch?v=zDCKZn4-dck
// starting version
// similate problem hitting Ctrl+C
// kill the program to exit

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	const grs = 10

	var d device
	var l log.Logger
	l.SetOutput(&d)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Println(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan
		d.problem = !d.problem
	}
}
