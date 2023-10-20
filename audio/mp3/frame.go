package mp3

const FRAME_SIZE = 384 // # samples

// GOAL - be able to parse a frame directly
type Frame struct {
	// Frame header (32 bits)
	Header Header

	// Optional error correction bits (0 | 16 bits)
	CRC CRC

	// uhhh
	SideInfo SideInfo

	// Audio data
	MainData MainData

	// uhhh
	AncillaryData AncillaryData
}

func ParseFrame(data int) *Frame {
	header := ParseHeader(data)

	return &Frame{
		Header:   *header,
		CRC:      *ParseCRC(data),
		SideInfo: *ParseSideInfo(data, header.isStereo()),
	}
}

func (f Frame) FrameLength() int {
	ratio := f.Header.CalculateBitrate() / f.Header.CalculateSampleRate()
	return 144*ratio + f.Header.PaddingBit
}
