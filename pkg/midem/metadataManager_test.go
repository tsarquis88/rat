package midem

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type MetadataGeneratorTestSuite struct {
	suite.Suite
	outputFolder string
	fileNameA    string
	fileNameB    string
	testFileA    string
	testFileB    string
}

func (suite *MetadataGeneratorTestSuite) SetupTest() {
	os.Mkdir(suite.outputFolder, 0755)
}

func (suite *MetadataGeneratorTestSuite) TearDownTest() {
	os.RemoveAll(suite.outputFolder)
}

// Generate()

func (suite *MetadataGeneratorTestSuite) TestGenerateOneFile() {
	os.WriteFile(suite.testFileA, []byte("12345"), 0755)

	expectedMetadatas := []Metadata{{suite.fileNameA, 5, 0755}}
	metadatas := Generate([]string{suite.testFileA})
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func (suite *MetadataGeneratorTestSuite) TestGenerateMultipleFiles() {
	os.WriteFile(suite.testFileA, []byte("12345"), 0755)
	os.WriteFile(suite.testFileB, []byte("ABC"), 0755)

	expectedMetadatas := []Metadata{{suite.fileNameA, 5, 0755}, {suite.fileNameB, 3, 0755}}
	metadatas := Generate([]string{suite.testFileA, suite.testFileB})
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func (suite *MetadataGeneratorTestSuite) TestGenerateInexistantFilePanics() {
	assert.Panics(suite.T(), func() { Generate([]string{suite.testFileA}) }, "Should panic")
}

// Dump()

func (suite *MetadataGeneratorTestSuite) TestDumpOneMetadata() {
	metadatas := []Metadata{{suite.fileNameA, 5, 0755}}

	expectedFileSize := make([]byte, 8)
	expectedFileSize[0] = 5

	expectedFileMode := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileMode, uint32(0755))

	var expectedDump []byte
	expectedDump = append(expectedDump, byte(1))
	expectedDump = append(expectedDump, byte(len(suite.fileNameA)))
	expectedDump = append(expectedDump, []byte(suite.fileNameA)...)
	expectedDump = append(expectedDump, expectedFileSize...)
	expectedDump = append(expectedDump, expectedFileMode...)

	dump := Dump(metadatas)
	assert.Equal(suite.T(), expectedDump, dump)
}

func (suite *MetadataGeneratorTestSuite) TestDumpMultipleFiles() {
	metadatas := []Metadata{{suite.fileNameA, 5, 0755}, {suite.fileNameB, 3, 0755}}

	expectedFileSizeA := make([]byte, 8)
	expectedFileSizeA[0] = 5
	expectedFileModeA := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileModeA, uint32(0755))

	expectedFileSizeB := make([]byte, 8)
	expectedFileSizeB[0] = 3
	expectedFileModeB := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileModeB, uint32(0755))

	var expectedDump []byte
	expectedDump = append(expectedDump, byte(2))
	expectedDump = append(expectedDump, byte(len(suite.fileNameA)))
	expectedDump = append(expectedDump, []byte(suite.fileNameA)...)
	expectedDump = append(expectedDump, expectedFileSizeA...)
	expectedDump = append(expectedDump, expectedFileModeA...)
	expectedDump = append(expectedDump, byte(len(suite.fileNameB)))
	expectedDump = append(expectedDump, []byte(suite.fileNameB)...)
	expectedDump = append(expectedDump, expectedFileSizeB...)
	expectedDump = append(expectedDump, expectedFileModeB...)

	dump := Dump(metadatas)
	assert.Equal(suite.T(), expectedDump, dump)
}

// Parse()

func (suite *MetadataGeneratorTestSuite) TestParseOneMetadata() {
	expectedMetadatas := []Metadata{{suite.fileNameA, 5, 0755}}

	expectedFileSize := make([]byte, 8)
	expectedFileSize[0] = 5
	expectedFileMode := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileMode, uint32(0755))

	var dump []byte
	dump = append(dump, byte(1))
	dump = append(dump, byte(len(suite.fileNameA)))
	dump = append(dump, []byte(suite.fileNameA)...)
	dump = append(dump, expectedFileSize...)
	dump = append(dump, expectedFileMode...)
	dump = append(dump, []byte("Unread data")...)

	dataBytesManager := NewDataBytesManagerMock(dump)

	metadatas := Parse(dataBytesManager)
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func (suite *MetadataGeneratorTestSuite) TestParseMultipleMetadatas() {
	expectedMetadatas := []Metadata{{suite.fileNameA, 5, 0755}, {suite.fileNameB, 3, 0755}}

	expectedFileSizeA := make([]byte, 8)
	expectedFileSizeA[0] = 5

	expectedFileSizeB := make([]byte, 8)
	expectedFileSizeB[0] = 3

	expectedFileModeA := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileModeA, uint32(0755))

	expectedFileModeB := make([]byte, 4)
	binary.LittleEndian.PutUint32(expectedFileModeB, uint32(0755))

	var dump []byte
	dump = append(dump, byte(2))
	dump = append(dump, byte(len(suite.fileNameA)))
	dump = append(dump, []byte(suite.fileNameA)...)
	dump = append(dump, expectedFileSizeA...)
	dump = append(dump, expectedFileModeA...)
	dump = append(dump, byte(len(suite.fileNameB)))
	dump = append(dump, []byte(suite.fileNameB)...)
	dump = append(dump, expectedFileSizeB...)
	dump = append(dump, expectedFileModeB...)
	dump = append(dump, []byte("Unread data")...)

	dataBytesManager := NewDataBytesManagerMock(dump)

	metadatas := Parse(dataBytesManager)
	assert.Equal(suite.T(), expectedMetadatas, metadatas)
}

func TestMetadataGeneratorTestSuite(t *testing.T) {
	const OutputFolder = "/tmp/MetadataGeneratorTestSuite"
	const FileNameA = "fileA"
	const FileNameB = "fileB"
	const TestFileA = OutputFolder + "/" + FileNameA
	const TestFileB = OutputFolder + "/" + FileNameB
	var testSuite MetadataGeneratorTestSuite
	testSuite.outputFolder = OutputFolder
	testSuite.fileNameA = FileNameA
	testSuite.fileNameB = FileNameB
	testSuite.testFileA = TestFileA
	testSuite.testFileB = TestFileB
	suite.Run(t, &testSuite)
}
