package main

import (
	"crypto/sha256"
	"fmt"
)

func checksumMatches(message string, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message))
	return checksum == fmt.Sprintf("%x", h.Sum(nil))
}
