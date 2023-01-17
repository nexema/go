package identity

import (
	"bytes"
	"fmt"
	"github.com/example/models"
	"github.com/nexema/go/runtime"
	"io"
	"strings"
)

type NexemaPrimitives struct {
	MyString  string
	MyBoolean bool
	MyUint8   uint8
	MyUint16  uint16
	MyUint32  uint32
	MyUint64  uint64
	MyInt8    int8
	MyInt16   int16
	MyInt32   int32
	MyInt64   int64
	MyFloat32 float32
	MyFloat64 float64
	MyBinary  []byte
}

type NexemaPrimitivesBuilder struct {
	MyString  string
	MyBoolean bool
	MyUint8   uint8
	MyUint16  uint16
	MyUint32  uint32
	MyUint64  uint64
	MyInt8    int8
	MyInt16   int16
	MyInt32   int32
	MyInt64   int64
	MyFloat32 float32
	MyFloat64 float64
	MyBinary  []byte
}

// NewNexemaPrimitives constructs a new instance of NexemaPrimitives.
// Calling this takes the adventage of default values and other runtime checks
func NewNexemaPrimitives(builder ...NexemaPrimitivesBuilder) *NexemaPrimitives {
	if len(builder) > 1 {
		panic("builder expects one argument")
	}

	var builderInst NexemaPrimitivesBuilder
	if len(builder) == 1 {
		builderInst = builder[0]
	}

	instance := &NexemaPrimitives{}
	instance.MyString = builderInst.MyString
	instance.MyBoolean = builderInst.MyBoolean
	instance.MyUint8 = builderInst.MyUint8
	instance.MyUint16 = builderInst.MyUint16
	instance.MyUint32 = builderInst.MyUint32
	instance.MyUint64 = builderInst.MyUint64
	instance.MyInt8 = builderInst.MyInt8
	instance.MyInt16 = builderInst.MyInt16
	instance.MyInt32 = builderInst.MyInt32
	instance.MyInt64 = builderInst.MyInt64
	instance.MyFloat32 = builderInst.MyFloat32
	instance.MyFloat64 = builderInst.MyFloat64
	instance.MyBinary = builderInst.MyBinary
	return instance
}

func (u *NexemaPrimitives) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	encoder.EncodeString(u.MyString)

	encoder.EncodeBool(u.MyBoolean)

	encoder.EncodeUint8(u.MyUint8)

	encoder.EncodeUint16(u.MyUint16)

	encoder.EncodeUint32(u.MyUint32)

	encoder.EncodeUint64(u.MyUint64)

	encoder.EncodeInt8(u.MyInt8)

	encoder.EncodeInt16(u.MyInt16)

	encoder.EncodeInt32(u.MyInt32)

	encoder.EncodeInt64(u.MyInt64)

	encoder.EncodeFloat32(u.MyFloat32)

	encoder.EncodeFloat64(u.MyFloat64)

	encoder.EncodeBinary(u.MyBinary)

	return encoder.TakeBytes(), nil
}

func (u *NexemaPrimitives) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *NexemaPrimitives) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	u.MyString, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	u.MyBoolean, err = decoder.DecodeBool()
	if err != nil {
		return err
	}

	u.MyUint8, err = decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.MyUint16, err = decoder.DecodeUint16()
	if err != nil {
		return err
	}

	u.MyUint32, err = decoder.DecodeUint32()
	if err != nil {
		return err
	}

	u.MyUint64, err = decoder.DecodeUint64()
	if err != nil {
		return err
	}

	u.MyInt8, err = decoder.DecodeInt8()
	if err != nil {
		return err
	}

	u.MyInt16, err = decoder.DecodeInt16()
	if err != nil {
		return err
	}

	u.MyInt32, err = decoder.DecodeInt32()
	if err != nil {
		return err
	}

	u.MyInt64, err = decoder.DecodeInt64()
	if err != nil {
		return err
	}

	u.MyFloat32, err = decoder.DecodeFloat32()
	if err != nil {
		return err
	}

	u.MyFloat64, err = decoder.DecodeFloat64()
	if err != nil {
		return err
	}

	u.MyBinary, err = decoder.DecodeBinary()
	if err != nil {
		return err
	}

	return nil
}

func (u *NexemaPrimitives) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *NexemaPrimitives) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *NexemaPrimitives) String() string {
	parts := make([]string, 13)
	parts[0] = fmt.Sprintf("MyString: %v", u.MyString)
	parts[1] = fmt.Sprintf("MyBoolean: %v", u.MyBoolean)
	parts[2] = fmt.Sprintf("MyUint8: %v", u.MyUint8)
	parts[3] = fmt.Sprintf("MyUint16: %v", u.MyUint16)
	parts[4] = fmt.Sprintf("MyUint32: %v", u.MyUint32)
	parts[5] = fmt.Sprintf("MyUint64: %v", u.MyUint64)
	parts[6] = fmt.Sprintf("MyInt8: %v", u.MyInt8)
	parts[7] = fmt.Sprintf("MyInt16: %v", u.MyInt16)
	parts[8] = fmt.Sprintf("MyInt32: %v", u.MyInt32)
	parts[9] = fmt.Sprintf("MyInt64: %v", u.MyInt64)
	parts[10] = fmt.Sprintf("MyFloat32: %v", u.MyFloat32)
	parts[11] = fmt.Sprintf("MyFloat64: %v", u.MyFloat64)
	parts[12] = fmt.Sprintf("MyBinary: %v", u.MyBinary)
	return fmt.Sprintf("NexemaPrimitives(%s)", strings.Join(parts, ", "))
}

