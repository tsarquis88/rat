package cmdLineParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNoArguments(t *testing.T) {
	args := []string{"binaryName"}
	assert.Panics(t, func() { Parse(args) })
}

func TestParseRat(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json"}
	rat, outputFolder, files := Parse(args)
	assert.Equal(t, true, rat)
	assert.Equal(t, "", outputFolder)
	assert.Equal(t, []string{"output.rat", "fileA.json"}, files)
}

func TestParseRatMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json", "fileB.xml"}
	rat, outputFolder, files := Parse(args)
	assert.Equal(t, true, rat)
	assert.Equal(t, "", outputFolder)
	assert.Equal(t, []string{"output.rat", "fileA.json", "fileB.xml"}, files)
}

func TestParseRatMissingArguments(t *testing.T) {
	args := []string{"binaryName", "fileA.json"}
	assert.Panics(t, func() { Parse(args) })
}

func TestParseDerat(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x"}
	rat, outputFolder, files := Parse(args)
	assert.Equal(t, false, rat)
	assert.Equal(t, "", outputFolder)
	assert.Equal(t, []string{"fileA.rat"}, files)
}

func TestParseDeratMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x", "fileB.rat"}
	rat, outputFolder, files := Parse(args)
	assert.Equal(t, false, rat)
	assert.Equal(t, "", outputFolder)
	assert.Equal(t, []string{"fileA.rat", "fileB.rat"}, files)
}

func TestParseRatWithOutputFolder(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json", "-C", "outputFolder"}
	rat, outputFolder, files := Parse(args)
	assert.Equal(t, true, rat)
	assert.Equal(t, "outputFolder", outputFolder)
	assert.Equal(t, []string{"output.rat", "fileA.json"}, files)
}

func TestParseDeratWithOutputFolder(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x", "fileB.rat", "-C", "outputFolder"}
	rat, outputFolder, files := Parse(args)
	assert.Equal(t, false, rat)
	assert.Equal(t, "outputFolder", outputFolder)
	assert.Equal(t, []string{"fileA.rat", "fileB.rat"}, files)
}
