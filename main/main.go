package main

import (
	"fmt"
	"os"
	"example.com/cmdLineParser"
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

	mixer.New(files).Mix()
}
