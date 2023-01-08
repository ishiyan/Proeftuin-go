package internal

import (
	"bufio"
	"compress/bzip2"
	"fmt"
	"io"
	"strings"

	"compressed/internal/mmap"
)

type Bz2FileScannerMmap struct {
	m  *mmap.Mmap
	br io.Reader
	s  *bufio.Scanner
}

func NewBz2FileScannerMmap(fileName string) (*Bz2FileScannerMmap, error) {
	const extBz2 = ".bz2"

	if !strings.HasSuffix(fileName, extBz2) {
		return nil, fmt.Errorf("input bzip2 file %q should have %q extension", fileName, extBz2)
	}

	m, err := mmap.OpenFile(fileName, mmap.ReadOnly)
	if err != nil {
		return nil, fmt.Errorf("cannot mmap %q file: %w", fileName, err)
	}

	br := bzip2.NewReader(m)

	s := bufio.NewScanner(br)
	return &Bz2FileScannerMmap{
		m:  m,
		br: br,
		s:  s,
	}, nil
}

// Data returns mapped memory.
func (m *Bz2FileScannerMmap) Data() []byte {
	return m.m.Data()
}

func (m *Bz2FileScannerMmap) Close() error {
	if err := m.m.Close(); err != nil {
		return fmt.Errorf("cannot close mmap: %w", err)
	}

	return nil
}

func (m *Bz2FileScannerMmap) Scan() bool {
	return m.s.Scan()
}

func (m *Bz2FileScannerMmap) Text() string {
	return m.s.Text()
}

func (m *Bz2FileScannerMmap) Bytes() []byte {
	return m.s.Bytes()
}

func (m *Bz2FileScannerMmap) Err() error {
	return m.s.Err()
}

func (s *Bz2FileScannerMmap) CopyTo(dst io.Writer) (int64, error) {
	return io.Copy(dst, s.br)
}
