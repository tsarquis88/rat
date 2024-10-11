package rat

import (
	"os"
	"path/filepath"
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

	expectedName := FillWith([]byte(suite.fileName), 0, NameLen)
	expectedMode := []byte{'0', '0', '0', '0', '7', '5', '5', 0}
	expectedSize := []byte{'0', '0', '0', '0', '0', '0', '0', '1', '7', '5', '0', 0}
	expectedTypeflag := uint8(RegulatFileType)
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
