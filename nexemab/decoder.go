package nexemab

import (
	"bufio"
	"errors"
	"io"
	"math"
)

// Decoder provides methods to decode a nexema binary input.
// Every method does not check if the decoded data is the type its being decoded, for example:
// if you try to decode a null, and the next 1 byte is not 0xc0, it will not report an error, but maybe the next
// data to read isn't valid.
// It's a design behaviour, searching for a better performance, every struct serialized and deserialized using nexemab
// is expected to not be corrupted or with a wrong format.
//
// todo: find a way to add some kind of control bytes
type Decoder struct {
	buf []byte
	r   *bufio.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		r:   bufio.NewReader(reader),
		buf: make([]byte, 9),
	}
}

func (d *Decoder) DecodeNull() error {
	_, err := d.r.ReadByte()
	return err
}

// IsNextNull returns true if the next byte to decode represents NULL
func (d *Decoder) IsNextNull() bool {
	buf, err := d.r.Peek(1)
	if err != nil {
		return false
	}

	return buf[0] == null
}

func (d *Decoder) DecodeBool() (bool, error) {
	b, err := d.r.ReadByte()
	if err != nil {
		return false, err
	}

	return b == boolTrue, nil
}

func (d *Decoder) DecodeString() (string, error) {
	// Read string length
	strLen, err := d.DecodeVarint()
	if err != nil {
		return "", err
	}

	buf := make([]byte, strLen)
	_, err = d.r.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (d *Decoder) DecodeBinary() ([]byte, error) {
	// Read binary length
	bufLen, err := d.DecodeVarint()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, bufLen)
	_, err = d.r.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (d *Decoder) DecodeUvarint() (uint64, error) {

	// code taken from binary.Uvarint but modified to read from underlying buffer
	var x uint64
	var s uint

	var b byte
	var i int = 0
	var err error

	for {
		b, err = d.r.ReadByte()
		if err != nil {
			return 0, nil
		}

		if i == maxVarintLen {
			return 0, errors.New("overflow")
		}
		if b < 0x80 {
			if i == maxVarintLen-1 && b > 1 {
				return 0, errors.New("overflow")
			}
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
		i++
	}
}

func (d *Decoder) DecodeVarint() (int64, error) {
	// code taken from binary.Varint but modified to read from underlying buffer
	ux, err := d.DecodeUvarint()
	x := int64(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x, err
}

func (d *Decoder) DecodeUint8() (uint8, error) {
	b, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}

	return b, nil
}

func (d *Decoder) DecodeInt8() (int8, error) {
	v, err := d.DecodeUint8()
	return int8(v), err
}

func (d *Decoder) DecodeUint16() (uint16, error) {
	d.buf = d.buf[:2]
	_, err := d.r.Read(d.buf)
	if err != nil {
		return 0, err
	}

	return (uint16(d.buf[0]) << 8) | uint16(d.buf[1]), nil
}

func (d *Decoder) DecodeInt16() (int16, error) {
	v, err := d.DecodeUint16()
	return int16(v), err
}

func (d *Decoder) DecodeUint32() (uint32, error) {
	d.buf = d.buf[:4]
	_, err := d.r.Read(d.buf)
	if err != nil {
		return 0, err
	}

	return (uint32(d.buf[0]) << 24) | (uint32(d.buf[1]) << 16) | (uint32(d.buf[2]) << 8) | uint32(d.buf[3]), nil
}

func (d *Decoder) DecodeInt32() (int32, error) {
	v, err := d.DecodeUint32()
	return int32(v), err
}

func (d *Decoder) DecodeUint64() (uint64, error) {
	d.buf = d.buf[:8]
	_, err := d.r.Read(d.buf)
	if err != nil {
		return 0, err
	}

	return (uint64(d.buf[0]) << 56) |
		(uint64(d.buf[1]) << 48) |
		(uint64(d.buf[2]) << 40) |
		(uint64(d.buf[3]) << 32) |
		(uint64(d.buf[4]) << 24) |
		(uint64(d.buf[5]) << 16) |
		(uint64(d.buf[6]) << 8) |
		uint64(d.buf[7]), nil
}

func (d *Decoder) DecodeInt64() (int64, error) {
	v, err := d.DecodeUint64()
	return int64(v), err
}

func (d *Decoder) DecodeFloat32() (float32, error) {
	v, err := d.DecodeUint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(v), err
}

func (d *Decoder) DecodeFloat64() (float64, error) {
	v, err := d.DecodeUint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(v), err
}
