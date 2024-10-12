package rat

import (
	"io/fs"
	"os"
)

type BlockWriter struct {
	fileHandle os.File
}

func NewBlockWriter(filename string, mode fs.FileMode) BlockWriter {
	fileHandle, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		panic(err)
	}
	return BlockWriter{*fileHandle}
}

func (manager *BlockWriter) WriteBlock(block []byte) {
	_, err := manager.fileHandle.Write(block)
	if err != nil {
		panic(err)
	}
}
