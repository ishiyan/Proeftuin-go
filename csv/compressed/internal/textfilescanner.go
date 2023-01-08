package internal

import (
	"bufio"
	"fmt"
	"os"
)

type TextFileScanner struct {
	f *os.File
	s *bufio.Scanner
}

func NewTextFileScanner(fileName string) (*TextFileScanner, error) {
	if f, err := os.Open(fileName); err != nil {
		return nil, fmt.Errorf("cannot open '%s' file: %w", fileName, err)
	} else {
		s := bufio.NewScanner(f)
		return &TextFileScanner{
			f: f,
			s: s,
		}, nil
	}
}

func (s *TextFileScanner) Close() error {
	if err := s.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (s *TextFileScanner) Scan() bool {
	return s.s.Scan()
}

func (s *TextFileScanner) Text() string {
	return s.s.Text()
}

func (s *TextFileScanner) Bytes() []byte {
	return s.s.Bytes()
}

func (s *TextFileScanner) Err() error {
	return s.s.Err()
}
