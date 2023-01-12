package nexemab

import (
	"bytes"
	"unsafe"
)

type Encoder struct {
	buf *bytes.Buffer
}

func NewEncoder() *Encoder {
	return &Encoder{buf: new(bytes.Buffer)}
}

func (e *Encoder) encodeBool(v bool) {
	if v {
		e.buf.WriteByte(BoolTrue)
	} else {
		e.buf.WriteByte(BoolFalse)
	}
}

func (e *Encoder) encodeNull() {
	e.buf.WriteByte(Null)
}

func (e *Encoder) encodeString(input string) {
	// todo: encode length
	// todo: found a better, performant safe way of doing this
	buf := *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{input, len(input)}))
	e.buf.Write(buf)
}

func (e *Encoder) encodeUint8(v uint8) {
	e.buf.WriteByte(v)
}

func (e *Encoder) encodeUint16(v uint16) {
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) encodeUint32(v uint32) {
	e.buf.WriteByte(byte(v >> 24))
	e.buf.WriteByte(byte(v >> 16))
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) Close() []byte {
	return e.buf.Bytes()
}
