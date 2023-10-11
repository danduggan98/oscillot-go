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
		given_bits := fmt.Sprintf("%d %d %d %d", Bool_to_int(first_bit), Bool_to_int(second_bit), Bool_to_int(third_bit), Bool_to_int(fourth_bit))
		t.Fatalf("Bits should be 1 0 1 0, but got " + given_bits)
	}
}

func TestBitSlice(t *testing.T) {

}
