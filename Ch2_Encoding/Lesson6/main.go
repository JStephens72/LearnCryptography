package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	final := []byte{}
	parts := strings.Split(s, ":")
	for _, v := range parts {
		dec, err := hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
		final = append(final, dec...)
	}
	return final, nil
}
