package dataBytesDumper

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const OutputFolder = "/tmp/DataBytesDumperTestSuite"
const TestFileA = OutputFolder + "/fileA"

type DataBytesDumperTestSuite struct {
    suite.Suite
}

func (suite *DataBytesDumperTestSuite) SetupTest() {
	os.Mkdir(OutputFolder, 0755)
}

func (suite *DataBytesDumperTestSuite) TearDownTest() {
	os.RemoveAll(OutputFolder)
}

func (suite *DataBytesDumperTestSuite) TestNew() {
	assert.NotPanics(suite.T(), func() {NewDataBytesDumper(TestFileA)}, "Should not panic")
}

func (suite *DataBytesDumperTestSuite) TestDump() {
	data := []byte {'1', '2', '3'}
	NewDataBytesDumper(TestFileA).Dump(data)

	file, err := os.OpenFile(TestFileA, os.O_RDONLY, 0755)
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
    suite.Run(t, new(DataBytesDumperTestSuite))
}
