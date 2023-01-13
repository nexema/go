package buffer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnsure(t *testing.T) {
	buffer := NewBuffer(BufferConfig{InitialBufferSize: 5, Capacity: 5})
	resized := buffer.ensure(5) // should not resize because initial size is 10
	require.False(t, resized)

	// write some data
	// buffer.writeByte('a')
	// buffer.writeByte('b')
	// buffer.writeByte('c')
	// buffer.writeByte('d')
	// buffer.writeByte('e')

	buffer.Write([]byte("abcde"))
	buffer.Write([]byte("hola")) // here it should be resized

	// but now should resize
	// resized = buffer.ensure(64)
	// require.True(t, resized)
	// require.Len(t, buffer.backstore, 2)
	// require.Equal(t, buffer.offset, 0)
	// require.Equal(t, buffer.size, 64*2)

	result := buffer.Bytes()
	_ = result
}

func BenchmarkBuffer(b *testing.B) {
	buffer := NewBuffer(BufferConfig{Capacity: 5, InitialBufferSize: 10})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer.Write([]byte("this is m"))
		buffer.Write([]byte("hello world this is only a test"))
	}

	buffer.Bytes()
}

func BenchmarkGoBuffer(b *testing.B) {
	buffer := new(bytes.Buffer)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer.Write([]byte("this is m"))
		buffer.Write([]byte("hello world this is only a test"))
	}

	buffer.Bytes()
}
