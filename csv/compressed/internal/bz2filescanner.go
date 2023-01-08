package internal

import (
	"bufio"
	"compress/bzip2"
	"fmt"
	"io"
	"os"
	"strings"
)

type Bz2FileScanner struct {
	f  *os.File
	br io.Reader
	s  *bufio.Scanner
}

func NewBz2FileScanner(fileName string) (*Bz2FileScanner, error) {
	const extBz2 = ".bz2"

	if !strings.HasSuffix(fileName, extBz2) {
		return nil, fmt.Errorf("input bzip2 file %q should have %q extension", fileName, extBz2)
	}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open %q file: %w", fileName, err)
	}

	r := bufio.NewReaderSize(f, 32768)
	br := bzip2.NewReader(r)

	s := bufio.NewScanner(br)
	return &Bz2FileScanner{
		f:  f,
		br: br,
		s:  s,
	}, nil
}

func (s *Bz2FileScanner) Close() error {
	if err := s.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (s *Bz2FileScanner) Scan() bool {
	return s.s.Scan()
}

func (s *Bz2FileScanner) Text() string {
	return s.s.Text()
}

func (s *Bz2FileScanner) Bytes() []byte {
	return s.s.Bytes()
}

func (s *Bz2FileScanner) Err() error {
	return s.s.Err()
}

func (s *Bz2FileScanner) CopyTo(dst io.Writer) (int64, error) {
	return io.Copy(dst, s.br)
}