func (u *NexemaPrimitives) Equals(other *NexemaPrimitives) bool {
	if u.MyString != other.MyString {
		return false
	}
	if u.MyBoolean != other.MyBoolean {
		return false
	}
	if u.MyUint8 != other.MyUint8 {
		return false
	}
	if u.MyUint16 != other.MyUint16 {
		return false
	}
	if u.MyUint32 != other.MyUint32 {
		return false
	}
	if u.MyUint64 != other.MyUint64 {
		return false
	}
	if u.MyInt8 != other.MyInt8 {
		return false
	}
	if u.MyInt16 != other.MyInt16 {
		return false
	}
	if u.MyInt32 != other.MyInt32 {
		return false
	}
	if u.MyInt64 != other.MyInt64 {
		return false
	}
	if u.MyFloat32 != other.MyFloat32 {
		return false
	}
	if u.MyFloat64 != other.MyFloat64 {
		return false
	}
	if !bytes.Equal(u.MyBinary, other.MyBinary) {
		return false
	}
	return true
}

type NexemaNullablePrimitives struct {
	MyString  runtime.Nullable[string]
	MyBoolean runtime.Nullable[bool]
	MyUint8   runtime.Nullable[uint8]
	MyUint16  runtime.Nullable[uint16]
	MyUint32  runtime.Nullable[uint32]
	MyUint64  runtime.Nullable[uint64]
	MyInt8    runtime.Nullable[int8]
	MyInt16   runtime.Nullable[int16]
	MyInt32   runtime.Nullable[int32]
	MyInt64   runtime.Nullable[int64]
	MyFloat32 runtime.Nullable[float32]
	MyFloat64 runtime.Nullable[float64]
	MyBinary  runtime.Nullable[[]byte]
}

type NexemaNullablePrimitivesBuilder struct {
	MyString  runtime.Nullable[string]
	MyBoolean runtime.Nullable[bool]
	MyUint8   runtime.Nullable[uint8]
	MyUint16  runtime.Nullable[uint16]
	MyUint32  runtime.Nullable[uint32]
	MyUint64  runtime.Nullable[uint64]
	MyInt8    runtime.Nullable[int8]
	MyInt16   runtime.Nullable[int16]
	MyInt32   runtime.Nullable[int32]
	MyInt64   runtime.Nullable[int64]
	MyFloat32 runtime.Nullable[float32]
	MyFloat64 runtime.Nullable[float64]
	MyBinary  runtime.Nullable[[]byte]
}

// NewNexemaNullablePrimitives constructs a new instance of NexemaNullablePrimitives.
// Calling this takes the adventage of default values and other runtime checks
func NewNexemaNullablePrimitives(builder ...NexemaNullablePrimitivesBuilder) *NexemaNullablePrimitives {
	if len(builder) > 1 {
		panic("builder expects one argument")
	}

	var builderInst NexemaNullablePrimitivesBuilder
	if len(builder) == 1 {
		builderInst = builder[0]
	}

	instance := &NexemaNullablePrimitives{}
	instance.MyString = builderInst.MyString
	instance.MyBoolean = builderInst.MyBoolean
	instance.MyUint8 = builderInst.MyUint8
	instance.MyUint16 = builderInst.MyUint16
	instance.MyUint32 = builderInst.MyUint32
	instance.MyUint64 = builderInst.MyUint64
	instance.MyInt8 = builderInst.MyInt8
	instance.MyInt16 = builderInst.MyInt16
	instance.MyInt32 = builderInst.MyInt32
	instance.MyInt64 = builderInst.MyInt64
	instance.MyFloat32 = builderInst.MyFloat32
	instance.MyFloat64 = builderInst.MyFloat64
	instance.MyBinary = builderInst.MyBinary
	return instance
}

func (u *NexemaNullablePrimitives) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	if u.MyString.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeString(*u.MyString.Value)

	}
	if u.MyBoolean.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeBool(*u.MyBoolean.Value)

	}
	if u.MyUint8.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeUint8(*u.MyUint8.Value)

	}
	if u.MyUint16.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeUint16(*u.MyUint16.Value)

	}
	if u.MyUint32.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeUint32(*u.MyUint32.Value)

	}
	if u.MyUint64.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeUint64(*u.MyUint64.Value)

	}
	if u.MyInt8.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeInt8(*u.MyInt8.Value)

	}
	if u.MyInt16.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeInt16(*u.MyInt16.Value)

	}
	if u.MyInt32.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeInt32(*u.MyInt32.Value)

	}
	if u.MyInt64.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeInt64(*u.MyInt64.Value)

	}
	if u.MyFloat32.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeFloat32(*u.MyFloat32.Value)

	}
	if u.MyFloat64.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeFloat64(*u.MyFloat64.Value)

	}
	if u.MyBinary.IsNull() {
		encoder.EncodeNull()
	} else {

		encoder.EncodeBinary(*u.MyBinary.Value)

	}
	return encoder.TakeBytes(), nil
}

