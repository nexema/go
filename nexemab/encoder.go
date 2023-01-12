package nexemab

import (
	"bytes"
	"math"
	"unsafe"
)

type Encoder struct {
	buf *bytes.Buffer
}

func NewEncoder() *Encoder {
	// todo: accept initial buffer size to avoid unneccessary grows
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
	// todo: found a better, performant safe way of doing this
	inputLen := len(input)
	buf := *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{input, inputLen}))
	e.encodeVarint(int64(inputLen))
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

func (e *Encoder) encodeUint64(v uint64) {
	e.buf.WriteByte(byte(v >> 56))
	e.buf.WriteByte(byte(v >> 48))
	e.buf.WriteByte(byte(v >> 40))
	e.buf.WriteByte(byte(v >> 32))
	e.buf.WriteByte(byte(v >> 24))
	e.buf.WriteByte(byte(v >> 16))
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) encodeInt64(v int64) {
	e.buf.WriteByte(byte(v >> 56))
	e.buf.WriteByte(byte(v >> 48))
	e.buf.WriteByte(byte(v >> 40))
	e.buf.WriteByte(byte(v >> 32))
	e.buf.WriteByte(byte(v >> 24))
	e.buf.WriteByte(byte(v >> 16))
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) encodeUvarint(v uint64) {
	// code internally used by binary.PutUvarint, but modified to write directly to the underlying buffer
	i := 0
	for v >= 0x80 {
		e.buf.WriteByte(byte(v) | 0x80)
		v >>= 7
		i++
	}
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) encodeVarint(v int64) {
	// code internally used by binary.PutVarint, but modified to write directly to the underlying buffer
	ux := uint64(v) << 1
	if v < 0 {
		ux = ^ux
	}

	e.encodeUvarint(ux)
}

func (e *Encoder) encodeFloat32(v float32) {
	bits := math.Float32bits(v)
	e.encodeUint32(bits)
}

func (e *Encoder) encodeFloat64(v float64) {
	bits := math.Float64bits(v)
	e.encodeUint64(bits)
}

func (e *Encoder) Close() []byte {
	return e.buf.Bytes()
}
