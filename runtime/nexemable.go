package runtime

import "io"

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
