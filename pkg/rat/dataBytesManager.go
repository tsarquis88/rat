package rat

type IDataBytesManager interface {
	Read(bytesQty uint) ([]byte, int)
}
