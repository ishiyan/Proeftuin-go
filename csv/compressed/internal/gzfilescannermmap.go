package internal

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"strings"

	"compressed/internal/mmap"
)

type GzFileScannerMmap struct {
	m  *mmap.Mmap
	gr *gzip.Reader
	s  *bufio.Scanner
}

func NewGzFileScannerMmap(fileName string) (*GzFileScannerMmap, error) {
	const extGz = ".gz"

	if !strings.HasSuffix(fileName, extGz) {
		return nil, fmt.Errorf("input gzip file %q should have %q extension", fileName, extGz)
	}

	m, err := mmap.OpenFile(fileName, mmap.ReadOnly)
	if err != nil {
		return nil, fmt.Errorf("cannot mmap %q file: %w", fileName, err)
	}

	gr, err := gzip.NewReader(m)
	if err != nil {
		m.Close()
		return nil, fmt.Errorf("cannot create gzip reader: %w", err)
	}

	s := bufio.NewScanner(gr)
	return &GzFileScannerMmap{
		m:  m,
		gr: gr,
		s:  s,
	}, nil
}

// Data returns mapped memory.
func (m *GzFileScannerMmap) Data() []byte {
	return m.m.Data()
}

func (m *GzFileScannerMmap) Close() error {
	if err := m.m.Close(); err != nil {
		return fmt.Errorf("cannot close mmap: %w", err)
	}

	return nil
}

func (m *GzFileScannerMmap) Scan() bool {
	return m.s.Scan()
}

func (m *GzFileScannerMmap) Text() string {
	return m.s.Text()
}

func (m *GzFileScannerMmap) Bytes() []byte {
	return m.s.Bytes()
}

func (m *GzFileScannerMmap) Err() error {
	return m.s.Err()
}

func (s *GzFileScannerMmap) CopyTo(dst io.Writer) (int64, error) {
	return io.Copy(dst, s.gr)
}
