package mmap

import (
	"errors"
	"fmt"
	"io"
	"os"
	"unsafe"
)

// Mode specifies how a mmap file should be opened.
type Mode int

const (
	ReadOnly  Mode = Mode(os.O_RDONLY) // ReadOnly enables read-access to a mmap file.
	WriteOnly Mode = Mode(os.O_WRONLY) // WriteOnly enables write-access to a mmap file.
	ReadWrite Mode = Mode(os.O_RDWR)   // ReadWrite enables read-write-access to a mmap file.
)

// Mmap represents a file mapped into memory.
type Mmap struct {
	data     []byte
	c        int
	f        *os.File
	fi       os.FileInfo
	writable bool
	readable bool
}

var errNotWritable = errors.New("not writable")
var errNotReadable = errors.New("not readable")
var errNotMapped = errors.New("not mapped")
var errInvalidWhence = errors.New("invalid whence")
var errNegativePosition = errors.New("negative position")

// OpenFile memory-maps the named file for reading/writing, depending on
// the mode value: ReadOnly, WriteOnly, ReadWrite.
func OpenFile(fileName string, mode Mode) (*Mmap, error) {
	return mmap(fileName, mode)
}

// Len returns the length of the underlying memory-mapped file.
func (m *Mmap) Len() int {
	return len(m.data)
}

// Data returns mapped memory.
func (m *Mmap) Data() []byte {
	return m.data
}

// Flush synchronizes the mapping's contents to the file's contents on disk.
func (m *Mmap) Flush() error {
	return m.sync()
}

// Lock keeps the mapped region in physical memory.
func (m *Mmap) Lock() error {
	return m.lock()
}

// Unlock reverses the effect of Lock.
func (m *Mmap) Unlock() error {
	return m.unlock()
}

// Close implements the io.Closer interface.
func (m *Mmap) Close() error {
	return m.unmap()
}

// Read implements the io.Reader interface.
func (m *Mmap) Read(p []byte) (int, error) {
	if !m.readable {
		return 0, errNotReadable
	}

	if m.c >= len(m.data) {
		return 0, io.EOF
	}

	n := copy(p, m.data[m.c:])
	m.c += n
	return n, nil
}

// Write implements the io.Writer interface.
func (m *Mmap) Write(p []byte) (int, error) {
	if !m.writable {
		return 0, errNotWritable
	}

	if m.c >= len(m.data) {
		return 0, io.ErrShortWrite
	}

	n := copy(m.data[m.c:], p)
	m.c += n

	if len(p) > n {
		return n, io.ErrShortWrite
	}

	return n, nil
}

// ReadByte implements the io.ByteReader interface.
func (m *Mmap) ReadByte() (byte, error) {
	if !m.readable {
		return 0, errNotReadable
	}

	if m.c >= len(m.data) {
		return 0, io.EOF
	}

	v := m.data[m.c]
	m.c++
	return v, nil
}

// WriteByte implements the io.ByteWriter interface.
func (m *Mmap) WriteByte(c byte) error {
	if !m.writable {
		return errNotWritable
	}

	if m.c >= len(m.data) {
		return io.ErrShortWrite
	}

	m.data[m.c] = c
	m.c++

	return nil
}

// ReadAt implements the io.ReaderAt interface.
func (m *Mmap) ReadAt(p []byte, off int64) (int, error) {
	if !m.readable {
		return 0, errNotReadable
	}

	if m.data == nil {
		return 0, errNotMapped
	}

	if m.c >= len(m.data) {
		return 0, io.EOF
	}

	if off < 0 || int64(len(m.data)) < off {
		return 0, fmt.Errorf("invalid ReadAt offset %d", off)
	}

	n := copy(p, m.data[off:])
	if n < len(p) {
		return n, io.EOF
	}

	return n, nil
}

// WriteAt implements the io.WriterAt interface.
func (m *Mmap) WriteAt(p []byte, off int64) (int, error) {
	if !m.writable {
		return 0, errNotWritable
	}

	if m.data == nil {
		return 0, errNotMapped
	}

	if off < 0 || int64(len(m.data)) < off {
		return 0, fmt.Errorf("invalid WriteAt offset %d", off)
	}

	n := copy(m.data[off:], p)
	if n < len(p) {
		return n, io.ErrShortWrite
	}

	return n, nil
}

// Seek implements the io.Seeker interface.
func (m *Mmap) Seek(offset int64, whence int) (int64, error) {
	c := m.c
	switch whence {
	case io.SeekStart:
		c = int(offset)
	case io.SeekCurrent:
		c += int(offset)
	case io.SeekEnd:
		c = len(m.data) - int(offset)
	default:
		return 0, errInvalidWhence
	}

	if c < 0 {
		return 0, errNegativePosition
	}

	m.c = c
	return int64(c), nil
}

func openFile(fileName string, mode Mode) (int, *Mmap, error) {
	f, err := os.OpenFile(fileName, int(mode), 0666)
	if err != nil {
		return 0, nil, fmt.Errorf("couldn't open %q: %w", fileName, err)
	}

	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return 0, nil, fmt.Errorf("couldn't stat %q: %w", fileName, err)
	}

	writable := mode == WriteOnly || mode == ReadWrite
	readable := mode == ReadOnly || mode == ReadWrite
	size := fi.Size()
	if size == 0 {
		return 0, &Mmap{f: f, fi: fi, writable: writable, readable: readable}, nil
	}
	if size < 0 {
		f.Close()
		return 0, nil, fmt.Errorf("file %q has negative size", fileName)
	}
	if size != int64(int(size)) {
		f.Close()
		return 0, nil, fmt.Errorf("file %q is too large", fileName)
	}

	return int(size), &Mmap{f: f, fi: fi, writable: writable, readable: readable}, nil
}

func (m *Mmap) addr() uintptr {
	data := m.data
	return uintptr(unsafe.Pointer(&data[0]))
}

func (m *Mmap) len() uintptr {
	return uintptr(len(m.data))
}

func (m *Mmap) fdescr() uintptr {
	return m.f.Fd()
}
