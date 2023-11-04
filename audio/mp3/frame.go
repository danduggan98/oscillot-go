package mp3

const FRAME_SIZE = 384 // # samples

// GOAL - be able to parse a frame directly
type Frame struct {
	// Frame header (32 bits)
	Header Header

	// Optional error correction bits (0 | 16 bits)
	CRC CRC

	// ... (17 | 32 bits)
	SideInfo SideInfo

	// Audio data
	MainData MainData

	// uhhh
	AncillaryData AncillaryData
}

func ParseFrame(data []byte) *Frame {
	header := ParseHeader(data[:4]) // 4 bytes
	crc := ParseCRC(data[4:6])      // 2 bytes

	var siEnd = 39 // 32 bytes
	if header.isStereo() {
		siEnd = 24 // 17 bytes
	}
	sideInfo := ParseSideInfo(data[6:siEnd], header.isStereo())

	return &Frame{
		Header:   *header,
		CRC:      *crc,
		SideInfo: *sideInfo,
	}
}

func (f Frame) FrameLength() uint32 {
	ratio := f.Header.CalculateBitrate() / f.Header.CalculateSampleRate()
	return 144*ratio + f.Header.PaddingBit
}
