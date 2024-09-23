package midem

import (
	"os"
)

type DataBytesDumper struct {
	fileHandle os.File
}

func NewDataBytesDumper(filename string) DataBytesDumper {
	mode := os.FileMode(438) // RW-RW-RW
	fileHandle, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		panic(err)
	}
	return DataBytesDumper{*fileHandle}
}

func (manager DataBytesDumper) Dump(data []byte) {
	_, err := manager.fileHandle.Write(data)
	if err != nil {
		panic(err)
	}
}
