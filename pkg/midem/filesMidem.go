package midem

import (
	"fmt"
	"os"
)

func MixFiles(filesList []string) {
	inputFiles := filesList[1:]
	outputFile := filesList[0]

	fmt.Printf("Mix: %s\n", inputFiles)

	var managers []IDataBytesManager
	for _, file := range inputFiles {
		managers = append(managers, NewDataBytesFileManager(file))
	}

	dumpData := Dump(Generate(inputFiles))
	dumpData = append(dumpData, NewMixer(managers).Mix()...)

	if FileExists(outputFile) {
		panic("Output file exists")
	}

	NewDataBytesDumper(outputFile).Dump(dumpData)

	fmt.Printf("Files mixed into: %s\n", outputFile)
}

func DemixFiles(filesList []string) {
	for _, inputFile := range filesList {
		fmt.Printf("Demix: %s\n", inputFile)

		fileManager := NewDataBytesFileManager(inputFile)

		for _, demixData := range Demix(fileManager) {
			fmt.Printf("Writing file %s\n", demixData.Filename)
			os.WriteFile(demixData.Filename, demixData.Data, os.FileMode(demixData.Mode))
		}
	}
}
