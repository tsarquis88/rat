package midem

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FilesMidemTestSuite struct {
	suite.Suite
	inputFolder  string
	outputFolder string
	outputFile   string
}

func (suite *FilesMidemTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *FilesMidemTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

func (suite *FilesMidemTestSuite) TestMixAndDemix() {
	filesInDir := GetFilesInDir(suite.inputFolder)

	originalFiles := make(map[string][]byte)
	var inputFiles []string
	for _, file := range filesInDir {
		filePath := suite.inputFolder + file
		originalFiles[file] = HashFile(filePath)
		inputFiles = append(inputFiles, filePath)
	}

	MixFiles(inputFiles, suite.outputFile)
	assert.Equal(suite.T(), true, FileExists(suite.outputFile))

	DemixFiles([]string{suite.outputFile}, suite.outputFolder)
	for file, hash := range originalFiles {
		filepath := suite.outputFolder + file
		assert.Equal(suite.T(), true, FileExists(filepath))
		assert.Equal(suite.T(), hash, HashFile(filepath))
	}
}

func (suite *FilesMidemTestSuite) TestMixAndDemixFolder() {
	filesInDir := GetFilesInDir(suite.inputFolder)

	originalFiles := make(map[string][]byte)
	for _, file := range filesInDir {
		filePath := suite.inputFolder + file
		originalFiles[file] = HashFile(filePath)
	}

	MixFiles([]string{suite.inputFolder}, suite.outputFile)
	assert.Equal(suite.T(), true, FileExists(suite.outputFile))

	DemixFiles([]string{suite.outputFile}, suite.outputFolder)
	assert.Equal(suite.T(), true, FileExists(suite.outputFolder+suite.inputFolder))
	for file, hash := range originalFiles {
		filepath := suite.outputFolder + suite.inputFolder + file
		assert.Equal(suite.T(), true, FileExists(filepath))
		assert.Equal(suite.T(), hash, HashFile(filepath))
	}
}

func TestFilesMidemTestSuite(t *testing.T) {
	const InputFolder = "test_files/"
	const OutputFolder = "/tmp/FilesMidemTestSuite/"
	var testSuite FilesMidemTestSuite
	testSuite.inputFolder = InputFolder
	testSuite.outputFolder = OutputFolder
	testSuite.outputFile = OutputFolder + "output.mix"
	suite.Run(t, &testSuite)
}
