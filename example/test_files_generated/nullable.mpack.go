package test_files

import "github.com/messagepack-schema/go/runtime"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type Nullables struct {
	A1  runtime.Nullable[bool]
	A2  runtime.Nullable[string]
	A3  runtime.Nullable[uint8]
	A4  runtime.Nullable[uint16]
	A5  runtime.Nullable[uint32]
	A6  runtime.Nullable[uint64]
	A7  runtime.Nullable[int8]
	A8  runtime.Nullable[int16]
	A9  runtime.Nullable[int32]
	A10 runtime.Nullable[int64]
	A11 runtime.Nullable[float32]
	A12 runtime.Nullable[float64]
	A13 runtime.Nullable[[]byte]
	A14 []runtime.Nullable[bool]
	A15 map[string]runtime.Nullable[float32]
}

func (u *Nullables) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	return buf.Bytes(), nil
}
