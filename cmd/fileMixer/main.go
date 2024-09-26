package main

import (
	"os"

	"github.com/tsarquis88/file_mixer/pkg/cmdLineParser"
	"github.com/tsarquis88/file_mixer/pkg/midem"
)

func main() {
	// Parse arguments
	performMix, files := cmdLineParser.Parse(os.Args)

	if performMix {
		midem.MixFiles(files[1:], files[0])
	} else {
		midem.DemixFiles(files, "")
	}
}
