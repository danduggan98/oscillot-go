package mp3

import "github.com/danduggan98/oscillot-go/util"

type CRC struct {
	// Used for error correction if enabled by the header protection bit (0 | 16 bits)
	CorrectionBits uint32
}

func ParseCRC(data []byte) *CRC {
	var bits uint32 = util.BytesToInts(data)[0]

	return &CRC{
		CorrectionBits: util.BitSlice(0, 16, bits),
	}
}
