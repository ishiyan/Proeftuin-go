package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"compressed/internal"
)

const (
	useScannerStrings = false
	extCsv            = ".csv"
	extBz2            = ".bz2"
)

func main() {
	mmapPtr := flag.String("mmap", "none", "use memory mapping: [none, scanner]")
	csvPtr := flag.String("csv", "overwrite", "what to do if csv file already exists: [overwrite, append, fail]")
	flag.Parse()

	fileName := flag.Arg(0)
	if fileName == "" {
		usage()
		fail("expecting input bzip2 file name as the positional argument")
		return
	}

	if !strings.HasSuffix(fileName, extBz2) {
		fail("input bzip2 file name should have '.bz2' extension")
		return
	}

	fmt.Printf("mmap=%s, csv=%s\n", *mmapPtr, *csvPtr)

	if *csvPtr == "fail" {
		if _, err := os.Stat(fileName + extCsv); !errors.Is(err, os.ErrNotExist) {
			fail("output csv file already exists")
			return
		}
	}

	append := *csvPtr == "append"
	start := time.Now()
	switch *mmapPtr {
	case "none":
		if err := bz2Scanner2Csv(fileName, append, useScannerStrings); err != nil {
			fail(err.Error())
			return
		}
	case "scanner":
		if err := bz2MmapScanner2Csv(fileName, append); err != nil {
			fail(err.Error())
			return
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("elapsed %s\n", elapsed)
}

func usage() {
	fmt.Println("usage:")
	fmt.Println("bz2csv {-mmap=[none, scanner]} {-csv=[overwrite, append, fail]} fileName")
	fmt.Println("-mmap     - whether to use memory mapping, possible values are")
	fmt.Println("            none    - use ordinary bzip2 scanner")
	fmt.Println("            scanner - use a bzip2 scanner on the memory mapped byte array")
	fmt.Println("-csv      - what to do if csv file already exists, possible values are")
	fmt.Println("            overwrite - overwrite existing file")
	fmt.Println("            append    - append to existing file")
	fmt.Println("            fail      - do nothing and exit")
	fmt.Println("-fileName - an input bzip2 file, file name should end with '.bz2'")
	fmt.Println("            an output csv file will have '.csv' extension appended")
	fmt.Println("            line terminations will be replaced with LF")
	fmt.Println("")
}

func fail(s string) {
	fmt.Println("panic: " + s)
}

func bz2Scanner2Csv(fileName string, append, useStrings bool) error {
	w, err := internal.NewTextFileWriter(fileName+extCsv, append)
	if err != nil {
		return fmt.Errorf("cannot create text file writer: %w", err)
	}
	defer w.Close()

	s, err := internal.NewBz2FileScanner(fileName)
	if err != nil {
		return fmt.Errorf("cannot create bz2 file scanner for the %q file: %w", fileName, err)
	}
	defer s.Close()

	newLine := []byte{'\n'}

	if useStrings {
		for s.Scan() {
			line := s.Text()
			if err := w.WriteString(line); err != nil {
				return fmt.Errorf("text file writer: %w", err)
			}

			if err := w.WriteBytes(newLine); err != nil {
				return fmt.Errorf("text file writer: %w", err)
			}
		}
	} else {
		for s.Scan() {
			bs := s.Bytes()
			if err := w.WriteBytes(bs); err != nil {
				return fmt.Errorf("text file writer: %w", err)
			}

			if err := w.WriteBytes(newLine); err != nil {
				return fmt.Errorf("text file writer: %w", err)
			}
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("cannot scan: %w", err)
	}

	return nil
}

func bz2MmapScanner2Csv(fileName string, append bool) error {
	w, err := internal.NewTextFileWriter(fileName+extCsv, append)
	if err != nil {
		return fmt.Errorf("cannot create text file writer: %w", err)
	}
	defer w.Close()

	s, err := internal.NewBz2FileScannerMmap(fileName)
	if err != nil {
		return fmt.Errorf("cannot create bz2 file scanner for the %q file: %w", fileName, err)
	}
	defer s.Close()

	newLine := []byte{'\n'}

	for s.Scan() {
		bs := s.Bytes()
		if err := w.WriteBytes(bs); err != nil {
			return fmt.Errorf("text file writer: %w", err)
		}

		if err := w.WriteBytes(newLine); err != nil {
			return fmt.Errorf("text file writer: %w", err)
		}
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("cannot scan: %w", err)
	}

	return nil
}