package nexemab

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeBool(t *testing.T) {
	encoder := NewEncoder()
	encoder.EncodeBool(true)
	require.Equal(t, []byte{boolTrue}, encoder.Close())

	encoder = NewEncoder()
	encoder.EncodeBool(false)
	require.Equal(t, []byte{boolFalse}, encoder.Close())

	encoder = NewEncoder()
	encoder.EncodeBool(true)
	encoder.EncodeBool(true)
	encoder.EncodeBool(false)
	encoder.EncodeBool(true)
	require.Equal(t, []byte{boolTrue, boolTrue, boolFalse, boolTrue}, encoder.Close())
}

func TestEncodeString(t *testing.T) {
	encoder := NewEncoder()
	const hw = "hello world"
	encoder.EncodeString(hw)
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

func TestEncode(t *testing.T) {
	encoder := NewEncoder()
	encoder.EncodeBool(true)
	encoder.EncodeBool(false)
	encoder.EncodeBool(true)
	encoder.EncodeBinary([]byte{5, 255, 98})
	encoder.EncodeFloat32(32.433)
	encoder.EncodeFloat64(1543.0998990)
	encoder.EncodeNull()
	encoder.EncodeInt8(-11)
	encoder.EncodeUint8(2)
	encoder.EncodeInt16(-2555)
	encoder.EncodeUint16(12222)
	encoder.EncodeInt32(-487574930)
	encoder.EncodeUint32(3413241124)
	encoder.EncodeInt64(-112414)
	encoder.EncodeUint64(111112321412414)
	encoder.EncodeVarint(999844)
	encoder.EncodeUvarint(812)
	encoder.EncodeNull()
	encoder.EncodeString("hello world")
	encoder.EncodeVarint(3333)
	buffer := encoder.Close()

	decoder := NewDecoder(bytes.NewBuffer(buffer))
	requires(decoder.DecodeBool()).toBe(t, true)
	requires(decoder.DecodeBool()).toBe(t, false)
	requires(decoder.DecodeBool()).toBe(t, true)
	requires(decoder.DecodeBinary()).toBe(t, []byte{5, 255, 98})
	requires(decoder.DecodeFloat32()).toBe(t, 32.433)
	requires(decoder.DecodeFloat64()).toBe(t, 1543.0998990)
	require.True(t, decoder.IsNextNull())
	require.Nil(t, decoder.DecodeNull()) // consume nil
	requires(decoder.DecodeInt8()).toBe(t, -11)
	requires(decoder.DecodeUint8()).toBe(t, 2)
	requires(decoder.DecodeInt16()).toBe(t, -2555)
	requires(decoder.DecodeUint16()).toBe(t, 12222)
	requires(decoder.DecodeInt32()).toBe(t, -487574930)
	requires(decoder.DecodeUint32()).toBe(t, 3413241124)
	requires(decoder.DecodeInt64()).toBe(t, -112414)
	requires(decoder.DecodeUint64()).toBe(t, 111112321412414)
	requires(decoder.DecodeVarint()).toBe(t, 999844)
	requires(decoder.DecodeUvarint()).toBe(t, 812)
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encoder := NewEncoder()
		encoder.EncodeBool(true)
		encoder.EncodeBool(false)
		encoder.EncodeBool(true)
		encoder.EncodeBinary([]byte{5, 255, 98})
		encoder.EncodeFloat32(32.433)
		encoder.EncodeFloat64(1543.0998990)
		encoder.EncodeNull()
		encoder.EncodeInt8(-11)
		encoder.EncodeUint8(2)
		encoder.EncodeInt16(-2555)
		encoder.EncodeUint16(12222)
		encoder.EncodeInt32(-487574930)
		encoder.EncodeUint32(3413241124)
		encoder.EncodeInt64(-112414)
		encoder.EncodeUint64(111112321412414)
		encoder.EncodeVarint(999844)
		encoder.EncodeUvarint(812)
		encoder.EncodeNull()
		encoder.EncodeString("hello world")
		encoder.EncodeVarint(3333)
		encoder.Close()
	}
}

type requeriment[T any] struct {
	given T
	err   error
}

func requires[T any](given T, err error) *requeriment[T] {
	return &requeriment[T]{given: given, err: err}
}

func (r *requeriment[T]) toBe(t *testing.T, want T) {
	require.Nil(t, r.err)
	require.Equal(t, want, r.given)
}
