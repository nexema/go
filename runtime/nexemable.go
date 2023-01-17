package runtime

import (
	"io"
	"sync"
)

// Nexemable interface provides methods to serialize and deserialize an struct to/from
// nexemab. Any struct which implements this interface becomes a Nexemable type.
type Nexemable interface {
	Encode() ([]byte, error)
	MustEncode() []byte

	Decode(reader io.Reader) error
	MustDecode(reader io.Reader)
	MergeUsing(other Nexemable) error
	MergeFrom(buffer []byte) error
	Clone() Nexemable
}

// NexemaType represents the base struct for any Nexema's type. This is not meant to be used at runtime
type NexemaType struct {
	NoCopy
}

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
type NoCopy [0]sync.Mutex

// Lock is a no-op used by -copylocks checker from `go vet`.
// func (*NoCopy) Lock()   {}
// func (*NoCopy) UnLock() {}
