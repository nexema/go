package test_files

import test_files "github.com/example/test_files/test_files_generated/sub_package/another"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type MyNestedType struct{ First test_files.AnotherType }

func (u *MyNestedType) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
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
	decoder := msgpack.NewDecoder(reader)
	var err error

	FirstBinary, err := decoder.DecodeBytes()
	if err != nil {
		return err
	}

	err = u.First.MergeFrom(FirstBinary)
	if err != nil {
		return err
	}

	return nil
}
