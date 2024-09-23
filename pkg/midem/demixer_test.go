package midem

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDemixOneFile(t *testing.T) {
	fileSize := make([]byte, 8)
	fileSize[0] = 10
	fileMode := []byte{0, 6, 6, 6}

	var data []byte
	data = append(data, byte(1))
	data = append(data, byte(8))
	data = append(data, []byte("filename")...)
	data = append(data, fileSize...)
	data = append(data, fileMode...)
	data = append(data, []byte("1234567890")...)

	var expectedDemixData []DemixData
	expectedDemixData = append(expectedDemixData, DemixData{"filename", []byte("1234567890"), uint32(binary.LittleEndian.Uint32(fileMode))})

	manager := NewDataBytesManagerMock(data)
	demixData := Demix(manager)

	assert.Equal(t, expectedDemixData, demixData)
}

func TestDemixMultipleFiles(t *testing.T) {
	var data []byte
	data = append(data, byte(3))
	fileSize := make([]byte, 8)

	fileSize[0] = 14
	fileModeA := []byte{0, 5, 5, 5}
	data = append(data, byte(9))
	data = append(data, []byte("file.json")...)
	data = append(data, fileSize...)
	data = append(data, fileModeA...)

	fileSize[0] = 25
	fileModeB := []byte{0, 6, 6, 6}
	data = append(data, byte(8))
	data = append(data, []byte("data.xml")...)
	data = append(data, fileSize...)
	data = append(data, fileModeB...)

	fileSize[0] = 19
	fileModeC := []byte{0, 7, 7, 7}
	data = append(data, byte(10))
	data = append(data, []byte("output.log")...)
	data = append(data, fileSize...)
	data = append(data, fileModeC...)

	data = append(data, []byte("{<S\"dotemetesattih\"li:snf>gab llhsaaebp}lpae<n/edde.tails>")...)

	var expectedDemixData []DemixData
	expectedDemixData = append(expectedDemixData, DemixData{"file.json", []byte("{\"test\":false}"), uint32(binary.LittleEndian.Uint32(fileModeA))})
	expectedDemixData = append(expectedDemixData, DemixData{"data.xml", []byte("<details>blabla</details>"), uint32(binary.LittleEndian.Uint32(fileModeB))})
	expectedDemixData = append(expectedDemixData, DemixData{"output.log", []byte("Something happened."), uint32(binary.LittleEndian.Uint32(fileModeC))})

	manager := NewDataBytesManagerMock(data)
	demixData := Demix(manager)

	assert.Equal(t, expectedDemixData, demixData)
}
