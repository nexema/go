package test_files

import (
	"bytes"

	another "github.com/example/test_files/test_files_generated/sub_package/another"
	v5 "github.com/vmihailenco/msgpack/v5"
)

type AnotherType struct {
	first  string
	second bool
	status another.AnotherEnum
}

func (u *AnotherType) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := v5.NewEncoder(buf)
	var err error
	err = writer.EncodeString(u.first)
	if err != nil {
		return nil, err
	}
	err = writer.EncodeBool(u.second)
	if err != nil {
		return nil, err
	}
	statusBinary, err := u.status.Serialize()
	if err != nil {
		return nil, err
	}
	err = writer.EncodeBytes(statusBinary)
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
	decoder := v5.NewDecoder(reader)
	var err error
	u.first, err = decoder.DecodeString()
	if err != nil {
		return err
	}
	u.second, err = decoder.DecodeBool()
	if err != nil {
		return err
	}
	return nil
}
func (u *AnotherType) MergeUsing(other *AnotherType) error {
	u.first = other.first
	u.second = other.second
	u.status = other.status
	return nil
}
