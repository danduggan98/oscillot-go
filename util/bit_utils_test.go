package util

import (
	"fmt"
	"testing"
)

func TestNthBit(t *testing.T) {

	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	first_bit := Nth_bit(0, data)
	second_bit := Nth_bit(7, data)
	third_bit := Nth_bit(8, data)
	fourth_bit := Nth_bit(15, data)

	if !first_bit || second_bit || !third_bit || fourth_bit {
		given_bits := fmt.Sprintf("%d %d %d %d", Bit_to_int(first_bit), Bit_to_int(second_bit), Bit_to_int(third_bit), Bit_to_int(fourth_bit))
		t.Fatalf("Bits should be 1 0 1 0, but got " + given_bits)
	}
}

func TestBitSlice(t *testing.T) {

	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	ones_slice := Bit_slice(4, 6, data)
	for _, v := range ones_slice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + Bit_array_to_string(ones_slice))
		}
	}

	zeros_slice := Bit_slice(13, 15, data)
	for _, v := range zeros_slice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + Bit_array_to_string(zeros_slice))
		}
	}
}
