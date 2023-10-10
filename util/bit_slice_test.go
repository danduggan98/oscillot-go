package util

import (
	"fmt"
	"testing"
)

func TestNthBit(t *testing.T) {

	data := 0xAAAA

	first_bit := Nth_bit(0, data)
	second_bit := Nth_bit(1, data)
	third_bit := Nth_bit(2, data)
	fourth_bit := Nth_bit(3, data)

	if first_bit != true ||
		second_bit != false ||
		third_bit != true ||
		fourth_bit != false {

		given_bits := fmt.Sprintf("%d%d%d%d", Bool_to_int(first_bit), Bool_to_int(second_bit), Bool_to_int(third_bit), Bool_to_int(fourth_bit))
		t.Fatalf("First four bits should be 1010, but got " + given_bits)
	}
}

func TestBitSlice(t *testing.T) {

}
