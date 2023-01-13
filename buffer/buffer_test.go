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

var value = []byte("hello world this is only a test")
var valueLen = len(value)

func BenchmarkBuffer(b *testing.B) {
	buffer := NewBuffer(BufferConfig{Capacity: 5, InitialBufferSize: 300})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// buffer.Write([]byte("this is m"))
		buffer.Write(value)
		buffer.Write(value)
		buffer.Write(value)
		buffer.Reset()
	}

}

func BenchmarkBufferGo(b *testing.B) {
	buffer := new(bytes.Buffer)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer.Write(value)
		buffer.Write(value)
		buffer.Write(value)
		buffer.Reset()
	}

}

func BenchmarkBufCopy(b *testing.B) {
	initialBuffer := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	for i := 0; i < b.N; i++ {
		newBuffer := make([]byte, 10, 50)
		copy(newBuffer, initialBuffer)
	}
}

func BenchmarkBufAppend(b *testing.B) {
	initialBuffer := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	for i := 0; i < b.N; i++ {
		newBuffer := make([]byte, 0, 50)
		newBuffer = append(newBuffer, initialBuffer...)
		_ = newBuffer
	}
}
