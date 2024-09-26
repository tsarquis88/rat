package midem

type IDataBytesManager interface {
	Read(bytesQty uint) ([]byte, int)
	Name() string
	Origin() string
}
