package rat

import (
	"encoding/binary"
	"fmt"
	"math"
)

type Header struct {
	name     string
	mode     uint32
	size     uint
	filetype uint8
}

func octalToDecimal(octal []byte) uint {
	decimal := uint(0)
	octalLen := 11
	for i := octalLen - 1; i >= 0; i-- {
		if octal[i] != 48 {
			exp := octalLen - 1 - i
			value := uint(octal[i]-48) * uint(math.Pow(float64(8), float64(exp)))
			decimal = decimal + value
		}
	}
	return decimal
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

func NewHeader(headerRaw HeaderRaw) Header {
	return Header{trimPadding(headerRaw.name), getMode(headerRaw.mode), octalToDecimal(headerRaw.size), headerRaw.typeflag}
}

func (header *Header) ToString() string {
	return fmt.Sprintf("%+v\n", header)
}
