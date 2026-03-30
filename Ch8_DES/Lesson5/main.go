package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = padMsg(plaintext, des.BlockSize)

	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	if len(plaintext)%blockSize != 0 {
		indexOfLastBlock := (len(plaintext) / blockSize) * blockSize
		fullBlocks := plaintext[:indexOfLastBlock]
		lastBlockPadded := padWithZeros(plaintext[indexOfLastBlock:], blockSize)
		fullBlocks = append(fullBlocks, lastBlockPadded...)
		return fullBlocks
	}
	return plaintext
}
