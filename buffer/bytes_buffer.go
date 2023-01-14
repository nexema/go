package buffer

import "bytes"

// Implements the Buffer interface using an underlying bytes.Buffer
type bytesBuffer struct {
	buf *bytes.Buffer
}

func NewBytesBuffer() *bytesBuffer {
	return &bytesBuffer{buf: new(bytes.Buffer)}
}

func (bb *bytesBuffer) WriteByte(v byte) error {
	return bb.buf.WriteByte(v)
}

func (bb *bytesBuffer) Write(v []byte) error {
	_, err := bb.buf.Write(v)
	return err
}

func (bb *bytesBuffer) Bytes() []byte {
	return bb.buf.Bytes()
}

func (bb *bytesBuffer) Reset() {
	bb.buf.Reset()
}
