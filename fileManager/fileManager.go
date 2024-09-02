package fileManager

import (
	"errors"
	"io"
	"os"
)

type FileManager struct {
	filename string
	fileHandle os.File
}

func NewFileManager(filename string, write bool) (FileManager) {
	var mode int
	if write {
		mode = os.O_CREATE | os.O_WRONLY
	} else {
		mode = os.O_RDONLY
	}

	fileHandle, err := os.OpenFile(filename, mode, 0755)
	if err != nil {
		panic(err)
	}
	return FileManager {filename, *fileHandle}
}

func (manager FileManager) Read() (byte, int) {
	buff := make([]byte, 1)
	readBytes , err := manager.fileHandle.Read(buff)
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	return buff[0], readBytes
}

func (manager FileManager) Write(data byte) (int) {
	bytesWriten, err := manager.fileHandle.Write([]byte{data})
	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}
	return bytesWriten
}
