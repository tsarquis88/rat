package demixer

import (
	"example.com/dataBytesManager"
	"example.com/metadataManager"
)

type DemixData struct {
	Filename string
	Data []byte
	Mode uint32
}

type demixDataInternal struct {
	demixData DemixData
	missingBytes int
}

func Demix(dataBytesSource dataBytesManager.IDataBytesManager) []DemixData {
	readData, _ := dataBytesSource.Read(1024*10)
	metadatas, metadataEndByte := metadataManager.Parse(readData)

	var filesDemixData []demixDataInternal
	for _, metadata := range metadatas {
		filesDemixData = append(filesDemixData, demixDataInternal{DemixData{metadata.Filename, []byte{}, metadata.Mode}, int(metadata.Size)})
	}

	fileIdx := 0
	byteIdx := metadataEndByte
	filesQty := len(metadatas)
	parsedFiles := 0
	for {
		if filesDemixData[fileIdx].missingBytes > 0 {
			filesDemixData[fileIdx].demixData.Data = append(filesDemixData[fileIdx].demixData.Data, readData[byteIdx])
			filesDemixData[fileIdx].missingBytes--
			
			if filesDemixData[fileIdx].missingBytes == 0 {
				parsedFiles++
				if parsedFiles >= filesQty {
					break
				}
			}
			byteIdx++	
		}

		fileIdx++
		if fileIdx >= filesQty {
			fileIdx = 0
		}
	}

	var demixDataFinal []DemixData
	for _, demixData := range filesDemixData {
		demixDataFinal = append(demixDataFinal, demixData.demixData)
	}
	return demixDataFinal
}

