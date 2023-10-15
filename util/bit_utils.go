package util

///// Bit manipulation \\\\\
// TODO - these only work with 32 bit ints - provide a more generic method for other types

// Extract a particular bit from an integer
func NthBit(n, data int) bool {
	if n > 31 {
		panic("index out of range")
	}

	mask := 1 << (31 - n) // Move a one onto the desired bit
	bit := data & mask    // Isolate the bit
	return bit != 0       // Convert to bool
}

// Extract a section of bits from an integer
func BitSlice(start, end, data int) []bool {
	if end < start {
		panic("start index exceeds end index")

	} else if start == end {
		return []bool{NthBit(start, data)}
	}

	sliceLen := (end - start) + 1
	slice := make([]bool, sliceLen)

	bits := data >> (31 - end)  // Shift the desired bits to the end
	mask := (1 << sliceLen) - 1 // Create a string of 1s the length of our section
	bits &= mask                // Isolate the desired bits

	for i := 0; i < sliceLen; i++ {
		nextBit := NthBit(32-sliceLen+i, bits)
		slice[i] = nextBit
	}
	return slice
}

func ToBits(data int) []bool {
	return BitSlice(0, 31, data)
}

func Equals(a, b []bool) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

///// Bit formatting \\\\\

func BitToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func BitToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func BitArrayToString(arr []bool) string {
	str := ""
	for _, v := range arr {
		str += BitToString(v)
	}
	return str
}
