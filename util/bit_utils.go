package util

// Extract a particular bit from an integer
func NthBit(n, data uint32) uint32 {
	if n > 31 {
		panic("index out of range")
	}

	bit := data >> (31 - n)
	return bit & 0b1
}

// Extract a section of bits from an integer
func BitSlice(start, len, data uint32) uint32 {
	end := start + len

	if end > 32 {
		panic("bit slice index out of range")
	} else if len == 1 {
		return NthBit(start, data)
	}

	bits := data >> (32 - end)
	var mask uint32 = (1 << len) - 1
	return bits & mask
}
