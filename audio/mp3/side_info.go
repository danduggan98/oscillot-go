package mp3

type SideInfoGranule struct {
	// Number of bits used for scalefactors and encoded data
	// - Mono = 12 bits
	// - Stereo = 24 bits
	Part23Length int

	// Total size of big_values partition
	// - Mono = 9 bits
	// - Stereo = 18 bits
	BigValues int

	// Quantization step size
	// - Mono = 8 bits
	// - Stereo = 16 bits
	GlobalGain int

	// Number of bits used for transmission of scalefactors
	// - Mono = 4 bits
	// - Stereo = 8 bits
	ScalefacCompress int

	// Indicates that a different window is used
	// - Mono = 1 bit
	// - Stereo = 2 bits
	WindowSwitchingFlag int

	// Indicates the type of window used
	// - Mono = 2 bits
	// - Stereo = 4 bits
	BlockType int

	// Used if windows_switching_flag is set
	// - Mono = 1 bit
	// - Stereo = 2 bits
	MixedBlockFlag int

	// Indicates which Huffman table to use when decoding the big_values field
	// - (Switching flag = 1)
	// - Mono = 10 bits
	// - Stereo = 20 bits
	//
	// - (Switching flag = 0)
	// - Mono = 15 bits
	// - Stereo = 30 bits
	TableSelect int

	// Used when windows_switching_flag is set and block_type = 10
	// - Mono = 9 bits
	// - Stereo = 18 bits
	SubblockGain int

	// The number of scalefactor bands in region 0
	// - Mono = 4 bits
	// - Stereo = 8 bits
	Region0Count int

	// The number of scalefactor bands in region 1
	// - Mono = 3 bits
	// - Stereo = 6 bits
	Region1Count int

	// Used to amplify high frequencies
	// - Mono = 1 bit
	// - Stereo = 2 bits
	Preflag int

	// Step size for quanitizing scale factors
	// - Mono = 1 bit
	// - Stereo = 2 bits
	ScalfacScale int

	// Specifies which Huffman table to use for the count1 region
	// - Mono = 1 bit
	// - Stereo = 2 bits
	Count1TableSelect int
}

type SideInfo struct {
	MainDataStart int

	PrivateBits int

	SCSFI int

	Group1 int

	Group2 int
}

// TODO
func ParseSideInfo(data int, stereo bool) *SideInfo {
	return nil
}
