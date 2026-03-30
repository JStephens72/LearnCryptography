package main

import (
	"crypto/sha256"
	"fmt"
)

func hmac(message, key string) string {
	key1 := key[:len(key)/2]
	key2 := key[len(key)/2:]

	inner := sha256.New()
	inner.Write([]byte(key2 + message))

	outer := sha256.New()
	combined := append([]byte(key1), inner.Sum(nil)...)
	outer.Write(combined)

	return fmt.Sprintf("%x", outer.Sum(nil))
}