func (u *NexemaNullablePrimitives) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *NexemaNullablePrimitives) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	if decoder.IsNextNull() {
		u.MyString.Clear()
	} else {

		var value string
		value, err = decoder.DecodeString()
		if err != nil {
			return err
		}

		u.MyString.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyBoolean.Clear()
	} else {

		var value bool
		value, err = decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.MyBoolean.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyUint8.Clear()
	} else {

		var value uint8
		value, err = decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.MyUint8.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyUint16.Clear()
	} else {

		var value uint16
		value, err = decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.MyUint16.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyUint32.Clear()
	} else {

		var value uint32
		value, err = decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.MyUint32.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyUint64.Clear()
	} else {

		var value uint64
		value, err = decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.MyUint64.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyInt8.Clear()
	} else {

		var value int8
		value, err = decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.MyInt8.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyInt16.Clear()
	} else {

		var value int16
		value, err = decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.MyInt16.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyInt32.Clear()
	} else {

		var value int32
		value, err = decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.MyInt32.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyInt64.Clear()
	} else {

		var value int64
		value, err = decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.MyInt64.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyFloat32.Clear()
	} else {

		var value float32
		value, err = decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.MyFloat32.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyFloat64.Clear()
	} else {

		var value float64
		value, err = decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.MyFloat64.SetValue(value)

	}
	if decoder.IsNextNull() {
		u.MyBinary.Clear()
	} else {

		var value []byte
		value, err = decoder.DecodeBinary()
		if err != nil {
			return err
		}

		u.MyBinary.SetValue(value)

	}
	return nil
}

func (u *NexemaNullablePrimitives) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *NexemaNullablePrimitives) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *NexemaNullablePrimitives) String() string {
	parts := make([]string, 13)
	parts[0] = fmt.Sprintf("MyString: %v", u.MyString)
	parts[1] = fmt.Sprintf("MyBoolean: %v", u.MyBoolean)
	parts[2] = fmt.Sprintf("MyUint8: %v", u.MyUint8)
	parts[3] = fmt.Sprintf("MyUint16: %v", u.MyUint16)
	parts[4] = fmt.Sprintf("MyUint32: %v", u.MyUint32)
	parts[5] = fmt.Sprintf("MyUint64: %v", u.MyUint64)
	parts[6] = fmt.Sprintf("MyInt8: %v", u.MyInt8)
	parts[7] = fmt.Sprintf("MyInt16: %v", u.MyInt16)
	parts[8] = fmt.Sprintf("MyInt32: %v", u.MyInt32)
	parts[9] = fmt.Sprintf("MyInt64: %v", u.MyInt64)
	parts[10] = fmt.Sprintf("MyFloat32: %v", u.MyFloat32)
	parts[11] = fmt.Sprintf("MyFloat64: %v", u.MyFloat64)
	parts[12] = fmt.Sprintf("MyBinary: %v", u.MyBinary)
	return fmt.Sprintf("NexemaNullablePrimitives(%s)", strings.Join(parts, ", "))
}

func (u *NexemaNullablePrimitives) Equals(other *NexemaNullablePrimitives) bool {
	if (u.MyString.IsNull() != other.MyString.IsNull()) || (*u.MyString.Value != *other.MyString.Value) {
		return false
	}
	if (u.MyBoolean.IsNull() != other.MyBoolean.IsNull()) || (*u.MyBoolean.Value != *other.MyBoolean.Value) {
		return false
	}
	if (u.MyUint8.IsNull() != other.MyUint8.IsNull()) || (*u.MyUint8.Value != *other.MyUint8.Value) {
		return false
	}
	if (u.MyUint16.IsNull() != other.MyUint16.IsNull()) || (*u.MyUint16.Value != *other.MyUint16.Value) {
		return false
	}
	if (u.MyUint32.IsNull() != other.MyUint32.IsNull()) || (*u.MyUint32.Value != *other.MyUint32.Value) {
		return false
	}
	if (u.MyUint64.IsNull() != other.MyUint64.IsNull()) || (*u.MyUint64.Value != *other.MyUint64.Value) {
		return false
	}
	if (u.MyInt8.IsNull() != other.MyInt8.IsNull()) || (*u.MyInt8.Value != *other.MyInt8.Value) {
		return false
	}
	if (u.MyInt16.IsNull() != other.MyInt16.IsNull()) || (*u.MyInt16.Value != *other.MyInt16.Value) {
		return false
	}
	if (u.MyInt32.IsNull() != other.MyInt32.IsNull()) || (*u.MyInt32.Value != *other.MyInt32.Value) {
		return false
	}
	if (u.MyInt64.IsNull() != other.MyInt64.IsNull()) || (*u.MyInt64.Value != *other.MyInt64.Value) {
		return false
	}
	if (u.MyFloat32.IsNull() != other.MyFloat32.IsNull()) || (*u.MyFloat32.Value != *other.MyFloat32.Value) {
		return false
	}
	if (u.MyFloat64.IsNull() != other.MyFloat64.IsNull()) || (*u.MyFloat64.Value != *other.MyFloat64.Value) {
		return false
	}
	if (u.MyBinary.IsNull() != other.MyBinary.IsNull()) || (!bytes.Equal(*u.MyBinary.Value, *other.MyBinary.Value)) {
		return false
	}
	return true
}

type NexemaList struct {
	List1 []string
	List2 [][]byte
	List3 []runtime.Nullable[string]
	List4 []runtime.Nullable[string]
	List5 [][]byte
	List6 []runtime.Nullable[[]byte]
	List7 []runtime.Nullable[[]byte]
}

type NexemaListBuilder struct {
	List1 []string
	List2 [][]byte
	List3 []runtime.Nullable[string]
	List4 []runtime.Nullable[string]
	List5 [][]byte
	List6 []runtime.Nullable[[]byte]
	List7 []runtime.Nullable[[]byte]
}

// NewNexemaList constructs a new instance of NexemaList.
// Calling this takes the adventage of default values and other runtime checks
func NewNexemaList(builder ...NexemaListBuilder) *NexemaList {
	if len(builder) > 1 {
		panic("builder expects one argument")
	}

	var builderInst NexemaListBuilder
	if len(builder) == 1 {
		builderInst = builder[0]
	}

	instance := &NexemaList{}
	instance.List1 = builderInst.List1
	instance.List2 = builderInst.List2
	instance.List3 = builderInst.List3
	instance.List4 = builderInst.List4
	instance.List5 = builderInst.List5
	instance.List6 = builderInst.List6
	instance.List7 = builderInst.List7
	return instance
}

