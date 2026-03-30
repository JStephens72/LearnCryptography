package main

import (
	"crypto/rand"
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
	p1 := new(big.Int).Sub(p, big.NewInt(1))
	q1 := new(big.Int).Sub(q, big.NewInt(1))
	return new(big.Int).Mul(p1, q1)
}

func getE(tot *big.Int) *big.Int {
	one := big.NewInt(1)
	two := big.NewInt(2)

	e := new(big.Int)

	for {
		e, _ = rand.Int(randReader, tot)

		if e.Cmp(two) <= 0 {
			continue
		}
		if e.Cmp(tot) >= 0 {
			continue
		}
		if gcd(e, tot).Cmp(one) == 0 {
			return e
		}
	}
}
