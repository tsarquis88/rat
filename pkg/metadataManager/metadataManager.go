package metadataManager

import (
	"encoding/binary"
	"os"
	"path/filepath"
)

type Metadata struct {
	Filename string
	Size int64
	Mode uint32
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
		metadatas = append(metadatas, Metadata {filepath.Base(filename), stat.Size(), uint32(stat.Mode())})
	}
	return metadatas
}

func Dump(metadatas []Metadata) []byte {
	var metadatasDump []byte
	metadatasDump = append(metadatasDump, byte(len(metadatas)))

	for _, metadata := range metadatas {
		fileSizeAsBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(fileSizeAsBytes, uint64(metadata.Size))
		fileModeAsBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(fileModeAsBytes, uint32(metadata.Mode))

		var metadataDump []byte
		metadataDump = append(metadataDump, byte(len(metadata.Filename)))
		metadataDump = append(metadataDump, []byte(metadata.Filename)...)
		metadataDump = append(metadataDump, fileSizeAsBytes...)
		metadataDump = append(metadataDump, fileModeAsBytes...)
		metadatasDump = append(metadatasDump, metadataDump...)
	}
	return metadatasDump
}

func Parse(dump []byte) ([]Metadata, int) {
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

		fileMode := dump[idx:idx+4]
		idx = idx + 4

		metadatas = append(metadatas, Metadata{string(filename), int64(binary.LittleEndian.Uint64(fileSize)), uint32(binary.LittleEndian.Uint32(fileMode))})

		parsedMetadatas++
		if parsedMetadatas == int(metadatasQty) {
			break
		}
	}
	return metadatas, idx
}
