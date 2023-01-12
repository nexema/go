package nexemab

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeBool(t *testing.T) {
	encoder := NewEncoder()
	encoder.encodeBool(true)
	require.Equal(t, []byte{BoolTrue}, encoder.Close())

	encoder = NewEncoder()
	encoder.encodeBool(false)
	require.Equal(t, []byte{BoolFalse}, encoder.Close())

	encoder = NewEncoder()
	encoder.encodeBool(true)
	encoder.encodeBool(true)
	encoder.encodeBool(false)
	encoder.encodeBool(true)
	require.Equal(t, []byte{BoolTrue, BoolTrue, BoolFalse, BoolTrue}, encoder.Close())
}

func TestEncodeString(t *testing.T) {
	encoder := NewEncoder()
	encoder.encodeString("hello world")
	require.Equal(t, []byte("hello world"), encoder.Close())

	var b strings.Builder
	b.Grow(1e+9) // 1gb string
	for i := 0; i < 1e+9; i++ {
		b.WriteByte('a')
	}
	s := b.String()
	encoder = NewEncoder()
	encoder.encodeString(s)
	require.Equal(t, []byte(s), encoder.Close())
}
