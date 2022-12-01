package test_files

import "github.com/messagepack-schema/go/runtime"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type MyNestedType struct{ First custom }

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
