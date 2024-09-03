package main

import (
	"fmt"
	"os"

	"example.com/cmdLineParser"
	"example.com/dataBytesManager"
	"example.com/dataBytesFileManager"
	"example.com/dataBytesDumper"
	"example.com/metadataManager"
	"example.com/mixer"
	"example.com/demixer"
)

func mix(outputFile string, inputFiles []string) {
	fmt.Printf("Mix: %s\n", inputFiles)

	var managers []dataBytesManager.IDataBytesManager
	for _, file := range inputFiles {
		managers = append(managers, dataBytesFileManager.NewDataBytesFileManager(file))
	}

	dumpData := metadataManager.Dump(metadataManager.Generate(inputFiles))
	dumpData = append(dumpData, mixer.NewMixer(managers).Mix()...)
	dataBytesDumper.NewDataBytesDumper(outputFile).Dump(dumpData)

	fmt.Printf("Files mixed into: %s\n", outputFile)
}

func demix(inputFile string) {
	fmt.Printf("Demix: %s\n", inputFile)

	fileManager := dataBytesFileManager.NewDataBytesFileManager(inputFile)

	for _, demixData := range demixer.Demix(fileManager) {
		fmt.Printf("Writing file %s\n", demixData.Filename)
		os.WriteFile(demixData.Filename, demixData.Data, os.FileMode(demixData.Mode))
	}
}

func main() {
	// Parse arguments
	file, list := cmdLineParser.Parse(os.Args)

	if list == nil {
		demix(file)
	} else
	{
		mix(file, list)
	}
}
