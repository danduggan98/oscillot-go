package mp3

import "github.com/danduggan98/oscillot-go/util"

type SCFSI struct {

	// If enabled, scale factors are shared between granules (1 bit)
	ShareBit uint32

	//
	Band uint32
}

type SideInfoGranule struct {
	// Number of bits used for scalefactors and encoded data
	//  - Mono = 12 bits
	//  - Stereo = 24 bits
	Part23Length uint32

	// Total size of big_values partition
	//  - Mono = 9 bits
	//  - Stereo = 18 bits
	BigValues uint32

	// Quantization step size
	//  - Mono = 8 bits
	//  - Stereo = 16 bits
	GlobalGain uint32

	// Number of bits used for transmission of scalefactors
	//  - Mono = 4 bits
	//  - Stereo = 8 bits
	ScalefacCompress uint32

	// Indicates that a different window is used
	//  - Mono = 1 bit
	//  - Stereo = 2 bits
	WindowSwitchingFlag uint32

	// Indicates the type of window used
	//  - Mono = 2 bits
	//  - Stereo = 4 bits
	BlockType uint32

	// Used if windows_switching_flag is set
	//  - Mono = 1 bit
	//  - Stereo = 2 bits
	MixedBlockFlag uint32

	// Indicates which Huffman table to use when decoding the big_values field
	//  - (Switching flag = 1)
	//    - Mono = 10 bits
	//    - Stereo = 20 bits
	//
	//  - (Switching flag = 0)
	//      - Mono = 15 bits
	//      - Stereo = 30 bits
	TableSelect uint32

	// Used when windows_switching_flag is set and block_type = 10
	//  - Mono = 9 bits
	//  - Stereo = 18 bits
	SubblockGain uint32

	// The number of scalefactor bands in region 0
	//  - Mono = 4 bits
	//  - Stereo = 8 bits
	Region0Count uint32

	// The number of scalefactor bands in region 1
	//  - Mono = 3 bits
	//  - Stereo = 6 bits
	Region1Count uint32

	// Used to amplify high frequencies
	//  - Mono = 1 bit
	//  - Stereo = 2 bits
	Preflag uint32

	// Step size for quanitizing scale factors
	//  - Mono = 1 bit
	//  - Stereo = 2 bits
	ScalfacScale uint32

	// Specifies which Huffman table to use for the count1 region
	//  - Mono = 1 bit
	//  - Stereo = 2 bits
	Count1TableSelect uint32
}

type SideInfo struct {

	// Indicates byte where main data begins (negative offset from first byte of sync word)
	MainDataStart uint32

	// Unused
	PrivateBits uint32

	// Scalefactor selection - determines which scalefactors are used in each granule
	SCFSI SCFSI

	Group1 SideInfoGranule

	Group2 SideInfoGranule
}

// TODO
func ParseSideInfo(data []byte, stereo bool) *SideInfo {
	bits := util.BytesToInts(data)
	mainDataPtr := util.BitSlice(0, 9, bits[0])

	// scfsi = scfsi bit + scfsci band
	if stereo {
		return &SideInfo{
			MainDataStart: mainDataPtr,
			PrivateBits:   util.BitSlice(9, 3, bits[0]),
			SCFSI: SCFSI{
				ShareBit: util.NthBit(12, bits[0]),
				Band:     util.BitSlice(13, 4, bits[0]), // ???
			},
			// TODO
		}
	} else {
		return &SideInfo{
			MainDataStart: mainDataPtr,
			PrivateBits:   util.BitSlice(9, 5, bits[0]),
			// TODO
		}
	}
}
