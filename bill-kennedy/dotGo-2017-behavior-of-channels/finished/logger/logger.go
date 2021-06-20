package logger

import (
	"fmt"
	"io"
	"sync"
)

// Logger implements our logger.
type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New is a Logger factory
func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	l.wg.Add(1)
	go func() {
		for v := range l.ch {
			fmt.Fprintf(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

// Stop closes the logger.
func (l *Logger) Stop() {
	close(l.ch)
	l.wg.Wait()
}

// Println prints the line.
func (l *Logger) Println(s string) {
	select {
	case l.ch <- s:
	default:
		fmt.Println("DROP")
	}
}
