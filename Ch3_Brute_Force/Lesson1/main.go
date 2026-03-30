package main

import (
	"math"
)

func alphabetSize(numBits int) float64 {
	return math.Pow(2, float64(numBits))
}
