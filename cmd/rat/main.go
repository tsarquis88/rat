package main

import (
	"os"

	"github.com/tsarquis88/rat/pkg/cmdLineParser"
	"github.com/tsarquis88/rat/pkg/rat"
)

func main() {
	// Parse arguments
	performRat, files := cmdLineParser.Parse(os.Args)

	if performRat {
		rat.Rat(files[1:], files[0])
	} else {
		rat.Derat(files, "")
	}
}