func (u *NexemaList) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	if u.List1 == nil {
		return nil, fmt.Errorf("field List1 is null but its not marked as nullable")
	}
	if u.List2 == nil {
		return nil, fmt.Errorf("field List2 is null but its not marked as nullable")
	}
	if u.List3 == nil {
		return nil, fmt.Errorf("field List3 is null but its not marked as nullable")
	}
	if u.List4 == nil {
		encoder.EncodeNull()
	} else {

		encoder.BeginArray(int64(len(u.List4)))
		for _, element := range u.List4 {
			if element.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeString(*element.Value)
			}
		}

	}
	if u.List5 == nil {
		encoder.EncodeNull()
	} else {

		encoder.BeginArray(int64(len(u.List5)))
		for _, element := range u.List5 {
			encoder.EncodeBinary(element)
		}

	}
	if u.List6 == nil {
		return nil, fmt.Errorf("field List6 is null but its not marked as nullable")
	}
	if u.List7 == nil {
		encoder.EncodeNull()
	} else {

		encoder.BeginArray(int64(len(u.List7)))
		for _, element := range u.List7 {
			if element.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeBinary(*element.Value)
			}
		}

	}
	return encoder.TakeBytes(), nil
}

func (u *NexemaList) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *NexemaList) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	list1ArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}
	u.List1 = make([]string, list1ArrayLen)
	for i := int64(0); i < list1ArrayLen; i++ {
		u.List1[i], err = decoder.DecodeString()
		if err != nil {
			return err
		}
	}

	list2ArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}
	u.List2 = make([][]byte, list2ArrayLen)
	for i := int64(0); i < list2ArrayLen; i++ {
		u.List2[i], err = decoder.DecodeBinary()
		if err != nil {
			return err
		}
	}

	list3ArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}
	u.List3 = make([]runtime.Nullable[string], list3ArrayLen)
	for i := int64(0); i < list3ArrayLen; i++ {
		if decoder.IsNextNull() {
			u.List3[i] = runtime.NewNull[string]()
		} else {
			list3, err := decoder.DecodeString()
			if err != nil {
				return err
			}

			u.List3[i] = runtime.NewNullable(list3)
		}
	}

	if decoder.IsNextNull() {
		u.List4 = nil
	} else {

		list4ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}
		u.List4 = make([]runtime.Nullable[string], list4ArrayLen)
		for i := int64(0); i < list4ArrayLen; i++ {
			if decoder.IsNextNull() {
				u.List4[i] = runtime.NewNull[string]()
			} else {
				list4, err := decoder.DecodeString()
				if err != nil {
					return err
				}

				u.List4[i] = runtime.NewNullable(list4)
			}
		}

	}
	if decoder.IsNextNull() {
		u.List5 = nil
	} else {

		list5ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}
		u.List5 = make([][]byte, list5ArrayLen)
		for i := int64(0); i < list5ArrayLen; i++ {
			u.List5[i], err = decoder.DecodeBinary()
			if err != nil {
				return err
			}
		}

	}

	list6ArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}
	u.List6 = make([]runtime.Nullable[[]byte], list6ArrayLen)
	for i := int64(0); i < list6ArrayLen; i++ {
		if decoder.IsNextNull() {
			u.List6[i] = runtime.NewNull[[]byte]()
		} else {
			list6, err := decoder.DecodeBinary()
			if err != nil {
				return err
			}

			u.List6[i] = runtime.NewNullable(list6)
		}
	}

	if decoder.IsNextNull() {
		u.List7 = nil
	} else {

		list7ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}
		u.List7 = make([]runtime.Nullable[[]byte], list7ArrayLen)
		for i := int64(0); i < list7ArrayLen; i++ {
			if decoder.IsNextNull() {
				u.List7[i] = runtime.NewNull[[]byte]()
			} else {
				list7, err := decoder.DecodeBinary()
				if err != nil {
					return err
				}

				u.List7[i] = runtime.NewNullable(list7)
			}
		}

	}
	return nil
}

func (u *NexemaList) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *NexemaList) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *NexemaList) String() string {
	parts := make([]string, 7)
	parts[0] = fmt.Sprintf("List1: %v", u.List1)
	parts[1] = fmt.Sprintf("List2: %v", u.List2)
	parts[2] = fmt.Sprintf("List3: %v", u.List3)
	parts[3] = fmt.Sprintf("List4: %v", u.List4)
	parts[4] = fmt.Sprintf("List5: %v", u.List5)
	parts[5] = fmt.Sprintf("List6: %v", u.List6)
	parts[6] = fmt.Sprintf("List7: %v", u.List7)
	return fmt.Sprintf("NexemaList(%s)", strings.Join(parts, ", "))
}

