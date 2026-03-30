package main

import (
	"math"
)

func nonceStrength(nonce []byte) int {
	return int(math.Pow(2, float64(len(nonce)*8)))
}
