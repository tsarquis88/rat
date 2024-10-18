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
	assert.Equal(t, Parameters{true, false, "", "output.rat", []string{"fileA.json"}, 1}, Parse(args))
}

func TestParseRatWithBlockingFactor(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json", "-b", "5"}
	assert.Equal(t, Parameters{true, false, "", "output.rat", []string{"fileA.json"}, 5}, Parse(args))
}

func TestParseRatMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json", "fileB.xml"}
	assert.Equal(t, Parameters{true, false, "", "output.rat", []string{"fileA.json", "fileB.xml"}, 1}, Parse(args))
}

func TestParseRatMissingArguments(t *testing.T) {
	args := []string{"binaryName", "fileA.json"}
	assert.Panics(t, func() { Parse(args) })
}

func TestParseDerat(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x"}
	assert.Equal(t, Parameters{false, false, "", "", []string{"fileA.rat"}, 1}, Parse(args))
}

func TestParseDeratMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x", "fileB.rat"}
	assert.Equal(t, Parameters{false, false, "", "", []string{"fileA.rat", "fileB.rat"}, 1}, Parse(args))
}

func TestParseRatWithOutputFolder(t *testing.T) {
	args := []string{"binaryName", "output.rat", "fileA.json", "-C", "outputFolder"}
	assert.Equal(t, Parameters{true, false, "outputFolder", "output.rat", []string{"fileA.json"}, 1}, Parse(args))
}

func TestParseDeratWithOutputFolder(t *testing.T) {
	args := []string{"binaryName", "fileA.rat", "-x", "fileB.rat", "-C", "outputFolder"}
	assert.Equal(t, Parameters{false, false, "outputFolder", "", []string{"fileA.rat", "fileB.rat"}, 1}, Parse(args))
}
