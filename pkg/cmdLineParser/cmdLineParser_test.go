package cmdLineParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNoArguments(t *testing.T) {
	args := []string{"binaryName"}
	assert.Panics(t, func() { Parse(args) }, "Should panic")
}

func TestParseMix(t *testing.T) {
	args := []string{"binaryName", "output.mix", "fileA.json"}
	mix, files := Parse(args)
	assert.Equal(t, true, mix, "Should be equal")
	assert.Equal(t, []string{"output.mix", "fileA.json"}, files, "Should be equal")
}

func TestParseMixMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "output.mix", "fileA.json", "fileB.xml"}
	mix, files := Parse(args)
	assert.Equal(t, true, mix, "Should be equal")
	assert.Equal(t, []string{"output.mix", "fileA.json", "fileB.xml"}, files, "Should be equal")
}

func TestParseMixMissingArguments(t *testing.T) {
	args := []string{"binaryName", "fileA.json"}
	assert.Panics(t, func() { Parse(args) }, "Should panic")
}

func TestParseDemix(t *testing.T) {
	args := []string{"binaryName", "fileA.mix", "-x"}
	mix, files := Parse(args)
	assert.Equal(t, false, mix, "Should be equal")
	assert.Equal(t, []string{"fileA.mix"}, files, "Should be equal")
}

func TestParseDemixMultipleFiles(t *testing.T) {
	args := []string{"binaryName", "fileA.mix", "-x", "fileB.mix"}
	mix, files := Parse(args)
	assert.Equal(t, false, mix, "Should be equal")
	assert.Equal(t, []string{"fileA.mix", "fileB.mix"}, files, "Should be equal")
}
