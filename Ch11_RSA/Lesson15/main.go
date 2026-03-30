package main

import (
	"math/big"
)

func decrypt(c, d, n *big.Int) *big.Int {
	m := new(big.Int)
	m.Exp(c, d, n)
	return m
}
