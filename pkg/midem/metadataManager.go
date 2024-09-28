package midem

import (
	"encoding/binary"
	"os"
	"path/filepath"
)

type Metadata struct {
	Filename string
	Size     int64
	Mode     uint32
}

type MetadataInput struct {
	filename  string
	originDir string
}

func GenerateMetadata(file MetadataInput) Metadata {
	fileHandle, err := os.OpenFile(filepath.Join(file.originDir, file.filename), os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	stat, statErr := fileHandle.Stat()
	if statErr != nil {
		panic(statErr)
	}
	return Metadata{file.filename, stat.Size(), uint32(stat.Mode())}
}

func DumpMetadata(metadata Metadata) []byte {
	fileSizeAsBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(fileSizeAsBytes, uint64(metadata.Size))
	fileModeAsBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(fileModeAsBytes, uint32(metadata.Mode))

	var metadataDump []byte
	metadataDump = append(metadataDump, byte(len(metadata.Filename)))
	metadataDump = append(metadataDump, []byte(metadata.Filename)...)
	metadataDump = append(metadataDump, fileSizeAsBytes...)
	metadataDump = append(metadataDump, fileModeAsBytes...)
	return metadataDump
}

func ParseMetadata(dataBytesSource IDataBytesManager) Metadata {
	filenameSize, _ := dataBytesSource.Read(1)
	filename, _ := dataBytesSource.Read(uint(filenameSize[0]))
	fileSize, _ := dataBytesSource.Read(8)
	fileMode, _ := dataBytesSource.Read(4)
	return Metadata{string(filename), int64(binary.LittleEndian.Uint64(fileSize)), uint32(binary.LittleEndian.Uint32(fileMode))}
}
