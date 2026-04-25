// enxdecryptchk fetches the current Euronext decryption passphrase
// and compares it to the compiled-in default.
package main

import (
	"fmt"
	"os"

	"euronext/euronext/intraday"
)

func main() {
	fetched, err := intraday.FetchPassphrase()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("fetched:  %s\n", fetched)
	fmt.Printf("default:  %s\n", intraday.DefaultPassphrase)

	if fetched == intraday.DefaultPassphrase {
		fmt.Println("result:   SAME")
	} else {
		fmt.Println("result:   DIFFERENT — update needed!")
		os.Exit(2)
	}
}
