package rat

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RatTestSuite struct {
	suite.Suite
	inputFolder  string
	outputFolder string
	outputFile   string
}

func (suite *RatTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *RatTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

func (suite *RatTestSuite) TestRatAndDerat() {
	filesInDir := GetFilesInDir(suite.inputFolder, false, false)

	originalFiles := make(map[string][]byte)
	var inputFiles []string
	for _, file := range filesInDir {
		originalFiles[file] = HashFile(file)
		inputFiles = append(inputFiles, file)
	}

	Rat(inputFiles, suite.outputFile)
	assert.Equal(suite.T(), true, FileExists(suite.outputFile))

	// List
	expectedFiles := make(map[string]string)
	expectedFiles[suite.outputFile] = ""
	for _, file := range filesInDir {
		expectedFiles[suite.outputFile] += file + "\n"
	}
	expectedFiles[suite.outputFile] = strings.TrimSuffix(expectedFiles[suite.outputFile], "\n")
	assert.Equal(suite.T(), expectedFiles, List([]string{suite.outputFile}))

	Derat([]string{suite.outputFile}, suite.outputFolder)
	for file, hash := range originalFiles {
		filepath := filepath.Join(suite.outputFolder, file)
		assert.Equal(suite.T(), true, FileExists(filepath))
		assert.Equal(suite.T(), hash, HashFile(filepath))
	}
}

func (suite *RatTestSuite) TestRatAndDeratFolder() {
	filesInDir := GetFilesInDir(suite.inputFolder, false, false)

	originalFiles := make(map[string][]byte)
	for _, file := range filesInDir {
		originalFiles[file] = HashFile(file)
	}

	Rat([]string{suite.inputFolder}, suite.outputFile)
	assert.Equal(suite.T(), true, FileExists(suite.outputFile))

	// List
	expectedFiles := make(map[string]string)
	expectedFiles[suite.outputFile] = suite.inputFolder + "\n"
	for _, file := range filesInDir {
		expectedFiles[suite.outputFile] += file + "\n"
	}
	expectedFiles[suite.outputFile] = strings.TrimSuffix(expectedFiles[suite.outputFile], "\n")
	assert.Equal(suite.T(), expectedFiles, List([]string{suite.outputFile}))

	Derat([]string{suite.outputFile}, suite.outputFolder)
	assert.Equal(suite.T(), true, FileExists(filepath.Join(suite.outputFolder, suite.inputFolder)))
	for file, hash := range originalFiles {
		filepath := filepath.Join(suite.outputFolder, file)
		assert.Equal(suite.T(), true, FileExists(filepath))
		assert.Equal(suite.T(), hash, HashFile(filepath))
	}
}

func (suite *RatTestSuite) TestRatOutputFileExists() {
	filesInDir := GetFilesInDir(suite.inputFolder, false, false)
	os.WriteFile(suite.outputFile, []byte("12345"), 0755)

	assert.Panics(suite.T(), func() { Rat(filesInDir, suite.outputFile) })
}

func TestRatTestSuite(t *testing.T) {
	const InputFolder = "test_files"
	const OutputFolder = "/tmp/RatTestSuite"
	var testSuite RatTestSuite
	testSuite.inputFolder = InputFolder
	testSuite.outputFolder = OutputFolder
	testSuite.outputFile = filepath.Join(OutputFolder, "output.rat")
	suite.Run(t, &testSuite)
}
