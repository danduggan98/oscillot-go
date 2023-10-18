package util

///// Bit manipulation \\\\\
// TODO - these only work with 32 bit ints - provide a more generic method for other types

// TODO - parse everything as a 32 bit int, forget about the bit nonsense

// Extract a particular bit from an integer
func NthBit(n, data int) int {
	if n > 31 {
		panic("index out of range")
	}

	mask := 1 << (31 - n) // Move a one onto the desired bit
	bit := data & mask    // Isolate the bit
	return bit >> n       // Convert to int
}

// Extract a section of bits from an integer
// TODO - back to original method:
// - shift to right
// - slice off front bits
// - return as int
func BitSlice(start, len, data int) int {
	end := start + len

	if start < 0 || end > 32 {
		panic("bit slice index out of range")
	} else if len == 1 {
		return NthBit(start, data)
	}

	bits := data >> (32 - end)
	mask := (1 << len) - 1
	return bits & mask
}
