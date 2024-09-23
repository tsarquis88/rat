package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/tsarquis88/file_mixer/pkg/cmdLineParser"
	"github.com/tsarquis88/file_mixer/pkg/midem"
)

func mix(filesList []string) {
	inputFiles := filesList[1:]
	outputFile := filesList[0]

	fmt.Printf("Mix: %s\n", inputFiles)

	var managers []midem.IDataBytesManager
	for _, file := range inputFiles {
		managers = append(managers, midem.NewDataBytesFileManager(file))
	}

	dumpData := midem.Dump(midem.Generate(inputFiles))
	dumpData = append(dumpData, midem.NewMixer(managers).Mix()...)

	if _, err := os.Stat(outputFile); !errors.Is(err, os.ErrNotExist) {
		panic("Output file exists")
	}

	midem.NewDataBytesDumper(outputFile).Dump(dumpData)

	fmt.Printf("Files mixed into: %s\n", outputFile)
}

func demix(filesList []string) {
	for _, inputFile := range filesList {
		fmt.Printf("Demix: %s\n", inputFile)

		fileManager := midem.NewDataBytesFileManager(inputFile)

		for _, demixData := range midem.Demix(fileManager) {
			fmt.Printf("Writing file %s\n", demixData.Filename)
			os.WriteFile(demixData.Filename, demixData.Data, os.FileMode(demixData.Mode))
		}
	}
}

func main() {
	// Parse arguments
	performMix, files := cmdLineParser.Parse(os.Args)

	if performMix {
		mix(files)
	} else {
		demix(files)
	}
}
