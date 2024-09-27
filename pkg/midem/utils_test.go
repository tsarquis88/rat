package midem

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilsTestSuite struct {
	suite.Suite
	outputFolder string
}

func (suite *UtilsTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *UtilsTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

// FileExists()

func (suite *UtilsTestSuite) TestFileExistsPositive() {
	testFile := suite.outputFolder + "test.file"
	os.WriteFile(testFile, []byte("12345"), 0755)

	assert.Equal(suite.T(), true, FileExists(testFile))
}

func (suite *UtilsTestSuite) TestFileExistsNegative() {
	testFile := suite.outputFolder + "test.file"

	assert.Equal(suite.T(), false, FileExists(testFile))
}

func (suite *UtilsTestSuite) TestFileExistsNegativeWithDir() {
	testFile := suite.outputFolder + "test_folder/" + "test.file"

	assert.Equal(suite.T(), false, FileExists(testFile))
}

// IsDir()

func (suite *UtilsTestSuite) TestIsDirPositive() {
	testFolder := suite.outputFolder + "test_folder"
	os.Mkdir(testFolder, 0755)

	assert.Equal(suite.T(), true, IsDir(testFolder))
}

func (suite *UtilsTestSuite) TestIsDirNegative() {
	testFolder := suite.outputFolder + "test_folder"
	os.WriteFile(testFolder, []byte("12345"), 0755)

	assert.Equal(suite.T(), false, IsDir(testFolder))
}

func (suite *UtilsTestSuite) TestIsDirNegativeInexistant() {
	testFolder := suite.outputFolder + "test_folder"

	assert.Panics(suite.T(), func() { IsDir(testFolder) }, "Should panic")
}

// GetFilesInDir()

func (suite *UtilsTestSuite) TestGetFilesInDirNoFiles() {
	assert.Equal(suite.T(), 0, len(GetFilesInDir(suite.outputFolder)))
}

func (suite *UtilsTestSuite) TestGetFilesInDirOneFile() {
	os.WriteFile(suite.outputFolder+"fileA", []byte("12345"), 0755)

	expectedFiles := []string{"fileA"}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder))
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFiles() {
	os.WriteFile(suite.outputFolder+"fileA", []byte("12345"), 0755)
	os.WriteFile(suite.outputFolder+"fileB", []byte("ABCDE"), 0755)
	os.WriteFile(suite.outputFolder+"fileC", []byte(",.-{}"), 0755)

	expectedFiles := []string{"fileA", "fileB", "fileC"}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder))
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFilesWithDir() {
	os.WriteFile(suite.outputFolder+"fileA", []byte("12345"), 0755)
	os.WriteFile(suite.outputFolder+"fileB", []byte("ABCDE"), 0755)
	os.WriteFile(suite.outputFolder+"fileC", []byte(",.-{}"), 0755)
	os.Mkdir(suite.outputFolder+"folder", 0755)

	expectedFiles := []string{"fileA", "fileB", "fileC", "folder"}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder))
}

func (suite *UtilsTestSuite) TestGetFilesInDirInexistantDir() {
	assert.Panics(suite.T(), func() { GetFilesInDir(suite.outputFolder + "folder") }, "Should panic")
}

// HashFile()

func (suite *UtilsTestSuite) TestHashFile() {
	os.WriteFile(suite.outputFolder+"fileA", []byte("12345"), 0755)

	const ExpectedHash = "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5"
	assert.Equal(suite.T(), ExpectedHash, fmt.Sprintf("%x", HashFile(suite.outputFolder+"fileA")))
}

func (suite *UtilsTestSuite) TestHashFileInexistantFile() {
	assert.Panics(suite.T(), func() { HashFile(suite.outputFolder + "file") }, "Should panic")
}

// TestSuite

func TestUtilsTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/UtilsTestSuite/"
	var testSuite UtilsTestSuite
	testSuite.outputFolder = OutputFolder
	suite.Run(t, &testSuite)
}
