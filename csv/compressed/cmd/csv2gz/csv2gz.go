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
	extGz             = ".gz"
)

func main() {
	mmapPtr := flag.String("mmap", "none", "use memory mapping: [none, direct, scanner]")
	gzPtr := flag.String("gz", "overwrite", "what to do if gzip file already exists: [overwrite, append, fail]")
	flag.Parse()

	fileName := flag.Arg(0)
	if fileName == "" {
		usage()
		fail("expecting input text file name as the positional argument")
		return
	}

	if strings.HasSuffix(fileName, extGz) {
		fail("input text file name shouldn't have '.gz' extension")
		return
	}

	fmt.Printf("mmap=%s, gz=%s\n", *mmapPtr, *gzPtr)

	if *gzPtr == "fail" {
		if _, err := os.Stat(fileName + extGz); !errors.Is(err, os.ErrNotExist) {
			fail("output gzip file already exists")
			return
		}
	}

	append := *gzPtr == "append"
	start := time.Now()
	switch *mmapPtr {
	case "none":
		if err := textScanner2Gzip(fileName, append, useScannerStrings); err != nil {
			fail(err.Error())
			return
		}
	case "scanner":
		if err := mmapScanner2Gzip(fileName, append); err != nil {
			fail(err.Error())
			return
		}
	case "direct":
		if err := mmap2Gzip(fileName, append); err != nil {
			fail(err.Error())
			return
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("csv2gz {-mmap=[none, direct, scanner]} {-gz=[overwrite, append, fail]} fileName")
	fmt.Println("-mmap     - whether to use memory mapping, possible values are")
	fmt.Println("            none    - use ordinary file scanner")
	fmt.Println("                      line terminations will be replaced with LF")
	fmt.Println("            direct  - pass the memory mapped byte array to gzip writer directly")
	fmt.Println("            scanner - use a scanner on the memory mapped byte array")
	fmt.Println("                      line terminations will be replaced with LF")
	fmt.Println("-gz       - what to do if gzip file already exists, possible values are")
	fmt.Println("            overwrite - overwrite existing file")
	fmt.Println("            append    - append to existing file")
	fmt.Println("            fail      - do nothing and exit")
	fmt.Println("-fileName - an input text file to compress")
	fmt.Println("            an output gzip file will have '.gz' extension appended")
	fmt.Println("")
}

func fail(s string) {
	fmt.Println("panic: " + s)
}

func mmap2Gzip(fileName string, append bool) error {
	m, err := mmap.OpenFile(fileName, mmap.ReadOnly)
	if err != nil {
		return fmt.Errorf("cannot mmap %q file: %w", fileName, err)
	}
	defer m.Close()

	w, err := internal.NewGzFileWriter(fileName+extGz, append)
	if err != nil {
		return fmt.Errorf("cannot create gz file writer: %w", err)
	}
	defer w.Close()

	if err := w.WriteBytes(m.Data()); err != nil {
		return fmt.Errorf("gz file writer: %w", err)
	}

	return nil
}

func textScanner2Gzip(fileName string, append, useStrings bool) error {
	s, err := internal.NewTextFileScanner(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file scanner for the %q file: %w", fileName, err)
	}
	defer s.Close()

	w, err := internal.NewGzFileWriter(fileName+extGz, append)
	if err != nil {
		return fmt.Errorf("cannot create gz file writer: %w", err)
	}
	defer w.Close()

	newLine := []byte{'\n'}

	if useStrings {
		for s.Scan() {
			line := s.Text()
			if err := w.WriteString(line); err != nil {
				return fmt.Errorf("gz file writer: %w", err)
			}

			if err := w.WriteBytes(newLine); err != nil {
				return fmt.Errorf("gz file writer: %w", err)
			}
		}
	} else {
		for s.Scan() {
			bs := s.Bytes()
			if err := w.WriteBytes(bs); err != nil {
				return fmt.Errorf("gz file writer: %w", err)
			}

			if err := w.WriteBytes(newLine); err != nil {
				return fmt.Errorf("gz file writer: %w", err)
			}
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("cannot scan: %w", err)
	}

	return nil
}

func mmapScanner2Gzip(fileName string, append bool) error {
	s, err := internal.NewTextFileScannerMmap(fileName)
	if err != nil {
		return fmt.Errorf("cannot create file scanner for the %q file: %w", fileName, err)
	}
	defer s.Close()

	w, err := internal.NewGzFileWriter(fileName+extGz, append)
	if err != nil {
		return fmt.Errorf("cannot create gz file writer: %w", err)
	}
	defer w.Close()

	newLine := []byte{'\n'}

	for s.Scan() {
		bs := s.Bytes()
		if err := w.WriteBytes(bs); err != nil {
			return fmt.Errorf("gz file writer: %w", err)
		}

		if err := w.WriteBytes(newLine); err != nil {
			return fmt.Errorf("gz file writer: %w", err)
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("cannot scan: %w", err)
	}

	return nil
}