func (u *NexemaList) Equals(other *NexemaList) bool {
	if (u.List1 == nil) != (other.List1 == nil) {
		return false
	}

	list1LenThis := len(u.List1)
	list1LenOther := len(other.List1)
	if list1LenThis != list1LenOther {
		return false
	}

	for i := 0; i < list1LenThis; i++ {
		a := u.List1[i]
		b := other.List1[i]

		if a != b {
			return false
		}
	}
	if (u.List2 == nil) != (other.List2 == nil) {
		return false
	}

	list2LenThis := len(u.List2)
	list2LenOther := len(other.List2)
	if list2LenThis != list2LenOther {
		return false
	}

	for i := 0; i < list2LenThis; i++ {
		a := u.List2[i]
		b := other.List2[i]

		if !bytes.Equal(a, b) {
			return false
		}
	}
	if (u.List3 == nil) != (other.List3 == nil) {
		return false
	}

	list3LenThis := len(u.List3)
	list3LenOther := len(other.List3)
	if list3LenThis != list3LenOther {
		return false
	}

	for i := 0; i < list3LenThis; i++ {
		a := u.List3[i]
		b := other.List3[i]

		if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
			return false
		}
	}
	if (u.List4 == nil) != (other.List4 == nil) {
		return false
	}

	list4LenThis := len(u.List4)
	list4LenOther := len(other.List4)
	if list4LenThis != list4LenOther {
		return false
	}

	for i := 0; i < list4LenThis; i++ {
		a := u.List4[i]
		b := other.List4[i]

		if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
			return false
		}
	}
	if (u.List5 == nil) != (other.List5 == nil) {
		return false
	}

	list5LenThis := len(u.List5)
	list5LenOther := len(other.List5)
	if list5LenThis != list5LenOther {
		return false
	}

	for i := 0; i < list5LenThis; i++ {
		a := u.List5[i]
		b := other.List5[i]

		if !bytes.Equal(a, b) {
			return false
		}
	}
	if (u.List6 == nil) != (other.List6 == nil) {
		return false
	}

	list6LenThis := len(u.List6)
	list6LenOther := len(other.List6)
	if list6LenThis != list6LenOther {
		return false
	}

	for i := 0; i < list6LenThis; i++ {
		a := u.List6[i]
		b := other.List6[i]

		if (a.IsNull() != b.IsNull()) || (!bytes.Equal(*a.Value, *b.Value)) {
			return false
		}
	}
	if (u.List7 == nil) != (other.List7 == nil) {
		return false
	}

	list7LenThis := len(u.List7)
	list7LenOther := len(other.List7)
	if list7LenThis != list7LenOther {
		return false
	}

	for i := 0; i < list7LenThis; i++ {
		a := u.List7[i]
		b := other.List7[i]

		if (a.IsNull() != b.IsNull()) || (!bytes.Equal(*a.Value, *b.Value)) {
			return false
		}
	}
	return true
}

type NexemaMap struct {
	Map0 map[string]string
	Map1 map[string]runtime.Nullable[string]
	Map2 map[string]string
	Map3 map[string]runtime.Nullable[string]
	Map4 map[string][]byte
	Map5 map[string]runtime.Nullable[[]byte]
	Map6 map[string]runtime.Nullable[[]byte]
}

type NexemaMapBuilder struct {
	Map0 map[string]string
	Map1 map[string]runtime.Nullable[string]
	Map2 map[string]string
	Map3 map[string]runtime.Nullable[string]
	Map4 map[string][]byte
	Map5 map[string]runtime.Nullable[[]byte]
	Map6 map[string]runtime.Nullable[[]byte]
}

// NewNexemaMap constructs a new instance of NexemaMap.
// Calling this takes the adventage of default values and other runtime checks
func NewNexemaMap(builder ...NexemaMapBuilder) *NexemaMap {
	if len(builder) > 1 {
		panic("builder expects one argument")
	}

	var builderInst NexemaMapBuilder
	if len(builder) == 1 {
		builderInst = builder[0]
	}

	instance := &NexemaMap{}
	instance.Map0 = builderInst.Map0
	instance.Map1 = builderInst.Map1
	instance.Map2 = builderInst.Map2
	instance.Map3 = builderInst.Map3
	instance.Map4 = builderInst.Map4
	instance.Map5 = builderInst.Map5
	instance.Map6 = builderInst.Map6
	return instance
}

func (u *NexemaMap) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	if u.Map0 == nil {
		return nil, fmt.Errorf("field Map0 is null but its not marked as nullable")
	}
	if u.Map1 == nil {
		return nil, fmt.Errorf("field Map1 is null but its not marked as nullable")
	}
	if u.Map2 == nil {
		encoder.EncodeNull()
	} else {

		for key, value := range u.Map2 {
			encoder.EncodeString(key)
			encoder.EncodeString(value)
		}

	}
	if u.Map3 == nil {
		encoder.EncodeNull()
	} else {

		for key, value := range u.Map3 {
			encoder.EncodeString(key)
			if value.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeString(*value.Value)
			}
		}

	}
	if u.Map4 == nil {
		return nil, fmt.Errorf("field Map4 is null but its not marked as nullable")
	}
	if u.Map5 == nil {
		return nil, fmt.Errorf("field Map5 is null but its not marked as nullable")
	}
	if u.Map6 == nil {
		encoder.EncodeNull()
	} else {

		for key, value := range u.Map6 {
			encoder.EncodeString(key)
			if value.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeBinary(*value.Value)
			}
		}

	}
	return encoder.TakeBytes(), nil
}

