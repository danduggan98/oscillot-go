package util

import (
	"fmt"
	"testing"
)

func TestNthBit(t *testing.T) {

	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	firstBit, err1 := NthBit(0, data)
	secondBit, err2 := NthBit(7, data)
	thirdBit, err3 := NthBit(8, data)
	fourthBit, err4 := NthBit(15, data)

	err := FirstError(err1, err2, err3, err4)
	if err != nil {
		t.Fatalf("NthBit threw error: %s", err.Error())
	}

	if !firstBit || secondBit || !thirdBit || fourthBit {
		givenBits := fmt.Sprintf("%d %d %d %d", BitToInt(firstBit), BitToInt(secondBit), BitToInt(thirdBit), BitToInt(fourthBit))
		t.Fatalf("Bits should be 1 0 1 0, but got " + givenBits)
	}
}

func TestBitSlice(t *testing.T) {

	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	firstBit, _ := BitSlice(0, 0, data)
	if len(firstBit) == 0 || !firstBit[0] {
		t.Fatalf("Bit should be one, but was empty or zero")
	}

	frontSlice, _ := BitSlice(0, 4, data)
	for _, v := range frontSlice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + BitArrayToString(frontSlice))
		}
	}

	onesSlice, _ := BitSlice(4, 6, data)
	for _, v := range onesSlice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + BitArrayToString(onesSlice))
		}
	}

	zerosSlice, _ := BitSlice(13, 15, data)
	for _, v := range zerosSlice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + BitArrayToString(zerosSlice))
		}
	}

	endSlice, _ := BitSlice(29, 31, data)
	for _, v := range endSlice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + BitArrayToString(endSlice))
		}
	}
}

func TestEquals(t *testing.T) {

	a := 0xFEC8FEC8
	b := 0xFEC8FEC8
	aBits, _ := ToBits(a)
	bBits, _ := ToBits(b)

	if !Equals(aBits, bBits) {
		t.Fatalf("bits should be equal")
	}

	a >>= 1
	aBits, _ = ToBits(a)

	if Equals(aBits, bBits) {
		t.Fatalf("bits should not be equal")
	}
}
