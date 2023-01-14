package buffer

import "bytes"

// Implements the Buffer interface using an underlying bytes.Buffer
type BytesBuffer struct {
	buf *bytes.Buffer
}

func NewBytesBuffer() *BytesBuffer {
	return &BytesBuffer{buf: new(bytes.Buffer)}
}

func (bb *BytesBuffer) WriteByte(v byte) error {
	return bb.buf.WriteByte(v)
}

func (bb *BytesBuffer) Write(v []byte) error {
	_, err := bb.buf.Write(v)
	return err
}

func (bb *BytesBuffer) Bytes() []byte {
	return bb.buf.Bytes()
}

func (bb *BytesBuffer) Reset() {
	bb.buf.Reset()
}
