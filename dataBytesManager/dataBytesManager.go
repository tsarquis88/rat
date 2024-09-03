package dataBytesManager

type IDataBytesManager interface {
	Read(bytesQty uint) ([]byte, int)
}
