package mixer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"example.com/dataBytesManager"
	"example.com/dataBytesManagerMock"
)

func TestMixOneManager(t *testing.T) {
	managerA := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'1', '2', '3'})
	
	managers := []dataBytesManager.IDataBytesManager{managerA}
    mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixTwoManagers(t *testing.T) {
	managerA := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'1', '3', '5'})
	managerB := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'2', '4', '6'})
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB}
    mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3', '4', '5', '6'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixTwoManagersDifferentSizes(t *testing.T) {
	managerA := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'1', '3', '5'})
	managerB := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'2', '4', '6', '7', '8'})
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB}
    mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixThreeManagers(t *testing.T) {
	managerA := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'A', 'D', 'G'})
	managerB := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'B', 'E', 'H'})
	managerC := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'C', 'F', 'I'})
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB, managerC}
    mixer := NewMixer(managers)

	expected := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixThreeManagersDifferentSizes(t *testing.T) {
	managerA := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'A', 'D', 'G', 'I'})
	managerB := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'B', 'E'})
	managerC := dataBytesManagerMock.NewDataBytesManagerMock([]byte{'C', 'F', 'H'})
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB, managerC}
    mixer := NewMixer(managers)

	expected := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	assert.Equal(t, expected, mixer.Mix())
}
