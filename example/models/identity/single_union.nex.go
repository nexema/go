package identity

import (
	"bytes"
	"github.com/nexema/go/runtime"
	"io"
)

type SingleUnion struct {
	value      interface{}
	fieldIndex int64
}

type SingleUnionWhichField int8

const (
	SingleUnion_NotSet SingleUnionWhichField = -1
	SingleUnion_Field1 SingleUnionWhichField = 0
	SingleUnion_Field2 SingleUnionWhichField = 1
	SingleUnion_Field3 SingleUnionWhichField = 2
)

type singleUnionBuilder struct{}

var SingleUnionBuilder *singleUnionBuilder = &singleUnionBuilder{}

func (u *SingleUnion) IsSet() bool {
	return u.fieldIndex != -1
}

func (u *SingleUnion) Clear() {
	u.value = nil
	u.fieldIndex = -1
}

func (u *SingleUnion) WhichField() SingleUnionWhichField {
	return SingleUnionWhichField(u.fieldIndex)
}

func (u *SingleUnion) CurrentValue() interface{} {
	return u.value
}

func (*singleUnionBuilder) Field1(value string) *SingleUnion {
	return &SingleUnion{
		value:      value,
		fieldIndex: 0,
	}
}

func (u *SingleUnion) SetField1(value string) {
	u.value = value
	u.fieldIndex = 0
}

func (*singleUnionBuilder) Field2(value bool) *SingleUnion {
	return &SingleUnion{
		value:      value,
		fieldIndex: 1,
	}
}

func (u *SingleUnion) SetField2(value bool) {
	u.value = value
	u.fieldIndex = 1
}

func (*singleUnionBuilder) Field3(value []runtime.Nullable[string]) *SingleUnion {
	return &SingleUnion{
		value:      value,
		fieldIndex: 2,
	}
}

func (u *SingleUnion) SetField3(value []runtime.Nullable[string]) {
	u.value = value
	u.fieldIndex = 2
}

func (u *SingleUnion) MergeUsing(other *SingleUnion) error {
	u.fieldIndex = other.fieldIndex
	u.value = other.value
	return nil
}

func (u *SingleUnion) Clone() *SingleUnion {
	return &SingleUnion{
		fieldIndex: u.fieldIndex,
		value:      u.value,
	}
}

func (u *SingleUnion) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	encoder.EncodeVarint(u.fieldIndex)
	switch u.fieldIndex {

	case 0:

		encoder.EncodeString(u.value.(string))

	case 1:

		encoder.EncodeBool(u.value.(bool))

	case 2:

	}

	return encoder.TakeBytes(), nil
}

func (u *SingleUnion) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *SingleUnion) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	u.fieldIndex, err = decoder.DecodeVarint()
	if err != nil {
		return err
	}

	switch u.fieldIndex {
	case -1:
		u.value = nil

	case 0:

		u.value, err = decoder.DecodeString()
		if err != nil {
			return err
		}

	case 1:

		u.value, err = decoder.DecodeBool()
		if err != nil {
			return err
		}

	case 2:

		field3ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}
		value := make([]runtime.Nullable[string], field3ArrayLen)
		for i := int64(0); i < field3ArrayLen; i++ {
			if decoder.IsNextNull() {
				value[i] = runtime.NewNull[string]()
			} else {
				field3, err := decoder.DecodeString()
				if err != nil {
					return err
				}

				value[i] = runtime.NewNullable(field3)
			}
		}
		u.value = value

	}

	return nil
}

func (u *SingleUnion) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *SingleUnion) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}
