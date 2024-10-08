package rat

import (
	"fmt"
)

type HeaderRaw struct {
	name     []byte
	mode     []byte
	uid      []byte
	gid      []byte
	size     []byte
	mtime    []byte
	chksum   []byte
	typeflag byte
	linkname []byte
	magic    []byte
	version  []byte
	uname    []byte
	gname    []byte
	devmajor []byte
	devminor []byte
	prefix   []byte
}

func readWrapper(manager IDataBytesManager, size uint) []byte {
	data, _ := manager.Read(size)
	return data
}

func NewHeaderRaw(headerRaw []byte) HeaderRaw {
	manager := NewDataBytesSliceManager(headerRaw)
	return HeaderRaw{readWrapper(manager, 100),
		readWrapper(manager, 8),
		readWrapper(manager, 8),
		readWrapper(manager, 8),
		readWrapper(manager, 12),
		readWrapper(manager, 12),
		readWrapper(manager, 8),
		readWrapper(manager, 1)[0],
		readWrapper(manager, 100),
		readWrapper(manager, 6),
		readWrapper(manager, 2),
		readWrapper(manager, 32),
		readWrapper(manager, 32),
		readWrapper(manager, 8),
		readWrapper(manager, 8),
		readWrapper(manager, 155)}
}

func (header *HeaderRaw) ToString() string {
	return fmt.Sprintf("%+v", header)
}
