package cmdLineParser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseNoFiles(t *testing.T) {
    args := []string{"binaryName"}
	result, err := Parse(args)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestParseOneFile(t *testing.T) {
    args := []string{"binaryName", "filenameA"}
	result, err := Parse(args)
	assert.Equal(t, result, []string{"filenameA"}, "They should be equal")
	assert.Nil(t, err)
}

func TestParseMultipleFiles(t *testing.T) {
    args := []string{"binaryName", "filenameA", "filenameB", "filenameC"}
	result, err := Parse(args)
	assert.Equal(t, result, []string{"filenameA", "filenameB", "filenameC"}, "They should be equal")
	assert.Nil(t, err)
}
