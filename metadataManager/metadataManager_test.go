package metadataManager

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
)

const OutputFolder = "/tmp/MetadataGeneratorTestSuite"
const FileNameA = "fileA"
const FileNameB = "fileB"
const TestFileA = OutputFolder + "/" + FileNameA
const TestFileB = OutputFolder + "/" + FileNameB

type MetadataGeneratorTestSuite struct {
    suite.Suite
}

func (suite *MetadataGeneratorTestSuite) SetupTest() {
	os.Mkdir(OutputFolder, 0755)
}

func (suite *MetadataGeneratorTestSuite) TearDownTest() {
	os.RemoveAll(OutputFolder)
}

// Generate()

func (suite *MetadataGeneratorTestSuite) TestGenerateOneFile() {
	os.WriteFile(TestFileA, []byte("12345"), 0755)

	expectedMetadatas := []Metadata{{FileNameA, 5}}
	metadatas := Generate([]string{TestFileA})
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func (suite *MetadataGeneratorTestSuite) TestGenerateMultipleFiles() {
	os.WriteFile(TestFileA, []byte("12345"), 0755)
	os.WriteFile(TestFileB, []byte("ABC"), 0755)

	expectedMetadatas := []Metadata{{FileNameA, 5}, {FileNameB, 3}}
	metadatas := Generate([]string{TestFileA, TestFileB})
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func (suite *MetadataGeneratorTestSuite) TestGenerateInexistantFilePanics() {
	assert.Panics(suite.T(), func() {Generate([]string{TestFileA})}, "Should panic")
}

// Dump()

func (suite *MetadataGeneratorTestSuite) TestDumpOneMetadata() {
	metadatas := []Metadata{{FileNameA, 5}}

	expectedFileSize := make([]byte, 8)
	expectedFileSize[0] = 5

	var expectedDump []byte
	expectedDump = append(expectedDump, byte(1))
	expectedDump = append(expectedDump, byte(len(FileNameA)))
	expectedDump = append(expectedDump, []byte(FileNameA)...)
	expectedDump = append(expectedDump, expectedFileSize...)

	dump := Dump(metadatas)
	assert.Equal(suite.T(), expectedDump, dump)
}

func (suite *MetadataGeneratorTestSuite) TestDumpMultipleFiles() {
	metadatas := []Metadata{{FileNameA, 5}, {FileNameB, 3}}

	expectedFileSizeA := make([]byte, 8)
	expectedFileSizeA[0] = 5

	expectedFileSizeB := make([]byte, 8)
	expectedFileSizeB[0] = 3
	
	var expectedDump []byte
	expectedDump = append(expectedDump, byte(2))
	expectedDump = append(expectedDump, byte(len(FileNameA)))
	expectedDump = append(expectedDump, []byte(FileNameA)...)
	expectedDump = append(expectedDump, expectedFileSizeA...)
	expectedDump = append(expectedDump, byte(len(FileNameB)))
	expectedDump = append(expectedDump, []byte(FileNameB)...)
	expectedDump = append(expectedDump, expectedFileSizeB...)

	dump := Dump(metadatas)
	assert.Equal(suite.T(), expectedDump, dump)
}

// Parse()

func (suite *MetadataGeneratorTestSuite) TestParseOneMetadata() {
	expectedMetadatas := []Metadata{{FileNameA, 5}}

	expectedFileSize := make([]byte, 8)
	expectedFileSize[0] = 5

	var dump []byte
	dump = append(dump, byte(1))
	dump = append(dump, byte(len(FileNameA)))
	dump = append(dump, []byte(FileNameA)...)
	dump = append(dump, expectedFileSize...)

	metadatas := Parse(dump)
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func (suite *MetadataGeneratorTestSuite) TestParseMultipleMetadatas() {
	expectedMetadatas := []Metadata{{FileNameA, 5}, {FileNameB, 3}}

	expectedFileSizeA := make([]byte, 8)
	expectedFileSizeA[0] = 5

	expectedFileSizeB := make([]byte, 8)
	expectedFileSizeB[0] = 3

	var dump []byte
	dump = append(dump, byte(2))
	dump = append(dump, byte(len(FileNameA)))
	dump = append(dump, []byte(FileNameA)...)
	dump = append(dump, expectedFileSizeA...)
	dump = append(dump, byte(len(FileNameB)))
	dump = append(dump, []byte(FileNameB)...)
	dump = append(dump, expectedFileSizeB...)

	metadatas := Parse(dump)
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func TestMetadataGeneratorTestSuite(t *testing.T) {
    suite.Run(t, new(MetadataGeneratorTestSuite))
}
