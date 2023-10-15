package audio

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func decode() {
	// TODO
}

func read_file(path string) {
	file, err := os.Open(path)
	check(err)

	read_frame(file, 0) // TODO - extract in chunks
	file.Close()
}

func read_frame(file *os.File, start int64) {
	//buffer := make([]byte, mp3.FRAME_SIZE_BYTES)

	_, err := file.Seek(start, 1)
	check(err)

	//chunk, err := file.Read(buffer)
	check(err)

	// TODO - read chunk as bool array
	// TODO - convert to Frame object
}
