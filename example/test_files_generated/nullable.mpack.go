package test_files

import (
	"bytes"

	runtime "github.com/messagepack-schema/go/runtime"
	msgpack "github.com/messagepack-schema/go/runtime/msgpack"
)

type Nullables struct {
	A1  runtime.Nullable[bool]
	A2  runtime.Nullable[string]
	A3  runtime.Nullable[uint8]
	A4  runtime.Nullable[uint16]
	A5  runtime.Nullable[uint32]
	A6  runtime.Nullable[uint64]
	A7  runtime.Nullable[int8]
	A8  runtime.Nullable[int16]
	A9  runtime.Nullable[int32]
	A10 runtime.Nullable[int64]
	A11 runtime.Nullable[float32]
	A12 runtime.Nullable[float64]
	A13 runtime.Nullable[[]byte]
	A14 []runtime.Nullable[bool]
	A15 map[string]runtime.Nullable[float32]
}

func (u *Nullables) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error
	if u.A1.HasValue() {
		err = writer.EncodeBool(u.A1.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A2.HasValue() {
		err = writer.EncodeString(u.A2.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A3.HasValue() {
		err = writer.EncodeUint8(u.A3.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A4.HasValue() {
		err = writer.EncodeUint16(u.A4.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A5.HasValue() {
		err = writer.EncodeUint32(u.A5.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A6.HasValue() {
		err = writer.EncodeUint64(u.A6.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A7.HasValue() {
		err = writer.EncodeInt8(u.A7.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A8.HasValue() {
		err = writer.EncodeInt16(u.A8.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A9.HasValue() {
		err = writer.EncodeInt32(u.A9.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A10.HasValue() {
		err = writer.EncodeInt64(u.A10.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A11.HasValue() {
		err = writer.EncodeFloat32(u.A11.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A12.HasValue() {
		err = writer.EncodeFloat64(u.A12.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	if u.A13.HasValue() {
		err = writer.EncodeBytes(u.A13.GetValue())
	} else {
		err = writer.EncodeNil()
	}
	if err != nil {
		return nil, err
	}
	err = writer.EncodeArrayLen(len(u.A14))
	if err != nil {
		return nil, err
	}
	for _, v := range u.A14 {
		if v.HasValue() {
			err = writer.EncodeBool(v.GetValue())
		} else {
			err = writer.EncodeNil()
		}
		if err != nil {
			return nil, err
		}
	}
	err = writer.EncodeMapLen(len(u.A15))
	if err != nil {
		return nil, err
	}
	for k, v := range u.A15 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}
		if v.HasValue() {
			err = writer.EncodeFloat32(v.GetValue())
		} else {
			err = writer.EncodeNil()
		}
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
	decoder := msgpack.NewDecoder(reader)
	var err error
	u.A1, err = decoder.DecodeBool()
	if err != nil {
		return err
	}
	u.A2, err = decoder.DecodeString()
	if err != nil {
		return err
	}
	u.A3, err = decoder.DecodeUint8()
	if err != nil {
		return err
	}
	u.A4, err = decoder.DecodeUint16()
	if err != nil {
		return err
	}
	u.A5, err = decoder.DecodeUint32()
	if err != nil {
		return err
	}
	u.A6, err = decoder.DecodeUint64()
	if err != nil {
		return err
	}
	u.A7, err = decoder.DecodeInt8()
	if err != nil {
		return err
	}
	u.A8, err = decoder.DecodeInt16()
	if err != nil {
		return err
	}
	u.A9, err = decoder.DecodeInt32()
	if err != nil {
		return err
	}
	u.A10, err = decoder.DecodeInt64()
	if err != nil {
		return err
	}
	u.A11, err = decoder.DecodeFloat32()
	if err != nil {
		return err
	}
	u.A12, err = decoder.DecodeFloat64()
	if err != nil {
		return err
	}
	u.A13, err = decoder.DecodeBytes()
	if err != nil {
		return err
	}
	A14Len, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}
	u.A14 = make([]bool, A14Len)
	for i := 0; i < A14Len; i++ {
		u.A14[i], err = decoder.DecodeBool()
		if err != nil {
			return err
		}
	}
	A15Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}
	u.A15 = make(map[string]float32)
	for i := 0; i < A15Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}
		u.A15[k] = v
	}
	return nil
}
func (u *Nullables) MergeUsing(other *Nullables) error {
	u.A1 = other.A1
	u.A2 = other.A2
	u.A3 = other.A3
	u.A4 = other.A4
	u.A5 = other.A5
	u.A6 = other.A6
	u.A7 = other.A7
	u.A8 = other.A8
	u.A9 = other.A9
	u.A10 = other.A10
	u.A11 = other.A11
	u.A12 = other.A12
	u.A13 = other.A13
	u.A14 = other.A14
	u.A15 = other.A15
	return nil
}
