package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type TextFileWriter struct {
	f  *os.File
	bw *bufio.Writer
}

func NewTextFileWriter(fileName string, append bool) (*TextFileWriter, error) {
	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}

	if f, err := os.OpenFile(fileName, flag, 0666); err != nil {
		return nil, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	} else {
		bw := bufio.NewWriterSize(f, 32768)
		return &TextFileWriter{
			f:  f,
			bw: bw,
		}, nil
	}
}

func (w *TextFileWriter) Writer() io.Writer {
	return w.bw
}

func (w *TextFileWriter) Flush() error {
	if err := w.bw.Flush(); err != nil {
		w.f.Sync()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.f.Sync(); err != nil {
		return fmt.Errorf("cannot sync file: %w", err)
	}

	return nil
}

func (w *TextFileWriter) Close() error {
	if err := w.bw.Flush(); err != nil {
		w.f.Close()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (w *TextFileWriter) WriteString(s string) error {
	if _, err := w.bw.WriteString(s); err != nil {
		return fmt.Errorf("cannot write string: %w", err)
	}

	return nil
}

func (w *TextFileWriter) WriteBytes(b []byte) error {
	if _, err := w.bw.Write(b); err != nil {
		return fmt.Errorf("cannot write bytes: %w", err)
	}

	return nil
}
