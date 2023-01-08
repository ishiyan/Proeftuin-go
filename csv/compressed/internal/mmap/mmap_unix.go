//go:build darwin || dragonfly || freebsd || linux || openbsd || solaris || netbsd
// +build darwin dragonfly freebsd linux openbsd solaris netbsd

package mmap

import (
	"os"
	"syscall"
	"unsafe"
)

func mmap(fileName string, mode Mode) (*Mmap, error) {
	size, m, err := openFile(fileName, mode)
	if err != nil {
		return nil, fmt.Errorf("mmap: %w", err)
	}

	if size == 0 {
		return m, nil
	}

	prot := syscall.PROT_READ
	if m.writable {
		prot |= syscall.PROT_WRITE
	}

	data, err := syscall.Mmap(int(m.fdescr()), 0, int(size), prot, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap: %w", os.NewSyscallError("Mmap", err))
	}

	m.data = data

	runtime.SetFinalizer(m, (*Mmap).unmap)
	return m, nil
}

func (m *Mmap) sync() error {
	if m.writable {
		if _, _, err := syscall.Syscall(syscall.SYS_MSYNC, m.addr(), m.len(), syscall.MS_SYNC); err != nil {
			return fmt.Errorf("sync: %w", os.NewSyscallError("SYS_MSYNC", err))
		}
	}

	return nil
}

func (m *Mmap) lock() error {
	if m.data == nil {
		if err := syscall.Mlock(m.data); err != nil {
			return fmt.Errorf("lock: %w", os.NewSyscallError("Mlock", err))
		}
	}

	return nil
}

func (m *Mmap) unlock() error {
	if m.data != nil {
		if err := syscall.Munlock(m.data); err != nil {
			return fmt.Errorf("unlock: %w", os.NewSyscallError("Munlock", err))
		}
	}

	return nil
}

func (m *Mmap) unmap() error {
	if m.data == nil {
		return nil
	}
	defer m.f.Close()

	runtime.SetFinalizer(m, nil)
	data := m.data
	m.data = nil
	if err := syscall.Munmap(data); err != nil {
		return fmt.Errorf("unmap: %w", os.NewSyscallError("Munmap", err))
	}

	return nil
}
