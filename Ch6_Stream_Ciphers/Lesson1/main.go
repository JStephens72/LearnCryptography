package main

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	for {
		textChar, textOk := <-textCh
		if !textOk {
			return
		}
		keyChar, keyOk := <-keyCh
		if !keyOk {
			return
		}
		result <- textChar ^ keyChar
	}
}
