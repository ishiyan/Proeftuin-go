package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"compressed/internal"
	"compressed/internal/mmap"
)

const (
	useScannerStrings = false
	extBz2            = ".bz2"
)

func main() {
	mmapPtr := flag.String("mmap", "none", "use memory mapping: [none, direct, scanner]")
	bz2Ptr := flag.String("bz2", "overwrite", "what to do if bzip2 file already exists: [overwrite, append, fail]")
	flag.Parse()

	fileName := flag.Arg(0)
	if fileName == "" {
		usage()
		fail("expecting input text file name as the positional argument")
		return
	}

	if strings.HasSuffix(fileName, extBz2) {
		fail("input text file name shouldn't have '.bz2' extension")
		return
	}

	fmt.Printf("mmap=%s, bz2=%s\n", *mmapPtr, *bz2Ptr)

	if *bz2Ptr == "fail" {
		if _, err := os.Stat(fileName + extBz2); !errors.Is(err, os.ErrNotExist) {
			fail("output bzip2 file already exists")
			return
		}
	}

	append := *bz2Ptr == "append"
	start := time.Now()
	switch *mmapPtr {
	case "none":
		if err := textScanner2Bz2(fileName, append, useScannerStrings); err != nil {
			fail(err.Error())
			return
		}
	case "scanner":
		if err := mmapScanner2Bz2(fileName, append); err != nil {
			fail(err.Error())
			return
		}
	case "direct":
		if err := mmap2Bz2(fileName, append); err != nil {
			fail(err.Error())
			return
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("csv2bz {-mmap=[none, direct, scanner]} {-bz2=[overwrite, append, fail]} fileName")
	fmt.Println("-mmap     - whether to use memory mapping, possible values are")
	fmt.Println("            none    - use ordinary file scanner")
	fmt.Println("                      line terminations will be replaced with LF")
	fmt.Println("            direct  - pass the memory mapped byte array to bzip2 writer directly")
	fmt.Println("                      line terminations will not be changed")
	fmt.Println("            scanner - use a scanner on the memory mapped byte array")
	fmt.Println("                      line terminations will be replaced with LF")
	fmt.Println("-bz2      - what to do if bzip2 file already exists, possible values are")
	fmt.Println("            overwrite - overwrite existing file")
	fmt.Println("            append    - append to existing file")
	fmt.Println("            fail      - do nothing and exit")
	fmt.Println("-fileName - an input text file to compress")
	fmt.Println("            an output bzip2 file will have '.bz2' extension appended")
	fmt.Println("")
}

func fail(s string) {
	fmt.Println("panic: " + s)
}

func mmap2Bz2(fileName string, append bool) error {
	m, err := mmap.OpenFile(fileName, mmap.ReadOnly)
	if err != nil {
		return fmt.Errorf("cannot mmap %q file: %w", fileName, err)
	}
	defer m.Close()

	w, err := internal.NewBz2FileWriter(fileName+extBz2, append)
	if err != nil {
		return fmt.Errorf("cannot create bz2 file writer: %w", err)
	}
	defer w.Close()

	if err := w.WriteBytes(m.Data()); err != nil {
		return fmt.Errorf("bz2 file writer: %w", err)
	}

	return nil
}

func textScanner2Bz2(fileName string, append, useStrings bool) error {
	s, err := internal.NewTextFileScanner(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file scanner for the %q file: %w", fileName, err)
	}
	defer s.Close()

	w, err := internal.NewBz2FileWriter(fileName+extBz2, append)
	if err != nil {
		return fmt.Errorf("cannot create bzip2 file writer: %w", err)
	}
	defer w.Close()

	newLine := []byte{'\n'}

	if useStrings {
		for s.Scan() {
			line := s.Text()
			if err := w.WriteString(line); err != nil {
				return fmt.Errorf("bz2 file writer: %w", err)
			}

			if err := w.WriteBytes(newLine); err != nil {
				return fmt.Errorf("bz2 file writer: %w", err)
			}
		}
	} else {
		for s.Scan() {
			bs := s.Bytes()
			if err := w.WriteBytes(bs); err != nil {
				return fmt.Errorf("bz2 file writer: %w", err)
			}

			if err := w.WriteBytes(newLine); err != nil {
				return fmt.Errorf("bz2 file writer: %w", err)
			}
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("cannot scan: %w", err)
	}

	return nil
}

func mmapScanner2Bz2(fileName string, append bool) error {
	s, err := internal.NewTextFileScannerMmap(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file scanner for the %q file: %w", fileName, err)
	}
	defer s.Close()

	w, err := internal.NewBz2FileWriter(fileName+extBz2, append)
	if err != nil {
		return fmt.Errorf("cannot create bz2 file writer: %w", err)
	}
	defer w.Close()

	newLine := []byte{'\n'}

	for s.Scan() {
		bs := s.Bytes()
		if err := w.WriteBytes(bs); err != nil {
			return fmt.Errorf("bz2 file writer: %w", err)
		}

		if err := w.WriteBytes(newLine); err != nil {
			return fmt.Errorf("bz2 file writer: %w", err)
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("cannot scan: %w", err)
	}

	return nil
}
