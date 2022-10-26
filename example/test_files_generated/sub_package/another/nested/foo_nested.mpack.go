package test_files

import (
	"bytes"
	another "github.com/example/test_files/test_files_generated/sub_package/another"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type MyNestedType struct {
	first another.AnotherEnum
}

func (u *MyNestedType) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	err = writer.EncodeUint8(u.first.Index())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (u *MyNestedType) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}
	return buf
}
func (u *MyNestedType) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := v5.NewDecoder(reader)
	var err error
	firstIdx, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}
	u.first = AnotherEnumPicker.ByIndex(firstIdx)
	return nil
}
func (u *MyNestedType) MergeUsing(other *MyNestedType) error {
	u.first = other.first
	return nil
}
