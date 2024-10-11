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

func (header *HeaderRaw) Dump() []byte {
	var dump []byte
	dump = append(dump, header.name...)
	dump = append(dump, header.mode...)
	dump = append(dump, header.uid...)
	dump = append(dump, header.gid...)
	dump = append(dump, header.size...)
	dump = append(dump, header.mtime...)
	dump = append(dump, header.chksum...)
	dump = append(dump, header.typeflag)
	dump = append(dump, header.linkname...)
	dump = append(dump, header.magic...)
	dump = append(dump, header.version...)
	dump = append(dump, header.uname...)
	dump = append(dump, header.gname...)
	dump = append(dump, header.devmajor...)
	dump = append(dump, header.devminor...)
	dump = append(dump, header.prefix...)
	return FillWith(dump, 0, 512)
}

func (header *HeaderRaw) ToString() string {
	return fmt.Sprintf("Name: %x (%s)", header.name, header.name) +
		fmt.Sprintf("\nMode: %s (%d)", header.mode, header.mode) +
		fmt.Sprintf("\nUID: %s (%d)", header.uid, header.uid) +
		fmt.Sprintf("\nGID: %s (%d)", header.gid, header.gid) +
		fmt.Sprintf("\nSize: %x (%d)", header.size, header.size) +
		fmt.Sprintf("\nMtime: %s (%d)", header.mtime, header.mtime) +
		fmt.Sprintf("\nChksum: %s (%d)", header.chksum, header.chksum) +
		fmt.Sprintf("\nTypeflag: %c", header.typeflag) +
		fmt.Sprintf("\nLinkname: %x", header.linkname) +
		fmt.Sprintf("\nMagic: %s (%d)", header.magic, header.magic) +
		fmt.Sprintf("\nVersion: %s (%d)", header.version, header.version) +
		fmt.Sprintf("\nUname: %x (%s)", header.uname, header.uname) +
		fmt.Sprintf("\nGname: %x (%s)", header.gname, header.gname) +
		fmt.Sprintf("\nDevmajor: %x", header.devmajor) +
		fmt.Sprintf("\nDevminor: %x", header.devminor) +
		fmt.Sprintf("\nPrefix: %x", header.prefix)
}
