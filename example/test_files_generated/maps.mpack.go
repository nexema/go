package test_files

import (
	"bytes"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type MapTest struct {
	a1  map[string]bool
	a2  map[string]string
	a3  map[string]uint8
	a4  map[string]uint16
	a5  map[string]uint32
	a6  map[string]uint64
	a7  map[string]int8
	a8  map[string]int16
	a9  map[string]int32
	a10 map[string]int64
	a11 map[string]float32
	a12 map[string]float64
	b1  map[uint8]bool
	b2  map[uint8]string
	b3  map[uint8]uint8
	b4  map[uint8]uint16
	b5  map[uint8]uint32
	b6  map[uint8]uint64
	b7  map[uint8]int8
	b8  map[uint8]int16
	b9  map[uint8]int32
	b10 map[uint8]int64
	b11 map[uint8]float32
	b12 map[uint8]float64
	c1  map[uint16]bool
	c2  map[uint16]string
	c3  map[uint16]uint8
	c4  map[uint16]uint16
	c5  map[uint16]uint32
	c6  map[uint16]uint64
	c7  map[uint16]int8
	c8  map[uint16]int16
	c9  map[uint16]int32
	c10 map[uint16]int64
	c11 map[uint16]float32
	c12 map[uint16]float64
	d1  map[uint32]bool
	d2  map[uint32]string
	d3  map[uint32]uint8
	d4  map[uint32]uint16
	d5  map[uint32]uint32
	d6  map[uint32]uint64
	d7  map[uint32]int8
	d8  map[uint32]int16
	d9  map[uint32]int32
	d10 map[uint32]int64
	d11 map[uint32]float32
	d12 map[uint32]float64
	e1  map[uint64]bool
	e2  map[uint64]string
	e3  map[uint64]uint8
	e4  map[uint64]uint16
	e5  map[uint64]uint32
	e6  map[uint64]uint64
	e7  map[uint64]int8
	e8  map[uint64]int16
	e9  map[uint64]int32
	e10 map[uint64]int64
	e11 map[uint64]float32
	e12 map[uint64]float64
	f1  map[int8]bool
	f2  map[int8]string
	f3  map[int8]uint8
	f4  map[int8]uint16
	f5  map[int8]uint32
	f6  map[int8]uint64
	f7  map[int8]int8
	f8  map[int8]int16
	f9  map[int8]int32
	f10 map[int8]int64
	f11 map[int8]float32
	f12 map[int8]float64
	g1  map[int16]bool
	g2  map[int16]string
	g3  map[int16]uint8
	g4  map[int16]uint16
	g5  map[int16]uint32
	g6  map[int16]uint64
	g7  map[int16]int8
	g8  map[int16]int16
	g9  map[int16]int32
	g10 map[int16]int64
	g11 map[int16]float32
	g12 map[int16]float64
	h1  map[int32]bool
	h2  map[int32]string
	h3  map[int32]uint8
	h4  map[int32]uint16
	h5  map[int32]uint32
	h6  map[int32]uint64
	h7  map[int32]int8
	h8  map[int32]int16
	h9  map[int32]int32
	h10 map[int32]int64
	h11 map[int32]float32
	h12 map[int32]float64
	i1  map[int64]bool
	i2  map[int64]string
	i3  map[int64]uint8
	i4  map[int64]uint16
	i5  map[int64]uint32
	i6  map[int64]uint64
	i7  map[int64]int8
	i8  map[int64]int16
	i9  map[int64]int32
	i10 map[int64]int64
	i11 map[int64]float32
	i12 map[int64]float64
	j1  map[float32]bool
	j2  map[float32]string
	j3  map[float32]uint8
	j4  map[float32]uint16
	j5  map[float32]uint32
	j6  map[float32]uint64
	j7  map[float32]int8
	j8  map[float32]int16
	j9  map[float32]int32
	j10 map[float32]int64
	j11 map[float32]float32
	j12 map[float32]float64
	k1  map[float64]bool
	k2  map[float64]string
	k3  map[float64]uint8
	k4  map[float64]uint16
	k5  map[float64]uint32
	k6  map[float64]uint64
	k7  map[float64]int8
	k8  map[float64]int16
	k9  map[float64]int32
	k10 map[float64]int64
	k11 map[float64]float32
	k12 map[float64]float64
}

func (u *MapTest) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	err = writer.EncodeMapLen(len(u.a1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a1 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a2 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a3 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a4 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a5 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a6 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a7 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a8 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a9 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a10 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a11 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a12 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b1 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b2 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b3 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b4 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b5 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b6 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b7 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b8 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b9 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b10 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b11 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.b12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.b12 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c1 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c2 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c3 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c4 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c5 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c6 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c7 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c8 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c9 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c10 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c11 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.c12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.c12 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d1 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d2 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d3 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d4 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d5 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d6 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d7 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d8 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d9 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d10 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d11 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.d12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.d12 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e1 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e2 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e3 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e4 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e5 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e6 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e7 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e8 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e9 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e10 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e11 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.e12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.e12 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f1 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f2 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f3 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f4 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f5 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f6 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f7 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f8 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f9 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f10 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f11 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.f12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.f12 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g1 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g2 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g3 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g4 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g5 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g6 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g7 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g8 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g9 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g10 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g11 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.g12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.g12 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h1 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h2 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h3 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h4 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h5 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h6 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h7 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h8 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h9 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h10 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h11 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.h12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.h12 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i1 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i2 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i3 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i4 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i5 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i6 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i7 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i8 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i9 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i10 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i11 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.i12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.i12 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j1 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j2 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j3 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j4 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j5 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j6 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j7 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j8 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j9 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j10 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j11 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.j12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.j12 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k1))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k1 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k2))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k2 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k3))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k3 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k4))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k4 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k5))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k5 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k6))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k6 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k7))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k7 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k8))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k8 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k9))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k9 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k10))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k10 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k11))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k11 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.k12))
	if err != nil {
		return nil, err
	}
	for k, v := range u.k12 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
