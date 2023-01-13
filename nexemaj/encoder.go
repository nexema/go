package nexemaj

import (
	"strconv"

	"github.com/nexema/go/buffer"
)

type Encoder struct {
	buf *buffer.Buffer
}

func NewEncoder() *Encoder {
	return &Encoder{
		buf: buffer.NewBuffer(buffer.BufferConfig{Capacity: 10, InitialBufferSize: 4096}),
	}
}

func (e *Encoder) WriteObjStart() {
	e.write(lBrace)
}

func (e *Encoder) WriteObjEnd() {
	if e.buf.At(-1) == comma {
		e.buf.Set(-1, rBrace)
		e.buf.WriteByte(comma)
	} else {
		e.buf.Write([]byte{rBrace, comma})
	}
}

func (e *Encoder) WriteArrayStart() {
	e.write(lBrack)
}

func (e *Encoder) WriteArrayEnd() {
	if e.buf.At(-1) == comma {
		e.buf.Set(-1, rBrack)
		e.buf.WriteByte(comma)
	} else {
		e.buf.Write([]byte{rBrack, comma})
	}
}

func (e *Encoder) WriteString(v string) {
	e.write(doubleQuote)
	e.writeBuf([]byte(v)) // todo: better and performant way of writing this
	e.write(doubleQuote)
}

func (e *Encoder) WriteUint64(v uint64) {
	str := strconv.FormatUint(v, 10)
	e.writeBuf([]byte(str))
}

func (e *Encoder) WriteInt64(v int64) {
	str := strconv.FormatInt(v, 10)
	e.writeBuf([]byte(str))
}

func (e *Encoder) WriteFloat64(v float64) {
	str := strconv.FormatFloat(v, 'G', -1, 64)
	e.writeBuf([]byte(str))
}

func (e *Encoder) WriteFloat32(v float32) {
	str := strconv.FormatFloat(float64(v), 'G', -1, 32)
	e.writeBuf([]byte(str))
}

func (e *Encoder) WriteStringKey(v string) {
	e.WriteString(v)
	e.WriteColon()
}

func (e *Encoder) WriteBool(v bool) {
	if v {
		e.writeBuf([]byte{letterT, letterR, letterU, letterE})
	} else {
		e.writeBuf([]byte{letterF, letterA, letterL, letterS, letterE})
	}
}

func (e *Encoder) WriteNull() {
	e.writeBuf([]byte{letterN, letterU, letterL, letterL})
}

func (e *Encoder) WriteColon() {
	e.buf.WriteByte(colon)
}

func (e *Encoder) WriteComma() {
	e.buf.WriteByte(comma)
}

func (e *Encoder) String() string {
	buffer := e.buf.Bytes()
	if buffer[len(buffer)-1] == comma {
		buffer = buffer[:len(buffer)-1]
	}
	return string(buffer)
}

func (e *Encoder) write(b byte) {
	e.buf.WriteByte(b)
}

func (e *Encoder) writeBuf(b []byte) {
	e.buf.Write(b)
}
