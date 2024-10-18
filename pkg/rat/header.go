package rat

import (
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

const NameLen = 100
const ModeLen = 8
const UidLen = 8
const GidLen = 8
const SizeLen = 12
const MtimeLen = 12
const ChksumLen = 8
const LinknameLen = 100
const MagicLen = 6
const VersionLen = 2
const UnameLen = 32
const GnameLen = 32
const DevmajorLen = 8
const DevminorLen = 8
const PrefixLen = 155

type Header struct {
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

func getIds(fileStat fs.FileInfo) (UID uint32, GID uint32) {
	if stat, ok := fileStat.Sys().(*syscall.Stat_t); ok {
		UID = stat.Uid
		GID = stat.Gid
		return
	}
	panic("Couldn't get UID and GID")
}

func dumpMode(value uint32) []byte {
	usersMode := byte(value & 0b00000111)
	groupMode := byte((value & 0b00111000) >> 3)
	ownerMode := byte((value & 0b111000000) >> 6)
	modeSlice := FillWith([]byte{}, 48, 8)
	modeSlice[7] = 0
	modeSlice[6] += usersMode
	modeSlice[5] += groupMode
	modeSlice[4] += ownerMode
	return modeSlice
}

func dumpValue(value uint, size uint) []byte {
	valueSlice := FillWith([]byte{}, 48, size)
	valueSlice[len(valueSlice)-1] = 0
	octal := DecimalToOctal(uint(value))
	for i := len(octal) - 1; i >= 0; i-- {
		valueSlice[len(valueSlice)-(len(octal)-i)-1] = octal[i]
	}
	return valueSlice
}

func getMtime(fileStat fs.FileInfo) uint {
	if stat, ok := fileStat.Sys().(*syscall.Stat_t); ok {
		return uint(stat.Mtim.Sec)
	}
	panic("Couldn't get mktime")
}

func dumpMagic() []byte {
	return []byte{'u', 's', 't', 'a', 'r', 32}
}

func getName(id uint32) string {
	user, err := user.LookupId(strconv.Itoa(int(id)))
	if err != nil {
		return ""
	}
	return user.Name
}

func NewHeaderFromFile(file string, blockSize uint) Header {
	fileHandle, err := os.OpenFile(file, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	fileStat, err := fileHandle.Stat()
	if err != nil {
		panic(err)
	}

	uid, gid := getIds(fileStat)

	typeflag := RegularFileType
	if IsDir(file) {
		typeflag = DirFileType
	}

	file = TrimPrefixRecursive(file, "/")
	file = TrimPrefixRecursive(file, "../")

	var header Header
	header.name = FillWith([]byte(file), 0, NameLen)
	header.mode = dumpMode(uint32(fileStat.Mode()))
	header.uid = dumpValue(uint(uid), UidLen)
	header.gid = dumpValue(uint(gid), GidLen)
	header.size = dumpValue(uint(fileStat.Size()), SizeLen)
	header.mtime = dumpValue(getMtime(fileStat), MtimeLen)
	header.chksum = FillWith([]byte{}, 32, ChksumLen)
	header.typeflag = byte(typeflag)
	header.linkname = FillWith([]byte{}, 0, LinknameLen)
	header.magic = dumpMagic()
	header.version = FillWith([]byte{32}, 0, VersionLen)
	header.uname = FillWith([]byte(getName(uid)), 0, UnameLen)
	header.gname = FillWith([]byte(getName(gid)), 0, GnameLen)
	header.devmajor = FillWith([]byte{}, 0, DevmajorLen)
	header.devminor = FillWith([]byte{}, 0, DevminorLen)
	header.prefix = FillWith([]byte{}, 0, PrefixLen)

	// Checksum always at the end.
	header.chksum = ShiftLeft(dumpValue(GetChecksum(header.Dump(blockSize)), ChksumLen), 1, 32)

	return header
}

func readWrapper(manager IDataBytesManager, size uint) []byte {
	data, _ := manager.Read(size)
	return data
}

func NewHeaderFromDump(headerDump []byte) Header {
	manager := NewDataBytesSliceManager(headerDump)
	return Header{readWrapper(manager, NameLen),
		readWrapper(manager, ModeLen),
		readWrapper(manager, UidLen),
		readWrapper(manager, GidLen),
		readWrapper(manager, SizeLen),
		readWrapper(manager, MtimeLen),
		readWrapper(manager, ChksumLen),
		readWrapper(manager, 1)[0],
		readWrapper(manager, LinknameLen),
		readWrapper(manager, MagicLen),
		readWrapper(manager, VersionLen),
		readWrapper(manager, UnameLen),
		readWrapper(manager, GnameLen),
		readWrapper(manager, DevmajorLen),
		readWrapper(manager, DevminorLen),
		readWrapper(manager, PrefixLen)}
}

func (header *Header) Dump(blockSize uint) []byte {
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
	return FillWith(dump, 0, blockSize)
}

func (header *Header) ToString() string {
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
