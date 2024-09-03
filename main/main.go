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
)

func main() {
	// Parse arguments
	output, files, err := cmdLineParser.Parse(os.Args)
	if (err != nil)	{
		fmt.Println("Input files missing")
		os.Exit(0)
	}
	fmt.Printf("Files to be mixed: %s\n", files)

	var managers []dataBytesManager.IDataBytesManager
	for _, file := range files {
		fmt.Printf("Reading file: %s\n", file)
		managers = append(managers, dataBytesFileManager.NewDataBytesFileManager(file))
	}
	fmt.Printf("Files read: %d\n", len(managers))

	fmt.Print("Generating metadata...\n")
	dumpData := metadataManager.Dump(metadataManager.Generate(files))

	fmt.Print("Mixing...\n")
	dumpData = append(dumpData, mixer.NewMixer(managers).Mix()...)

	fmt.Printf("Dumping to %s\n", output)
	dataBytesDumper.NewDataBytesDumper(output).Dump(dumpData)
}
