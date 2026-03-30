package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	timeCost    uint32 = 3
	memoryCost  uint32 = 32 * 1024
	parallelism uint8  = 4
	keyLen             = 32
	saltLen            = 16
)

func hashPassword(password string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("generate salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, timeCost, memoryCost, parallelism, keyLen)

	b64 := base64.RawStdEncoding
	phc := "$argon2id$v=19" +
		"$m=" + strconv.FormatUint(uint64(memoryCost), 10) +
		",t=" + strconv.FormatUint(uint64(timeCost), 10) +
		",p=" + strconv.FormatUint(uint64(parallelism), 10) +
		"$" + b64.EncodeToString(salt) +
		"$" + b64.EncodeToString(hash)

	return phc, nil
}

func checkPasswordHash(password, hash string) bool {
	parts := strings.Split(hash, "$")
	if len(parts) != 6 || parts[1] != "argon2id" || parts[2] != "v=19" {
		return false
	}

	var (
		memoryU32 uint32
		timeU32   uint32
		paraU8    uint8
	)

	paramParts := strings.Split(parts[3], ",")
	if len(paramParts) != 3 {
		return false
	}
	for _, kv := range paramParts {
		k, v, ok := strings.Cut(kv, "=")
		if !ok {
			return false
		}
		switch k {
		case "m":
			val, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return false
			}
			memoryU32 = uint32(val)
		case "t":
			val, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return false
			}
			timeU32 = uint32(val)
		case "p":
			val, err := strconv.ParseUint(v, 10, 8)
			if err != nil || val == 0 {
				return false
			}
			paraU8 = uint8(val)
		default:
			return false
		}
	}

	b64 := base64.RawStdEncoding
	salt, err := b64.DecodeString(parts[4])
	if err != nil {
		return false
	}
	want, err := b64.DecodeString(parts[5])
	if err != nil {
		return false
	}

	got := argon2.IDKey([]byte(password), salt, timeU32, memoryU32, paraU8, uint32(len(want)))

	return subtle.ConstantTimeCompare(got, want) == 1
}
