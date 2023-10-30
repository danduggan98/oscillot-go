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

// Convert an array of bytes to an array of ints
func BytesToInt(data []byte) []uint32 {
	result := make([]uint32, 0)
	var next uint32

	for i := 0; i < len(data); i += 4 {
		next = 0

		for j := 0; j < 4 && i+j < len(data); j++ {
			next <<= 8
			next |= uint32(data[i+j])
		}
		result = append(result, next)
	}

	return result
}
