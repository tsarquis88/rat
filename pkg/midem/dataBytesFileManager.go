package midem

import (
	"errors"
	"io"
	"os"
)

type DataBytesFileManager struct {
	filename   string
	originDir  string
	fileHandle os.File
}

func NewDataBytesFileManager(filename string, originDir string) DataBytesFileManager {
	fileHandle, err := os.OpenFile(filename, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	return DataBytesFileManager{filename, originDir, *fileHandle}
}

func (manager DataBytesFileManager) Read(bytesQty uint) ([]byte, int) {
	buff := make([]byte, bytesQty)
	readBytes, err := manager.fileHandle.Read(buff)
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	return buff, readBytes
}

func (manager DataBytesFileManager) Name() string {
	return manager.filename
}

func (manager DataBytesFileManager) Origin() string {
	return manager.originDir
}
