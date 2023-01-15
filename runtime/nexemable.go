package runtime

// Nexemable interface provides methods to serialize and deserialize an struct to/from
// nexemab. Any struct which implements this interface becomes a Nexemable type.
type Nexemable interface {
	Encode() ([]byte, error)
	MustEncode() []byte

	Decode(buffer []byte) error
	MustDecode(buffer []byte)
}
