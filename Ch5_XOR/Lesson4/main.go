package main

func crypt(plaintext, key []byte) []byte {
	final := []byte{}
	for i := range plaintext {
		final = append(final, plaintext[i]^key[i])
	}
	return final
}