func (u *MapTest) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}
	return buf
}
func (u *MapTest) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := v5.NewDecoder(reader)
	var err error
	a1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a1 = make(map[string]bool)
	for i := 0; i < a1Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.a1[k] = v
	}
	a2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a2 = make(map[string]string)
	for i := 0; i < a2Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.a2[k] = v
	}
	a3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a3 = make(map[string]uint8)
	for i := 0; i < a3Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.a3[k] = v
	}
	a4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a4 = make(map[string]uint16)
	for i := 0; i < a4Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.a4[k] = v
	}
	a5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a5 = make(map[string]uint32)
	for i := 0; i < a5Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.a5[k] = v
	}
	a6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a6 = make(map[string]uint64)
	for i := 0; i < a6Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.a6[k] = v
	}
	a7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a7 = make(map[string]int8)
	for i := 0; i < a7Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.a7[k] = v
	}
	a8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a8 = make(map[string]int16)
	for i := 0; i < a8Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.a8[k] = v
	}
	a9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a9 = make(map[string]int32)
	for i := 0; i < a9Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.a9[k] = v
	}
	a10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a10 = make(map[string]int64)
	for i := 0; i < a10Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.a10[k] = v
	}
	a11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a11 = make(map[string]float32)
	for i := 0; i < a11Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.a11[k] = v
	}
	a12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a12 = make(map[string]float64)
	for i := 0; i < a12Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.a12[k] = v
	}
	b1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b1 = make(map[uint8]bool)
	for i := 0; i < b1Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.b1[k] = v
	}
	b2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b2 = make(map[uint8]string)
	for i := 0; i < b2Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.b2[k] = v
	}
	b3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b3 = make(map[uint8]uint8)
	for i := 0; i < b3Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.b3[k] = v
	}
	b4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b4 = make(map[uint8]uint16)
	for i := 0; i < b4Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.b4[k] = v
	}
	b5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b5 = make(map[uint8]uint32)
	for i := 0; i < b5Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.b5[k] = v
	}
	b6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b6 = make(map[uint8]uint64)
	for i := 0; i < b6Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.b6[k] = v
	}
	b7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b7 = make(map[uint8]int8)
	for i := 0; i < b7Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.b7[k] = v
	}
	b8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b8 = make(map[uint8]int16)
	for i := 0; i < b8Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.b8[k] = v
	}
	b9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b9 = make(map[uint8]int32)
	for i := 0; i < b9Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.b9[k] = v
	}
	b10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b10 = make(map[uint8]int64)
	for i := 0; i < b10Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.b10[k] = v
	}
	b11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b11 = make(map[uint8]float32)
	for i := 0; i < b11Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.b11[k] = v
	}
	b12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.b12 = make(map[uint8]float64)
	for i := 0; i < b12Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.b12[k] = v
	}
	c1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c1 = make(map[uint16]bool)
	for i := 0; i < c1Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.c1[k] = v
	}
	c2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c2 = make(map[uint16]string)
	for i := 0; i < c2Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.c2[k] = v
	}
	c3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c3 = make(map[uint16]uint8)
	for i := 0; i < c3Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.c3[k] = v
	}
	c4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c4 = make(map[uint16]uint16)
	for i := 0; i < c4Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.c4[k] = v
	}
	c5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c5 = make(map[uint16]uint32)
	for i := 0; i < c5Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.c5[k] = v
	}
	c6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c6 = make(map[uint16]uint64)
	for i := 0; i < c6Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.c6[k] = v
	}
	c7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c7 = make(map[uint16]int8)
	for i := 0; i < c7Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.c7[k] = v
	}
	c8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c8 = make(map[uint16]int16)
	for i := 0; i < c8Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.c8[k] = v
	}
	c9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c9 = make(map[uint16]int32)
	for i := 0; i < c9Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.c9[k] = v
	}
	c10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c10 = make(map[uint16]int64)
	for i := 0; i < c10Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.c10[k] = v
	}
	c11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c11 = make(map[uint16]float32)
	for i := 0; i < c11Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.c11[k] = v
	}
	c12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.c12 = make(map[uint16]float64)
	for i := 0; i < c12Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.c12[k] = v
	}
	d1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d1 = make(map[uint32]bool)
	for i := 0; i < d1Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.d1[k] = v
	}
	d2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d2 = make(map[uint32]string)
	for i := 0; i < d2Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.d2[k] = v
	}
	d3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d3 = make(map[uint32]uint8)
	for i := 0; i < d3Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.d3[k] = v
	}
	d4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d4 = make(map[uint32]uint16)
	for i := 0; i < d4Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.d4[k] = v
	}
	d5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d5 = make(map[uint32]uint32)
	for i := 0; i < d5Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.d5[k] = v
	}
	d6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d6 = make(map[uint32]uint64)
	for i := 0; i < d6Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.d6[k] = v
	}
	d7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d7 = make(map[uint32]int8)
	for i := 0; i < d7Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.d7[k] = v
	}
	d8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d8 = make(map[uint32]int16)
	for i := 0; i < d8Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.d8[k] = v
	}
	d9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d9 = make(map[uint32]int32)
	for i := 0; i < d9Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.d9[k] = v
	}
	d10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d10 = make(map[uint32]int64)
	for i := 0; i < d10Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.d10[k] = v
	}
	d11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d11 = make(map[uint32]float32)
	for i := 0; i < d11Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.d11[k] = v
	}
	d12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.d12 = make(map[uint32]float64)
	for i := 0; i < d12Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.d12[k] = v
	}
	e1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e1 = make(map[uint64]bool)
	for i := 0; i < e1Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.e1[k] = v
	}
	e2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e2 = make(map[uint64]string)
	for i := 0; i < e2Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.e2[k] = v
	}
	e3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e3 = make(map[uint64]uint8)
	for i := 0; i < e3Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.e3[k] = v
	}
	e4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e4 = make(map[uint64]uint16)
	for i := 0; i < e4Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.e4[k] = v
	}
	e5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e5 = make(map[uint64]uint32)
	for i := 0; i < e5Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.e5[k] = v
	}
	e6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e6 = make(map[uint64]uint64)
	for i := 0; i < e6Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.e6[k] = v
	}
	e7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e7 = make(map[uint64]int8)
	for i := 0; i < e7Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.e7[k] = v
	}
	e8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e8 = make(map[uint64]int16)
	for i := 0; i < e8Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.e8[k] = v
	}
	e9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e9 = make(map[uint64]int32)
	for i := 0; i < e9Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.e9[k] = v
	}
	e10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e10 = make(map[uint64]int64)
	for i := 0; i < e10Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.e10[k] = v
	}
	e11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e11 = make(map[uint64]float32)
	for i := 0; i < e11Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.e11[k] = v
	}
	e12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.e12 = make(map[uint64]float64)
	for i := 0; i < e12Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.e12[k] = v
	}
	f1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f1 = make(map[int8]bool)
	for i := 0; i < f1Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.f1[k] = v
	}
	f2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f2 = make(map[int8]string)
	for i := 0; i < f2Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.f2[k] = v
	}
	f3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f3 = make(map[int8]uint8)
	for i := 0; i < f3Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.f3[k] = v
	}
	f4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f4 = make(map[int8]uint16)
	for i := 0; i < f4Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.f4[k] = v
	}
	f5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f5 = make(map[int8]uint32)
	for i := 0; i < f5Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.f5[k] = v
	}
	f6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f6 = make(map[int8]uint64)
	for i := 0; i < f6Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.f6[k] = v
	}
	f7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f7 = make(map[int8]int8)
	for i := 0; i < f7Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.f7[k] = v
	}
	f8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f8 = make(map[int8]int16)
	for i := 0; i < f8Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.f8[k] = v
	}
	f9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f9 = make(map[int8]int32)
	for i := 0; i < f9Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.f9[k] = v
	}
	f10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f10 = make(map[int8]int64)
	for i := 0; i < f10Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.f10[k] = v
	}
	f11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f11 = make(map[int8]float32)
	for i := 0; i < f11Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.f11[k] = v
	}
	f12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.f12 = make(map[int8]float64)
	for i := 0; i < f12Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.f12[k] = v
	}
	g1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g1 = make(map[int16]bool)
	for i := 0; i < g1Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.g1[k] = v
	}
	g2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g2 = make(map[int16]string)
	for i := 0; i < g2Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.g2[k] = v
	}
	g3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g3 = make(map[int16]uint8)
	for i := 0; i < g3Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.g3[k] = v
	}
	g4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g4 = make(map[int16]uint16)
	for i := 0; i < g4Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.g4[k] = v
	}
	g5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g5 = make(map[int16]uint32)
	for i := 0; i < g5Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.g5[k] = v
	}
	g6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g6 = make(map[int16]uint64)
	for i := 0; i < g6Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.g6[k] = v
	}
	g7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g7 = make(map[int16]int8)
	for i := 0; i < g7Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.g7[k] = v
	}
	g8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g8 = make(map[int16]int16)
	for i := 0; i < g8Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.g8[k] = v
	}
	g9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g9 = make(map[int16]int32)
	for i := 0; i < g9Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.g9[k] = v
	}
	g10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g10 = make(map[int16]int64)
	for i := 0; i < g10Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.g10[k] = v
	}
	g11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g11 = make(map[int16]float32)
	for i := 0; i < g11Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.g11[k] = v
	}
	g12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.g12 = make(map[int16]float64)
	for i := 0; i < g12Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.g12[k] = v
	}
	h1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h1 = make(map[int32]bool)
	for i := 0; i < h1Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.h1[k] = v
	}
	h2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h2 = make(map[int32]string)
	for i := 0; i < h2Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.h2[k] = v
	}
	h3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h3 = make(map[int32]uint8)
	for i := 0; i < h3Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.h3[k] = v
	}
	h4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h4 = make(map[int32]uint16)
	for i := 0; i < h4Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.h4[k] = v
	}
	h5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h5 = make(map[int32]uint32)
	for i := 0; i < h5Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.h5[k] = v
	}
	h6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h6 = make(map[int32]uint64)
	for i := 0; i < h6Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.h6[k] = v
	}
	h7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h7 = make(map[int32]int8)
	for i := 0; i < h7Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.h7[k] = v
	}
	h8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h8 = make(map[int32]int16)
	for i := 0; i < h8Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.h8[k] = v
	}
	h9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h9 = make(map[int32]int32)
	for i := 0; i < h9Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.h9[k] = v
	}
	h10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h10 = make(map[int32]int64)
	for i := 0; i < h10Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.h10[k] = v
	}
	h11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h11 = make(map[int32]float32)
	for i := 0; i < h11Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.h11[k] = v
	}
	h12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.h12 = make(map[int32]float64)
	for i := 0; i < h12Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.h12[k] = v
	}
	i1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i1 = make(map[int64]bool)
	for i := 0; i < i1Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.i1[k] = v
	}
	i2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i2 = make(map[int64]string)
	for i := 0; i < i2Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.i2[k] = v
	}
	i3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i3 = make(map[int64]uint8)
	for i := 0; i < i3Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.i3[k] = v
	}
	i4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i4 = make(map[int64]uint16)
	for i := 0; i < i4Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.i4[k] = v
	}
	i5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i5 = make(map[int64]uint32)
	for i := 0; i < i5Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.i5[k] = v
	}
	i6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i6 = make(map[int64]uint64)
	for i := 0; i < i6Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.i6[k] = v
	}
	i7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i7 = make(map[int64]int8)
	for i := 0; i < i7Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.i7[k] = v
	}
	i8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i8 = make(map[int64]int16)
	for i := 0; i < i8Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.i8[k] = v
	}
	i9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i9 = make(map[int64]int32)
	for i := 0; i < i9Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.i9[k] = v
	}
	i10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i10 = make(map[int64]int64)
	for i := 0; i < i10Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.i10[k] = v
	}
	i11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i11 = make(map[int64]float32)
	for i := 0; i < i11Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.i11[k] = v
	}
	i12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.i12 = make(map[int64]float64)
	for i := 0; i < i12Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.i12[k] = v
	}
	j1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j1 = make(map[float32]bool)
	for i := 0; i < j1Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.j1[k] = v
	}
	j2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j2 = make(map[float32]string)
	for i := 0; i < j2Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.j2[k] = v
	}
	j3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j3 = make(map[float32]uint8)
	for i := 0; i < j3Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.j3[k] = v
	}
	j4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j4 = make(map[float32]uint16)
	for i := 0; i < j4Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.j4[k] = v
	}
	j5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j5 = make(map[float32]uint32)
	for i := 0; i < j5Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.j5[k] = v
	}
	j6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j6 = make(map[float32]uint64)
	for i := 0; i < j6Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.j6[k] = v
	}
	j7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j7 = make(map[float32]int8)
	for i := 0; i < j7Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.j7[k] = v
	}
	j8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j8 = make(map[float32]int16)
	for i := 0; i < j8Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.j8[k] = v
	}
	j9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j9 = make(map[float32]int32)
	for i := 0; i < j9Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.j9[k] = v
	}
	j10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j10 = make(map[float32]int64)
	for i := 0; i < j10Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.j10[k] = v
	}
	j11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j11 = make(map[float32]float32)
	for i := 0; i < j11Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.j11[k] = v
	}
	j12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.j12 = make(map[float32]float64)
	for i := 0; i < j12Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.j12[k] = v
	}
	k1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k1 = make(map[float64]bool)
	for i := 0; i < k1Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}
		u.k1[k] = v
	}
	k2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k2 = make(map[float64]string)
	for i := 0; i < k2Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		u.k2[k] = v
	}
	k3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k3 = make(map[float64]uint8)
	for i := 0; i < k3Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}
		u.k3[k] = v
	}
	k4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k4 = make(map[float64]uint16)
	for i := 0; i < k4Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}
		u.k4[k] = v
	}
	k5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k5 = make(map[float64]uint32)
	for i := 0; i < k5Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}
		u.k5[k] = v
	}
	k6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k6 = make(map[float64]uint64)
	for i := 0; i < k6Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}
		u.k6[k] = v
	}
	k7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k7 = make(map[float64]int8)
	for i := 0; i < k7Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}
		u.k7[k] = v
	}
	k8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k8 = make(map[float64]int16)
	for i := 0; i < k8Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}
		u.k8[k] = v
	}
	k9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k9 = make(map[float64]int32)
	for i := 0; i < k9Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}
		u.k9[k] = v
	}
	k10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k10 = make(map[float64]int64)
	for i := 0; i < k10Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}
		u.k10[k] = v
	}
	k11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k11 = make(map[float64]float32)
	for i := 0; i < k11Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.k11[k] = v
	}
	k12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.k12 = make(map[float64]float64)
	for i := 0; i < k12Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}
		u.k12[k] = v
	}
	return nil
}
func (u *MapTest) MergeUsing(other *MapTest) error {
	u.a1 = other.a1
	u.a2 = other.a2
	u.a3 = other.a3
	u.a4 = other.a4
	u.a5 = other.a5
	u.a6 = other.a6
	u.a7 = other.a7
	u.a8 = other.a8
	u.a9 = other.a9
	u.a10 = other.a10
	u.a11 = other.a11
	u.a12 = other.a12
	u.b1 = other.b1
	u.b2 = other.b2
	u.b3 = other.b3
	u.b4 = other.b4
	u.b5 = other.b5
	u.b6 = other.b6
	u.b7 = other.b7
	u.b8 = other.b8
	u.b9 = other.b9
	u.b10 = other.b10
	u.b11 = other.b11
	u.b12 = other.b12
	u.c1 = other.c1
	u.c2 = other.c2
	u.c3 = other.c3
	u.c4 = other.c4
	u.c5 = other.c5
	u.c6 = other.c6
	u.c7 = other.c7
	u.c8 = other.c8
	u.c9 = other.c9
	u.c10 = other.c10
	u.c11 = other.c11
	u.c12 = other.c12
	u.d1 = other.d1
	u.d2 = other.d2
	u.d3 = other.d3
	u.d4 = other.d4
	u.d5 = other.d5
	u.d6 = other.d6
	u.d7 = other.d7
	u.d8 = other.d8
	u.d9 = other.d9
	u.d10 = other.d10
	u.d11 = other.d11
	u.d12 = other.d12
	u.e1 = other.e1
	u.e2 = other.e2
	u.e3 = other.e3
	u.e4 = other.e4
	u.e5 = other.e5
	u.e6 = other.e6
	u.e7 = other.e7
	u.e8 = other.e8
	u.e9 = other.e9
	u.e10 = other.e10
	u.e11 = other.e11
	u.e12 = other.e12
	u.f1 = other.f1
	u.f2 = other.f2
	u.f3 = other.f3
	u.f4 = other.f4
	u.f5 = other.f5
	u.f6 = other.f6
	u.f7 = other.f7
	u.f8 = other.f8
	u.f9 = other.f9
	u.f10 = other.f10
	u.f11 = other.f11
	u.f12 = other.f12
	u.g1 = other.g1
	u.g2 = other.g2
	u.g3 = other.g3
	u.g4 = other.g4
	u.g5 = other.g5
	u.g6 = other.g6
	u.g7 = other.g7
	u.g8 = other.g8
	u.g9 = other.g9
	u.g10 = other.g10
	u.g11 = other.g11
	u.g12 = other.g12
	u.h1 = other.h1
	u.h2 = other.h2
	u.h3 = other.h3
	u.h4 = other.h4
	u.h5 = other.h5
	u.h6 = other.h6
	u.h7 = other.h7
	u.h8 = other.h8
	u.h9 = other.h9
	u.h10 = other.h10
	u.h11 = other.h11
	u.h12 = other.h12
	u.i1 = other.i1
	u.i2 = other.i2
	u.i3 = other.i3
	u.i4 = other.i4
	u.i5 = other.i5
	u.i6 = other.i6
	u.i7 = other.i7
	u.i8 = other.i8
	u.i9 = other.i9
	u.i10 = other.i10
	u.i11 = other.i11
	u.i12 = other.i12
	u.j1 = other.j1
	u.j2 = other.j2
	u.j3 = other.j3
	u.j4 = other.j4
	u.j5 = other.j5
	u.j6 = other.j6
	u.j7 = other.j7
	u.j8 = other.j8
	u.j9 = other.j9
	u.j10 = other.j10
	u.j11 = other.j11
	u.j12 = other.j12
	u.k1 = other.k1
	u.k2 = other.k2
	u.k3 = other.k3
	u.k4 = other.k4
	u.k5 = other.k5
	u.k6 = other.k6
	u.k7 = other.k7
	u.k8 = other.k8
	u.k9 = other.k9
	u.k10 = other.k10
	u.k11 = other.k11
	u.k12 = other.k12
	return nil
}
