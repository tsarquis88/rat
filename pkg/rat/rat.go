package rat

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const BlockSize = 512

const DirFileType = 53

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

func validateHeader(data []byte) bool {
	for _, value := range data {
		if value != 0 {
			return true
		}
	}
	return false
}

func getPaddingIndex(data []byte) uint {
	var i int
	for i = len(data) - 1; i >= 0; i = i - 2 {
		if data[i] != 0 || data[i-1] != 0 {
			break
		}
	}
	return uint(i)
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
			fmt.Printf("Done\n")
		} else {
			dataBytesManager = NewDataBytesFileManager(inputFile)
		}

		for {
			data, _ := dataBytesManager.Read(BlockSize)
			if !validateHeader(data) {
				break
			}

			headerRaw := NewHeaderRaw(data)
			header := NewHeader(headerRaw)

			if header.filetype == DirFileType {
				err := os.MkdirAll(filepath.Join(outputFolder, filepath.Dir(header.name)), 0755)
				if err != nil {
					panic(err)
				}
				continue
			}

			fmt.Printf("Reading file %s (%d)... ", header.name, header.size)
			var fileData []byte
			var bytesRead uint
			for {
				data, _ := dataBytesManager.Read(BlockSize)
				fileData = append(fileData, data...)
				bytesRead = bytesRead + BlockSize
				if bytesRead >= header.size {
					break
				}
			}
			fileData = fileData[:getPaddingIndex(fileData)]
			fmt.Printf("Done.\n")

			outputFile := filepath.Join(outputFolder, header.name)
			fmt.Printf("Writing output file %s... ", outputFile)
			NewDataBytesDumper(outputFile, os.FileMode(0664)).Dump(fileData)
			fmt.Printf("Done.\n")
		}
	}
}
