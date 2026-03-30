package main

import (
	"math/big"
)

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	d := new(big.Int)
	d.ModInverse(e, tot)
	return d
}
