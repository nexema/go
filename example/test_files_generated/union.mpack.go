package test_files

import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type MyUnion struct {
	value      interface{}
	fieldIndex int8
}
type MyUnionWhichField int8

const (
	MyUnion_NotSet MyUnionWhichField = -1
	MyUnion_A      MyUnionWhichField = 1
	MyUnion_B      MyUnionWhichField = 2
	MyUnion_C      MyUnionWhichField = 3
	MyUnion_D      MyUnionWhichField = 4
	MyUnion_E      MyUnionWhichField = 5
	MyUnion_F      MyUnionWhichField = 6
	MyUnion_G      MyUnionWhichField = 7
	MyUnion_H      MyUnionWhichField = 8
	MyUnion_I      MyUnionWhichField = 9
	MyUnion_J      MyUnionWhichField = 10
	MyUnion_K      MyUnionWhichField = 11
	MyUnion_L      MyUnionWhichField = 12
	MyUnion_M      MyUnionWhichField = 13
)

type myUnionBuilder struct{}

var MyUnionBuilder *myUnionBuilder = &myUnionBuilder{}

func (*myUnionBuilder) A(value bool) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 1,
	}
}

func (*myUnionBuilder) B(value string) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 2,
	}
}

func (*myUnionBuilder) C(value uint8) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 3,
	}
}

func (*myUnionBuilder) D(value uint16) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 4,
	}
}

func (*myUnionBuilder) E(value uint32) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 5,
	}
}

func (*myUnionBuilder) F(value uint64) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 6,
	}
}

func (*myUnionBuilder) G(value int8) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 7,
	}
}

func (*myUnionBuilder) H(value int16) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 8,
	}
}

func (*myUnionBuilder) I(value int32) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 9,
	}
}

func (*myUnionBuilder) J(value int64) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 10,
	}
}

func (*myUnionBuilder) K(value float32) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 11,
	}
}

func (*myUnionBuilder) L(value float64) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 12,
	}
}

func (*myUnionBuilder) M(value []byte) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 13,
	}
}

func (u *MyUnion) IsSet() bool {
	return u.fieldIndex != -1
}

func (u *MyUnion) Clear() {
	u.value = nil
	u.fieldIndex = -1
}

func (u *MyUnion) WhichField() MyUnionWhichField {
	return MyUnionWhichField(u.fieldIndex)
}

func (u *MyUnion) CurrentValue() interface{} {
	return u.value
}

func (u MyUnion) SetA(value bool) {
	u.value = value
	u.fieldIndex = 1
}

func (u MyUnion) GetA() bool {
	return u.value.(bool)
}

func (u MyUnion) SetB(value string) {
	u.value = value
	u.fieldIndex = 2
}

func (u MyUnion) GetB() string {
	return u.value.(string)
}

func (u MyUnion) SetC(value uint8) {
	u.value = value
	u.fieldIndex = 3
}

func (u MyUnion) GetC() uint8 {
	return u.value.(uint8)
}

func (u MyUnion) SetD(value uint16) {
	u.value = value
	u.fieldIndex = 4
}

func (u MyUnion) GetD() uint16 {
	return u.value.(uint16)
}

func (u MyUnion) SetE(value uint32) {
	u.value = value
	u.fieldIndex = 5
}

func (u MyUnion) GetE() uint32 {
	return u.value.(uint32)
}

func (u MyUnion) SetF(value uint64) {
	u.value = value
	u.fieldIndex = 6
}

func (u MyUnion) GetF() uint64 {
	return u.value.(uint64)
}

func (u MyUnion) SetG(value int8) {
	u.value = value
	u.fieldIndex = 7
}

func (u MyUnion) GetG() int8 {
	return u.value.(int8)
}

func (u MyUnion) SetH(value int16) {
	u.value = value
	u.fieldIndex = 8
}

func (u MyUnion) GetH() int16 {
	return u.value.(int16)
}

func (u MyUnion) SetI(value int32) {
	u.value = value
	u.fieldIndex = 9
}

func (u MyUnion) GetI() int32 {
	return u.value.(int32)
}

func (u MyUnion) SetJ(value int64) {
	u.value = value
	u.fieldIndex = 10
}

func (u MyUnion) GetJ() int64 {
	return u.value.(int64)
}

func (u MyUnion) SetK(value float32) {
	u.value = value
	u.fieldIndex = 11
}

func (u MyUnion) GetK() float32 {
	return u.value.(float32)
}

func (u MyUnion) SetL(value float64) {
	u.value = value
	u.fieldIndex = 12
}

func (u MyUnion) GetL() float64 {
	return u.value.(float64)
}

func (u MyUnion) SetM(value []byte) {
	u.value = value
	u.fieldIndex = 13
}

