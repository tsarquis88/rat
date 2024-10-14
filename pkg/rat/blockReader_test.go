package rat

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BlockReaderTestSuite struct {
	suite.Suite
	outputFolder string
	testFileA    string
	testFileB    string
}

func (suite *BlockReaderTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *BlockReaderTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

// NewBlockReader()

func (suite *BlockReaderTestSuite) TestNewBlockReader() {
	os.WriteFile(suite.testFileA, []byte("123"), 0755)
	assert.NotPanics(suite.T(), func() { NewBlockReader(suite.testFileA, 512) }, "Should not panic")
}

func (suite *BlockReaderTestSuite) TestNewBlockReaderInexistantFile() {
	assert.Panics(suite.T(), func() { NewBlockReader(suite.testFileB, 512) }, "Should panic")
}

// ReadBlock()

func (suite *BlockReaderTestSuite) TestReadBlockExactSize() {
	const BlockSize = 512
	const FileDataSize = 512

	var fileData []byte
	for i := 0; i < FileDataSize; i++ {
		fileData = append(fileData, byte(i))
	}
	os.WriteFile(suite.testFileA, fileData, 0755)

	reader := NewBlockReader(suite.testFileA, BlockSize)
	block, more := reader.ReadBlock()

	assert.Equal(suite.T(), fileData, block)
	assert.True(suite.T(), more)

}

func (suite *BlockReaderTestSuite) TestReadBlockLessSize() {
	const BlockSize = 512
	const FileDataSize = 100

	var fileData []byte
	for i := 0; i < FileDataSize; i++ {
		fileData = append(fileData, byte(i))
	}
	os.WriteFile(suite.testFileA, fileData, 0755)

	reader := NewBlockReader(suite.testFileA, BlockSize)
	block, more := reader.ReadBlock()

	expectedBlock := FillWith(fileData, 0, BlockSize)
	assert.Equal(suite.T(), expectedBlock, block)
	assert.False(suite.T(), more)
}

func (suite *BlockReaderTestSuite) TestReadBlockMoreSize() {
	const BlockSize = 512
	const FileDataSize = 1000

	var fileData []byte
	for i := 0; i < FileDataSize; i++ {
		fileData = append(fileData, byte(i))
	}
	os.WriteFile(suite.testFileA, fileData, 0755)

	reader := NewBlockReader(suite.testFileA, BlockSize)
	block, more := reader.ReadBlock()

	expectedBlock := fileData[:BlockSize]
	assert.Equal(suite.T(), expectedBlock, block)
	assert.True(suite.T(), more)
}

func (suite *BlockReaderTestSuite) TestReadBlockMultipleBlocks() {
	const BlockSize = 512
	const FileDataSize = 1500

	var fileData []byte
	for i := 0; i < FileDataSize; i++ {
		fileData = append(fileData, byte(i))
	}
	os.WriteFile(suite.testFileA, fileData, 0755)

	reader := NewBlockReader(suite.testFileA, BlockSize)

	block, more := reader.ReadBlock()
	expectedBlock := fileData[:BlockSize]
	assert.Equal(suite.T(), expectedBlock, block)
	assert.True(suite.T(), more)

	block, more = reader.ReadBlock()
	expectedBlock = fileData[BlockSize : BlockSize*2]
	assert.Equal(suite.T(), expectedBlock, block)
	assert.True(suite.T(), more)

	block, more = reader.ReadBlock()
	expectedBlock = FillWith(fileData[BlockSize*2:], 0, BlockSize)
	assert.Equal(suite.T(), expectedBlock, block)
	assert.False(suite.T(), more)
}

// TestBlockReaderTestSuite()

func TestBlockReaderTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/BlockReaderTestSuite"
	var testSuite BlockReaderTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.testFileA = OutputFolder + "/fileA"
	testSuite.testFileB = OutputFolder + "/fileB"

	suite.Run(t, &testSuite)
}
