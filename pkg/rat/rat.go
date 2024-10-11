package rat

import (
	"fmt"
	"os"
	"path/filepath"
)

const BlockSize = 512

const RegulatFileType = 48
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

func trimPadding(value []byte) string {
	var paddingIdx int
	for paddingIdx = 0; paddingIdx < len(value); paddingIdx++ {
		if value[paddingIdx] == 0 {
			break
		}
	}
	return string(value[:paddingIdx])
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

	if FileExists(outputFile) {
		panic("Output file exists")
	}
	outputDumper := NewDataBytesDumper(outputFile, 438)

	for _, file := range inputFiles {

		if IsDir(file) {
			panic("Folder rating not yet supported")
		}

		header := NewHeaderFromFile(file)
		outputDumper.Dump(header.Dump())

		dataManager := NewDataBytesFileManager(file)
		missingBytes := OctalToDecimal(header.size, 11)
		for {
			fileData, _ := dataManager.Read(BlockSize)
			outputDumper.Dump(FillWith(fileData, 0, BlockSize))

			if int(missingBytes-BlockSize) <= 0 {
				break
			}
			missingBytes -= BlockSize
		}
	}

	fileSuffix := FillWith([]byte{}, 0, BlockSize)
	for i := 0; i < 18; i++ {
		outputDumper.Dump(fileSuffix)
	}

	// outExtension := filepath.Ext(outputFile)
	// if outExtension == ".gz" {
	// 	fmt.Print("Compressing... ")
	// 	ratDump = GzipCompress(ratDump)
	// 	fmt.Printf("Done.\n")
	// }
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

			header := NewHeaderFromDump(data)
			fmt.Println(header.ToString())

			filename := trimPadding(header.name)
			if header.typeflag == DirFileType {
				err := os.MkdirAll(filepath.Join(outputFolder, filepath.Dir(filename)), 0755)
				if err != nil {
					panic(err)
				}
				continue
			}

			size := OctalToDecimal(header.size, 11)
			fmt.Printf("Reading file %s (%d)... ", header.name, size)
			var fileData []byte
			var bytesRead uint
			for {
				data, _ := dataBytesManager.Read(BlockSize)
				fileData = append(fileData, data...)
				bytesRead = bytesRead + BlockSize
				if bytesRead >= size {
					break
				}
			}
			fileData = fileData[:getPaddingIndex(fileData)]
			fmt.Printf("Done.\n")

			outputFile := filepath.Join(outputFolder, filename)
			fmt.Printf("Writing output file %s... ", outputFile)
			err := os.MkdirAll(filepath.Dir(filename), 0755)
			if err != nil {
				panic(err)
			}
			NewDataBytesDumper(outputFile, os.FileMode(0664)).Dump(fileData)
			fmt.Printf("Done.\n")
		}
	}
}
