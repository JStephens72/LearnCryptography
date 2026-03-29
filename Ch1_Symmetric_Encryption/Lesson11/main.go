package main

import (
	"fmt"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	randReader := rand.New(rand.NewSource(0))

	randomBytes := make([]byte, length)
	randReader.Read(randomBytes)
	return fmt.Sprintf("%x", randomBytes), nil
}
