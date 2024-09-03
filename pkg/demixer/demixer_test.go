package demixer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"example.com/dataBytesManagerMock"
)

func TestDemixOneFile(t *testing.T) {
	fileSize := make([]byte, 8)
	fileSize[0] = 10

	var data []byte
	data = append(data, byte(1))
	data = append(data, byte(8))
	data = append(data, []byte("filename")...)
	data = append(data, fileSize...)
	data = append(data, []byte("1234567890")...)

	var expectedDemixData []DemixData
	expectedDemixData = append(expectedDemixData, DemixData{"filename", []byte("1234567890")})

	manager := dataBytesManagerMock.NewDataBytesManagerMock(data)
	demixData := Demix(manager)

	assert.Equal(t, expectedDemixData, demixData)
}

func TestDemixMultipleFiles(t *testing.T) {
	var data []byte
	data = append(data, byte(3))
	fileSize := make([]byte, 8)
	
	fileSize[0] = 14
	data = append(data, byte(9))
	data = append(data, []byte("file.json")...)
	data = append(data, fileSize...)
	
	fileSize[0] = 25
	data = append(data, byte(8))
	data = append(data, []byte("data.xml")...)
	data = append(data, fileSize...)
	
	fileSize[0] = 19
	data = append(data, byte(10))
	data = append(data, []byte("output.log")...)
	data = append(data, fileSize...)

	data = append(data, []byte("{<S\"dotemetesattih\"li:snf>gab llhsaaebp}lpae<n/edde.tails>")...)

	var expectedDemixData []DemixData
	expectedDemixData = append(expectedDemixData, DemixData{"file.json", []byte("{\"test\":false}")})
	expectedDemixData = append(expectedDemixData, DemixData{"data.xml", []byte("<details>blabla</details>")})
	expectedDemixData = append(expectedDemixData, DemixData{"output.log", []byte("Something happened.")})

	manager := dataBytesManagerMock.NewDataBytesManagerMock(data)
	demixData := Demix(manager)

	assert.Equal(t, expectedDemixData, demixData)
}
