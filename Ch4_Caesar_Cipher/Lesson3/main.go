package main

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	final := ""
	for _, c := range text {
		offset := getOffsetChar(c, key)
		final += offset
	}
	return final
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	for i, curr := range alphabet {
		if curr == c {
			modI := (i + offset) % len(alphabet)
			if modI < 0 {
				modI = len(alphabet) + modI
			}
			return string(alphabet[modI])
		}
	}
	return ""
}
