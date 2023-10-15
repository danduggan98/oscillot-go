package util_test

import (
	"fmt"
	"testing"

	"github.com/danduggan98/oscillot-go/util"
)

func TestNthBit(t *testing.T) {
	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	firstBit := util.NthBit(0, data)
	secondBit := util.NthBit(7, data)
	thirdBit := util.NthBit(8, data)
	fourthBit := util.NthBit(15, data)

	if !firstBit || secondBit || !thirdBit || fourthBit {
		givenBits := fmt.Sprintf("%d %d %d %d", util.BitToInt(firstBit), util.BitToInt(secondBit), util.BitToInt(thirdBit), util.BitToInt(fourthBit))
		t.Fatalf("Bits should be 1 0 1 0, but got " + givenBits)
	}
}

func TestBitSlice(t *testing.T) {
	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	firstBit := util.BitSlice(0, 1, data)
	if len(firstBit) == 0 || !firstBit[0] {
		t.Fatalf("Bit should be one, but was %v", util.BitArrayToString(firstBit))
	}

	frontSlice := util.BitSlice(0, 5, data)
	for _, v := range frontSlice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + util.BitArrayToString(frontSlice))
		}
	}

	onesSlice := util.BitSlice(4, 3, data)
	for _, v := range onesSlice {
		if !v {
			t.Fatalf("Bits should be all ones, but found a zero: " + util.BitArrayToString(onesSlice))
		}
	}

	zerosSlice := util.BitSlice(13, 3, data)
	for _, v := range zerosSlice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + util.BitArrayToString(zerosSlice))
		}
	}

	endSlice := util.BitSlice(29, 3, data)
	for _, v := range endSlice {
		if v {
			t.Fatalf("Bits should be all zeros, but found a one: " + util.BitArrayToString(endSlice))
		}
	}
}

func TestEquals(t *testing.T) {
	a := 0xFEC8FEC8
	b := 0xFEC8FEC8
	aBits := util.ToBits(a)
	bBits := util.ToBits(b)

	if !util.Equals(aBits, bBits) {
		t.Fatalf("bits should be equal")
	}

	a >>= 1
	aBits = util.ToBits(a)

	if util.Equals(aBits, bBits) {
		t.Fatalf("bits should not be equal")
	}
}
