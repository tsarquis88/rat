package rat

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const CompressionRaw = 00
const CompressionGzip = 01

func GetDataFromManager(fileManager IDataBytesManager, size int64) []byte {
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

func Rat(inputFiles []string, outputFile string) {
	fmt.Printf("Rat: %s\n", inputFiles)

	var filesToRat []MetadataInput
	for _, file := range inputFiles {
		originDir := filepath.Dir(file)
		if IsDir(file) {
			dirFiles := GetFilesInDir(file, true)
			for _, dirFile := range dirFiles {
				filenameWithoutOriginDir := strings.TrimPrefix(dirFile, originDir)
				if filenameWithoutOriginDir[0] == '/' {
					filenameWithoutOriginDir = filenameWithoutOriginDir[1:]
				}
				filesToRat = append(filesToRat, MetadataInput{filenameWithoutOriginDir, originDir})
			}
		} else {
			filesToRat = append(filesToRat, MetadataInput{filepath.Base(file), originDir})
		}
	}

	ratDump := DumpRatMetadata(GenerateRatMetadata(len(filesToRat)))

	for _, file := range filesToRat {
		fmt.Printf("File to rat: (%s) %s\n", file.originDir, file.filename)

		metadata := GenerateMetadata(file)
		fileManager := NewDataBytesFileManager(filepath.Join(file.originDir, file.filename))
		fileData := GetDataFromManager(fileManager, metadata.Size)
		ratDump = append(ratDump, DumpMetadata(metadata)...)
		ratDump = append(ratDump, fileData...)
	}

	outExtension := filepath.Ext(outputFile)
	if outExtension == ".gz" {
		fmt.Print("Compressing... ")
		ratDump = GzipCompress(ratDump)
		fmt.Printf("Done\n")
	}

	if FileExists(outputFile) {
		panic("Output file exists")
	}

	NewDataBytesDumper(outputFile, 438).Dump(ratDump)
	fmt.Printf("Files rated into: %s\n", outputFile)
}

func Derat(filesList []string, outputFolder string) {
	for _, inputFile := range filesList {
		fmt.Printf("Derat: %s\n", inputFile)

		var dataBytesManager IDataBytesManager
		inExtension := filepath.Ext(inputFile)
		if inExtension == ".gz" {
			fmt.Print("Decompressing... ")
			dataBytesManager = NewDataBytesSliceManager(GzipDecompress(FileRead(inputFile)))
			fmt.Printf("Done\n")
		} else {
			dataBytesManager = NewDataBytesFileManager(inputFile)
		}

		ratMetadata := ParseRatDump(dataBytesManager)
		fmt.Printf("Files found: %d\n", ratMetadata.filesQty)

		for i := 0; i < ratMetadata.filesQty; i++ {
			metadata := ParseDump(dataBytesManager)
			fmt.Printf("File found: %s (size = %d, mode = %d)\n", metadata.Filename, metadata.Size, metadata.Mode)

			fileData := GetDataFromManager(dataBytesManager, metadata.Size)
			err := os.MkdirAll(filepath.Join(outputFolder, filepath.Dir(metadata.Filename)), 0755)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Writing file %s\n", filepath.Join(outputFolder, metadata.Filename))
			NewDataBytesDumper(filepath.Join(outputFolder, metadata.Filename), os.FileMode(metadata.Mode)).Dump(fileData)
		}
	}
}
