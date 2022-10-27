package test_files

import (
	"bytes"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type MyBarNested struct {
	First MyNestedType
}

func (u *MyBarNested) Serialize() ([]byte, error) {
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
	FirstBinary, err := decoder.DecodeBytes()
	if err != nil {
		return err
	}
	u.First = MyNestedType{}
	u.First.MergeFrom(FirstBinary)
	return nil
}
func (u *MyBarNested) MergeUsing(other *MyBarNested) error {
	u.First = other.First
	return nil
}
