package rat

import (
	"os"
	"path/filepath"
	"strings"
)

const DefaultBlockSize = 512

const RegularFileType = '0'
const DirFileType = '5'

type RatDerat struct {
	blockSize uint
}

func NewRatDerat(blockingFactor uint) RatDerat {
	return RatDerat{DefaultBlockSize * blockingFactor}
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

func isBlockEmpty(data []byte) bool {
	for _, value := range data {
		if value != 0 {
			return false
		}
	}
	return true
}

func convertMode(value []byte) uint32 {
	var mode uint32
	mode += (uint32(value[6]) - 48)
	mode += (uint32(value[5]) - 48) << 3
	mode += (uint32(value[4]) - 48) << 6
	return mode
}

func (ratDerat *RatDerat) Rat(inputFiles []string, outputFile string) {
	if FileExists(outputFile) {
		panic("Output file exists")
	}
	outExtension := filepath.Ext(outputFile)
	if outExtension == ".gz" {
		panic("Rat compression not yet supported")
	}
	writer := NewBlockWriter(outputFile, 438)

	var filesToRat []string
	for _, file := range inputFiles {
		if IsDir(file) {
			filesToRat = append(filesToRat, GetFilesInDir(file, true, true)...)
		} else {
			filesToRat = append(filesToRat, file)
		}
	}

	for _, file := range filesToRat {
		header := NewHeaderFromFile(file, ratDerat.blockSize)
		writer.WriteBlock(header.Dump(ratDerat.blockSize))

		if header.typeflag == DirFileType {
			continue
		}

		blockReader := NewBlockReader(file, ratDerat.blockSize)
		for {
			block, more := blockReader.ReadBlock()
			writer.WriteBlock(block)
			if !more {
				break
			}
		}
	}
}

func (ratDerat *RatDerat) Derat(filesList []string, outputFolder string) {
	for _, inputFile := range filesList {
		inExtension := filepath.Ext(inputFile)
		if inExtension == ".gz" {
			panic("Rat compression not yet supported")
		}

		reader := NewBlockReader(inputFile, ratDerat.blockSize)
		for {
			headerBlock, _ := reader.ReadBlock()
			if isBlockEmpty(headerBlock) {
				break
			}
			header := NewHeaderFromDump(headerBlock)

			filename := trimPadding(header.name)
			if header.typeflag == DirFileType {
				err := os.MkdirAll(filepath.Join(outputFolder, filepath.Dir(filename)), 0755)
				if err != nil {
					panic(err)
				}
				continue
			}

			size := OctalToDecimal(header.size, 11)
			outputFile := filepath.Join(outputFolder, filename)
			err := os.MkdirAll(filepath.Dir(outputFile), 0755)
			if err != nil {
				panic(err)
			}
			writer := NewBlockWriter(outputFile, os.FileMode(convertMode(header.mode)))

			var bytesRead uint
			for {
				data, _ := reader.ReadBlock()
				bytesRead = bytesRead + ratDerat.blockSize
				if bytesRead >= size {
					writer.WriteBlock(data[:ratDerat.blockSize-(bytesRead-size)]) // Trim remaining zeroes.
					break
				}
				writer.WriteBlock(data)
			}
		}
	}
}

func (ratDerat *RatDerat) List(filesList []string) map[string]string {
	filesMap := make(map[string]string)
	for _, inputFile := range filesList {
		inExtension := filepath.Ext(inputFile)
		if inExtension == ".gz" {
			panic("Rat compression not yet supported")
		}

		filesMap[inputFile] = ""

		reader := NewBlockReader(inputFile, ratDerat.blockSize)
		for {
			headerBlock, _ := reader.ReadBlock()
			if isBlockEmpty(headerBlock) {
				if filesMap[inputFile] != "" {
					filesMap[inputFile] = strings.TrimSuffix(filesMap[inputFile], "\n")
				}
				break
			}
			header := NewHeaderFromDump(headerBlock)
			size := OctalToDecimal(header.size, 11)
			filesMap[inputFile] += trimPadding(header.name) + "\n"

			if header.typeflag != DirFileType {
				newOffset := uint(0)
				for {
					newOffset += ratDerat.blockSize
					if newOffset >= size {
						break
					}
				}
				reader.AdjustOffset(newOffset)
			}
		}
	}
	return filesMap
}
