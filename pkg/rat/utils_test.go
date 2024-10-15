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

// FileRead()

func (suite *UtilsTestSuite) TestFileRead() {
	fileData := []byte("12345")
	testFile := filepath.Join(suite.outputFolder, "test.file")
	os.WriteFile(testFile, fileData, 0755)

	assert.Equal(suite.T(), fileData, FileRead(testFile))
}

func (suite *UtilsTestSuite) TestFileReadInexistantFile() {
	testFile := filepath.Join(suite.outputFolder, "test.file")

	assert.Panics(suite.T(), func() { FileRead(testFile) })
}

func (suite *UtilsTestSuite) TestFileReadBigFile() {
	var fileData []byte
	for i := 0; i < 1024*10000; i++ {
		fileData = append(fileData, byte(i))
	}

	testFile := filepath.Join(suite.outputFolder, "test.file")
	os.WriteFile(testFile, fileData, 0755)

	assert.Equal(suite.T(), fileData, FileRead(testFile))
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

	assert.Panics(suite.T(), func() { IsDir(testFolder) })
}

// GetFilesInDir()

func (suite *UtilsTestSuite) TestGetFilesInDirNoFiles() {
	assert.Equal(suite.T(), 0, len(GetFilesInDir(suite.outputFolder, false, false)))
}

func (suite *UtilsTestSuite) TestGetFilesInDirOneFile() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	os.WriteFile(filePathA, []byte("12345"), 0755)

	expectedFiles := []string{filePathA}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false, false))
}

func (suite *UtilsTestSuite) TestGetFilesInDirOneFileIncludeDirs() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	os.WriteFile(filePathA, []byte("12345"), 0755)

	expectedFiles := []string{suite.outputFolder, filePathA}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false, true))
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFiles() {
	filePathA := filepath.Join(suite.outputFolder, "fileA")
	filePathB := filepath.Join(suite.outputFolder, "fileB")
	filePathC := filepath.Join(suite.outputFolder, "fileC")
	os.WriteFile(filePathA, []byte("12345"), 0755)
	os.WriteFile(filePathB, []byte("ABCDE"), 0755)
	os.WriteFile(filePathC, []byte(",.-{}"), 0755)

	expectedFiles := []string{filePathA, filePathB, filePathC}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false, false))
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
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, false, false))
}

func (suite *UtilsTestSuite) TestGetFilesInDirInexistantDir() {
	assert.Panics(suite.T(), func() { GetFilesInDir(suite.outputFolder+"folder", false, false) })
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
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, true, false))
}

func (suite *UtilsTestSuite) TestGetFilesInDirMultipleFilesWithDirRecursiveIncludeDirs() {
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

	expectedFiles := []string{suite.outputFolder, filePathA, filePathB, filePathC, folderPath, filePathD}
	assert.Equal(suite.T(), expectedFiles, GetFilesInDir(suite.outputFolder, true, true))
}

// HashFile()

func (suite *UtilsTestSuite) TestHashFile() {
	os.WriteFile(suite.outputFolder+"fileA", []byte("12345"), 0755)

	const ExpectedHash = "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5"
	assert.Equal(suite.T(), ExpectedHash, fmt.Sprintf("%x", HashFile(suite.outputFolder+"fileA")))
}

func (suite *UtilsTestSuite) TestHashFileInexistantFile() {
	assert.Panics(suite.T(), func() { HashFile(suite.outputFolder + "file") })
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

// OctalToDecimal()

func (suite *UtilsTestSuite) TestOctalToDecimal() {
	assert.Equal(suite.T(), uint(294), OctalToDecimal([]byte{'4', '4', '6'}, 3))
	assert.Equal(suite.T(), uint(2054353), OctalToDecimal([]byte{'7', '6', '5', '4', '3', '2', '1'}, 7))
}

func (suite *UtilsTestSuite) TestOctalToDecimalReducedLen() {
	assert.Equal(suite.T(), uint(36), OctalToDecimal([]byte{'4', '4', '6'}, 2))
	assert.Equal(suite.T(), uint(501), OctalToDecimal([]byte{'7', '6', '5', '4', '3', '2', '1'}, 3))
}

func (suite *UtilsTestSuite) TestOctalToDecimalOutOfIndex() {
	assert.Panics(suite.T(), func() { OctalToDecimal([]byte{'7', '6', '5', '4', '3', '2', '1'}, 9) })
}

// DecimalToOctal()

func (suite *UtilsTestSuite) TestDecimalToOctal() {
	assert.Equal(suite.T(), []byte{'1', '4', '4'}, DecimalToOctal(100))
	assert.Equal(suite.T(), []byte{'7', '2', '6', '7', '4', '6', '4', '2', '6', '1'}, DecimalToOctal(987654321))
}

// FillWith()

func (suite *UtilsTestSuite) TestFillWith() {
	assert.Equal(suite.T(), []byte{1, 2, 3, 4}, FillWith([]byte{1, 2, 3}, 4, 4))
	assert.Equal(suite.T(), []byte{'A', 'X', 'X', 'X', 'X'}, FillWith([]byte{'A'}, 'X', 5))
}

func (suite *UtilsTestSuite) TestFillWithReducedLen() {
	assert.Equal(suite.T(), []byte{1, 2, 3}, FillWith([]byte{1, 2, 3}, 4, 1))
}

func (suite *UtilsTestSuite) TestFillWithEmpty() {
	assert.Equal(suite.T(), []byte{4, 4, 4}, FillWith([]byte{}, 4, 3))
}

// TrimPrefixRecursive()

func (suite *UtilsTestSuite) TestTrimPrefixRecursive() {
	assert.Equal(suite.T(), "folder/file.json", TrimPrefixRecursive("../../folder/file.json", "../"))
	assert.Equal(suite.T(), "folder/file.json", TrimPrefixRecursive("../folder/file.json", "../"))
	assert.Equal(suite.T(), "folder/file.json", TrimPrefixRecursive("folder/file.json", "../"))
}

// GetChecksum()

func (suite *UtilsTestSuite) TestGetChecksum() {
	data := []byte{0, 1, 2, 3}
	assert.Equal(suite.T(), uint(6), GetChecksum(data))
}

func (suite *UtilsTestSuite) TestGetChecksumEmpty() {
	data := []byte{}
	assert.Equal(suite.T(), uint(0), GetChecksum(data))
}

func (suite *UtilsTestSuite) TestGetChecksumBigSlice() {
	const DataSize = 1000
	data := FillWith([]byte{}, 3, DataSize)
	assert.Equal(suite.T(), uint(3000), GetChecksum(data))
}

// ShiftLeft()

func (suite *UtilsTestSuite) TestShiftLeft() {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(suite.T(), []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 99}, ShiftLeft(data, 1, 99))
	assert.Equal(suite.T(), []byte{6, 7, 8, 9, 99, 99, 99, 99, 99, 99}, ShiftLeft(data, 5, 99))
}

func (suite *UtilsTestSuite) TestShiftLeftBigAmount() {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Panics(suite.T(), func() { ShiftLeft(data, 100, 99) })
}

// TestUtilsTestSuite()

func TestUtilsTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/UtilsTestSuite/"
	var testSuite UtilsTestSuite
	testSuite.outputFolder = OutputFolder
	suite.Run(t, &testSuite)
}
