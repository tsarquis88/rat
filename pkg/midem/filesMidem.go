package midem

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetDataFromManager(fileManager DataBytesFileManager, size int64) []byte {
	const ReadSize = 1 << (10 * 2) // 1MB
	var data []byte
	if size <= ReadSize {
		buffer, _ := fileManager.Read(uint(size))
		data = append(data, buffer...)
	} else {
		missingBytes := size
		for {
			bytesToRead := min(ReadSize, missingBytes)
			buffer, _ := fileManager.Read(uint(bytesToRead))
			data = append(data, buffer...)
			missingBytes -= int64(bytesToRead)

			if missingBytes <= 0 {
				break
			}
		}
	}
	return data
}

func MixFiles(inputFiles []string, outputFile string) {
	fmt.Printf("Mix: %s\n", inputFiles)

	var filesToMix []MetadataInput
	for _, file := range inputFiles {
		originDir := filepath.Dir(file)
		if IsDir(file) {
			dirFiles := GetFilesInDir(file, true)
			for _, dirFile := range dirFiles {
				filenameWithoutOriginDir := strings.TrimPrefix(dirFile, originDir)
				if filenameWithoutOriginDir[0] == '/' {
					filenameWithoutOriginDir = filenameWithoutOriginDir[1:]
				}
				filesToMix = append(filesToMix, MetadataInput{filenameWithoutOriginDir, originDir})
			}
		} else {
			filesToMix = append(filesToMix, MetadataInput{filepath.Base(file), originDir})
		}
	}

	var mixDump []byte
	mixDump = append(mixDump, byte(len(filesToMix)))

	for _, file := range filesToMix {
		fmt.Printf("File to mix: (%s) %s\n", file.originDir, file.filename)

		metadata := GenerateMetadata(file)
		mixDump = append(mixDump, DumpMetadata(metadata)...)
		fileManager := NewDataBytesFileManager(filepath.Join(file.originDir, file.filename))
		mixDump = append(mixDump, GetDataFromManager(fileManager, metadata.Size)...)
	}

	if FileExists(outputFile) {
		panic("Output file exists")
	}

	NewDataBytesDumper(outputFile).Dump(mixDump)
	fmt.Printf("Files mixed into: %s\n", outputFile)
}

func DemixFiles(filesList []string, outputFolder string) {
	for _, inputFile := range filesList {
		fmt.Printf("Demix: %s\n", inputFile)

		fileManager := NewDataBytesFileManager(inputFile)

		metadatasQtyRaw, _ := fileManager.Read(1)
		metadatasQty := int(metadatasQtyRaw[0])
		fmt.Printf("Files found: %d\n", metadatasQty)

		for i := 0; i < metadatasQty; i++ {
			metadata := ParseMetadata(fileManager)
			fmt.Printf("File found: %s (size = %d, mode = %d)\n", metadata.Filename, metadata.Size, metadata.Mode)

			fileData := GetDataFromManager(fileManager, metadata.Size)

			err := os.MkdirAll(filepath.Join(outputFolder, filepath.Dir(metadata.Filename)), 0755)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Writing file %s\n", filepath.Join(outputFolder, metadata.Filename))
			err = os.WriteFile(filepath.Join(outputFolder, metadata.Filename), fileData, os.FileMode(metadata.Mode))
			if err != nil {
				panic(err)
			}
		}
	}
}
