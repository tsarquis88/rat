package rat

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type DataBytesSliceManagerTestSuite struct {
	suite.Suite
}

// NewDataBytesSliceManager()

func (suite *DataBytesSliceManagerTestSuite) TestNew() {
	testData := []byte("12345")
	newManager := NewDataBytesSliceManager(testData)
	assert.Equal(suite.T(), testData, newManager.dataBytes)
	assert.Equal(suite.T(), uint(0), newManager.dataIdx)
}

// Read()

func (suite *DataBytesSliceManagerTestSuite) TestRead() {
	testData := []byte("12345")
	newManager := NewDataBytesSliceManager(testData)

	data, n := newManager.Read(1)
	assert.Equal(suite.T(), 1, n)
	assert.Equal(suite.T(), []byte("1"), data)

	data, n = newManager.Read(2)
	assert.Equal(suite.T(), 2, n)
	assert.Equal(suite.T(), []byte("23"), data)

	data, n = newManager.Read(3)
	assert.Equal(suite.T(), 2, n)
	assert.Equal(suite.T(), []byte("45"), data)

	data, n = newManager.Read(4)
	assert.Equal(suite.T(), 0, n)
	assert.Equal(suite.T(), []byte(""), data)
}

func (suite *DataBytesSliceManagerTestSuite) TestReadOneRead() {
	testData := []byte("12345")
	newManager := NewDataBytesSliceManager(testData)

	data, n := newManager.Read(5)
	assert.Equal(suite.T(), 5, n)
	assert.Equal(suite.T(), testData, data)
}

func (suite *DataBytesSliceManagerTestSuite) TestReadOneReadPass() {
	testData := []byte("12345")
	newManager := NewDataBytesSliceManager(testData)

	data, n := newManager.Read(10)
	assert.Equal(suite.T(), 5, n)
	assert.Equal(suite.T(), testData, data)
}

// TestDataBytesSliceManagerTestSuite

func TestDataBytesSliceManagerTestSuite(t *testing.T) {
	var testSuite DataBytesSliceManagerTestSuite
	suite.Run(t, &testSuite)
}
