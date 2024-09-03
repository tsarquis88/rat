package demixer

import (
	"fmt"
	"example.com/dataBytesManager"
	"example.com/metadataManager"
)

func Demix(dataBytesSource dataBytesManager.IDataBytesManager) {
	readData, _ := dataBytesSource.Read(1024)
	metadatas := metadataManager.Parse(readData)

	fmt.Printf("Files found: %d\n", len(metadatas))
}

