package mp3

import "github.com/danduggan98/oscillot-go/util"

// Bit rate -> MPEG version -> Layer
var bitRateTable = [][][]int{
	{{1, 2, 3}, {1, 2}}, // 0000
	{{1, 2, 3}, {1, 2}}, // 0001
	{{1, 2, 3}, {1, 2}}, // 0010
	{{1, 2, 3}, {1, 2}}, // 0011
	{{1, 2, 3}, {1, 2}}, // 0100
	{{1, 2, 3}, {1, 2}}, // 0101
	{{1, 2, 3}, {1, 2}}, // 0111
	{{1, 2, 3}, {1, 2}}, // 1000
	{{1, 2, 3}, {1, 2}}, // 1001
	{{1, 2, 3}, {1, 2}}, // 1010
	{{1, 2, 3}, {1, 2}}, // 1011
	{{1, 2, 3}, {1, 2}}, // 1100
	{{1, 2, 3}, {1, 2}}, // 1101
	{{1, 2, 3}, {1, 2}}, // 1110
	{{1, 2, 3}, {1, 2}}, // 1111
}

// Sample rate -> MPEG version -> Frequency
var sampleRateTable = [][]int{
	{44100, 22050, 11025},
	{48000, 24000, 12000},
	{32000, 16000, 8000},
	{0, 0, 0}, // reserved
}

// TODO - cross check this with the other sources
type Header struct {
	// Fixed value used as a searchable entry point to the stream (12 bits)
	SyncWord int

	// Specifies the MPEG version (1 bit)
	//  - 1 = MPEG-1
	//	- 0 = MPEG-2
	Version int

	// Specifies ..... something (2 bits)
	//  - 00 = Reserved
	//  - 01 = Layer 3
	//  - 10 = Layer 2
	//  - 11 = Layer 1
	Layer int

	// If set, CRC is used for correction (1 bit)
	ErrorProtection int

	// The bit rate used to encode the frame (4 bits)
	BitRate int

	// Sampling frequency (2 bits)
	Frequency int

	// Fills extra frame space for streams with 128 kbit/s and 44100 Hz (1 bit)
	PaddingBit int

	// Triggers application-specific behavior (1 bit)
	PrivacyBit int

	// Channel mode (2 bits)
	//  - 00 = Stereo
	//  - 01 = Joint stereo
	//  - 10 = Dual channel
	//  - 11 = Single channel
	Mode int

	// Specifies which methods to use in joint stereo mode (2 bits)
	//  - First bit = MS stereo on/off
	//  - Second bit = Intensity stereo on/off
	ModeExtension int

	// Specifies if the file is copyrighted (1 bit)
	Copyrighted int

	// Indicates something ... (1 bit)
	Original int

	// Indicates that the file requires equalization (2 bits)
	//  - 00 = None
	//  - 01 = 50/15 ms
	//  - 10 = Reserved
	//  - 11 = CCITT J.17
	Emphasis int
}

func ParseHeader(bits int) *Header {
	return &Header{
		SyncWord:        util.BitSlice(0, 12, bits),
		Version:         util.NthBit(12, bits),
		Layer:           util.BitSlice(13, 2, bits),
		ErrorProtection: util.NthBit(15, bits),
		BitRate:         util.BitSlice(16, 4, bits),
		Frequency:       util.BitSlice(20, 2, bits),
		PaddingBit:      util.NthBit(22, bits),
		PrivacyBit:      util.NthBit(23, bits),
		Mode:            util.BitSlice(24, 2, bits),
		ModeExtension:   util.BitSlice(26, 2, bits),
		Copyrighted:     util.NthBit(28, bits),
		Original:        util.NthBit(29, bits),
		Emphasis:        util.BitSlice(30, 2, bits),
	}
}

func (h Header) CalculateBitrate() int {
	layer_idx := 0

	if h.Version == 1 {
		layer_idx = 4 - h.Layer
	} else {
		if h.Layer < 0b11 {
			layer_idx = 1
		} else {
			layer_idx = 0
		}
	}

	return bitRateTable[h.BitRate][h.Version][layer_idx]
}

func (h Header) CalculateSampleRate() int {
	return sampleRateTable[h.Frequency][h.Version^1]
}
