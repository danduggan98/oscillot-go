package mp3

import "github.com/danduggan98/oscillot-go/util"

type CRC struct {
	// Used for error correction if enabled by the header protection bit (0 | 16 bits)
	CorrectionBits uint32
}

func ParseCRC(bits uint32) *CRC { // TODO
	return &CRC{
		CorrectionBits: util.BitSlice(0, 16, bits),
	}
}
