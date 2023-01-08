package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/larzconwell/bzip2"
)

type Bz2FileWriter struct {
	f  *os.File
	zw *bzip2.Writer
	bw *bufio.Writer
}

func NewBz2FileWriter(fileName string, append bool) (*Bz2FileWriter, error) {
	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	}

	if f, err := os.OpenFile(fileName, flag, 0666); err != nil {
		return nil, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	} else {
		if zw, err := bzip2.NewWriterLevel(f, bzip2.BestSpeed); err != nil {
			f.Close()
			return nil, fmt.Errorf("cannot create bzip2 writer: %w", err)
		} else {
			bw := bufio.NewWriterSize(zw, 32768)
			return &Bz2FileWriter{
				f:  f,
				zw: zw,
				bw: bw,
			}, nil
		}
	}
}

func (w *Bz2FileWriter) Flush() error {
	if err := w.bw.Flush(); err != nil {
		w.zw.Flush()
		w.f.Sync()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.zw.Flush(); err != nil {
		w.f.Sync()
		return fmt.Errorf("cannot flush bzip2 writer: %w", err)
	}

	if err := w.f.Sync(); err != nil {
		return fmt.Errorf("cannot sync file: %w", err)
	}

	return nil
}

func (w *Bz2FileWriter) Close() error {
	if err := w.bw.Flush(); err != nil {
		w.zw.Close()
		w.f.Close()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.zw.Close(); err != nil {
		w.f.Close()
		return fmt.Errorf("cannot close bzip2 writer: %w", err)
	}

	if err := w.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (w *Bz2FileWriter) WriteString(s string) error {
	if _, err := w.bw.WriteString(s); err != nil {
		return fmt.Errorf("cannot write string: %w", err)
	}

	return nil
}

func (w *Bz2FileWriter) WriteBytes(b []byte) error {
	if _, err := w.bw.Write(b); err != nil {
		return fmt.Errorf("cannot write bytes: %w", err)
	}

	return nil
}
