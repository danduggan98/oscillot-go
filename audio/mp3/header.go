package mp3

import "github.com/danduggan98/oscillot-go/util"

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

// TODO - helper methods

// TODO - make this work
// IDEA - this could be a tree, similar to a huffman code
// 0 = left child, 1 = right
// traverse tree by reading bits left to right
// store values at equvalent places
// initial split at root is between versions
// nah jk it's a 3D array
var bitRateTable = [][]int{ // TODO
	{1, 2}, // 0000
}

// TODO
func (h Header) CalculateBitrate() int {
	// [bitrate bits][version][layer] = bitrate
	//////////// TODO - convert bits, layer and version to index
	// bits = 0-15, layer = 0-3, ver = 0-1

	return 0
	//return bitRateTable[][]
}

var sampleRateTable = [][]int{
	{44100, 22050, 11025},
	{48000, 24000, 12000},
	{32000, 16000, 8000},
	{0, 0, 0}, // reserved
}

// TODO
func (h Header) CalculateSampleRate() int {
	// [frequency][layer] = sample rate
	return 0
	//return sampleRateTable[][]
}
