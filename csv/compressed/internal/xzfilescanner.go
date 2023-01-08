package internal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ulikunitz/xz"
)

type XzFileScanner struct {
	f  *os.File
	xr *xz.Reader
	s  *bufio.Scanner
}

func NewXzFileScanner(fileName string) (*XzFileScanner, error) {
	const extXz = ".xz"

	if !strings.HasSuffix(fileName, extXz) {
		return nil, fmt.Errorf("input xz file %q should have %q extension", fileName, extXz)
	}

	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("cannot open %q file: %w", fileName, err)
	}

	conf := xz.ReaderConfig{SingleStream: true}
	if err = conf.Verify(); err != nil {
		f.Close()
		return nil, fmt.Errorf("invalid reader configuration: %w", err)
	}

	br := bufio.NewReaderSize(f, 32768)
	xr, err := conf.NewReader(br)
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("cannot create xz reader: %w", err)
	}

	s := bufio.NewScanner(xr)
	return &XzFileScanner{
		f:  f,
		xr: xr,
		s:  s,
	}, nil
}

func (s *XzFileScanner) Close() error {
	if err := s.f.Close(); err != nil {
		return fmt.Errorf("cannot close file: %w", err)
	}

	return nil
}

func (s *XzFileScanner) Scan() bool {
	return s.s.Scan()
}

func (s *XzFileScanner) Text() string {
	return s.s.Text()
}

func (s *XzFileScanner) Bytes() []byte {
	return s.s.Bytes()
}

func (s *XzFileScanner) Err() error {
	return s.s.Err()
}

func (s *XzFileScanner) CopyTo(dst io.Writer) (int64, error) {
	return io.Copy(dst, s.xr)
}
