package util

///// Bit manipulation \\\\\
// TODO - these only work with 32 bit ints - provide a more generic method for other types

// Extract a particular bit from an integer
func NthBit(n, data int) bool {
	mask := 1 << (31 - n) // Move a single 1 to the position of our bit
	bit := data & mask    // Isolate the bit
	return bit != 0       // Convert to bool
}

// Extract a section of bits from an integer
func BitSlice(start, end, data int) []bool {
	if end < start {
		panic("Invalid slice indices")

	} else if start == end {
		return []bool{NthBit(start, data)}
	}

	shifted_bits := data >> (31 - end) // Shift the desired bits to the end
	mask := (1 << start) - 1           // Create a string of 1s the length of our section
	slice := shifted_bits & mask       // Isolate the desired bits

	// TODO - THIS MASK BECOMES ALL ZEROS WHEN START IS 0
	slice_len := (end - start) + 1
	slice_idx := 32 - slice_len
	bit_array := make([]bool, slice_len)

	for i := 0; i < slice_len; i++ {
		next_bit := NthBit(slice_idx+i, slice)
		bit_array[i] = next_bit
	}
	return bit_array
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
