package test_files

import (
	"bytes"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type Nullables struct {
	a1  bool
	a2  string
	a3  uint8
	a4  uint16
	a5  uint32
	a6  uint64
	a7  int8
	a8  int16
	a9  int32
	a10 int64
	a11 float32
	a12 float64
	a13 []byte
	a14 []bool
	a15 map[string]float32
}

func (u *Nullables) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	err = writer.EncodeBool(u.a1)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeString(u.a2)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint8(u.a3)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint16(u.a4)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint32(u.a5)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeUint64(u.a6)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt8(u.a7)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt16(u.a8)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt32(u.a9)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeInt64(u.a10)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeFloat32(u.a11)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeFloat64(u.a12)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeBytes(u.a13)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeArrayLen(len(u.a14))
	if err != nil {
		return nil, err
	}
	for _, v := range u.a14 {
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.a15))
	if err != nil {
		return nil, err
	}
	for k, v := range u.a15 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}
func (u *Nullables) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}
	return buf
}
func (u *Nullables) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := v5.NewDecoder(reader)
	var err error
	u.a1, err = decoder.DecodeBool()
	if err != nil {
		return err
	}
	u.a2, err = decoder.DecodeString()
	if err != nil {
		return err
	}
	u.a3, err = decoder.DecodeUint8()
	if err != nil {
		return err
	}
	u.a4, err = decoder.DecodeUint16()
	if err != nil {
		return err
	}
	u.a5, err = decoder.DecodeUint32()
	if err != nil {
		return err
	}
	u.a6, err = decoder.DecodeUint64()
	if err != nil {
		return err
	}
	u.a7, err = decoder.DecodeInt8()
	if err != nil {
		return err
	}
	u.a8, err = decoder.DecodeInt16()
	if err != nil {
		return err
	}
	u.a9, err = decoder.DecodeInt32()
	if err != nil {
		return err
	}
	u.a10, err = decoder.DecodeInt64()
	if err != nil {
		return err
	}
	u.a11, err = decoder.DecodeFloat32()
	if err != nil {
		return err
	}
	u.a12, err = decoder.DecodeFloat64()
	if err != nil {
		return err
	}
	u.a13, err = decoder.DecodeBytes()
	if err != nil {
		return err
	}
	a14Len, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.a14 = make([]bool, a14Len)
	for i := 0; i < a14Len; i++ {
		u.a14[i], err = decoder.DecodeBool()
		if err != nil {
			return err
		}
	}
	a15Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.a15 = make(map[string]float32)
	for i := 0; i < a15Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.a15[k] = v
	}
	return nil
}
func (u *Nullables) MergeUsing(other *Nullables) error {
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
	u.a13 = other.a13
	u.a14 = other.a14
	u.a15 = other.a15
	return nil
}
