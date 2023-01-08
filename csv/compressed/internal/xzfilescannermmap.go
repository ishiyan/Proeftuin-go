package internal

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/ulikunitz/xz"

	"compressed/internal/mmap"
)

type XzFileScannerMmap struct {
	m  *mmap.Mmap
	xr *xz.Reader
	s  *bufio.Scanner
}

func NewXzFileScannerMmap(fileName string) (*XzFileScannerMmap, error) {
	const extXz = ".xz"

	if !strings.HasSuffix(fileName, extXz) {
		return nil, fmt.Errorf("input xz file %q should have %q extension", fileName, extXz)
	}

	m, err := mmap.OpenFile(fileName, mmap.ReadOnly)
	if err != nil {
		return nil, fmt.Errorf("cannot mmap %q file: %w", fileName, err)
	}

	conf := xz.ReaderConfig{SingleStream: true}
	if err = conf.Verify(); err != nil {
		m.Close()
		return nil, fmt.Errorf("invalid reader configuration: %w", err)
	}

	xr, err := conf.NewReader(m)
	if err != nil {
		m.Close()
		return nil, fmt.Errorf("cannot create xz reader: %w", err)
	}

	s := bufio.NewScanner(xr)
	return &XzFileScannerMmap{
		m:  m,
		xr: xr,
		s:  s,
	}, nil
}

// Data returns mapped memory.
func (m *XzFileScannerMmap) Data() []byte {
	return m.m.Data()
}

func (m *XzFileScannerMmap) Close() error {
	if err := m.m.Close(); err != nil {
		return fmt.Errorf("cannot close mmap: %w", err)
	}

	return nil
}

func (m *XzFileScannerMmap) Scan() bool {
	return m.s.Scan()
}

func (m *XzFileScannerMmap) Text() string {
	return m.s.Text()
}

func (m *XzFileScannerMmap) Bytes() []byte {
	return m.s.Bytes()
}

func (m *XzFileScannerMmap) Err() error {
	return m.s.Err()
}

func (s *XzFileScannerMmap) CopyTo(dst io.Writer) (int64, error) {
	return io.Copy(dst, s.xr)
}
