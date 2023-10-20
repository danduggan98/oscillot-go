package mp3

import "github.com/danduggan98/oscillot-go/util"

// Bit rate -> MPEG version -> Layer
var bitRateTable = [][][]int{
	{{0, 0, 0}, {0, 0}},           // 0000 (free)
	{{32, 32, 32}, {32, 8}},       // 0001
	{{64, 48, 40}, {48, 16}},      // 0010
	{{96, 56, 48}, {56, 24}},      // 0011
	{{128, 64, 56}, {64, 32}},     // 0100
	{{160, 80, 64}, {80, 40}},     // 0101
	{{192, 96, 80}, {96, 48}},     // 0110
	{{224, 112, 96}, {112, 56}},   // 0111
	{{256, 128, 112}, {128, 64}},  // 1000
	{{288, 160, 128}, {144, 80}},  // 1001
	{{320, 192, 160}, {160, 96}},  // 1010
	{{352, 224, 192}, {176, 112}}, // 1011
	{{384, 256, 224}, {192, 128}}, // 1100
	{{416, 320, 256}, {224, 144}}, // 1101
	{{448, 384, 320}, {256, 160}}, // 1110
	{{0, 0, 0}, {0, 0}},           // 1111 (reserved)
}

// Sample rate -> MPEG version -> Frequency
var sampleRateTable = [][]int{
	{44100, 22050, 11025}, // 00
	{48000, 24000, 12000}, // 01
	{32000, 16000, 8000},  // 10
	{0, 0, 0},             // 11 (reserved)
}

var modeExtensionTable = [][]bool{
	{false, false},
	{true, false},
	{false, true},
	{true, true},
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
		layer_idx = 3 - h.Layer
	} else {
		if h.Layer < 0b11 {
			layer_idx = 1
		} else {
			layer_idx = 0
		}
	}

	return bitRateTable[h.BitRate][h.Version^1][layer_idx]
}

func (h Header) CalculateSampleRate() int {
	return sampleRateTable[h.Frequency][h.Version^1]
}

func (h Header) CalculateModeExtension() (bool, bool) {
	tableEntry := modeExtensionTable[h.ModeExtension]
	return tableEntry[0], tableEntry[1]
}

func (h Header) isStereo() bool {
	return h.Mode < 0b11
}
