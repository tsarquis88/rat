package rat

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MetadataGeneratorTestSuite struct {
	suite.Suite
	outputFolder string
	fileName     string
	testFile     string
}

func (suite *MetadataGeneratorTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *MetadataGeneratorTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

// GenerateRatMetadata()

func (suite *MetadataGeneratorTestSuite) TestGenerateRatMetadata() {
	assert.Equal(suite.T(), RatMetadata{3, 1}, GenerateRatMetadata(3, 1))
}

// DumpRatMetadata()

func (suite *MetadataGeneratorTestSuite) TestDumpRatMetadata() {
	metadata := RatMetadata{5, 0}
	assert.Equal(suite.T(), []byte{5, 0}, DumpRatMetadata(metadata))
}

// ParseRatDump()

func (suite *MetadataGeneratorTestSuite) TestParseRatDump() {
	dataBytesManager := NewDataBytesManagerMock([]byte{1, 2})
	assert.Equal(suite.T(), RatMetadata{1, 2}, ParseRatDump(dataBytesManager))
}

// GenerateMetadata()

func (suite *MetadataGeneratorTestSuite) TestGenerateMetadata() {
	os.WriteFile(suite.testFile, []byte("12345"), 0755)

	expectedMetadata := Metadata{suite.fileName, 5, 0755}
	inputFiles := MetadataInput{suite.fileName, suite.outputFolder}
	metadata := GenerateMetadata(inputFiles)
	assert.Equal(suite.T(), expectedMetadata, metadata)
}

func (suite *MetadataGeneratorTestSuite) TestGenerateMetadataInexistantFile() {
	inputFile := MetadataInput{suite.testFile, suite.outputFolder}
	assert.Panics(suite.T(), func() { GenerateMetadata(inputFile) }, "Should panic")
}

// DumpMetadata()

func (suite *MetadataGeneratorTestSuite) TestDumpMetadata() {
	metadata := Metadata{suite.fileName, 5, 0755}

	expectedFileSize := make([]byte, 8)
	expectedFileSize[0] = 5

	expectedFileMode := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileMode, uint32(0755))

	var expectedDump []byte
	expectedDump = append(expectedDump, byte(len(suite.fileName)))
	expectedDump = append(expectedDump, []byte(suite.fileName)...)
	expectedDump = append(expectedDump, expectedFileSize...)
	expectedDump = append(expectedDump, expectedFileMode...)

	assert.Equal(suite.T(), expectedDump, DumpMetadata(metadata))
}

// ParseDump()

func (suite *MetadataGeneratorTestSuite) TestParseDump() {
	expectedMetadata := Metadata{suite.fileName, 5, 0755}

	expectedFileSize := make([]byte, 8)
	expectedFileSize[0] = 5
	expectedFileMode := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileMode, uint32(0755))

	var dump []byte
	dump = append(dump, byte(len(suite.fileName)))
	dump = append(dump, []byte(suite.fileName)...)
	dump = append(dump, expectedFileSize...)
	dump = append(dump, expectedFileMode...)
	dump = append(dump, byte(0))
	dump = append(dump, []byte("Unread data")...)

	dataBytesManager := NewDataBytesManagerMock(dump)

	assert.Equal(suite.T(), expectedMetadata, ParseDump(dataBytesManager))
}

// TestMetadataGeneratorTestSuite

func TestMetadataGeneratorTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/MetadataGeneratorTestSuite/"
	const FileName = "file.test"
	var testSuite MetadataGeneratorTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.fileName = FileName
	testSuite.testFile = filepath.Join(OutputFolder, FileName)
	suite.Run(t, &testSuite)
}
