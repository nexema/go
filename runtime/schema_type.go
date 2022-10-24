package runtime

type SchemaType[T any] interface {
	Serialize() ([]byte, error)
	MergeFrom(buffer []byte) error
	MergeUsing(other *T) error
}
