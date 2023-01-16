package identity

import (
	"bytes"
	"io"

	"github.com/example/models"
	"github.com/nexema/go/runtime"
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

func (u NexemaPrimitives) Decode(reader io.Reader) error {
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

func (u NexemaPrimitives) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u NexemaPrimitives) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
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

func (u *NexemaNullablePrimitives) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	if u.MyString.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyString.Value

		encoder.EncodeString(value)

	}

	if u.MyBoolean.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyBoolean.Value

		encoder.EncodeBool(value)

	}

	if u.MyUint8.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyUint8.Value

		encoder.EncodeUint8(value)

	}

	if u.MyUint16.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyUint16.Value

		encoder.EncodeUint16(value)

	}

	if u.MyUint32.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyUint32.Value

		encoder.EncodeUint32(value)

	}

	if u.MyUint64.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyUint64.Value

		encoder.EncodeUint64(value)

	}

	if u.MyInt8.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyInt8.Value

		encoder.EncodeInt8(value)

	}

	if u.MyInt16.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyInt16.Value

		encoder.EncodeInt16(value)

	}

	if u.MyInt32.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyInt32.Value

		encoder.EncodeInt32(value)

	}

	if u.MyInt64.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyInt64.Value

		encoder.EncodeInt64(value)

	}

	if u.MyFloat32.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyFloat32.Value

		encoder.EncodeFloat32(value)

	}

	if u.MyFloat64.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyFloat64.Value

		encoder.EncodeFloat64(value)

	}

	if u.MyBinary.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.MyBinary.Value

		encoder.EncodeBinary(value)

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

func (u NexemaNullablePrimitives) Decode(reader io.Reader) error {
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

func (u NexemaNullablePrimitives) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u NexemaNullablePrimitives) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

type NexemaList struct {
	List1 []string
	List2 [][]byte
	List3 []runtime.Nullable[string]
	List4 runtime.Nullable[[]runtime.Nullable[string]]
	List5 runtime.Nullable[[][]byte]
	List6 []runtime.Nullable[[]byte]
	List7 runtime.Nullable[[]runtime.Nullable[[]byte]]
}

func (u *NexemaList) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	if u.List4.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.List4.Value

		encoder.BeginArray(int64(len(value)))
		for _, element := range value {
			if element.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeString(*element.Value)
			}
		}

	}

	if u.List5.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.List5.Value

		encoder.BeginArray(int64(len(value)))
		for _, element := range value {
			encoder.EncodeBinary(element)
		}

	}

	if u.List7.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.List7.Value

		encoder.BeginArray(int64(len(value)))
		for _, element := range value {
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

func (u NexemaList) Decode(reader io.Reader) error {
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

			u.List3[i] = runtime.NewNullable[string](list3)
		}
	}

	if decoder.IsNextNull() {
		u.List4.Clear()
	} else {

		list4ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}

		u.List4.SetValue(make([]runtime.Nullable[string], list4ArrayLen))
		for i := int64(0); i < list4ArrayLen; i++ {
			if decoder.IsNextNull() {
				(*u.List4.Value)[i] = runtime.NewNull[string]()
			} else {
				list4, err := decoder.DecodeString()
				if err != nil {
					return err
				}
				(*u.List4.Value)[i] = runtime.NewNullable[string](list4)
			}
		}

	}

	if decoder.IsNextNull() {
		u.List5.Clear()
	} else {

		list5ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}

		u.List5.SetValue(make([][]byte, list5ArrayLen))
		for i := int64(0); i < list5ArrayLen; i++ {
			(*u.List5.Value)[i], err = decoder.DecodeBinary()
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

			u.List6[i] = runtime.NewNullable[[]byte](list6)
		}
	}

	if decoder.IsNextNull() {
		u.List7.Clear()
	} else {

		list7ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}

		u.List7.SetValue(make([]runtime.Nullable[[]byte], list7ArrayLen))
		for i := int64(0); i < list7ArrayLen; i++ {
			if decoder.IsNextNull() {
				(*u.List7.Value)[i] = runtime.NewNull[[]byte]()
			} else {
				list7, err := decoder.DecodeBinary()
				if err != nil {
					return err
				}
				(*u.List7.Value)[i] = runtime.NewNullable[[]byte](list7)
			}
		}

	}

	return nil
}

func (u NexemaList) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u NexemaList) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
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

func (u *NexemaMap) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	encoder.BeginMap(int64(len(u.Map0)))
	for key, value := range u.Map0 {
		encoder.EncodeString(key)
		encoder.EncodeString(value)
	}

	encoder.BeginMap(int64(len(u.Map1)))
	for key, value := range u.Map1 {
		encoder.EncodeString(key)
		if value.IsNull() {
			encoder.EncodeNull()
		} else {
			encoder.EncodeString(*value.Value)
		}
	}

	if u.Map2.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.Map2.Value

		encoder.BeginMap(int64(len(u.Map2)))
		for key, value := range u.Map2 {
			encoder.EncodeString(key)
			encoder.EncodeString(value)
		}

	}

	if u.Map3.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.Map3.Value

		encoder.BeginMap(int64(len(u.Map3)))
		for key, value := range u.Map3 {
			encoder.EncodeString(key)
			if value.IsNull() {
				encoder.EncodeNull()
			} else {
				encoder.EncodeString(*value.Value)
			}
		}

	}

	encoder.BeginMap(int64(len(u.Map4)))
	for key, value := range u.Map4 {
		encoder.EncodeString(key)
		encoder.EncodeBinary(value)
	}

	encoder.BeginMap(int64(len(u.Map5)))
	for key, value := range u.Map5 {
		encoder.EncodeString(key)
		if value.IsNull() {
			encoder.EncodeNull()
		} else {
			encoder.EncodeBinary(*value.Value)
		}
	}

	if u.Map6.IsNull() {
		encoder.EncodeNull()
	} else {
		value := *u.Map6.Value

		encoder.BeginMap(int64(len(u.Map6)))
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

func (u NexemaMap) Decode(reader io.Reader) error {
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

			u.Map1[key] = runtime.NewNullable[string](value)
		}

	}

	if decoder.IsNextNull() {
		u.Map2.Clear()
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
		u.Map3.Clear()
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

				u.Map3[key] = runtime.NewNullable[string](value)
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

			u.Map5[key] = runtime.NewNullable[[]byte](value)
		}

	}

	if decoder.IsNextNull() {
		u.Map6.Clear()
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

				u.Map6[key] = runtime.NewNullable[[]byte](value)
			}

		}

	}

	return nil
}

func (u NexemaMap) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u NexemaMap) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

type EmbeddedType struct {
	MyUnion SingleUnion
	MyEnum  models.Colors
}

func (u *EmbeddedType) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()

	myUnionBytes, err := myUnion.Encode()
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

func (u EmbeddedType) Decode(reader io.Reader) error {
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

func (u EmbeddedType) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u EmbeddedType) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}
