package main

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	for i, v := range masterKey {
		masterKey[i] = v ^ byte(roundNumber)
	}
	return masterKey
}
