package main

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	i := int(bits)
	if i >= len(base8Alphabet) || i < 0 {
		return ""
	}
	return string(base8Alphabet[i])
}
