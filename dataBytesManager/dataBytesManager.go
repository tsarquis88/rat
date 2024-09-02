package dataBytesManager

type IDataBytesManager interface {
	Read() (byte, int)
	Write(data byte) (int)
}
