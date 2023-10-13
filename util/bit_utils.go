package util

///// Bit manipulation \\\\\

// Extract a particular bit from an integer
// TODO - this only works with 32 bit ints - provide a more generic method for other types
func Nth_bit(n, data int) bool {
	mask := 1 << (31 - n) // Move a single 1 to the position of our bit
	bit := data & mask    // Isolate the bit
	return bit != 0       // Convert to bool
}

// Extract a section of bits from an integer
func Bit_slice(start, end, data int) []bool {
	shifted_bits := data >> (31 - end) // Shift the desired bits to the end
	mask := (1 << start) - 1           // Create a string of 1s the length of our section
	slice := shifted_bits & mask       // Isolate the desired bits

	slice_len := (end - start) + 1
	slice_idx := 32 - slice_len
	bit_array := make([]bool, slice_len)

	for i := 0; i < slice_len; i++ {
		next_bit := Nth_bit(slice_idx+i, slice)
		bit_array[i] = next_bit
	}
	return bit_array
}

///// Bit formatting \\\\\

func Bit_to_string(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func Bit_to_int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Bit_array_to_string(arr []bool) string {
	str := ""
	for _, v := range arr {
		str += Bit_to_string(v)
	}
	return str
}
