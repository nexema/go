package runtime

type SchemaType[T any] interface {
	Serialize() ([]byte, error)
	MustSerialize() []byte
	MergeFrom(buffer []byte) error
	MergeUsing(other *T) error
}
