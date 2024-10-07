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
	var filesToRat []MetadataInput
	for _, file := range inputFiles {
		originDir := filepath.Dir(strings.TrimSuffix(file, "/"))
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
		fmt.Printf("Rating file: (%s) %s... ", file.originDir, file.filename)

		metadata := GenerateMetadata(file)
		fileManager := NewDataBytesFileManager(filepath.Join(file.originDir, file.filename))
		fileData := GetDataFromManager(fileManager, metadata.Size)
		ratDump = append(ratDump, DumpMetadata(metadata)...)
		ratDump = append(ratDump, fileData...)
		fmt.Printf("Done.\n")
	}

	outExtension := filepath.Ext(outputFile)
	if outExtension == ".gz" {
		fmt.Print("Compressing... ")
		ratDump = GzipCompress(ratDump)
		fmt.Printf("Done.\n")
	}

	if FileExists(outputFile) {
		panic("Output file exists")
	}

	fmt.Printf("Writing output file: %s... ", outputFile)
	NewDataBytesDumper(outputFile, 438).Dump(ratDump)
	fmt.Printf("Done.\n")
}

func Derat(filesList []string, outputFolder string) {
	for _, inputFile := range filesList {
		var dataBytesManager IDataBytesManager
		inExtension := filepath.Ext(inputFile)
		if inExtension == ".gz" {
			fmt.Print("Decompressing... ")
			dataBytesManager = NewDataBytesSliceManager(GzipDecompress(FileRead(inputFile)))
			fmt.Printf("Done.\n")
		} else {
			dataBytesManager = NewDataBytesFileManager(inputFile)
		}

		ratMetadata := ParseRatDump(dataBytesManager)
		for i := 0; i < ratMetadata.filesQty; i++ {
			metadata := ParseDump(dataBytesManager)

			fmt.Printf("Derating file: %s... ", metadata.Filename)
			fileData := GetDataFromManager(dataBytesManager, metadata.Size)
			fmt.Printf("Done.\n")

			err := os.MkdirAll(filepath.Join(outputFolder, filepath.Dir(metadata.Filename)), 0755)
			if err != nil {
				panic(err)
			}

			fmt.Printf("Writing output file %s... ", filepath.Join(outputFolder, metadata.Filename))
			NewDataBytesDumper(filepath.Join(outputFolder, metadata.Filename), os.FileMode(metadata.Mode)).Dump(fileData)
			fmt.Printf("Done.\n")
		}
	}
}
