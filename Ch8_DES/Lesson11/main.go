package main

var sBoxTable = [][]byte{
	{0b00, 0b10, 0b01, 0b11},
	{0b10, 0b00, 0b11, 0b01},
	{0b01, 0b11, 0b00, 0b10},
	{0b11, 0b01, 0b10, 0b00},
}

func sBox(b byte) (byte, error) {
	lowNibble := b & 0x0F
	row := (lowNibble & 0b1100) >> 2
	column := (lowNibble & 0b0011)

	return sBoxTable[row][column], nil
}
