package main

func feistel(msg []byte, roundKeys [][]byte) []byte {
	lhs := msg[:len(msg)/2]
	rhs := msg[len(msg)/2:]
	for _, roundKey := range roundKeys {
		nextLeft := rhs
		hashedData := hash(rhs, roundKey, len(roundKey))
		rhs = xor(lhs, hashedData)
		lhs = nextLeft
	}
	return append(rhs, lhs...)
}
