package main

import (
	"fmt"
	"strings"
)

func getHexString(b []byte) string {
	result := []string{}
	for _, v := range b {
		result = append(result, fmt.Sprintf("%02x", v))
	}
	return strings.Join(result, ":")
}

func getBinaryString(b []byte) string {
	result := []string{}
	for _, v := range b {
		result = append(result, fmt.Sprintf("%08b", v)) // Binary representation, padded to 8 bits
	}
	return strings.Join(result, ":") // Join with colon delimiter
}
