package midem

type DemixData struct {
	Filename string
	Data     []byte
	Mode     uint32
}

type demixDataInternal struct {
	demixData    DemixData
	missingBytes int
}

func Demix(dataBytesSource IDataBytesManager) []DemixData {
	metadatas := Parse(dataBytesSource)

	var filesDemixData []demixDataInternal
	for _, metadata := range metadatas {
		filesDemixData = append(filesDemixData, demixDataInternal{DemixData{metadata.Filename, []byte{}, metadata.Mode}, int(metadata.Size)})
	}

	fileIdx := 0
	filesQty := len(metadatas)
	parsedFiles := 0
	for {
		if filesDemixData[fileIdx].missingBytes > 0 {
			newByte, _ := dataBytesSource.Read(1)
			filesDemixData[fileIdx].demixData.Data = append(filesDemixData[fileIdx].demixData.Data, newByte[0])
			filesDemixData[fileIdx].missingBytes--

			if filesDemixData[fileIdx].missingBytes == 0 {
				parsedFiles++
				if parsedFiles >= filesQty {
					break
				}
			}
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
