package internal

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ulikunitz/xz"
)

type XzFileWriter struct {
	f  *os.File
	xw *xz.Writer
	bw *bufio.Writer
}

func NewXzFileWriter(fileName string, append bool) (*XzFileWriter, error) {
	flag := os.O_WRONLY | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	}

	if f, err := os.OpenFile(fileName, flag, 0666); err != nil {
		return nil, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	} else {
		if xw, err := xz.NewWriter(f); err != nil {
			f.Close()
			return nil, fmt.Errorf("cannot create xz writer: %w", err)
		} else {
			bw := bufio.NewWriterSize(xw, 32768)
			return &XzFileWriter{
				f:  f,
				xw: xw,
				bw: bw,
			}, nil
		}
	}
}

func (w *XzFileWriter) Flush() error {
	if err := w.bw.Flush(); err != nil {
		w.f.Sync()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.f.Sync(); err != nil {
		return fmt.Errorf("cannot sync file: %w", err)
	}

	return nil
}

func (w *XzFileWriter) Close() error {
	if err := w.bw.Flush(); err != nil {
		w.xw.Close()
		w.f.Close()
		return fmt.Errorf("cannot flush bufio writer: %w", err)
	}

	if err := w.xw.Close(); err != nil {
		w.f.Close()
		return fmt.Errorf("cannot close xz writer: %w", err)
	}

	if err := w.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (w *XzFileWriter) WriteString(s string) error {
	if _, err := w.bw.WriteString(s); err != nil {
		return fmt.Errorf("cannot write string: %w", err)
	}

	return nil
}

func (w *XzFileWriter) WriteBytes(b []byte) error {
	if _, err := w.bw.Write(b); err != nil {
		return fmt.Errorf("cannot write bytes: %w", err)
	}

	return nil
}
