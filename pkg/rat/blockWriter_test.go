package rat

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BlockWriterTestSuite struct {
	suite.Suite
	outputFolder string
	testFileA    string
}

func (suite *BlockWriterTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *BlockWriterTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

// NewBlockWriter()

func (suite *BlockWriterTestSuite) TestNewBlockWriter() {
	assert.NotPanics(suite.T(), func() { NewBlockWriter(suite.testFileA, 438) }, "Should not panic")
}

func (suite *BlockWriterTestSuite) TestNewBlockWriterInexistantFolder() {
	assert.Panics(suite.T(), func() { NewBlockWriter(filepath.Join(suite.outputFolder, "other_folder", "test_file.json"), 438) }, "Should panic")
}

// WriteBlock()

func (suite *BlockWriterTestSuite) TestWriteBlock() {
	const BlockSize = 512

	var fileData []byte
	for i := 0; i < BlockSize; i++ {
		fileData = append(fileData, byte(i))
	}
	writer := NewBlockWriter(suite.testFileA, 438)
	writer.WriteBlock(fileData)

	file, err := os.OpenFile(suite.testFileA, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	buff := make([]byte, BlockSize)
	_, readErr := file.Read(buff)
	if readErr != nil {
		panic(readErr)
	}

	assert.Equal(suite.T(), buff, fileData)
}

func (suite *BlockWriterTestSuite) TestWriteBlockMultipleWrites() {
	const BlockSize = 512

	var fileData []byte
	for i := 0; i < BlockSize; i++ {
		fileData = append(fileData, byte(i))
	}
	writer := NewBlockWriter(suite.testFileA, 438)
	writer.WriteBlock(fileData)
	writer.WriteBlock(fileData)

	file, err := os.OpenFile(suite.testFileA, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	buff := make([]byte, BlockSize*2)
	_, readErr := file.Read(buff)
	if readErr != nil {
		panic(readErr)
	}

	expectedData := fileData
	expectedData = append(expectedData, fileData...)
	assert.Equal(suite.T(), buff, expectedData)
}

// TestBlockWriterTestSuite()

func TestBlockWriterTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/BlockWriterTestSuite"
	var testSuite BlockWriterTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.testFileA = OutputFolder + "/fileA"
	suite.Run(t, &testSuite)
}
