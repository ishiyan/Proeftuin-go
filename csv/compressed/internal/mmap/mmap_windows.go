//go:build windows
// +build windows

package mmap

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

// mmap on Windows is a two-step process.
// First, we call CreateFileMapping to get a handle.
// Then, we call MapviewToFile to get an actual pointer into memory.
// Because we want to emulate a POSIX-style mmap, we don't want to expose
// the handle -- only the pointer (a byte slice).
func mmap(fileName string, mode Mode) (*Mmap, error) {
	size, m, err := openFile(fileName, mode)
	if err != nil {
		return nil, fmt.Errorf("mmap: %w", err)
	}

	if size == 0 {
		return m, nil
	}

	prot := uint32(syscall.PAGE_READONLY)
	view := uint32(syscall.FILE_MAP_READ)
	if m.writable {
		prot = syscall.PAGE_READWRITE
		view = syscall.FILE_MAP_WRITE
	}

	// The maximum size is the area of the file, starting from 0,
	// that we wish to allow to be mappable.
	maxSizeLow, maxSizeHigh := uint32(size), uint32(size>>32)
	h, err := syscall.CreateFileMapping(m.hndl(), nil, prot, maxSizeHigh, maxSizeLow, nil)
	if err != nil {
		m.f.Close()
		return nil, fmt.Errorf("mmap: %w", os.NewSyscallError("CreateFileMapping", err))
	}

	// Actually map a view of the data into memory.
	ptr, err := syscall.MapViewOfFile(h, view, 0, 0, uintptr(size))
	if err != nil {
		syscall.CloseHandle(h)
		m.f.Close()
		return nil, fmt.Errorf("mmap: %w", os.NewSyscallError("MapViewOfFile", err))
	}

	const maxBytes = 1<<50 - 1

	//nolint:unsafeptr
	m.data = (*[maxBytes]byte)(unsafe.Pointer(ptr))[:size]

	syscall.CloseHandle(h)
	runtime.SetFinalizer(m, (*Mmap).unmap)
	return m, nil
}

func (m *Mmap) sync() error {
	if m.writable {
		if err := syscall.FlushViewOfFile(m.addr(), m.len()); err != nil {
			return fmt.Errorf("sync: couldn't flush view: %w", os.NewSyscallError("FlushViewOfFile", err))
		}

		if err := syscall.FlushFileBuffers(m.hndl()); err != nil {
			return fmt.Errorf("sync: couldn't flush buffers: %w", os.NewSyscallError("FlushFileBuffers", err))
		}
	}

	return nil
}

func (m *Mmap) lock() error {
	if err := syscall.VirtualLock(m.addr(), m.len()); err != nil {
		return fmt.Errorf("lock: %w", os.NewSyscallError("VirtualLock", err))
	}

	return nil
}

func (m *Mmap) unlock() error {
	if err := syscall.VirtualUnlock(m.addr(), m.len()); err != nil {
		return fmt.Errorf("unlock: %w", os.NewSyscallError("VirtualUnlock", err))
	}

	return nil
}

func (m *Mmap) unmap() error {
	if m.data == nil {
		return nil
	}
	defer m.f.Close()

	runtime.SetFinalizer(m, nil)
	addr := m.addr()
	m.data = nil

	err := syscall.UnmapViewOfFile(addr)
	if err != nil {
		return fmt.Errorf("unmap: %w", os.NewSyscallError("UnmapViewOfFile", err))
	}

	return nil
}

func (m *Mmap) hndl() syscall.Handle {
	return syscall.Handle(m.fdescr())
}
