package test_files

import (
	"bytes"
	another "github.com/example/test_files/test_files_generated/sub_package/another"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type MyNestedType struct {
	First another.AnotherType
}

func (u *MyNestedType) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	FirstBinary, err := u.First.Serialize()
	if err != nil {
		return nil, err
	}
	err = writer.EncodeBytes(FirstBinary)
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
	FirstBinary, err := decoder.DecodeBytes()
	if err != nil {
		return err
	}
	u.First = another.AnotherType{}
	u.First.MergeFrom(FirstBinary)
	return nil
}
func (u *MyNestedType) MergeUsing(other *MyNestedType) error {
	u.First = other.First
	return nil
}
