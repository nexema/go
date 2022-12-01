package test_files

import "github.com/messagepack-schema/go/runtime"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type MyBarNested struct {
	First runtime.Nullable[MyNestedType]
}

func (u *MyBarNested) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	if u.First.HasValue() {
		FirstBinary, err := u.First.Value.Serialize()
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBytes(FirstBinary)
		if err != nil {
			return nil, err
		}
	} else {
		err = writer.EncodeNil()
		if err != nil {
			return nil, err
		}
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
	decoder := msgpack.NewDecoder(reader)
	var err error
	var isNextNil bool

	isNextNil, err = decoder.IsNextNil()
	if err != nil {
		return err
	}

	if isNextNil {
		u.First = runtime.NewNull[MyNestedType]()
	} else {
		FirstBinary, err := decoder.DecodeBytes()
		if err != nil {
			return err
		}

		value := MyNestedType{}
		err = value.MergeFrom(FirstBinary)
		if err != nil {
			return err
		}
		u.First = runtime.NewNullable(value)
	}

	return nil
}
