package main

import (
	"fmt"
	"os"

	"github.com/tsarquis88/rat/pkg/cmdLineParser"
	"github.com/tsarquis88/rat/pkg/rat"
)

func main() {
	params := cmdLineParser.Parse(os.Args)

	ratDerat := rat.NewRatDerat(params.BlockingFactor)

	if params.List {
		for _, files := range ratDerat.List(params.InputFiles) {
			fmt.Println(files)
		}
	} else if params.Rat {
		ratDerat.Rat(params.InputFiles, params.OutputFile)
	} else {
		ratDerat.Derat(params.InputFiles, params.OutputFolder)
	}
}
