package util

import (
	"fmt"
	"testing"
)

func TestNthBit(t *testing.T) {

	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	first_bit := NthBit(0, data)
	second_bit := NthBit(7, data)
	third_bit := NthBit(8, data)
	fourth_bit := NthBit(15, data)

	if !first_bit || second_bit || !third_bit || fourth_bit {
		given_bits := fmt.Sprintf("%d %d %d %d", BitToInt(first_bit), BitToInt(second_bit), BitToInt(third_bit), BitToInt(fourth_bit))
		t.Fatalf("Bits should be 1 0 1 0, but got " + given_bits)
	}
}

func TestBitSlice(t *testing.T) {

	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	if !BitSlice(0, 0, data)[0] {
		t.Fatalf("Bits should be one, but was zero")
	}

	front_slice := BitSlice(0, 4, data)
	for _, v := range front_slice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + BitArrayToString(front_slice))
		}
	}

	ones_slice := BitSlice(4, 6, data)
	for _, v := range ones_slice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + BitArrayToString(ones_slice))
		}
	}

	zeros_slice := BitSlice(13, 15, data)
	for _, v := range zeros_slice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + BitArrayToString(zeros_slice))
		}
	}

	end_slice := BitSlice(29, 31, data)
	for _, v := range end_slice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + BitArrayToString(end_slice))
		}
	}
}

func TestEquals(t *testing.T) {

	a := 0xFEC8FEC8
	b := 0xFEC8FEC8

	//fmt.Println(BitArrayToString(ToBits(a)))
	//fmt.Println(BitArrayToString(ToBits(b)))

	if !Equals(ToBits(a), ToBits(b)) {
		t.Fatalf("bits should be equal")
	}

	a >>= 1

	//fmt.Println(BitArrayToString(ToBits(a)))
	//fmt.Println(BitArrayToString(ToBits(b)))

	if Equals(ToBits(a), ToBits(b)) {
		t.Fatalf("bits should not be equal")
	}
}
