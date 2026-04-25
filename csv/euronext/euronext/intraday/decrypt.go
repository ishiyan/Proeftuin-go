package intraday

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"
)

const DefaultPassphrase = "24ayqVo7yJma"

// encryptedResponse represents the CryptoJS AES encrypted JSON format.
type encryptedResponse struct {
	CT string `json:"ct"` // ciphertext (base64)
	IV string `json:"iv"` // initialization vector (hex)
	S  string `json:"s"`  // salt (hex)
}

// IsEncryptedResponse checks whether the raw bytes look like a CryptoJS encrypted response.
func IsEncryptedResponse(data []byte) bool {
	var er encryptedResponse
	if err := json.Unmarshal(data, &er); err != nil {
		return false
	}
	return er.CT != "" && er.IV != "" && er.S != ""
}

// DecryptResponse decrypts a CryptoJS AES encrypted response using the given passphrase.
// It returns the decrypted plaintext bytes.
func DecryptResponse(data []byte, passphrase string) ([]byte, error) {
	var er encryptedResponse
	if err := json.Unmarshal(data, &er); err != nil {
		return nil, fmt.Errorf("cannot parse encrypted response: %w", err)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(er.CT)
	if err != nil {
		return nil, fmt.Errorf("cannot base64-decode ciphertext: %w", err)
	}

	salt, err := hex.DecodeString(er.S)
	if err != nil {
		return nil, fmt.Errorf("cannot hex-decode salt: %w", err)
	}

	// Derive key and IV using OpenSSL's EVP_BytesToKey (MD5-based KDF).
	key, iv := evpBytesToKey([]byte(passphrase), salt, 32, 16)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("cannot create AES cipher: %w", err)
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext length %d is not a multiple of block size %d", len(ciphertext), aes.BlockSize)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, fmt.Errorf("cannot remove PKCS7 padding: %w", err)
	}

	return plaintext, nil
}

// evpBytesToKey derives key and IV from passphrase and salt using the OpenSSL
// EVP_BytesToKey algorithm with MD5 (used by CryptoJS when given a string passphrase).
func evpBytesToKey(passphrase, salt []byte, keyLen, ivLen int) ([]byte, []byte) {
	totalLen := keyLen + ivLen
	var derived []byte
	var block []byte

	for len(derived) < totalLen {
		h := md5.New()
		if len(block) > 0 {
			h.Write(block)
		}
		h.Write(passphrase)
		h.Write(salt)
		block = h.Sum(nil)
		derived = append(derived, block...)
	}

	return derived[:keyLen], derived[keyLen : keyLen+ivLen]
}

// pkcs7Unpad removes PKCS#7 padding from plaintext.
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}
	padLen := int(data[len(data)-1])
	if padLen == 0 || padLen > blockSize || padLen > len(data) {
		return nil, fmt.Errorf("invalid padding length %d", padLen)
	}
	for i := len(data) - padLen; i < len(data); i++ {
		if data[i] != byte(padLen) {
			return nil, fmt.Errorf("invalid padding byte at position %d", i)
		}
	}
	return data[:len(data)-padLen], nil
}

const ajaxSecureURL = "https://live.euronext.com/modules/custom/ajax_secure/js/ajax-secure.js"

var kyeRegexp = regexp.MustCompile(`var\s+kye\s*=\s*'([^']+)'`)

// FetchPassphrase downloads the ajax-secure.js script from Euronext and extracts
// the default passphrase. It returns the extracted passphrase or an error.
func FetchPassphrase() (string, error) {
	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(ajaxSecureURL)
	if err != nil {
		return "", fmt.Errorf("cannot fetch ajax-secure.js: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ajax-secure.js: unexpected status %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read ajax-secure.js body: %w", err)
	}

	matches := kyeRegexp.FindAllStringSubmatch(string(body), -1)
	if len(matches) == 0 {
		return "", fmt.Errorf("cannot find passphrase (kye) in ajax-secure.js")
	}

	// The last match is the else-branch default value.
	return matches[len(matches)-1][1], nil
}
