package runtime

// Nullable represents T but allowing to set it as "null"
type Nullable[T any] struct {
	Value *T
}

// NewNull creates a new Nullable[T] without a value
func NewNull[T any]() Nullable[T] {
	return Nullable[T]{}
}

// NewNullalbe creates a new Nullable[T] setting its value
// to the pointer of v
func NewNullable[T any](v T) Nullable[T] {
	return Nullable[T]{Value: &v}
}

// ValueOrZero returns a zero value of T if its null, otherwise, its actual value.
func (n *Nullable[T]) ValueOrZero() T {
	if n.Value == nil {
		var zero T
		return zero
	}

	return *n.Value
}

// IsNull returns true if the current nullable does not have a value
func (n *Nullable[T]) IsNull() bool {
	return n.Value == nil
}

// SetValues sets the value of the nullable
func (n *Nullable[T]) SetValue(v T) {
	n.Value = &v
}
