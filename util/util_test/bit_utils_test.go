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

	if firstBit == 0 || secondBit == 1 || thirdBit == 0 || fourthBit == 1 {
		givenBits := fmt.Sprintf("%d %d %d %d", firstBit, secondBit, thirdBit, fourthBit)
		t.Fatalf("Bits should be 1 0 1 0, but got " + givenBits)
	}
}

func TestBitSlice(t *testing.T) {
	data := 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	firstBit := util.BitSlice(0, 1, data)
	if firstBit == 0b0 {
		t.Fatalf("Bit should be 1, but was %b", firstBit)
	}

	firstSlice := util.BitSlice(0, 5, data)
	if firstSlice != 0b11111 {
		t.Fatalf("Slice should be 11111, but was %b", firstSlice)
	}

	secondSlice := util.BitSlice(4, 4, data)
	if secondSlice != 0b1110 {
		t.Fatalf("Slice should be 1110, but was %b", secondSlice)
	}

	thirdSlice := util.BitSlice(13, 6, data)
	if thirdSlice != 0b000111 {
		t.Fatalf("Slice should be 000111, but was %b", thirdSlice)
	}

	endSlice := util.BitSlice(29, 3, data)
	if endSlice != 0b000 {
		t.Fatalf("Slice should be 000, but was %b ", endSlice)
	}
}