func (u *NexemaMap) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *NexemaMap) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	map0MapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.Map0 = make(map[string]string, map0MapLen)
	for i := int64(0); i < map0MapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		value, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.Map0[key] = value
	}

	map1MapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.Map1 = make(map[string]runtime.Nullable[string], map1MapLen)
	for i := int64(0); i < map1MapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		if decoder.IsNextNull() {
			u.Map1[key] = runtime.NewNull[string]()
		} else {
			value, err := decoder.DecodeString()
			if err != nil {
				return err
			}

			u.Map1[key] = runtime.NewNullable(value)
		}
	}

	if decoder.IsNextNull() {
		u.Map2 = nil
	} else {

		map2MapLen, err := decoder.BeginDecodeMap()
		if err != nil {
			return err
		}

		u.Map2 = make(map[string]string, map2MapLen)
		for i := int64(0); i < map2MapLen; i++ {
			key, err := decoder.DecodeString()
			if err != nil {
				return err
			}
			value, err := decoder.DecodeString()
			if err != nil {
				return err
			}

			u.Map2[key] = value
		}

	}
	if decoder.IsNextNull() {
		u.Map3 = nil
	} else {

		map3MapLen, err := decoder.BeginDecodeMap()
		if err != nil {
			return err
		}

		u.Map3 = make(map[string]runtime.Nullable[string], map3MapLen)
		for i := int64(0); i < map3MapLen; i++ {
			key, err := decoder.DecodeString()
			if err != nil {
				return err
			}
			if decoder.IsNextNull() {
				u.Map3[key] = runtime.NewNull[string]()
			} else {
				value, err := decoder.DecodeString()
				if err != nil {
					return err
				}

				u.Map3[key] = runtime.NewNullable(value)
			}
		}

	}

	map4MapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.Map4 = make(map[string][]byte, map4MapLen)
	for i := int64(0); i < map4MapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		value, err := decoder.DecodeBinary()
		if err != nil {
			return err
		}

		u.Map4[key] = value
	}

	map5MapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.Map5 = make(map[string]runtime.Nullable[[]byte], map5MapLen)
	for i := int64(0); i < map5MapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		if decoder.IsNextNull() {
			u.Map5[key] = runtime.NewNull[[]byte]()
		} else {
			value, err := decoder.DecodeBinary()
			if err != nil {
				return err
			}

			u.Map5[key] = runtime.NewNullable(value)
		}
	}

	if decoder.IsNextNull() {
		u.Map6 = nil
	} else {

		map6MapLen, err := decoder.BeginDecodeMap()
		if err != nil {
			return err
		}

		u.Map6 = make(map[string]runtime.Nullable[[]byte], map6MapLen)
		for i := int64(0); i < map6MapLen; i++ {
			key, err := decoder.DecodeString()
			if err != nil {
				return err
			}
			if decoder.IsNextNull() {
				u.Map6[key] = runtime.NewNull[[]byte]()
			} else {
				value, err := decoder.DecodeBinary()
				if err != nil {
					return err
				}

				u.Map6[key] = runtime.NewNullable(value)
			}
		}

	}
	return nil
}

func (u *NexemaMap) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *NexemaMap) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *NexemaMap) String() string {
	parts := make([]string, 7)
	parts[0] = fmt.Sprintf("Map0: %v", u.Map0)
	parts[1] = fmt.Sprintf("Map1: %v", u.Map1)
	parts[2] = fmt.Sprintf("Map2: %v", u.Map2)
	parts[3] = fmt.Sprintf("Map3: %v", u.Map3)
	parts[4] = fmt.Sprintf("Map4: %v", u.Map4)
	parts[5] = fmt.Sprintf("Map5: %v", u.Map5)
	parts[6] = fmt.Sprintf("Map6: %v", u.Map6)
	return fmt.Sprintf("NexemaMap(%s)", strings.Join(parts, ", "))
}

func (u *NexemaMap) Equals(other *NexemaMap) bool {

	if (u.Map0 == nil) != (other.Map0 == nil) {
		return false
	}

	map0LenThis := len(u.Map0)
	map0LenOther := len(other.Map0)
	if map0LenThis != map0LenOther {
		return false
	}

	for k, a := range u.Map0 {
		b, ok := other.Map0[k]
		if !ok {
			return false
		}

		if a != b {
			return false
		}
	}

	if (u.Map1 == nil) != (other.Map1 == nil) {
		return false
	}

	map1LenThis := len(u.Map1)
	map1LenOther := len(other.Map1)
	if map1LenThis != map1LenOther {
		return false
	}

	for k, a := range u.Map1 {
		b, ok := other.Map1[k]
		if !ok {
			return false
		}

		if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
			return false
		}
	}

	if (u.Map2 == nil) != (other.Map2 == nil) {
		return false
	}

	map2LenThis := len(u.Map2)
	map2LenOther := len(other.Map2)
	if map2LenThis != map2LenOther {
		return false
	}

	for k, a := range u.Map2 {
		b, ok := other.Map2[k]
		if !ok {
			return false
		}

		if a != b {
			return false
		}
	}

	if (u.Map3 == nil) != (other.Map3 == nil) {
		return false
	}

	map3LenThis := len(u.Map3)
	map3LenOther := len(other.Map3)
	if map3LenThis != map3LenOther {
		return false
	}

	for k, a := range u.Map3 {
		b, ok := other.Map3[k]
		if !ok {
			return false
		}

		if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
			return false
		}
	}

	if (u.Map4 == nil) != (other.Map4 == nil) {
		return false
	}

	map4LenThis := len(u.Map4)
	map4LenOther := len(other.Map4)
	if map4LenThis != map4LenOther {
		return false
	}

	for k, a := range u.Map4 {
		b, ok := other.Map4[k]
		if !ok {
			return false
		}

		if !bytes.Equal(a, b) {
			return false
		}
	}

	if (u.Map5 == nil) != (other.Map5 == nil) {
		return false
	}

	map5LenThis := len(u.Map5)
	map5LenOther := len(other.Map5)
	if map5LenThis != map5LenOther {
		return false
	}

	for k, a := range u.Map5 {
		b, ok := other.Map5[k]
		if !ok {
			return false
		}

		if (a.IsNull() != b.IsNull()) || (!bytes.Equal(*a.Value, *b.Value)) {
			return false
		}
	}

	if (u.Map6 == nil) != (other.Map6 == nil) {
		return false
	}

	map6LenThis := len(u.Map6)
	map6LenOther := len(other.Map6)
	if map6LenThis != map6LenOther {
		return false
	}

	for k, a := range u.Map6 {
		b, ok := other.Map6[k]
		if !ok {
			return false
		}

		if (a.IsNull() != b.IsNull()) || (!bytes.Equal(*a.Value, *b.Value)) {
			return false
		}
	}
	return true
}

