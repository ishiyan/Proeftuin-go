package internal

import (
	"bufio"
	"fmt"

	"compressed/internal/mmap"
)

type TextFileScannerMmap struct {
	m *mmap.Mmap
	s *bufio.Scanner
}

func NewTextFileScannerMmap(fileName string) (*TextFileScannerMmap, error) {
	if m, err := mmap.OpenFile(fileName, mmap.ReadOnly); err != nil {
		return nil, fmt.Errorf("cannot mmap %q file: %w", fileName, err)
	} else {
		s := bufio.NewScanner(m)
		return &TextFileScannerMmap{
			m: m,
			s: s,
		}, nil
	}
}

// Data returns mapped memory.
func (m *TextFileScannerMmap) Data() []byte {
	return m.m.Data()
}

func (m *TextFileScannerMmap) Close() error {
	if err := m.m.Close(); err != nil {
		return fmt.Errorf("cannot close mmap: %w", err)
	}

	return nil
}

func (m *TextFileScannerMmap) Scan() bool {
	return m.s.Scan()
}

func (m *TextFileScannerMmap) Text() string {
	return m.s.Text()
}

func (m *TextFileScannerMmap) Bytes() []byte {
	return m.s.Bytes()
}

func (m *TextFileScannerMmap) Err() error {
	return m.s.Err()
}