func (u MyUnion) GetM() []byte {
	return u.value.([]byte)
}

func (u MyUnion) IsA() bool {
	return u.fieldIndex == 1
}

func (u MyUnion) IsB() bool {
	return u.fieldIndex == 2
}

func (u MyUnion) IsC() bool {
	return u.fieldIndex == 3
}

func (u MyUnion) IsD() bool {
	return u.fieldIndex == 4
}

func (u MyUnion) IsE() bool {
	return u.fieldIndex == 5
}

func (u MyUnion) IsF() bool {
	return u.fieldIndex == 6
}

func (u MyUnion) IsG() bool {
	return u.fieldIndex == 7
}

func (u MyUnion) IsH() bool {
	return u.fieldIndex == 8
}

func (u MyUnion) IsI() bool {
	return u.fieldIndex == 9
}

func (u MyUnion) IsJ() bool {
	return u.fieldIndex == 10
}

func (u MyUnion) IsK() bool {
	return u.fieldIndex == 11
}

func (u MyUnion) IsL() bool {
	return u.fieldIndex == 12
}

func (u MyUnion) IsM() bool {
	return u.fieldIndex == 13
}

func (u *MyUnion) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	err = writer.EncodeInt8(u.fieldIndex)
	if err != nil {
		return nil, err
	}

	if u.fieldIndex > -1 {
		switch u.fieldIndex {

		case 1:
			err = writer.EncodeBool(u.value.(bool))
			if err != nil {
				return nil, err
			}

		case 2:
			err = writer.EncodeString(u.value.(string))
			if err != nil {
				return nil, err
			}

		case 3:
			err = writer.EncodeUint8(u.value.(uint8))
			if err != nil {
				return nil, err
			}

		case 4:
			err = writer.EncodeUint16(u.value.(uint16))
			if err != nil {
				return nil, err
			}

		case 5:
			err = writer.EncodeUint32(u.value.(uint32))
			if err != nil {
				return nil, err
			}

		case 6:
			err = writer.EncodeUint64(u.value.(uint64))
			if err != nil {
				return nil, err
			}

		case 7:
			err = writer.EncodeInt8(u.value.(int8))
			if err != nil {
				return nil, err
			}

		case 8:
			err = writer.EncodeInt16(u.value.(int16))
			if err != nil {
				return nil, err
			}

		case 9:
			err = writer.EncodeInt32(u.value.(int32))
			if err != nil {
				return nil, err
			}

		case 10:
			err = writer.EncodeInt64(u.value.(int64))
			if err != nil {
				return nil, err
			}

		case 11:
			err = writer.EncodeFloat32(u.value.(float32))
			if err != nil {
				return nil, err
			}

		case 12:
			err = writer.EncodeFloat64(u.value.(float64))
			if err != nil {
				return nil, err
			}

		case 13:
			err = writer.EncodeBytes(u.value.([]byte))
			if err != nil {
				return nil, err
			}

		}
	}

	return buf.Bytes(), nil
}

func (u *MyUnion) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}

	return buf
}

func (u *MyUnion) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := msgpack.NewDecoder(reader)
	var err error

	fieldIndex, err := decoder.DecodeInt8()
	if err != nil {
		return err
	}

	u.fieldIndex = fieldIndex
	switch fieldIndex {
	case -1:
		u.value = nil

	case 1:
		u.value, err = decoder.DecodeBool()
		if err != nil {
			return err
		}

	case 2:
		u.value, err = decoder.DecodeString()
		if err != nil {
			return err
		}

	case 3:
		u.value, err = decoder.DecodeUint8()
		if err != nil {
			return err
		}

	case 4:
		u.value, err = decoder.DecodeUint16()
		if err != nil {
			return err
		}

	case 5:
		u.value, err = decoder.DecodeUint32()
		if err != nil {
			return err
		}

	case 6:
		u.value, err = decoder.DecodeUint64()
		if err != nil {
			return err
		}

	case 7:
		u.value, err = decoder.DecodeInt8()
		if err != nil {
			return err
		}

	case 8:
		u.value, err = decoder.DecodeInt16()
		if err != nil {
			return err
		}

	case 9:
		u.value, err = decoder.DecodeInt32()
		if err != nil {
			return err
		}

	case 10:
		u.value, err = decoder.DecodeInt64()
		if err != nil {
			return err
		}

	case 11:
		u.value, err = decoder.DecodeFloat32()
		if err != nil {
			return err
		}

	case 12:
		u.value, err = decoder.DecodeFloat64()
		if err != nil {
			return err
		}

	case 13:
		u.value, err = decoder.DecodeBytes()
		if err != nil {
			return err
		}

	}

	return nil
}
