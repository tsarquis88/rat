package rat

import (
	"io/fs"
	"os"
)

type DataBytesDumper struct {
	fileHandle os.File
}

func NewDataBytesDumper(filename string, mode fs.FileMode) DataBytesDumper {
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
