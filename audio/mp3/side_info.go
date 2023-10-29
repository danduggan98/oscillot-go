package mp3

type SideInfoGranule struct {
	// Number of bits used for scalefactors and encoded data
	// - Mono = 12 bits
	// - Stereo = 24 bits
	Part23Length uint32

	// Total size of big_values partition
	// - Mono = 9 bits
	// - Stereo = 18 bits
	BigValues uint32

	// Quantization step size
	// - Mono = 8 bits
	// - Stereo = 16 bits
	GlobalGain uint32

	// Number of bits used for transmission of scalefactors
	// - Mono = 4 bits
	// - Stereo = 8 bits
	ScalefacCompress uint32

	// Indicates that a different window is used
	// - Mono = 1 bit
	// - Stereo = 2 bits
	WindowSwitchingFlag uint32

	// Indicates the type of window used
	// - Mono = 2 bits
	// - Stereo = 4 bits
	BlockType uint32

	// Used if windows_switching_flag is set
	// - Mono = 1 bit
	// - Stereo = 2 bits
	MixedBlockFlag uint32

	// Indicates which Huffman table to use when decoding the big_values field
	// - (Switching flag = 1)
	// - Mono = 10 bits
	// - Stereo = 20 bits
	//
	// - (Switching flag = 0)
	// - Mono = 15 bits
	// - Stereo = 30 bits
	TableSelect uint32

	// Used when windows_switching_flag is set and block_type = 10
	// - Mono = 9 bits
	// - Stereo = 18 bits
	SubblockGain uint32

	// The number of scalefactor bands in region 0
	// - Mono = 4 bits
	// - Stereo = 8 bits
	Region0Count uint32

	// The number of scalefactor bands in region 1
	// - Mono = 3 bits
	// - Stereo = 6 bits
	Region1Count uint32

	// Used to amplify high frequencies
	// - Mono = 1 bit
	// - Stereo = 2 bits
	Preflag uint32

	// Step size for quanitizing scale factors
	// - Mono = 1 bit
	// - Stereo = 2 bits
	ScalfacScale uint32

	// Specifies which Huffman table to use for the count1 region
	// - Mono = 1 bit
	// - Stereo = 2 bits
	Count1TableSelect uint32
}

type SideInfo struct {
	MainDataStart uint32

	PrivateBits uint32

	SCSFI uint32

	Group1 uint32

	Group2 uint32
}

// TODO
func ParseSideInfo(data uint32, stereo bool) *SideInfo {
	if stereo {
		return &SideInfo{}
	} else {
		return &SideInfo{}
	}
}
