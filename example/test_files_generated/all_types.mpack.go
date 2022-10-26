package test_files

import (
	"bytes"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type AllTypes struct {
	a bool
	b string
	c uint8
	d uint16
	e uint32
	f uint64
	g int8
	h int16
	i int32
	j int64
	k float32
	l float64
	m []byte
	n []bool
	o []string
	p []uint8
	q []uint16
	r []uint32
	s []uint64
	t []int8
	u []int16
	v []int32
	w []int64
	x []float32
	y []float64
}

func (u *AllTypes) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	err = writer.EncodeBool(u.a)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeString(u.b)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint8(u.c)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint16(u.d)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint32(u.e)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint64(u.f)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt8(u.g)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt16(u.h)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt32(u.i)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt64(u.j)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeFloat32(u.k)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeFloat64(u.l)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeBytes(u.m)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeArrayLen(len(u.n))
	if err != nil {
		return nil, err
	}
	for _, v := range u.n {
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.o))
	if err != nil {
		return nil, err
	}
	for _, v := range u.o {
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.p))
	if err != nil {
		return nil, err
	}
	for _, v := range u.p {
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.q))
	if err != nil {
		return nil, err
	}
	for _, v := range u.q {
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.r))
	if err != nil {
		return nil, err
	}
	for _, v := range u.r {
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.s))
	if err != nil {
		return nil, err
	}
	for _, v := range u.s {
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.t))
	if err != nil {
		return nil, err
	}
	for _, v := range u.t {
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.u))
	if err != nil {
		return nil, err
	}
	for _, v := range u.u {
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.v))
	if err != nil {
		return nil, err
	}
	for _, v := range u.v {
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.w))
	if err != nil {
		return nil, err
	}
	for _, v := range u.w {
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.x))
	if err != nil {
		return nil, err
	}
	for _, v := range u.x {
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeArrayLen(len(u.y))
	if err != nil {
		return nil, err
	}
	for _, v := range u.y {
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
func (u *AllTypes) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}
	return buf
}
func (u *AllTypes) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := v5.NewDecoder(reader)
	var err error
	u.a, err = decoder.DecodeBool()
	if err != nil {
		return err
	}
	u.b, err = decoder.DecodeString()
	if err != nil {
		return err
	}
	u.c, err = decoder.DecodeUint8()
	if err != nil {
		return err
	}
	u.d, err = decoder.DecodeUint16()
	if err != nil {
		return err
	}
	u.e, err = decoder.DecodeUint32()
	if err != nil {
		return err
	}
	u.f, err = decoder.DecodeUint64()
	if err != nil {
		return err
	}
	u.g, err = decoder.DecodeInt8()
	if err != nil {
		return err
	}
	u.h, err = decoder.DecodeInt16()
	if err != nil {
		return err
	}
	u.i, err = decoder.DecodeInt32()
	if err != nil {
		return err
	}
	u.j, err = decoder.DecodeInt64()
	if err != nil {
		return err
	}
	u.k, err = decoder.DecodeFloat32()
	if err != nil {
		return err
	}
	u.l, err = decoder.DecodeFloat64()
	if err != nil {
		return err
	}
	u.m, err = decoder.DecodeBytes()
	if err != nil {
		return err
	}
	nLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.n = make([]bool, nLen)
	for i := 0; i < nLen; i++ {
		u.n[i], err = decoder.DecodeBool()
		if err != nil {
			return err
		}
	}
	oLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.o = make([]string, oLen)
	for i := 0; i < oLen; i++ {
		u.o[i], err = decoder.DecodeString()
		if err != nil {
			return err
		}
	}
	pLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.p = make([]uint8, pLen)
	for i := 0; i < pLen; i++ {
		u.p[i], err = decoder.DecodeUint8()
		if err != nil {
			return err
		}
	}
	qLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.q = make([]uint16, qLen)
	for i := 0; i < qLen; i++ {
		u.q[i], err = decoder.DecodeUint16()
		if err != nil {
			return err
		}
	}
	rLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.r = make([]uint32, rLen)
	for i := 0; i < rLen; i++ {
		u.r[i], err = decoder.DecodeUint32()
		if err != nil {
			return err
		}
	}
	sLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.s = make([]uint64, sLen)
	for i := 0; i < sLen; i++ {
		u.s[i], err = decoder.DecodeUint64()
		if err != nil {
			return err
		}
	}
	tLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.t = make([]int8, tLen)
	for i := 0; i < tLen; i++ {
		u.t[i], err = decoder.DecodeInt8()
		if err != nil {
			return err
		}
	}
	uLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.u = make([]int16, uLen)
	for i := 0; i < uLen; i++ {
		u.u[i], err = decoder.DecodeInt16()
		if err != nil {
			return err
		}
	}
	vLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.v = make([]int32, vLen)
	for i := 0; i < vLen; i++ {
		u.v[i], err = decoder.DecodeInt32()
		if err != nil {
			return err
		}
	}
	wLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.w = make([]int64, wLen)
	for i := 0; i < wLen; i++ {
		u.w[i], err = decoder.DecodeInt64()
		if err != nil {
			return err
		}
	}
	xLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.x = make([]float32, xLen)
	for i := 0; i < xLen; i++ {
		u.x[i], err = decoder.DecodeFloat32()
		if err != nil {
			return err
		}
	}
	yLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.y = make([]float64, yLen)
	for i := 0; i < yLen; i++ {
		u.y[i], err = decoder.DecodeFloat64()
		if err != nil {
			return err
		}
	}
	return nil
}
func (u *AllTypes) MergeUsing(other *AllTypes) error {
	u.a = other.a
	u.b = other.b
	u.c = other.c
	u.d = other.d
	u.e = other.e
	u.f = other.f
	u.g = other.g
	u.h = other.h
	u.i = other.i
	u.j = other.j
	u.k = other.k
	u.l = other.l
	u.m = other.m
	u.n = other.n
	u.o = other.o
	u.p = other.p
	u.q = other.q
	u.r = other.r
	u.s = other.s
	u.t = other.t
	u.u = other.u
	u.v = other.v
	u.w = other.w
	u.x = other.x
	u.y = other.y
	return nil
}
