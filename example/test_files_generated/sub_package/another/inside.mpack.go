package test_files

import "github.com/messagepack-schema/go/runtime"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type AnotherType struct {
	First  string
	Second bool
	Status custom
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
