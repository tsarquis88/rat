package mixer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"example.com/dataBytesManager"
)

type DataBytesManagerMock struct {
	idx int
	dataBytes []byte
	writen []byte
}

func (manager *DataBytesManagerMock) Read() (byte, int) {
	if len(manager.dataBytes) <= manager.idx {
		return 0, 0
	}
	data := manager.dataBytes[manager.idx]
	manager.idx++
	return data, 1
}

func (manager *DataBytesManagerMock) Write(data byte) (int) {
	manager.writen = append(manager.writen, data)
	return 1
}

func TestMixOneManager(t *testing.T) {
	managerA := &DataBytesManagerMock{0, []byte{'1', '2', '3'}, []byte{}}
	
	managers := []dataBytesManager.IDataBytesManager{managerA}
    mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3'}
	assert.Equal(t, mixer.Mix(), expected)
}

func TestMixTwoManagers(t *testing.T) {
	managerA := &DataBytesManagerMock{0, []byte{'1', '3', '5'}, []byte{}}
	managerB := &DataBytesManagerMock{0, []byte{'2', '4', '6'}, []byte{}}
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB}
    mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3', '4', '5', '6'}
	assert.Equal(t, mixer.Mix(), expected)
}

func TestMixTwoManagersDifferentSizes(t *testing.T) {
	managerA := &DataBytesManagerMock{0, []byte{'1', '3', '5'}, []byte{}}
	managerB := &DataBytesManagerMock{0, []byte{'2', '4', '6', '7', '8'}, []byte{}}
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB}
    mixer := NewMixer(managers)

	expected := []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
	assert.Equal(t, mixer.Mix(), expected)
}

func TestMixThreeManagers(t *testing.T) {
	managerA := &DataBytesManagerMock{0, []byte{'A', 'D', 'G'}, []byte{}}
	managerB := &DataBytesManagerMock{0, []byte{'B', 'E', 'H'}, []byte{}}
	managerC := &DataBytesManagerMock{0, []byte{'C', 'F', 'I'}, []byte{}}
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB, managerC}
    mixer := NewMixer(managers)

	expected := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	assert.Equal(t, mixer.Mix(), expected)
}

func TestMixThreeManagersDifferentSizes(t *testing.T) {
	managerA := &DataBytesManagerMock{0, []byte{'A', 'D', 'G', 'I'}, []byte{}}
	managerB := &DataBytesManagerMock{0, []byte{'B', 'E'}, []byte{}}
	managerC := &DataBytesManagerMock{0, []byte{'C', 'F', 'H'}, []byte{}}
	
	managers := []dataBytesManager.IDataBytesManager{managerA, managerB, managerC}
    mixer := NewMixer(managers)

	expected := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I'}
	assert.Equal(t, mixer.Mix(), expected)
}
