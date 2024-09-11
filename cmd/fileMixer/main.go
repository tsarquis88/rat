package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/tsarquis88/file_mixer/pkg/cmdLineParser"
	"github.com/tsarquis88/file_mixer/pkg/dataBytesDumper"
	"github.com/tsarquis88/file_mixer/pkg/dataBytesFileManager"
	"github.com/tsarquis88/file_mixer/pkg/dataBytesManager"
	"github.com/tsarquis88/file_mixer/pkg/demixer"
	"github.com/tsarquis88/file_mixer/pkg/metadataManager"
	"github.com/tsarquis88/file_mixer/pkg/mixer"
)

func mix(filesList []string) {
	inputFiles := filesList[1:]
	outputFile := filesList[0]

	fmt.Printf("Mix: %s\n", inputFiles)

	var managers []dataBytesManager.IDataBytesManager
	for _, file := range inputFiles {
		managers = append(managers, dataBytesFileManager.NewDataBytesFileManager(file))
	}

	dumpData := metadataManager.Dump(metadataManager.Generate(inputFiles))
	dumpData = append(dumpData, mixer.NewMixer(managers).Mix()...)

	if _, err := os.Stat(outputFile); !errors.Is(err, os.ErrNotExist) {
		panic("Output file exists")
	}

	dataBytesDumper.NewDataBytesDumper(outputFile).Dump(dumpData)

	fmt.Printf("Files mixed into: %s\n", outputFile)
}

func demix(filesList []string) {
	for _, inputFile := range filesList {
		fmt.Printf("Demix: %s\n", inputFile)

		fileManager := dataBytesFileManager.NewDataBytesFileManager(inputFile)

		for _, demixData := range demixer.Demix(fileManager) {
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
