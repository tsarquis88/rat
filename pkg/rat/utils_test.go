package rat

import (
	"fmt"
	"os"
	"path/filepath"
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
	assert.Equal(suite.T(), 0, len(GetFilesInDir(suite.outputFolder, false)))
}

func (suite *UtilsTestSuite) TestGetFilesInDirOneFile() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	os.WriteFile(filePathA, []byte("12345"), 0755)

	expectedFiles := []string{filePathA}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false))
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFiles() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	filePathB := filepath.Join(suite.outputFolder, "fileB")
	filePathC := filepath.Join(suite.outputFolder, "fileC")
	os.WriteFile(filePathA, []byte("12345"), 0755)
	os.WriteFile(filePathB, []byte("ABCDE"), 0755)
	os.WriteFile(filePathC, []byte(",.-{}"), 0755)

	expectedFiles := []string{filePathA, filePathB, filePathC}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false))
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFilesWithDir() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	filePathB := filepath.Join(suite.outputFolder, "fileB")
	filePathC := filepath.Join(suite.outputFolder, "fileC")
	folderPath := filepath.Join(suite.outputFolder, "folder")
	filePathD := filepath.Join(folderPath, "fileD")
	os.WriteFile(filePathA, []byte("12345"), 0755)
	os.WriteFile(filePathB, []byte("ABCDE"), 0755)
	os.WriteFile(filePathC, []byte(",.-{}"), 0755)
	os.Mkdir(folderPath, 0755)
	os.WriteFile(filePathD, []byte("___"), 0755)

	expectedFiles := []string{filePathA, filePathB, filePathC}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false))
}

func (suite *UtilsTestSuite) TestGetFilesInDirInexistantDir() {
	assert.Panics(suite.T(), func() { GetFilesInDir(suite.outputFolder+"folder", false) }, "Should panic")
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFilesWithDirRecursive() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	filePathB := filepath.Join(suite.outputFolder, "fileB")
	filePathC := filepath.Join(suite.outputFolder, "fileC")
	folderPath := filepath.Join(suite.outputFolder, "folder")
	filePathD := filepath.Join(folderPath, "fileD")
	os.WriteFile(filePathA, []byte("12345"), 0755)
	os.WriteFile(filePathB, []byte("ABCDE"), 0755)
	os.WriteFile(filePathC, []byte(",.-{}"), 0755)
	os.Mkdir(folderPath, 0755)
	os.WriteFile(filePathD, []byte("___"), 0755)

	expectedFiles := []string{filePathA, filePathB, filePathC, filePathD}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, true))
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

// GzipCompress()

func (suite *UtilsTestSuite) TestGzipCompressDecompress() {
	var dataOrigin []byte
	for i := 0; i < 10000000; i++ {
		dataOrigin = append(dataOrigin, byte(i))
	}

	dataCompressed := GzipCompress(dataOrigin)
	assert.Greater(suite.T(), len(dataOrigin), len(dataCompressed))

	dataDecompressed := GzipDecompress(dataCompressed)
	assert.Equal(suite.T(), dataOrigin, dataDecompressed)
}

// TestUtilsTestSuite()

func TestUtilsTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/UtilsTestSuite/"
	var testSuite UtilsTestSuite
	testSuite.outputFolder = OutputFolder
	suite.Run(t, &testSuite)
}
