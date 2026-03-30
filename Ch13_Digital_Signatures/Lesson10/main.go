package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(message))
	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`%v.%x`, message, sig), nil
}
