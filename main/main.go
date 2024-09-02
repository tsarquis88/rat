package main

import (
	"fmt"
	"os"

	"example.com/cmdLineParser"
	"example.com/dataBytesManager"
	"example.com/fileManager"
	"example.com/mixer"
)

func main() {
	// Parse arguments
	files, err := cmdLineParser.Parse(os.Args)
	if (err != nil)	{
		fmt.Println("Input files missing")
		os.Exit(0)
	}
	fmt.Printf("Files to be mixed: %s\n", files)

	var managers []dataBytesManager.IDataBytesManager
	for _, file := range files {
		fmt.Printf("Reading file: %s\n", file)
		managers = append(managers, fileManager.NewFileManager(file, false))
	}
	fmt.Printf("Files read: %d\n", len(managers))

	mixer.NewMixer(managers).Mix()
}
