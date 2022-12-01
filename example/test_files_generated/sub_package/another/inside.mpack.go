package test_files

import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"
import "fmt"

type AnotherType struct {
	First  string
	Second bool
	Status AnotherEnum
}

func (u *AnotherType) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	err = writer.EncodeString(u.First)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeBool(u.Second)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeUint8(u.Status.Index())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (u *AnotherType) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}

	return buf
}

func (u *AnotherType) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := msgpack.NewDecoder(reader)
	var err error

	u.First, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	u.Second, err = decoder.DecodeBool()
	if err != nil {
		return err
	}

	StatusIdx, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.Status = AnotherEnumPicker.ByIndex(StatusIdx)

	return nil
}

func (u *AnotherType) MergeUsing(other *AnotherType) {
	u.First = other.First
	u.Second = other.Second
	u.Status = other.Status
}

type AnotherEnum struct {
	index uint8
	name  string
}

type anotherEnumPicker struct{}

var AnotherEnumPicker *anotherEnumPicker = &anotherEnumPicker{}
var anotherEnumUnknown AnotherEnum = AnotherEnum{index: 0, name: "unknown"}
var anotherEnumSuccess AnotherEnum = AnotherEnum{index: 1, name: "success"}
var anotherEnumFailed AnotherEnum = AnotherEnum{index: 2, name: "failed"}

func (*anotherEnumPicker) Unknown() AnotherEnum {
	return anotherEnumUnknown
}

func (*anotherEnumPicker) Success() AnotherEnum {
	return anotherEnumSuccess
}

func (*anotherEnumPicker) Failed() AnotherEnum {
	return anotherEnumFailed
}

func (e AnotherEnum) Index() uint8 {
	return e.index
}

func (e AnotherEnum) Name() string {
	return e.name
}

func (*anotherEnumPicker) ByIndex(index uint8) AnotherEnum {
	switch index {

	case 0:
		return anotherEnumUnknown

	case 1:
		return anotherEnumSuccess

	case 2:
		return anotherEnumFailed

	default:
		panic(fmt.Sprintf("AnotherEnum with index %v not found", index))
	}
}

func (*anotherEnumPicker) ByName(name string) AnotherEnum {
	switch name {

	case "unknown":
		return anotherEnumUnknown

	case "success":
		return anotherEnumSuccess

	case "failed":
		return anotherEnumFailed

	default:
		panic(fmt.Sprintf("AnotherEnum with name %v not found", name))
	}
}
