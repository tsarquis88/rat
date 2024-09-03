package metadataManager

import (
	"encoding/binary"
	"os"
	"path/filepath"
)

type Metadata struct {
	filename string
	size int64
}

func Generate(files []string) ([]Metadata) {
	var metadatas []Metadata
	for _, filename := range files {
		fileHandle, err := os.OpenFile(filename, os.O_RDONLY, 0755)
		if err != nil {
			panic(err)
		}
		stat, statErr := fileHandle.Stat()
		if statErr != nil {
			panic(statErr)
		}
		metadatas = append(metadatas, Metadata {filepath.Base(filename), stat.Size()})
	}
	return metadatas
}

func Dump(metadatas []Metadata) []byte {
	var metadatasDump []byte
	metadatasDump = append(metadatasDump, byte(len(metadatas)))

	for _, metadata := range metadatas {
		fileSizeAsBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(fileSizeAsBytes, uint64(metadata.size))

		var metadataDump []byte
		metadataDump = append(metadataDump, byte(len(metadata.filename)))
		metadataDump = append(metadataDump, []byte(metadata.filename)...)
		metadataDump = append(metadataDump, fileSizeAsBytes...)
		metadatasDump = append(metadatasDump, metadataDump...)
	}
	return metadatasDump
}

// q fs f f f f s s s s s s s s
// 0 1  2 3 4 5 6 7 8 9 0 1 2 3

func Parse(dump []byte) []Metadata {
	var metadatas []Metadata

	metadatasQty := dump[0]
	parsedMetadatas := 0 

	idx := 1
	for {
		filenameSize := int(dump[idx])
		idx = idx + 1

		filename := dump[idx:idx+filenameSize]
		idx = idx + filenameSize

		fileSize := dump[idx:idx+8]
		idx = idx + 8

		metadatas = append(metadatas, Metadata{string(filename), int64(binary.LittleEndian.Uint64(fileSize))})

		parsedMetadatas++
		if parsedMetadatas == int(metadatasQty) {
			break
		}
	}
	return metadatas
}
