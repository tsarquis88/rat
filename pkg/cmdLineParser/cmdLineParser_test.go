package cmdLineParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNoArguments(t *testing.T) {
	args := []string{"binaryName"}
	assert.Panics(t, func() { Parse(args) }, "Should panic")
}

func TestParseRat(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json"}
	rat, files := Parse(args)
	assert.Equal(t, true, rat, "Should be equal")
	assert.Equal(t, []string{"output.rat", "fileA.json"}, files, "Should be equal")
}

func TestParseRatMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json", "fileB.xml"}
	rat, files := Parse(args)
	assert.Equal(t, true, rat, "Should be equal")
	assert.Equal(t, []string{"output.rat", "fileA.json", "fileB.xml"}, files, "Should be equal")
}

func TestParseRatMissingArguments(t *testing.T) {
	args := []string{"binaryName", "fileA.json"}
	assert.Panics(t, func() { Parse(args) }, "Should panic")
}

func TestParseDerat(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x"}
	rat, files := Parse(args)
	assert.Equal(t, false, rat, "Should be equal")
	assert.Equal(t, []string{"fileA.rat"}, files, "Should be equal")
}

func TestParseDeratMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x", "fileB.rat"}
	rat, files := Parse(args)
	assert.Equal(t, false, rat, "Should be equal")
	assert.Equal(t, []string{"fileA.rat", "fileB.rat"}, files, "Should be equal")
}
