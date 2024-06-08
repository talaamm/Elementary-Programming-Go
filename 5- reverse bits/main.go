package main

// ReverseBits reverses the bits of a given byte.
func ReverseBits(oct byte) byte {
	var reversed byte = 0
	for i := 0; i < 8; i++ {
		// Shift reversed left by 1 to make space for the next bit
		reversed <<= 1
		// If the least significant bit of oct is 1, set the least significant bit of reversed to 1
		if oct&1 == 1 {
			reversed |= 1
		}
		// Shift oct right by 1 to process the next bit
		oct >>= 1
	}
	return reversed
}

///////////////	 OR	///////////////////

// ReverseBits reverses the bits of a given byte.
func ReverseBits1(oct byte) byte {
	var reversed byte = 0

	// Iterate over each bit in the byte
	for i := 0; i < 8; i++ {
		// Extract the i-th bit from the input byte
		bit := (oct >> i) & 1
		// Set the (7 - i)-th bit in the reversed byte
		reversed |= bit << (7 - i)
	}

	return reversed
}

///////////////	 OR	///////////////////

func ReverseBits3(oct byte) byte {
	var reversed byte
	for i := 0; i < 8; i++ {
		reversed = (reversed << 1) | (oct & 1)
		oct >>= 1
	}
	return reversed
}
