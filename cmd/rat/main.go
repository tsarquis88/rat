package main

import (
	"fmt"
	"os"

	"github.com/tsarquis88/rat/pkg/cmdLineParser"
	"github.com/tsarquis88/rat/pkg/rat"
)

func main() {
	params := cmdLineParser.Parse(os.Args)

	if params.List {
		for _, files := range rat.List(params.InputFiles) {
			fmt.Println(files)
		}
	} else if params.Rat {
		rat.Rat(params.InputFiles, params.OutputFile)
	} else {
		rat.Derat(params.InputFiles, params.OutputFolder)
	}
}
