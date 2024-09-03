package cmdLineParser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseNoFiles(t *testing.T) {
    args := []string{"binaryName"}
	output, input, err := Parse(args)
	assert.Empty(t, output)
	assert.Nil(t, input)
	assert.NotNil(t, err)
}

func TestParseOneFile(t *testing.T) {
    args := []string{"binaryName", "outputFile", "filenameA"}
	output, input, err := Parse(args)
	assert.Equal(t, output, "outputFile", "They should be equal")
	assert.Equal(t, input, []string{"filenameA"}, "They should be equal")
	assert.Nil(t, err)
}

func TestParseMultipleFiles(t *testing.T) {
    args := []string{"binaryName", "outputFile", "filenameA", "filenameB", "filenameC"}
	output, input, err := Parse(args)
	assert.Equal(t, output, "outputFile", "They should be equal")
	assert.Equal(t, input, []string{"filenameA", "filenameB", "filenameC"}, "They should be equal")
	assert.Nil(t, err)
}
