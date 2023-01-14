package buffer

// Buffer provides methods to write to a underlying byte array in a performant way
type Buffer interface {
	WriteByte(v byte) error
	Write(v []byte) error
	Bytes() []byte
	Reset()
}
