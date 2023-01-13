package nexemab

import (
	"bytes"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeUvarint(t *testing.T) {
	tests := []struct {
		input []byte
		want  uint64
		err   error
	}{
		{
			input: []byte{1},
			want:  1,
		},
		{
			input: []byte{232, 7},
			want:  1000,
		},
		{
			input: []byte{139, 2},
			want:  267,
		},
		{
			input: []byte{216, 212, 3},
			want:  59992,
		},
		{
			input: []byte{255, 255, 255, 255, 15},
			want:  math.MaxUint32,
		},
		{
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
			want:  math.MaxUint64,
		},
		{
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
			want:  math.MaxInt64,
		},
		{
			input: []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f, 'b', 'c', 'd'},
			want:  math.MaxInt64,
		},
		{
			input: []byte{224, 11, 2, 'a', 'd'},
			want:  1504,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			decoder := NewDecoder(bytes.NewBuffer(tc.input))
			val, err := decoder.DecodeUvarint()

			require.Equal(t, tc.err, err)
			require.Equal(t, tc.want, val)
		})
	}
}

func TestDecodeVarint(t *testing.T) {
	tests := []struct {
		input []byte
		want  int64
		err   error
	}{
		{
			want:  1,
			input: []byte{0x2},
		},
		{
			want:  267,
			input: []byte{0x96, 0x4},
		},
		{
			want:  392,
			input: []byte{0x90, 0x6},
		},
		{
			want:  59992,
			input: []byte{0xb0, 0xa9, 0x7},
		},
		{
			want:  math.MaxUint32,
			input: []byte{0xfe, 0xff, 0xff, 0xff, 0x1f},
		},
		{
			want:  math.MaxInt32,
			input: []byte{0xfe, 0xff, 0xff, 0xff, 0xf},
		},
		{
			want:  math.MaxInt64,
			input: []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1},
		},
		{
			want:  math.MaxInt64,
			input: []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 5, 'g', '`'},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			decoder := NewDecoder(bytes.NewBuffer(tc.input))
			val, err := decoder.DecodeVarint()
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.want, val)
		})
	}
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
		err   error
	}{
		{
			want:  "hello world",
			input: []byte{22, 104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			decoder := NewDecoder(bytes.NewBuffer(tc.input))
			val, err := decoder.DecodeString()
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.want, val)
		})
	}
}

func TestDecodeNull(t *testing.T) {
	decoder := NewDecoder(bytes.NewBuffer([]byte{null, null}))
	err := decoder.DecodeNull()
	require.Nil(t, err)

	err = decoder.DecodeNull()
	require.Nil(t, err)
}

func TestDecodeUint8(t *testing.T) {
	tests := []struct {
		input []byte
		want  uint8
	}{
		{
			input: []byte{1},
			want:  1,
		},
		{
			input: []byte{122},
			want:  122,
		},
		{
			input: []byte{255},
			want:  255,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			decoder := NewDecoder(bytes.NewBuffer(tc.input))
			val, err := decoder.DecodeUint8()
			require.Nil(t, err)
			require.Equal(t, tc.want, val)
		})
	}
}

func TestDecodeInt8(t *testing.T) {
	tests := []struct {
		input []byte
		want  int8
	}{
		{
			input: []byte{1},
			want:  1,
		},
		{
			input: []byte{122},
			want:  122,
		},
		{
			input: []byte{127},
			want:  127,
		},
		{
			want:  -127,
			input: []byte{129},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			buf := tc.input

			decoder := NewDecoder(bytes.NewBuffer(buf))
			val, err := decoder.DecodeInt8()
			require.Nil(t, err)
			require.Equal(t, tc.want, val)
		})
	}
}
