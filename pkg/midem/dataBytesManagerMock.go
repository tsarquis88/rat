package midem

type DataBytesManagerMock struct {
	idx       uint
	dataBytes []byte
}

func NewDataBytesManagerMock(data []byte) *DataBytesManagerMock {
	return &DataBytesManagerMock{0, data}
}

func (manager *DataBytesManagerMock) Read(bytesQty uint) ([]byte, int) {
	dataLen := uint(len(manager.dataBytes))
	if dataLen <= manager.idx {
		return []byte{}, 0
	}
	upperLimit := manager.idx + bytesQty
	if upperLimit > dataLen {
		upperLimit = dataLen
	}
	retData := manager.dataBytes[manager.idx:upperLimit]
	readBytes := upperLimit - manager.idx
	manager.idx = upperLimit
	return retData, int(readBytes)
}

func (manager *DataBytesManagerMock) Name() string {
	return "DataBytesManagerMock"
}

func (manager *DataBytesManagerMock) Origin() string {
	return ""
}
