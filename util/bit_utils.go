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
func BitSlice(start, len, data int) []bool {
	if start < 0 || start+len > 32 {
		panic("bit slice index out of range")
	} else if len == 1 {
		return []bool{NthBit(start, data)}
	}

	bits := data << start

	slice := make([]bool, len)
	for i := 0; i < len; i++ {
		slice[i] = NthBit(i, bits)
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
