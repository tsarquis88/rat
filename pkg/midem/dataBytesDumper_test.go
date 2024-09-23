package midem

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type DataBytesDumperTestSuite struct {
	suite.Suite
	outputFolder string
	testFileA    string
}

func (suite *DataBytesDumperTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *DataBytesDumperTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

func (suite *DataBytesDumperTestSuite) TestNew() {
	assert.NotPanics(suite.T(), func() { NewDataBytesDumper(suite.testFileA) }, "Should not panic")
}

func (suite *DataBytesDumperTestSuite) TestDump() {
	data := []byte{'1', '2', '3'}
	NewDataBytesDumper(suite.testFileA).Dump(data)

	file, err := os.OpenFile(suite.testFileA, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	buff := make([]byte, 3)
	_, readErr := file.Read(buff)
	if readErr != nil {
		panic(readErr)
	}

	assert.Equal(suite.T(), buff, data)
}

func TestDataBytesDumperTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/DataBytesDumperTestSuite"
	var testSuite DataBytesDumperTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.testFileA = OutputFolder + "/fileA"
	suite.Run(t, &testSuite)
}
