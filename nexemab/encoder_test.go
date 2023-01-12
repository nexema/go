package nexemab

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeBool(t *testing.T) {
	encoder := NewEncoder()
	encoder.EncodeBool(true)
	require.Equal(t, []byte{BoolTrue}, encoder.Close())

	encoder = NewEncoder()
	encoder.EncodeBool(false)
	require.Equal(t, []byte{BoolFalse}, encoder.Close())

	encoder = NewEncoder()
	encoder.EncodeBool(true)
	encoder.EncodeBool(true)
	encoder.EncodeBool(false)
	encoder.EncodeBool(true)
	require.Equal(t, []byte{BoolTrue, BoolTrue, BoolFalse, BoolTrue}, encoder.Close())
}

func TestEncodeString(t *testing.T) {
	encoder := NewEncoder()
	const hw = "hello world"
	encoder.encodeString(hw)
	outBuf := encoder.Close()

	hwLen := int64(len(hw))

	encoder = NewEncoder()
	encoder.EncodeVarint(hwLen)
	testBuf := make([]byte, 0)
	testBuf = append(testBuf, encoder.Close()...)
	testBuf = append(testBuf, []byte(hw)...)
	require.Equal(t, testBuf, outBuf)

	// var b strings.Builder
	// b.Grow(1e+9) // 1gb string, keep in mind while testing
	// for i := 0; i < 1e+9; i++ {
	// 	b.WriteByte('a')
	// }
	// s := b.String()
	// encoder = NewEncoder()
	// encoder.encodeString(s)
	// require.Equal(t, []byte(s), encoder.Close())
}

func TestEncodeUvarint(t *testing.T) {
	tests := []struct {
		input uint64
		want  []byte
	}{
		{
			input: 1,
			want:  []byte{1},
		},
		{
			input: 267,
			want:  []byte{139, 2},
		},
		{
			input: 392,
			want:  []byte{136, 3},
		},
		{
			input: 59992,
			want:  []byte{216, 212, 3},
		},
		{
			input: math.MaxUint32,
			want:  []byte{255, 255, 255, 255, 15},
		},
		{
			input: math.MaxInt32,
			want:  []byte{255, 255, 255, 255, 7},
		},
		{
			input: math.MaxUint64,
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			input: math.MaxInt64,
			want:  []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			encoder := NewEncoder()
			encoder.EncodeUvarint(tc.input)
			got := encoder.Close()
			require.Equal(t, tc.want, got)
		})
	}
}

func TestEncodeVarint(t *testing.T) {
	tests := []struct {
		input int64
		want  []byte
	}{
		{
			input: 1,
			want:  []byte{0x2},
		},
		{
			input: 267,
			want:  []byte{0x96, 0x4},
		},
		{
			input: 392,
			want:  []byte{0x90, 0x6},
		},
		{
			input: 59992,
			want:  []byte{0xb0, 0xa9, 0x7},
		},
		{
			input: math.MaxUint32,
			want:  []byte{0xfe, 0xff, 0xff, 0xff, 0x1f},
		},
		{
			input: math.MaxInt32,
			want:  []byte{0xfe, 0xff, 0xff, 0xff, 0xf},
		},
		{
			input: math.MaxInt64,
			want:  []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			encoder := NewEncoder()
			encoder.EncodeVarint(tc.input)
			got := encoder.Close()
			require.Equal(t, tc.want, got)
		})
	}
}
