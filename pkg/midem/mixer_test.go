package midem

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMixOneManager(t *testing.T) {
	managerA := NewDataBytesManagerMock([]byte{'1', '2', '3'})

	managers := []IDataBytesManager{managerA}
	mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixTwoManagers(t *testing.T) {
	managerA := NewDataBytesManagerMock([]byte{'1', '3', '5'})
	managerB := NewDataBytesManagerMock([]byte{'2', '4', '6'})

	managers := []IDataBytesManager{managerA, managerB}
	mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3', '4', '5', '6'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixTwoManagersDifferentSizes(t *testing.T) {
	managerA := NewDataBytesManagerMock([]byte{'1', '3', '5'})
	managerB := NewDataBytesManagerMock([]byte{'2', '4', '6', '7', '8'})

	managers := []IDataBytesManager{managerA, managerB}
	mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixThreeManagers(t *testing.T) {
	managerA := NewDataBytesManagerMock([]byte{'A', 'D', 'G'})
	managerB := NewDataBytesManagerMock([]byte{'B', 'E', 'H'})
	managerC := NewDataBytesManagerMock([]byte{'C', 'F', 'I'})

	managers := []IDataBytesManager{managerA, managerB, managerC}
	mixer := NewMixer(managers)

	expected := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	assert.Equal(t, expected, mixer.Mix())
}

func TestMixThreeManagersDifferentSizes(t *testing.T) {
	managerA := NewDataBytesManagerMock([]byte{'A', 'D', 'G', 'I'})
	managerB := NewDataBytesManagerMock([]byte{'B', 'E'})
	managerC := NewDataBytesManagerMock([]byte{'C', 'F', 'H'})

	managers := []IDataBytesManager{managerA, managerB, managerC}
	mixer := NewMixer(managers)

	expected := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	assert.Equal(t, expected, mixer.Mix())
}
