package rat

type DataBytesSliceManager struct {
	dataBytes []byte
	dataIdx   uint
}

func NewDataBytesSliceManager(data []byte) *DataBytesSliceManager {
	return &DataBytesSliceManager{data, 0}
}

func (manager *DataBytesSliceManager) Read(bytesQty uint) ([]byte, int) {
	currentIdx := manager.dataIdx
	manager.dataIdx = min(manager.dataIdx+bytesQty, uint(len(manager.dataBytes)))
	return manager.dataBytes[currentIdx:manager.dataIdx], int(manager.dataIdx - currentIdx)
}
