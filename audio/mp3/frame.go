package mp3

const FRAME_SIZE_BYTES = 418

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
