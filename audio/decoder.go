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
func FindFrame(file *os.File) int { return 0 }

// Parse the frame starting at a particular byte
// TODO
func ReadFrame(file *os.File, start int) *mp3.Frame {
	return nil
}
