package mp3

// TODO - change case of all these to eliminate underscores
// TODO - clean up doc comments
type Header struct {
	// Fixed value used as a searchable entry point to the stream (12 bits)
	SyncWord [12]bool

	// Specifies the MPEG version (1 bit)
	//  - 1 = MPEG-1
	//	- 0 = MPEG-2
	Version bool

	// Specifies ..... something (2 bits)
	//  - 00 = Reserved
	//  - 01 = Layer 3
	//  - 10 = Layer 2
	//  - 11 = Layer 1
	Layer [2]bool

	// If set, CRC is used for correction (1 bit)
	ErrorProtection bool

	// The bit rate used to encode the frame (4 bits)
	BitRate [4]bool

	// Sampling frequency (2 bits)
	Frequency [2]bool

	// Fills extra frame space for streams with 128 kbit/s and 44100 Hz (1 bit)
	PaddingBit bool

	// Triggers application-specific behavior (1 bit)
	PrivacyBit bool

	// Channel mode (2 bits)
	//  - 00 = Stereo
	//  - 01 = Joint stereo
	//  - 10 = Dual channel
	//  - 11 = Single channel
	Mode [2]bool

	// Specifies which methods to use in joint stereo mode (2 bits)
	//  - First bit = MS stereo on/off
	//  - Second bit = Intensity stereo on/off
	ModeExtension [2]bool

	// Specifies if the file is copyrighted (1 bit)
	Copyrighted bool

	// Indicates something ... (1 bit)
	Original bool

	// Indicates that the file requires equalization (2 bits)
	//  - 00 = None
	//  - 01 = 50/15 ms
	//  - 10 = Reserved
	//  - 11 = CCITT J.17
	Emphasis [2]bool
}

// TODO - extract individual bits with bit_utils
//func parse(bits uint32) Header {}