// This is my embedded type
type EmbeddedType struct {
	MyUnion *SingleUnion
	MyEnum  models.Colors
}

type EmbeddedTypeBuilder struct {
	MyUnion *SingleUnion
	MyEnum  models.Colors
}

// NewEmbeddedType constructs a new instance of EmbeddedType.
// Calling this takes the adventage of default values and other runtime checks
func NewEmbeddedType(builder ...EmbeddedTypeBuilder) *EmbeddedType {
	if len(builder) > 1 {
		panic("builder expects one argument")
	}

	var builderInst EmbeddedTypeBuilder
	if len(builder) == 1 {
		builderInst = builder[0]
	}

	instance := &EmbeddedType{}
	instance.MyUnion = builderInst.MyUnion
	instance.MyEnum = builderInst.MyEnum
	return instance
}

func (u *EmbeddedType) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	myUnionBytes, err := u.MyUnion.Encode()
	if err != nil {
		return nil, err
	}

	encoder.EncodeBinary(myUnionBytes)

	encoder.EncodeUint8(u.MyEnum.Index())

	return encoder.TakeBytes(), nil
}

func (u *EmbeddedType) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *EmbeddedType) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	myUnionBytes, err := decoder.DecodeBinary()
	if err != nil {
		return err
	}

	err = u.MyUnion.MergeFrom(myUnionBytes)
	if err != nil {
		return err
	}

	myEnumEnumIndex, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.MyEnum = models.ColorsPicker.ByIndex(myEnumEnumIndex)

	return nil
}

func (u *EmbeddedType) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *EmbeddedType) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *EmbeddedType) String() string {
	parts := make([]string, 2)
	parts[0] = fmt.Sprintf("MyUnion: %v", u.MyUnion)
	parts[1] = fmt.Sprintf("MyEnum: %v", u.MyEnum)
	return fmt.Sprintf("EmbeddedType(%s)", strings.Join(parts, ", "))
}

func (u *EmbeddedType) Equals(other *EmbeddedType) bool {
	if !u.MyUnion.Equals(other.MyUnion) {
		return false
	}
	if u.MyEnum.Index() != other.MyEnum.Index() {
		return false
	}
	return true
}

// Nexema type with fields that declare default values
type NexemaDefaultType struct {
	MyString       string
	MyBool         bool
	MyInt32        int32
	MyFloat64      float64
	MyList         []string
	MyMap          map[string]bool
	MyNullablemap  map[string]runtime.Nullable[bool]
	MyNullablelist []runtime.Nullable[float64]
}

type NexemaDefaultTypeBuilder struct {
	MyString       string
	MyBool         bool
	MyInt32        int32
	MyFloat64      float64
	MyList         []string
	MyMap          map[string]bool
	MyNullablemap  map[string]runtime.Nullable[bool]
	MyNullablelist []runtime.Nullable[float64]
}

// NewNexemaDefaultType constructs a new instance of NexemaDefaultType.
// Calling this takes the adventage of default values and other runtime checks
func NewNexemaDefaultType(builder ...NexemaDefaultTypeBuilder) *NexemaDefaultType {
	if len(builder) > 1 {
		panic("builder expects one argument")
	}

	var builderInst NexemaDefaultTypeBuilder
	if len(builder) == 1 {
		builderInst = builder[0]
	}

	instance := &NexemaDefaultType{}
	if len(builderInst.MyString) == 0 {
		instance.MyString = "hola"
	} else {
		instance.MyString = builderInst.MyString
	}
	if true {
		instance.MyBool = true
	} else {
		instance.MyBool = builderInst.MyBool
	}
	if builderInst.MyInt32 == 0 {
		instance.MyInt32 = -2.314123e+06
	} else {
		instance.MyInt32 = builderInst.MyInt32
	}
	if builderInst.MyFloat64 == 0 {
		instance.MyFloat64 = 2324
	} else {
		instance.MyFloat64 = builderInst.MyFloat64
	}
	if true {
		instance.MyList = []string{"first", "second", "last"}
	} else {
		instance.MyList = builderInst.MyList
	}
	if true {
		instance.MyMap = map[string]bool{"another": false, "key": true}
	} else {
		instance.MyMap = builderInst.MyMap
	}
	if true {
		instance.MyNullablemap = map[string]runtime.Nullable[bool]{"a": runtime.NewNull[bool]()}
	} else {
		instance.MyNullablemap = builderInst.MyNullablemap
	}
	if true {
		instance.MyNullablelist = []runtime.Nullable[float64]{runtime.NewNullable[float64](25), runtime.NewNull[float64](), runtime.NewNullable[float64](43.23)}
	} else {
		instance.MyNullablelist = builderInst.MyNullablelist
	}
	return instance
}

func (u *NexemaDefaultType) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	encoder.EncodeString(u.MyString)

	encoder.EncodeBool(u.MyBool)

	encoder.EncodeInt32(u.MyInt32)

	encoder.EncodeFloat64(u.MyFloat64)

	if u.MyList == nil {
		return nil, fmt.Errorf("field MyList is null but its not marked as nullable")
	}
	if u.MyMap == nil {
		return nil, fmt.Errorf("field MyMap is null but its not marked as nullable")
	}
	if u.MyNullablemap == nil {
		encoder.EncodeNull()
	} else {

		for key, value := range u.MyNullablemap {
			encoder.EncodeString(key)
			if value.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeBool(*value.Value)
			}
		}

	}
	if u.MyNullablelist == nil {
		encoder.EncodeNull()
	} else {

		encoder.BeginArray(int64(len(u.MyNullablelist)))
		for _, element := range u.MyNullablelist {
			if element.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeFloat64(*element.Value)
			}
		}

	}
	return encoder.TakeBytes(), nil
}

