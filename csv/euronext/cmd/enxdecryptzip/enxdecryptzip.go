// enxdecryptzip decrypts all encrypted JSON files inside a ZIP archive in-place.
//
// Usage:
//
//	enxdecryptzip [-p passphrase] archive.zip
package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"euronext/euronext/intraday"
)

func main() {
	passphrase := flag.String("p", intraday.DefaultPassphrase, "decryption passphrase")
	checkKey := flag.Bool("checkkey", false, "fetch the latest passphrase from Euronext before decrypting")
	flag.Parse()

	if *checkKey {
		fetched, err := intraday.FetchPassphrase()
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: cannot fetch passphrase: %v (continuing with current passphrase)\n", err)
		} else if fetched != *passphrase {
			fmt.Fprintf(os.Stderr, "passphrase changed: %s\n", fetched)
			*passphrase = fetched
		}
	}

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: enxdecryptzip [-p passphrase] archive.zip\n")
		os.Exit(1)
	}

	zipPath := flag.Arg(0)

	if err := processZip(zipPath, *passphrase); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func processZip(zipPath string, passphrase string) error {
	// Read the source ZIP.
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("cannot open zip '%s': %w", zipPath, err)
	}
	defer reader.Close()

	// Build the new ZIP in memory.
	var buf bytes.Buffer
	writer := zip.NewWriter(&buf)

	decrypted := 0
	skipped := 0

	for _, f := range reader.File {
		rc, err := f.Open()
		if err != nil {
			return fmt.Errorf("cannot open entry '%s': %w", f.Name, err)
		}

		data, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return fmt.Errorf("cannot read entry '%s': %w", f.Name, err)
		}

		// Decrypt JSON files that contain encrypted responses.
		if strings.EqualFold(filepath.Ext(f.Name), ".json") && intraday.IsEncryptedResponse(data) {
			plain, err := intraday.DecryptResponse(data, passphrase)
			if err != nil {
				return fmt.Errorf("cannot decrypt '%s': %w", f.Name, err)
			}
			data = plain
			decrypted++
		} else {
			skipped++
		}

		// Preserve the original file header (name, timestamps, etc.).
		header := f.FileHeader
		header.Method = zip.Deflate

		w, err := writer.CreateHeader(&header)
		if err != nil {
			return fmt.Errorf("cannot create entry '%s' in new zip: %w", f.Name, err)
		}
		if _, err := w.Write(data); err != nil {
			return fmt.Errorf("cannot write entry '%s' to new zip: %w", f.Name, err)
		}
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("cannot finalize new zip: %w", err)
	}

	// Write the new ZIP over the original file.
	if err := os.WriteFile(zipPath, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("cannot write zip file '%s': %w", zipPath, err)
	}

	log.Printf("done: %d decrypted, %d unchanged, written to %s", decrypted, skipped, zipPath)
	return nil
}
