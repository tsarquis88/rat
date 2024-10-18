package rat

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HeaderTestSuite struct {
	suite.Suite
	outputFolder string
	fileName     string
	testFile     string
}

func (suite *HeaderTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *HeaderTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

// NewHeaderFromFile()

func (suite *HeaderTestSuite) TestNewHeaderFromFile() {
	const FileSize = 1000
	fileData := FillWith([]byte{}, 50, FileSize)
	os.WriteFile(suite.testFile, fileData, 0755)

	header := NewHeaderFromFile(suite.testFile)

	expectedName := FillWith([]byte(strings.TrimPrefix(suite.testFile, "/")), 0, NameLen)
	expectedMode := []byte{'0', '0', '0', '0', '7', '5', '5', 0}
	expectedSize := []byte{'0', '0', '0', '0', '0', '0', '0', '1', '7', '5', '0', 0}
	expectedTypeflag := uint8(RegularFileType)
	expectedMagic := []byte{'u', 's', 't', 'a', 'r', 32}
	expectedVersion := []byte{32, 0}

	assert.Equal(suite.T(), expectedName, header.name)
	assert.Equal(suite.T(), expectedMode, header.mode)
	assert.Equal(suite.T(), expectedSize, header.size)
	assert.Equal(suite.T(), expectedTypeflag, header.typeflag)
	assert.Equal(suite.T(), expectedMagic, header.magic)
	assert.Equal(suite.T(), expectedVersion, header.version)
}
func (suite *HeaderTestSuite) TestNewHeaderFromFileInexistantFile() {
	assert.Panics(suite.T(), func() { NewHeaderFromFile(suite.testFile) })
}

// NewHeaderFromDump()

func (suite *HeaderTestSuite) TestNewHeaderFromDump() {
	dumpData := []byte{102, 105, 108, 101, 46, 116, 101, 115, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48, 48, 48, 48, 55, 53, 53, 0, 48, 48, 48, 49, 55, 53, 48, 0, 48, 48, 48, 49, 55, 53, 48, 0, 48, 48, 48, 48, 48, 48, 48, 49, 55, 53, 48, 0, 49, 52, 55, 48, 50, 50, 48, 50, 48, 50, 55, 0, 0, 0, 0, 0, 0, 0, 0, 0, 48, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 117, 115, 116, 97, 114, 32, 32, 0, 116, 111, 109, 115, 97, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 116, 111, 109, 115, 97, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	header := NewHeaderFromDump(dumpData)

	expectedName := FillWith([]byte(suite.fileName), 0, NameLen)
	expectedMode := []byte{'0', '0', '0', '0', '7', '5', '5', 0}
	expectedSize := []byte{'0', '0', '0', '0', '0', '0', '0', '1', '7', '5', '0', 0}
	expectedTypeflag := uint8(RegularFileType)
	expectedMagic := []byte{'u', 's', 't', 'a', 'r', 32}
	expectedVersion := []byte{32, 0}

	assert.Equal(suite.T(), expectedName, header.name)
	assert.Equal(suite.T(), expectedMode, header.mode)
	assert.Equal(suite.T(), expectedSize, header.size)
	assert.Equal(suite.T(), expectedTypeflag, header.typeflag)
	assert.Equal(suite.T(), expectedMagic, header.magic)
	assert.Equal(suite.T(), expectedVersion, header.version)
}

// TestHeaderTestSuite

func TestHeaderTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/HeaderTestSuite/"
	const FileName = "file.test"
	var testSuite HeaderTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.fileName = FileName
	testSuite.testFile = filepath.Join(OutputFolder, FileName)
	suite.Run(t, &testSuite)
}
