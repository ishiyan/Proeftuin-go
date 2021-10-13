package fitbit

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"golang.org/x/oauth2"
)

// getClient retrieves a token and generates OAuth2-enabled client.
// The 'token.json' file stores the user's access and refresh tokens, and is
// created automatically when the authorization flow completes for the first time.
func getClient(config *oauth2.Config) (*http.Client, error) {
	const (
		tokenFile = "token.json"
		stateLen  = 42
	)

	tok, err := tokenFromJSONFile(tokenFile)
	if err != nil {
		var code string

		state := randomString(stateLen)
		if code, err = codeFromWebBrowser(state, config.AuthCodeURL(state), config.RedirectURL); err != nil {
			return nil, fmt.Errorf("unable to get code from web browser: %w", err)
		}

		if tok, err = config.Exchange(context.TODO(), code); err != nil {
			return nil, fmt.Errorf("unable to exchange code for token: %w", err)
		}

		if err = tokenToJSONFile(tokenFile, tok); err != nil {
			return nil, fmt.Errorf("unable to save token to file: %w", err)
		}
	}

	return config.Client(context.Background(), tok), nil
}

// tokenFromFile retrieves a token from a local JSON file.
func tokenFromJSONFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)

	return tok, err
}

// tokenToJSONFile saves a token to a file path.
func tokenToJSONFile(path string, token *oauth2.Token) error {
	const perm = 0600

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("unable to create token file '%s': %w", path, err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(token); err != nil {
		return fmt.Errorf("unable to encode token: %w", err)
	}

	return nil
}

// codeFromWebBrowser opens a default browser window to authCodeURL for the user
// to authorize the application, and returns the resulting OAuth2 code.
// It rejects requests where the "state" param does not match expectedState.
func codeFromWebBrowser(expectedState, authCodeURL, redirectURL string) (string, error) {
	redirect, err := url.Parse(redirectURL)
	if err != nil {
		return "", err
	}

	ln, err := net.Listen("tcp", redirect.Host)
	if err != nil {
		return "", err
	}
	defer ln.Close()

	ch := make(chan string)
	chErr := make(chan error)

	go func() {
		handler := func(w http.ResponseWriter, r *http.Request) {
			state := r.FormValue("state")
			code := r.FormValue("code")

			path := redirect.Path
			if path == "" {
				path = "/"
			}

			if r.Method != "GET" || r.URL.Path != path || state == "" || code == "" {
				http.Error(w, errorMsg, http.StatusNotFound)
				chErr <- fmt.Errorf(
					"invalid OAuth2 redirect request: method '%s' ('GET'), path '%s' ('%s'), state '%s' ('%s'), code '%s'",
					r.Method, r.URL.Path, redirect.Path, state, expectedState, code)

				return
			}

			if state != expectedState {
				http.Error(w, "invalid state", http.StatusUnauthorized)
				chErr <- fmt.Errorf("invalid OAuth2 state; expected '%s' but got '%s'",
					expectedState, state)

				return
			}

			fmt.Fprint(w, successBody)
			ch <- code
		}

		// Must disable keep-alives, otherwise repeated calls to this method can block indefinitely.
		srv := http.Server{Handler: http.HandlerFunc(handler)}
		srv.SetKeepAlivesEnabled(false)
		_ = srv.Serve(ln)
	}()

	err = openURL(authCodeURL)
	if err != nil {
		fmt.Printf("Can't open web browser: %s\n", err)
		fmt.Printf("Go to the following link in your browser to complete an authorization:\n%v\n", authCodeURL)
		fmt.Printf("Type the received authorization code: ")

		var code string
		if _, err = fmt.Scan(&code); err != nil {
			return "", fmt.Errorf("unable to read authorization code: %w", err)
		}

		fmt.Printf("\nScanned the authorization code: %s\n", code)

		return code, nil
	}

	select {
	case code := <-ch:
		return code, nil
	case err := <-chErr:
		return "", err
	}
}

// openURL opens the given url in the default browser.
func openURL(targetURL string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// Escape characters not allowed by cmd.
		u := strings.ReplaceAll(targetURL, "&", `^&`)

		cmd = exec.Command("cmd", "/c", "start", u)
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

// randomString generates a random state string.
func randomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	biglen := big.NewInt(int64(len(letterBytes)))

	b := make([]byte, n)
	for i := range b {
		j, _ := rand.Int(rand.Reader, biglen)
		b[i] = letterBytes[j.Int64()]
	}

	return string(b)
}

const errorMsg = "This endpoint is for OAuth2 callbacks only. Please close this page and return to the application."

const successBody = `<!DOCTYPE html>
<html>
	<head>
		<title>OAuth2 Success</title>
		<meta charset="utf-8">
		<style>
			body { text-align: center; padding: 5%; font-family: sans-serif; }
			h1 { font-size: 20px; }
			p { font-size: 16px; color: #444; }
		</style>
	</head>
	<body>
		<h1>Code obtained, thank you!</h1>
		<p>
			You may now close this page and return to the application.
		</p>
	</body>
</html>
`
