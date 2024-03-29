package nexemab

import (
	"math"
	"unsafe"

	"github.com/nexema/go/buffer"
)

type Encoder struct {
	buf buffer.Buffer
}

func NewEncoder(buffer buffer.Buffer, cap ...int) *Encoder {
	// capacity := 24
	// if len(cap) > 0 {
	// 	capacity = cap[0]
	// }

	// todo: accept initial buffer size to avoid unneccessary grows
	// buffer := make([]byte, 0, capacity)c
	// return &Encoder{buf: new(buffer.Buffer)}
	return &Encoder{buf: buffer}
}

func (e *Encoder) EncodeBool(v bool) {
	if v {
		e.buf.WriteByte(boolTrue)
	} else {
		e.buf.WriteByte(boolFalse)
	}
}

func (e *Encoder) EncodeNull() {
	e.buf.WriteByte(null)
}

func (e *Encoder) EncodeString(input string) {
	// todo: found a better, performant safe way of doing this
	inputLen := len(input)
	buf := *(*[]byte)(unsafe.Pointer(&struct {
		string
		int
	}{input, inputLen}))
	e.EncodeVarint(int64(inputLen))
	e.buf.Write(buf)
}

func (e *Encoder) EncodeUint8(v uint8) {
	e.buf.WriteByte(v)
}

func (e *Encoder) EncodeInt8(v int8) {
	e.EncodeUint8(uint8(v))
}

func (e *Encoder) EncodeUint16(v uint16) {
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) EncodeInt16(v int16) {
	e.EncodeUint16(uint16(v))
}

func (e *Encoder) EncodeUint32(v uint32) {
	e.buf.WriteByte(byte(v >> 24))
	e.buf.WriteByte(byte(v >> 16))
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) EncodeInt32(v int32) {
	e.EncodeUint32(uint32(v))
}

func (e *Encoder) EncodeUint64(v uint64) {
	e.buf.WriteByte(byte(v >> 56))
	e.buf.WriteByte(byte(v >> 48))
	e.buf.WriteByte(byte(v >> 40))
	e.buf.WriteByte(byte(v >> 32))
	e.buf.WriteByte(byte(v >> 24))
	e.buf.WriteByte(byte(v >> 16))
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) EncodeInt64(v int64) {
	e.buf.WriteByte(byte(v >> 56))
	e.buf.WriteByte(byte(v >> 48))
	e.buf.WriteByte(byte(v >> 40))
	e.buf.WriteByte(byte(v >> 32))
	e.buf.WriteByte(byte(v >> 24))
	e.buf.WriteByte(byte(v >> 16))
	e.buf.WriteByte(byte(v >> 8))
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) EncodeUvarint(v uint64) {
	i := 0
	for v >= 0x80 {
		e.buf.WriteByte(byte(v) | 0x80)
		v >>= 7
		i++
	}
	e.buf.WriteByte(byte(v))
}

func (e *Encoder) EncodeVarint(v int64) {
	// code internally used by binary.PutVarint, but modified to write directly to the underlying buffer
	ux := uint64(v) << 1
	if v < 0 {
		ux = ^ux
	}

	e.EncodeUvarint(ux)
}

func (e *Encoder) EncodeFloat32(v float32) {
	bits := math.Float32bits(v)
	e.EncodeUint32(bits)
}

func (e *Encoder) EncodeFloat64(v float64) {
	bits := math.Float64bits(v)
	e.EncodeUint64(bits)
}

func (e *Encoder) EncodeBinary(buffer []byte) {
	e.EncodeVarint(int64(len(buffer)))
	e.buf.Write(buffer)
}

// BeginArray writes a byte that represents that the next value will be an array, followed by its length.
func (e *Encoder) BeginArray(length int64) {
	e.buf.WriteByte(arrayBegin)
	e.EncodeVarint(length)
}

// BeginMap writes a byte that represents that the next value will be a map, followed by its length.
func (e *Encoder) BeginMap(length int64) {
	e.buf.WriteByte(mapBegin)
	e.EncodeVarint(length)
}
func (e *Encoder) Close() []byte {
	return e.buf.Bytes()
}
