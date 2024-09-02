package fileManager

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const OutputFolder = "/tmp/FileManagerTestSuite"
const TestFileA = OutputFolder + "/fileA"
const TestFileB = OutputFolder + "/fileB" // Won't be created

type FileManagerTestSuite struct {
    suite.Suite
}

func (suite *FileManagerTestSuite) SetupTest() {
	os.Mkdir(OutputFolder, 0755)
	os.WriteFile(TestFileA, []byte("123"), 0755)
}

func (suite *FileManagerTestSuite) TearDownTest() {
	os.RemoveAll(OutputFolder)
}

func (suite *FileManagerTestSuite) TestNewReadMode() {
	assert.NotPanics(suite.T(), func() {New(TestFileA, false)}, "Should not panic")
}

func (suite *FileManagerTestSuite) TestNewReadModeInexistantFile() {
	assert.Panics(suite.T(), func() {New(TestFileB, false)}, "Should panic")
}

func (suite *FileManagerTestSuite) TestNewWriteMode() {
	assert.NotPanics(suite.T(), func() {New(TestFileA, true)}, "Should not panic")
}

func (suite *FileManagerTestSuite) TestNewWriteModeNewFile() {
	assert.NotPanics(suite.T(), func() {New(TestFileB, true)}, "Should not panic")
}

func (suite *FileManagerTestSuite) TestRead() {
	manager := New(TestFileA, false)

	dataA, bytesReadA := manager.Read()
	assert.Equal(suite.T(), bytesReadA, 1)
	assert.Equal(suite.T(), dataA, uint8('1'))

	dataB, bytesReadB := manager.Read()
	assert.Equal(suite.T(), bytesReadB, 1)
	assert.Equal(suite.T(), dataB, uint8('2'))	

	dataC, bytesReadC := manager.Read()
	assert.Equal(suite.T(), bytesReadC, 1)
	assert.Equal(suite.T(), dataC, uint8('3'))

	_, bytesReadD := manager.Read()
	assert.Equal(suite.T(), bytesReadD, 0)
}

func (suite *FileManagerTestSuite) TestWrite() {
	manager := New(TestFileB, true)

	assert.Equal(suite.T(), manager.write(byte('X')), 1)
	assert.Equal(suite.T(), manager.write(byte('Y')), 1)
	assert.Equal(suite.T(), manager.write(byte('Z')), 1)

	file, err := os.OpenFile(TestFileB, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	buff := make([]byte, 3)
	_, readErr := file.Read(buff)
	if readErr != nil {
		panic(readErr)
	}

	assert.Equal(suite.T(), len(buff), 3)
	assert.Equal(suite.T(), buff, []byte("XYZ"))
}

func TestFileManagerTestSuite(t *testing.T) {
    suite.Run(t, new(FileManagerTestSuite))
}
