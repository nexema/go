package runtime

import (
	"bytes"
	"encoding/json"
)

type Nullable[T any] struct {
	Value *T
}

func (n *Nullable[T]) GetValue() T {
	return *n.Value
}

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{
		Value: &value,
	}
}

func NewNull[T any]() Nullable[T] {
	return Nullable[T]{}
}

func (n *Nullable[T]) SetValue(value T) {
	n.Value = &value
}

func (n *Nullable[T]) HasValue() bool {
	return n.Value != nil
}

func (n *Nullable[T]) Clear() {
	n.Value = nil
}

func (n *Nullable[T]) ValueOrDefault() *T {
	if n.Value == nil {
		var zero T
		return &zero
	}

	return n.Value
}

var jsonNullValue = []byte{110, 117, 108, 108}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.Value == nil {
		return jsonNullValue, nil
	}

	return json.Marshal(n.Value)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNullValue) {
		n.Value = nil
		return nil
	}

	var value T
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	n.Value = &value
	return nil
}
