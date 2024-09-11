package dataBytesFileManager

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const OutputFolder = "/tmp/DataBytesFileManagerTestSuite"
const TestFileA = OutputFolder + "/fileA"
const TestFileB = OutputFolder + "/fileB" // Won't be created

type DataBytesFileManagerTestSuite struct {
	suite.Suite
}

func (suite *DataBytesFileManagerTestSuite) SetupTest() {
	os.Mkdir(OutputFolder, 0755)
	os.WriteFile(TestFileA, []byte("123"), 0755)
}

func (suite *DataBytesFileManagerTestSuite) TearDownTest() {
	os.RemoveAll(OutputFolder)
}

func (suite *DataBytesFileManagerTestSuite) TestNew() {
	assert.NotPanics(suite.T(), func() { NewDataBytesFileManager(TestFileA) }, "Should not panic")
}

func (suite *DataBytesFileManagerTestSuite) TestNewInexistantFile() {
	assert.Panics(suite.T(), func() { NewDataBytesFileManager(TestFileB) }, "Should panic")
}

func (suite *DataBytesFileManagerTestSuite) TestRead() {
	manager := NewDataBytesFileManager(TestFileA)

	dataA, bytesReadA := manager.Read(1)
	assert.Equal(suite.T(), 1, bytesReadA)
	assert.Equal(suite.T(), []byte{'1'}, dataA)

	dataB, bytesReadB := manager.Read(1)
	assert.Equal(suite.T(), 1, bytesReadB)
	assert.Equal(suite.T(), []byte{'2'}, dataB)

	dataC, bytesReadC := manager.Read(1)
	assert.Equal(suite.T(), 1, bytesReadC)
	assert.Equal(suite.T(), []byte{'3'}, dataC)

	_, bytesReadD := manager.Read(1)
	assert.Equal(suite.T(), 0, bytesReadD)
}

func TestDataBytesFileManagerTestSuite(t *testing.T) {
	suite.Run(t, new(DataBytesFileManagerTestSuite))
}
