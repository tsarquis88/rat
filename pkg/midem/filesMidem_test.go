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

func (suite *FilesMidemTestSuite) TestMix() {
	inputFiles, err := os.ReadDir(suite.inputFolder)
	if err != nil {
		panic(err)
	}

	files := []string{suite.outputFile}
	for _, file := range inputFiles {
		files = append(files, suite.inputFolder+"/"+file.Name())
	}

	MixFiles(files)
	assert.Equal(suite.T(), true, FileExists(suite.outputFile))
}

func TestFilesMidemTestSuite(t *testing.T) {
	const InputFolder = "./test_files"
	const OutputFolder = "/tmp/FilesMidemTestSuite"
	var testSuite FilesMidemTestSuite
	testSuite.inputFolder = InputFolder
	testSuite.outputFolder = OutputFolder
	testSuite.outputFile = OutputFolder + "/output.mix"
	suite.Run(t, &testSuite)
}
