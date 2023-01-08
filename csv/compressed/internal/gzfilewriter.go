package internal

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

type GzFileWriter struct {
	f  *os.File
	gw *gzip.Writer
	bw *bufio.Writer
}

func NewGzFileWriter(fileName string, append bool) (*GzFileWriter, error) {
	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	}

	if f, err := os.OpenFile(fileName, flag, 0666); err != nil {
		return nil, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	} else {
		if gw, err := gzip.NewWriterLevel(f, gzip.BestCompression); err != nil {
			f.Close()
			return nil, fmt.Errorf("cannot create gzip writer: %w", err)
		} else {
			// Gzip has a 32Kb window so smaller buffer makes no sense.
			bw := bufio.NewWriterSize(gw, 32768)
			return &GzFileWriter{
				f:  f,
				gw: gw,
				bw: bw,
			}, nil
		}
	}
}

func (w *GzFileWriter) Flush() error {
	if err := w.bw.Flush(); err != nil {
		w.gw.Flush()
		w.f.Sync()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.gw.Flush(); err != nil {
		w.f.Sync()
		return fmt.Errorf("cannot flush gzip writer: %w", err)
	}

	if err := w.f.Sync(); err != nil {
		return fmt.Errorf("cannot sync file: %w", err)
	}

	return nil
}

func (w *GzFileWriter) Close() error {
	if err := w.bw.Flush(); err != nil {
		w.gw.Close()
		w.f.Close()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.gw.Close(); err != nil {
		w.f.Close()
		return fmt.Errorf("cannot close gzip writer: %w", err)
	}

	if err := w.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (w *GzFileWriter) WriteString(s string) error {
	if _, err := w.bw.WriteString(s); err != nil {
		return fmt.Errorf("cannot write string: %w", err)
	}

	return nil
}

func (w *GzFileWriter) WriteBytes(b []byte) error {
	if _, err := w.bw.Write(b); err != nil {
		return fmt.Errorf("cannot write bytes: %w", err)
	}

	return nil
}
