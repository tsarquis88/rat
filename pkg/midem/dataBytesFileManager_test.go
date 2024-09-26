package midem

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type DataBytesFileManagerTestSuite struct {
	suite.Suite
	outputFolder string
	testFileA    string
	testFileB    string
}

func (suite *DataBytesFileManagerTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
	os.WriteFile(suite.testFileA, []byte("123"), 0755)
}

func (suite *DataBytesFileManagerTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

func (suite *DataBytesFileManagerTestSuite) TestNew() {
	assert.NotPanics(suite.T(), func() { NewDataBytesFileManager(suite.testFileA, "") }, "Should not panic")
}

func (suite *DataBytesFileManagerTestSuite) TestNewInexistantFile() {
	assert.Panics(suite.T(), func() { NewDataBytesFileManager(suite.testFileB, "") }, "Should panic")
}

func (suite *DataBytesFileManagerTestSuite) TestRead() {
	manager := NewDataBytesFileManager(suite.testFileA, "")

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
	const OutputFolder = "/tmp/DataBytesFileManagerTestSuite"
	var testSuite DataBytesFileManagerTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.testFileA = OutputFolder + "/fileA"
	testSuite.testFileB = OutputFolder + "/fileB"

	suite.Run(t, &testSuite)
}
