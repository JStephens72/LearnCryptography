package main

import (
	"crypto/aes"
	"crypto/des"
	"errors"
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	k := make([]byte, keyLen)

	switch cipherType {
	case typeAES:
		blockCipher, err := aes.NewCipher(k)
		if err != nil {
			return 0, err
		}
		return blockCipher.BlockSize(), nil
	case typeDES:
		blockCipher, err := des.NewCipher(k)
		if err != nil {
			return 0, err
		}
		return blockCipher.BlockSize(), nil
	}

	return 0, errors.New("invalid cipher type")
}
