package test_files

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

func (u *AllTypes) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}

	return buf
}

func (u *AllTypes) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := msgpack.NewDecoder(reader)
	var err error

	u.A, err = decoder.DecodeBool()
	if err != nil {
		return err
	}

	u.B, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	u.C, err = decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.D, err = decoder.DecodeUint16()
	if err != nil {
		return err
	}

	u.E, err = decoder.DecodeUint32()
	if err != nil {
		return err
	}

	u.F, err = decoder.DecodeUint64()
	if err != nil {
		return err
	}

	u.G, err = decoder.DecodeInt8()
	if err != nil {
		return err
	}

	u.H, err = decoder.DecodeInt16()
	if err != nil {
		return err
	}

	u.I, err = decoder.DecodeInt32()
	if err != nil {
		return err
	}

	u.J, err = decoder.DecodeInt64()
	if err != nil {
		return err
	}

	u.K, err = decoder.DecodeFloat32()
	if err != nil {
		return err
	}

	u.L, err = decoder.DecodeFloat64()
	if err != nil {
		return err
	}

	u.M, err = decoder.DecodeBytes()
	if err != nil {
		return err
	}

	NLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.N = make([]bool, NLen)
	for i := 0; i < NLen; i++ {
		u.N[i], err = decoder.DecodeBool()
		if err != nil {
			return err
		}
	}

	OLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.O = make([]string, OLen)
	for i := 0; i < OLen; i++ {
		u.O[i], err = decoder.DecodeString()
		if err != nil {
			return err
		}
	}

	PLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.P = make([]uint8, PLen)
	for i := 0; i < PLen; i++ {
		u.P[i], err = decoder.DecodeUint8()
		if err != nil {
			return err
		}
	}

	QLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.Q = make([]uint16, QLen)
	for i := 0; i < QLen; i++ {
		u.Q[i], err = decoder.DecodeUint16()
		if err != nil {
			return err
		}
	}

	RLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.R = make([]uint32, RLen)
	for i := 0; i < RLen; i++ {
		u.R[i], err = decoder.DecodeUint32()
		if err != nil {
			return err
		}
	}

	SLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.S = make([]uint64, SLen)
	for i := 0; i < SLen; i++ {
		u.S[i], err = decoder.DecodeUint64()
		if err != nil {
			return err
		}
	}

	TLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.T = make([]int8, TLen)
	for i := 0; i < TLen; i++ {
		u.T[i], err = decoder.DecodeInt8()
		if err != nil {
			return err
		}
	}

	ULen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.U = make([]int16, ULen)
	for i := 0; i < ULen; i++ {
		u.U[i], err = decoder.DecodeInt16()
		if err != nil {
			return err
		}
	}

	VLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.V = make([]int32, VLen)
	for i := 0; i < VLen; i++ {
		u.V[i], err = decoder.DecodeInt32()
		if err != nil {
			return err
		}
	}

	WLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.W = make([]int64, WLen)
	for i := 0; i < WLen; i++ {
		u.W[i], err = decoder.DecodeInt64()
		if err != nil {
			return err
		}
	}

	XLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.X = make([]float32, XLen)
	for i := 0; i < XLen; i++ {
		u.X[i], err = decoder.DecodeFloat32()
		if err != nil {
			return err
		}
	}

	YLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.Y = make([]float64, YLen)
	for i := 0; i < YLen; i++ {
		u.Y[i], err = decoder.DecodeFloat64()
		if err != nil {
			return err
		}
	}

	return nil
}
