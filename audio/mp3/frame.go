package mp3

const FRAME_SIZE_BYTES = 418 // 384?

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
	return &Frame{
		Header:   *ParseHeader(data),
		CRC:      *ParseCRC(data),
		SideInfo: *ParseSideInfo(data),
	}
}

func (f Frame) FrameLength() int {
	ratio := f.Header.CalculateBitrate() / f.Header.CalculateSampleRate()
	return 144*ratio + f.Header.PaddingBit
}
