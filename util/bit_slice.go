package util

func Bool_to_int(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Extract a particular bit from an integer
func Nth_bit(n, data int) bool {
	return false
}

// Extract a section of bits from an integer
func Bit_slice(start, end, data int) []bool {

	slice_size := end - start + 1
	bit_array := make([]bool, slice_size)

	mask := ((1 << start) - 1) // Create a string of 1s
	mask <<= end               // Shift it to our starting position
	slice := data & mask       // Isolate the desired bits
	slice >>= start            // Move the slice back to the end

	return bit_array
}
