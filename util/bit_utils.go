package util

import "fmt"

///// Bit manipulation \\\\\

// Extract a particular bit from an integer
// TODO - this only works with 32 bit ints - provide a more generic method for other types
func Nth_bit(n, data int) bool {
	mask := 1 << (31 - n) // Move a single 1 to the position of our bit
	bit := data & mask    // Isolate the bit
	return bit > 0        // Convert to bool
}

// Extract a section of bits from an integer
func Bit_slice(start, end, data int) []bool {
	shifted_bits := data >> (31 - end) // Shift the desired bits to the end
	mask := (1 << start) - 1           // Create a string of 1s the length of our section
	slice := shifted_bits & mask       // Isolate the desired bits

	fmt.Printf("data: %b\n", data)
	fmt.Printf("shifted: %b\n", shifted_bits)
	fmt.Printf("mask: %b\n", mask)
	fmt.Printf("slice: %b\n", slice)

	slice_length := (end - start) + 1
	fmt.Printf("slice_length: %d\n", slice_length)

	bit_array := make([]bool, slice_length)

	for i := slice_length - 1; i >= 0; i-- {
		next_bit := Nth_bit(31-i, slice)
		bit_array[slice_length-i-1] = next_bit
		//bit_array = append(bit_array, next_bit)
		fmt.Printf("bit #%d: %v\n", i, next_bit)
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
