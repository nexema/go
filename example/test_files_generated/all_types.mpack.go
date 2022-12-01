package test_files

import "github.com/messagepack-schema/go/runtime"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

type AllTypes struct {
	A bool
	B string
	C uint8
	D uint16
	E uint32
	F uint64
	G int8
	H int16
	I int32
	J int64
	K float32
	L float64
	M []byte
	N []bool
	O []string
	P []uint8
	Q []uint16
	R []uint32
	S []uint64
	T []int8
	U []int16
	V []int32
	W []int64
	X []float32
	Y []float64
}

func (u *AllTypes) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	err = writer.EncodeBool(u.A)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeString(u.B)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeUint8(u.C)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeUint16(u.D)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeUint32(u.E)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeUint64(u.F)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeInt8(u.G)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeInt16(u.H)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeInt32(u.I)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeInt64(u.J)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeFloat32(u.K)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeFloat64(u.L)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeBytes(u.M)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeArrayLen(len(u.N))
	if err != nil {
		return nil, err
	}

	for _, v := range u.N {
		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.O))
	if err != nil {
		return nil, err
	}

	for _, v := range u.O {
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.P))
	if err != nil {
		return nil, err
	}

	for _, v := range u.P {
		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.Q))
	if err != nil {
		return nil, err
	}

	for _, v := range u.Q {
		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.R))
	if err != nil {
		return nil, err
	}

	for _, v := range u.R {
		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.S))
	if err != nil {
		return nil, err
	}

	for _, v := range u.S {
		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.T))
	if err != nil {
		return nil, err
	}

	for _, v := range u.T {
		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.U))
	if err != nil {
		return nil, err
	}

	for _, v := range u.U {
		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.V))
	if err != nil {
		return nil, err
	}

	for _, v := range u.V {
		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.W))
	if err != nil {
		return nil, err
	}

	for _, v := range u.W {
		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.X))
	if err != nil {
		return nil, err
	}

	for _, v := range u.X {
		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeArrayLen(len(u.Y))
	if err != nil {
		return nil, err
	}

	for _, v := range u.Y {
		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}