func (u *NexemaDefaultType) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u *NexemaDefaultType) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error

	u.MyString, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	u.MyBool, err = decoder.DecodeBool()
	if err != nil {
		return err
	}

	u.MyInt32, err = decoder.DecodeInt32()
	if err != nil {
		return err
	}

	u.MyFloat64, err = decoder.DecodeFloat64()
	if err != nil {
		return err
	}

	myListArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}
	u.MyList = make([]string, myListArrayLen)
	for i := int64(0); i < myListArrayLen; i++ {
		u.MyList[i], err = decoder.DecodeString()
		if err != nil {
			return err
		}
	}

	myMapMapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.MyMap = make(map[string]bool, myMapMapLen)
	for i := int64(0); i < myMapMapLen; i++ {
		key, err := decoder.DecodeString()
		if err != nil {
			return err
		}
		value, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.MyMap[key] = value
	}

	if decoder.IsNextNull() {
		u.MyNullablemap = nil
	} else {

		myNullablemapMapLen, err := decoder.BeginDecodeMap()
		if err != nil {
			return err
		}

		u.MyNullablemap = make(map[string]runtime.Nullable[bool], myNullablemapMapLen)
		for i := int64(0); i < myNullablemapMapLen; i++ {
			key, err := decoder.DecodeString()
			if err != nil {
				return err
			}
			if decoder.IsNextNull() {
				u.MyNullablemap[key] = runtime.NewNull[bool]()
			} else {
				value, err := decoder.DecodeBool()
				if err != nil {
					return err
				}

				u.MyNullablemap[key] = runtime.NewNullable(value)
			}
		}

	}
	if decoder.IsNextNull() {
		u.MyNullablelist = nil
	} else {

		myNullablelistArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}
		u.MyNullablelist = make([]runtime.Nullable[float64], myNullablelistArrayLen)
		for i := int64(0); i < myNullablelistArrayLen; i++ {
			if decoder.IsNextNull() {
				u.MyNullablelist[i] = runtime.NewNull[float64]()
			} else {
				myNullablelist, err := decoder.DecodeFloat64()
				if err != nil {
					return err
				}

				u.MyNullablelist[i] = runtime.NewNullable(myNullablelist)
			}
		}

	}
	return nil
}

func (u *NexemaDefaultType) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *NexemaDefaultType) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *NexemaDefaultType) String() string {
	parts := make([]string, 8)
	parts[0] = fmt.Sprintf("MyString: %v", u.MyString)
	parts[1] = fmt.Sprintf("MyBool: %v", u.MyBool)
	parts[2] = fmt.Sprintf("MyInt32: %v", u.MyInt32)
	parts[3] = fmt.Sprintf("MyFloat64: %v", u.MyFloat64)
	parts[4] = fmt.Sprintf("MyList: %v", u.MyList)
	parts[5] = fmt.Sprintf("MyMap: %v", u.MyMap)
	parts[6] = fmt.Sprintf("MyNullablemap: %v", u.MyNullablemap)
	parts[7] = fmt.Sprintf("MyNullablelist: %v", u.MyNullablelist)
	return fmt.Sprintf("NexemaDefaultType(%s)", strings.Join(parts, ", "))
}

func (u *NexemaDefaultType) Equals(other *NexemaDefaultType) bool {
	if u.MyString != other.MyString {
		return false
	}
	if u.MyBool != other.MyBool {
		return false
	}
	if u.MyInt32 != other.MyInt32 {
		return false
	}
	if u.MyFloat64 != other.MyFloat64 {
		return false
	}
	if (u.MyList == nil) != (other.MyList == nil) {
		return false
	}

	myListLenThis := len(u.MyList)
	myListLenOther := len(other.MyList)
	if myListLenThis != myListLenOther {
		return false
	}

	for i := 0; i < myListLenThis; i++ {
		a := u.MyList[i]
		b := other.MyList[i]

		if a != b {
			return false
		}
	}

	if (u.MyMap == nil) != (other.MyMap == nil) {
		return false
	}

	myMapLenThis := len(u.MyMap)
	myMapLenOther := len(other.MyMap)
	if myMapLenThis != myMapLenOther {
		return false
	}

	for k, a := range u.MyMap {
		b, ok := other.MyMap[k]
		if !ok {
			return false
		}

		if a != b {
			return false
		}
	}

	if (u.MyNullablemap == nil) != (other.MyNullablemap == nil) {
		return false
	}

	myNullablemapLenThis := len(u.MyNullablemap)
	myNullablemapLenOther := len(other.MyNullablemap)
	if myNullablemapLenThis != myNullablemapLenOther {
		return false
	}

	for k, a := range u.MyNullablemap {
		b, ok := other.MyNullablemap[k]
		if !ok {
			return false
		}

		if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
			return false
		}
	}
	if (u.MyNullablelist == nil) != (other.MyNullablelist == nil) {
		return false
	}

	myNullablelistLenThis := len(u.MyNullablelist)
	myNullablelistLenOther := len(other.MyNullablelist)
	if myNullablelistLenThis != myNullablelistLenOther {
		return false
	}

	for i := 0; i < myNullablelistLenThis; i++ {
		a := u.MyNullablelist[i]
		b := other.MyNullablelist[i]

		if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
			return false
		}
	}
	return true
}
