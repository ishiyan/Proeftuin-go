package internal

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

type GzFileScanner struct {
	f  *os.File
	gr *gzip.Reader
	s  *bufio.Scanner
}

func NewGzFileScanner(fileName string) (*GzFileScanner, error) {
	const extGz = ".gz"

	if !strings.HasSuffix(fileName, extGz) {
		return nil, fmt.Errorf("input gzip file %q should have %q extension", fileName, extGz)
	}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open %q file: %w", fileName, err)
	}

	br := bufio.NewReaderSize(f, 32768)
	gr, err := gzip.NewReader(br)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("cannot create gzip reader: %w", err)
	}

	s := bufio.NewScanner(gr)
	return &GzFileScanner{
		f:  f,
		gr: gr,
		s:  s,
	}, nil
}

func (s *GzFileScanner) Close() error {
	if err := s.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (s *GzFileScanner) Scan() bool {
	return s.s.Scan()
}

func (s *GzFileScanner) Text() string {
	return s.s.Text()
}

func (s *GzFileScanner) Bytes() []byte {
	return s.s.Bytes()
}

func (s *GzFileScanner) Err() error {
	return s.s.Err()
}

func (s *GzFileScanner) CopyTo(dst io.Writer) (int64, error) {
	return io.Copy(dst, s.gr)
}
