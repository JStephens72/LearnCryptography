package main

import (
	"fmt"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	for i := 0; i < int(math.Pow(2, 24)); i++ {
		potentialKey := intToBytes(i)
		decryptedMessage := crypt(encrypted, potentialKey)
		if string(decryptedMessage) == decrypted {
			return potentialKey, nil
		}
	}
	return nil, fmt.Errorf("key not found")
}
