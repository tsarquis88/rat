package rat

import (
	"encoding/binary"
	"fmt"
)

type Header struct {
	name     string
	mode     uint32
	size     uint
	filetype uint8
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

func NewHeader(headerRaw HeaderRaw) Header {
	return Header{trimPadding(headerRaw.name), getMode(headerRaw.mode), OctalToDecimal(headerRaw.size, 11), headerRaw.typeflag}
}

func (header *Header) ToString() string {
	return fmt.Sprintf("%+v\n", header)
}

func (header *Header) ToRaw() HeaderRaw {
	var rawHeader HeaderRaw
	rawHeader.name = make([]byte, 100)
	rawHeader.mode = make([]byte, 8)
	rawHeader.uid = make([]byte, 8)
	rawHeader.gid = make([]byte, 8)
	rawHeader.size = make([]byte, 12)
	rawHeader.mtime = make([]byte, 12)
	rawHeader.chksum = make([]byte, 8)
	rawHeader.typeflag = byte(0)
	rawHeader.linkname = make([]byte, 100)
	rawHeader.magic = make([]byte, 6)
	rawHeader.version = make([]byte, 2)
	rawHeader.uname = make([]byte, 32)
	rawHeader.gname = make([]byte, 32)
	rawHeader.devmajor = make([]byte, 8)
	rawHeader.devminor = make([]byte, 8)
	rawHeader.prefix = make([]byte, 155)

	rawHeader.name = getNameRaw(header.name)
	rawHeader.mode = getModeRaw(header.mode)
	rawHeader.size = getSizeRaw(header.size)
	rawHeader.typeflag = header.filetype
	return rawHeader
}
