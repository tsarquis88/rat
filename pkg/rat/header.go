package rat

import (
	"encoding/binary"
	"fmt"
	"io/fs"
	"os/user"
	"strconv"
	"syscall"
)

type Header struct {
	name     string
	mode     uint32
	uid      uint32
	gid      uint32
	size     uint
	filetype uint8
	uname    string
	gname    string
	mktime   uint
}

func NewHeader(fileStat fs.FileInfo, fileType uint8) Header {
	uid, gid := getIds(fileStat)
	uname := getName(uid)
	gname := getName(gid)

	return Header{
		fileStat.Name(),
		uint32(fileStat.Mode()),
		uid,
		gid,
		uint(fileStat.Size()),
		RegulatFileType,
		uname,
		gname,
		getMktime(fileStat),
	}
}

func getName(id uint32) string {

	user, err := user.LookupId(strconv.Itoa(int(id)))
	if err != nil {
		panic(err)
	}
	return user.Name
}

func getMktime(fileStat fs.FileInfo) uint {
	if stat, ok := fileStat.Sys().(*syscall.Stat_t); ok {
		return uint(stat.Mtim.Sec)
	}
	panic("Couldn't get mktime")
}

func getIds(fileStat fs.FileInfo) (UID uint32, GID uint32) {
	if stat, ok := fileStat.Sys().(*syscall.Stat_t); ok {
		UID = stat.Uid
		GID = stat.Gid
		return
	}
	panic("Couldn't get UID and GID")
}

func trimPadding(value []byte) string {
	var paddingIdx int
	for paddingIdx = 0; paddingIdx < len(value); paddingIdx++ {
		if value[paddingIdx] == 0 {
			break
		}
	}
	return string(value[:paddingIdx])
}

func getMode(value []byte) uint32 {
	return binary.LittleEndian.Uint32(value)
}

func getModeRaw(value uint32) []byte {
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

func getSizeRaw(value uint) []byte {
	sizeSlice := FillWith([]byte{}, 48, 12)
	sizeSlice[len(sizeSlice)-1] = 0
	octal := DecimalToOctal(value)
	for i := len(octal) - 1; i >= 0; i-- {
		sizeSlice[len(sizeSlice)-(len(octal)-i)-1] = octal[i]
	}
	return sizeSlice
}

func getNameRaw(value string) []byte {
	return FillWith([]byte(value), 0, 100)
}

func getUnameRaw(value string) []byte {
	return FillWith([]byte(value), 0, 32)
}

func getIdRaw(value uint32) []byte {
	idSlice := FillWith([]byte{}, 48, 8)
	idSlice[len(idSlice)-1] = 0
	octal := DecimalToOctal(uint(value))
	for i := len(octal) - 1; i >= 0; i-- {
		idSlice[len(idSlice)-(len(octal)-i)-1] = octal[i]
	}
	return idSlice
}

func getMagicRaw() []byte {
	return []byte{'u', 's', 't', 'a', 'r', 32}
}

func getVersionRaw() []byte {
	return []byte{32, 0}
}

func NewHeaderFromRaw(headerRaw HeaderRaw) Header {
	return Header{
		trimPadding(headerRaw.name),
		getMode(headerRaw.mode),
		getMode(headerRaw.uid),
		getMode(headerRaw.gid),
		OctalToDecimal(headerRaw.size, 11),
		headerRaw.typeflag,
		trimPadding(headerRaw.uname),
		trimPadding(headerRaw.gname),
		uint(binary.LittleEndian.Uint64(headerRaw.mtime)),
	}
}

func getMktimeRaw(value uint) []byte {
	timeSlice := FillWith([]byte{}, 48, 12)
	timeSlice[len(timeSlice)-1] = 0
	octal := DecimalToOctal(value)
	for i := len(octal) - 1; i >= 0; i-- {
		timeSlice[len(timeSlice)-(len(octal)-i)-1] = octal[i]
	}
	return timeSlice
}

func getChksum(header HeaderRaw) []byte {
	return []byte{48, 49, 50, 52, 54, 55, 0, 32}
	// return FillWith([]byte{}, 0, 8)
	// dump := header.Dump()
	// fmt.Println(len(dump))
	// acc := uint64(0)
	// for idx := 0; idx < 512; idx = idx+8{
	// 	acc = acc + binary.LittleEndian.Uint64(dump[idx:idx+8])
	// }
	// fmt.Println(acc)

	// chksum := []byte{}
	// binary.LittleEndian.PutUint64(chksum, acc)
	// return chksum
}

func (header *Header) ToString() string {
	return fmt.Sprintf("%+v\n", header)
}

func (header *Header) ToRaw() HeaderRaw {
	var rawHeader HeaderRaw
	rawHeader.chksum = FillWith([]byte{}, 32, 8)
	rawHeader.linkname = make([]byte, 100)
	rawHeader.devmajor = make([]byte, 8)
	rawHeader.devminor = make([]byte, 8)
	rawHeader.prefix = make([]byte, 155)

	rawHeader.name = getNameRaw(header.name)
	rawHeader.mode = getModeRaw(header.mode)
	rawHeader.uid = getIdRaw(header.uid)
	rawHeader.gid = getIdRaw(header.gid)
	rawHeader.size = getSizeRaw(header.size)
	rawHeader.uname = getUnameRaw(header.uname)
	rawHeader.gname = getUnameRaw(header.gname)
	rawHeader.magic = getMagicRaw()
	rawHeader.version = getVersionRaw()
	rawHeader.mtime = getMktimeRaw(header.mktime)
	rawHeader.typeflag = header.filetype
	rawHeader.chksum = getChksum(rawHeader)
	return rawHeader
}
