package audio

import (
	"os"

	"github.com/danduggan98/oscillot-go/audio/mp3"
)

// TODO
func ReadFile(path string) (*os.File, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	return file, nil
}

// Seek to the first frame in the file
// TODO
func FindFrame(file *os.File) uint32 { return 0 }

// Parse the frame starting at a particular byte
// TODO - read a chunk of bytes
func ReadFrame(file *os.File, start uint32) *mp3.Frame {
	/* TODO
	basic process:
	1. read 32 bytes from starting pointer into buffer
	2. parse bytes as frame header
	3. depending on is_stereo, etc. read either 17 or 32 bits from there
	4. parse into side info
	5. create frame from bytes
	*/
	return nil
}
