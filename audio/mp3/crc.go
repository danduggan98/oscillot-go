package mp3

import "github.com/danduggan98/oscillot-go/util"

type CRC struct {
	// Used for error correction if enabled by the header protection bit (0 | 16 bits)
	CorrectionBits []bool
}

func ParseCRC(bits int) CRC { // TODO - extract from uint16
	return CRC{
		CorrectionBits: util.BitSlice(0, 16, bits),
	}
}
