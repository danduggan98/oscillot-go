package util_test

import (
	"fmt"
	"testing"

	"github.com/danduggan98/oscillot-go/util"
)

func TestNthBit(t *testing.T) {
	var data uint32 = 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

	firstBit := util.NthBit(0, data)
	secondBit := util.NthBit(7, data)
	thirdBit := util.NthBit(8, data)
	fourthBit := util.NthBit(15, data)

	if firstBit == 0 || secondBit == 1 || thirdBit == 0 || fourthBit == 1 {
		givenBits := fmt.Sprintf("%d %d %d %d", firstBit, secondBit, thirdBit, fourthBit)
		t.Fatalf("Bits should be 1 0 1 0, but got " + givenBits)
	}
}

func TestNthBitSingleByte(t *testing.T) {
	var data uint32 = 0xAA // 1010 1011

	firstBit := util.NthBit(0, data)
	secondBit := util.NthBit(1, data)
	thirdBit := util.NthBit(4, data)
	fourthBit := util.NthBit(5, data)

	if firstBit == 0 || secondBit == 1 || thirdBit == 0 || fourthBit == 1 {
		givenBits := fmt.Sprintf("%d %d %d %d", firstBit, secondBit, thirdBit, fourthBit)
		t.Fatalf("Bits should be 1 0 1 0, but got " + givenBits)
	}
}

func TestBitSlice(t *testing.T) {
	var data uint32 = 0xFEC8FEC8 // 1111 1110 1100 1000 1111 1110 1100 1000

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

func TestBytesToInt(t *testing.T) {
	data := []byte{0xFE, 0xC8}
	var result []uint32 = util.BytesToInt(data)

	if len(result) != 1 {
		t.Fatalf("Slice should have 1 int, but had %d ", len(result))
	}
	if result[0] != 0xFEC8 {
		t.Fatalf("Slice should be [65224], but was %v ", result)
	}

	data = []byte{0xFE, 0xC8, 0xBA, 0x04}
	result = util.BytesToInt(data)

	if len(result) != 1 {
		t.Fatalf("Slice should have 1 int, but had %d ", len(result))
	}
	if result[0] != 0xFEC8BA04 {
		t.Fatalf("Slice should be [4274567684], but was %v ", result)
	}

	data = []byte{0xFE, 0xC8, 0xBA, 0x04, 0xAA, 0xBB, 0xCC, 0xDD}
	result = util.BytesToInt(data)

	if len(result) != 2 {
		t.Fatalf("Slice should have 2 ints, but had %d ", len(result))
	}
	if result[0] != 0xFEC8BA04 || result[1] != 0xAABBCCDD {
		t.Fatalf("Slice should be [4274567684 2864434397], but was %v ", result)
	}

	data = []byte{0xFE, 0xC8, 0xBA, 0x04, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE}
	result = util.BytesToInt(data)

	if len(result) != 3 {
		t.Fatalf("Slice should have 3 ints, but had %d ", len(result))
	}
	if result[0] != 0xFEC8BA04 || result[1] != 0xAABBCCDD || result[2] != 0xEE {
		t.Fatalf("Slice should be [4274567684 2864434397 238], but was %v ", result)
	}
}
