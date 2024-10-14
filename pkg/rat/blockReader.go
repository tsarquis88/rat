package rat

import (
	"errors"
	"io"
	"os"
)

type BlockReader struct {
	fileHandle os.File
	blockSize  uint
}

func NewBlockReader(filename string, blockSize uint) BlockReader {
	fileHandle, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	return BlockReader{*fileHandle, blockSize}
}

func (manager *BlockReader) ReadBlock() ([]byte, bool) {
	buff := FillWith([]byte{}, 0, manager.blockSize)
	n, err := manager.fileHandle.Read(buff)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return []byte{}, false
		}
		panic(err)
	}

	return FillWith(buff, 0, manager.blockSize), (n == BlockSize)
}
