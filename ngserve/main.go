package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	browserPtr := flag.Bool("browser", true, "open the web browser window on start")
	portPtr := flag.String("port", "3000", "port to serve")
	verbosePtr := flag.Bool("verbose", true, "do logging")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./dist")))
	if *browserPtr {
		go func() {
			if err := openURL("http://localhost:" + *portPtr); err != nil {
				fmt.Printf("Can't open web browser: %s\n", err)
			}
		}()
	}

	if *verbosePtr {
		fmt.Printf("Serving :" + *portPtr + " ...\n")
	}

	http.ListenAndServe(":"+*portPtr, nil)
}

// openURL opens the given url in the default browser.
func openURL(targetURL string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", targetURL)
	case "linux", "freebsd", "netbsd", "openbsd":
		cmd = exec.Command("xdg-open", targetURL)
	case "darwin":
		cmd = exec.Command("open", targetURL)
	default:
		return fmt.Errorf("unknown GOOS %s", runtime.GOOS)
	}

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	cmd.Stderr = buf

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%w: %s", err, buf.String())
	}

	return nil
}
