package main

func padWithZeros(block []byte, desiredSize int) []byte {
	return append(block, make([]byte, desiredSize-len(block))...)
}
