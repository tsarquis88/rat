package cmdLineParser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseNoArguments(t *testing.T) {
    args := []string{"binaryName"}
	assert.Panics(t, func() {Parse(args)}, "Should panic")
}

func TestParseDemix(t *testing.T) {
    args := []string{"binaryName", "inputFile.mix"}
	inputFile, list := Parse(args)
	assert.Equal(t, inputFile, "inputFile.mix", "Should be equal")
	assert.Nil(t, list, "Should be null")
}

func TestParseMix(t *testing.T) {
    args := []string{"binaryName", "outputFile.mix", "filenameA", "filenameB", "filenameC"}
	outputFile, list := Parse(args)
	assert.Equal(t, outputFile, "outputFile.mix", "Should be equal")
	assert.Equal(t, list, []string{"filenameA", "filenameB", "filenameC"}, "Should be equal")
}
