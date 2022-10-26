package test_files

import (
	"bytes"
	nested "github.com/example/test_files/test_files_generated/sub_package/another/nested"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type MyBarNested struct {
	first nested.MyNestedType
}

func (u *MyBarNested) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	firstBinary, err := u.first.Serialize()
	if err != nil {
		return nil, err
	}
	err = writer.EncodeBytes(firstBinary)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (u *MyBarNested) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}
	return buf
}
func (u *MyBarNested) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := v5.NewDecoder(reader)
	var err error
	firstBinary, err := decoder.DecodeBytes()
	if err != nil {
		return err
	}
	u.first = MyNestedType{}
	u.first.MergeFrom(firstBinary)
	return nil
}
func (u *MyBarNested) MergeUsing(other *MyBarNested) error {
	u.first = other.first
	return nil
}
