package test_files

import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

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
