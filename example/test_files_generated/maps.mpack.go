package test_files

import "bytes"
import "github.com/messagepack-schema/go/runtime/msgpack"

type MapTest struct {
	A1  map[string]bool
	A2  map[string]string
	A3  map[string]uint8
	A4  map[string]uint16
	A5  map[string]uint32
	A6  map[string]uint64
	A7  map[string]int8
	A8  map[string]int16
	A9  map[string]int32
	A10 map[string]int64
	A11 map[string]float32
	A12 map[string]float64
	B1  map[uint8]bool
	B2  map[uint8]string
	B3  map[uint8]uint8
	B4  map[uint8]uint16
	B5  map[uint8]uint32
	B6  map[uint8]uint64
	B7  map[uint8]int8
	B8  map[uint8]int16
	B9  map[uint8]int32
	B10 map[uint8]int64
	B11 map[uint8]float32
	B12 map[uint8]float64
	C1  map[uint16]bool
	C2  map[uint16]string
	C3  map[uint16]uint8
	C4  map[uint16]uint16
	C5  map[uint16]uint32
	C6  map[uint16]uint64
	C7  map[uint16]int8
	C8  map[uint16]int16
	C9  map[uint16]int32
	C10 map[uint16]int64
	C11 map[uint16]float32
	C12 map[uint16]float64
	D1  map[uint32]bool
	D2  map[uint32]string
	D3  map[uint32]uint8
	D4  map[uint32]uint16
	D5  map[uint32]uint32
	D6  map[uint32]uint64
	D7  map[uint32]int8
	D8  map[uint32]int16
	D9  map[uint32]int32
	D10 map[uint32]int64
	D11 map[uint32]float32
	D12 map[uint32]float64
	E1  map[uint64]bool
	E2  map[uint64]string
	E3  map[uint64]uint8
	E4  map[uint64]uint16
	E5  map[uint64]uint32
	E6  map[uint64]uint64
	E7  map[uint64]int8
	E8  map[uint64]int16
	E9  map[uint64]int32
	E10 map[uint64]int64
	E11 map[uint64]float32
	E12 map[uint64]float64
	F1  map[int8]bool
	F2  map[int8]string
	F3  map[int8]uint8
	F4  map[int8]uint16
	F5  map[int8]uint32
	F6  map[int8]uint64
	F7  map[int8]int8
	F8  map[int8]int16
	F9  map[int8]int32
	F10 map[int8]int64
	F11 map[int8]float32
	F12 map[int8]float64
	G1  map[int16]bool
	G2  map[int16]string
	G3  map[int16]uint8
	G4  map[int16]uint16
	G5  map[int16]uint32
	G6  map[int16]uint64
	G7  map[int16]int8
	G8  map[int16]int16
	G9  map[int16]int32
	G10 map[int16]int64
	G11 map[int16]float32
	G12 map[int16]float64
	H1  map[int32]bool
	H2  map[int32]string
	H3  map[int32]uint8
	H4  map[int32]uint16
	H5  map[int32]uint32
	H6  map[int32]uint64
	H7  map[int32]int8
	H8  map[int32]int16
	H9  map[int32]int32
	H10 map[int32]int64
	H11 map[int32]float32
	H12 map[int32]float64
	I1  map[int64]bool
	I2  map[int64]string
	I3  map[int64]uint8
	I4  map[int64]uint16
	I5  map[int64]uint32
	I6  map[int64]uint64
	I7  map[int64]int8
	I8  map[int64]int16
	I9  map[int64]int32
	I10 map[int64]int64
	I11 map[int64]float32
	I12 map[int64]float64
	J1  map[float32]bool
	J2  map[float32]string
	J3  map[float32]uint8
	J4  map[float32]uint16
	J5  map[float32]uint32
	J6  map[float32]uint64
	J7  map[float32]int8
	J8  map[float32]int16
	J9  map[float32]int32
	J10 map[float32]int64
	J11 map[float32]float32
	J12 map[float32]float64
	K1  map[float64]bool
	K2  map[float64]string
	K3  map[float64]uint8
	K4  map[float64]uint16
	K5  map[float64]uint32
	K6  map[float64]uint64
	K7  map[float64]int8
	K8  map[float64]int16
	K9  map[float64]int32
	K10 map[float64]int64
	K11 map[float64]float32
	K12 map[float64]float64
}

func (u *MapTest) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	err = writer.EncodeMapLen(len(u.A1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A1 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A2 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A3 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A4 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A5 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A6 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A7 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A8 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A9 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A10 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A11 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.A12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.A12 {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B1 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B2 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B3 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B4 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B5 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B6 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B7 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B8 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B9 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B10 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B11 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.B12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.B12 {
		err = writer.EncodeUint8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C1 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C2 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C3 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C4 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C5 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C6 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C7 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C8 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C9 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C10 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C11 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.C12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.C12 {
		err = writer.EncodeUint16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D1 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D2 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D3 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D4 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D5 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D6 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D7 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D8 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D9 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D10 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D11 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.D12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.D12 {
		err = writer.EncodeUint32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E1 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E2 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E3 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E4 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E5 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E6 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E7 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E8 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E9 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E10 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E11 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.E12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.E12 {
		err = writer.EncodeUint64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F1 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F2 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F3 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F4 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F5 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F6 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F7 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F8 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F9 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F10 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F11 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.F12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.F12 {
		err = writer.EncodeInt8(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G1 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G2 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G3 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G4 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G5 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G6 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G7 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G8 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G9 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G10 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G11 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.G12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.G12 {
		err = writer.EncodeInt16(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H1 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H2 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H3 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H4 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H5 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H6 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H7 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H8 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H9 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H10 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H11 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.H12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.H12 {
		err = writer.EncodeInt32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I1 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I2 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I3 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I4 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I5 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I6 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I7 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I8 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I9 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I10 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I11 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.I12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.I12 {
		err = writer.EncodeInt64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J1 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J2 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J3 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J4 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J5 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J6 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J7 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J8 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J9 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J10 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J11 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.J12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.J12 {
		err = writer.EncodeFloat32(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K1))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K1 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeBool(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K2))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K2 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K3))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K3 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K4))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K4 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K5))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K5 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K6))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K6 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K7))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K7 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt8(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K8))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K8 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt16(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K9))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K9 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K10))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K10 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeInt64(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K11))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K11 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat32(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.K12))
	if err != nil {
		return nil, err
	}

	for k, v := range u.K12 {
		err = writer.EncodeFloat64(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeFloat64(v)
		if err != nil {
			return nil, err
		}
	}

	return buf.Bytes(), nil
}

func (u *MapTest) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}

	return buf
}

func (u *MapTest) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := msgpack.NewDecoder(reader)
	var err error

	A1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A1 = make(map[string]bool)
	for i := 0; i < A1Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.A1[k] = v
	}

	A2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A2 = make(map[string]string)
	for i := 0; i < A2Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.A2[k] = v
	}

	A3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A3 = make(map[string]uint8)
	for i := 0; i < A3Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.A3[k] = v
	}

	A4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A4 = make(map[string]uint16)
	for i := 0; i < A4Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.A4[k] = v
	}

	A5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A5 = make(map[string]uint32)
	for i := 0; i < A5Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.A5[k] = v
	}

	A6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A6 = make(map[string]uint64)
	for i := 0; i < A6Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.A6[k] = v
	}

	A7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A7 = make(map[string]int8)
	for i := 0; i < A7Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.A7[k] = v
	}

	A8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A8 = make(map[string]int16)
	for i := 0; i < A8Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.A8[k] = v
	}

	A9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A9 = make(map[string]int32)
	for i := 0; i < A9Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.A9[k] = v
	}

	A10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A10 = make(map[string]int64)
	for i := 0; i < A10Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.A10[k] = v
	}

	A11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A11 = make(map[string]float32)
	for i := 0; i < A11Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.A11[k] = v
	}

	A12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.A12 = make(map[string]float64)
	for i := 0; i < A12Len; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.A12[k] = v
	}

	B1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B1 = make(map[uint8]bool)
	for i := 0; i < B1Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.B1[k] = v
	}

	B2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B2 = make(map[uint8]string)
	for i := 0; i < B2Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.B2[k] = v
	}

	B3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B3 = make(map[uint8]uint8)
	for i := 0; i < B3Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.B3[k] = v
	}

	B4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B4 = make(map[uint8]uint16)
	for i := 0; i < B4Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.B4[k] = v
	}

	B5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B5 = make(map[uint8]uint32)
	for i := 0; i < B5Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.B5[k] = v
	}

	B6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B6 = make(map[uint8]uint64)
	for i := 0; i < B6Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.B6[k] = v
	}

	B7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B7 = make(map[uint8]int8)
	for i := 0; i < B7Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.B7[k] = v
	}

	B8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B8 = make(map[uint8]int16)
	for i := 0; i < B8Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.B8[k] = v
	}

	B9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B9 = make(map[uint8]int32)
	for i := 0; i < B9Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.B9[k] = v
	}

	B10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B10 = make(map[uint8]int64)
	for i := 0; i < B10Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.B10[k] = v
	}

	B11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B11 = make(map[uint8]float32)
	for i := 0; i < B11Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.B11[k] = v
	}

	B12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.B12 = make(map[uint8]float64)
	for i := 0; i < B12Len; i++ {
		k, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.B12[k] = v
	}

	C1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C1 = make(map[uint16]bool)
	for i := 0; i < C1Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.C1[k] = v
	}

	C2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C2 = make(map[uint16]string)
	for i := 0; i < C2Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.C2[k] = v
	}

	C3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C3 = make(map[uint16]uint8)
	for i := 0; i < C3Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.C3[k] = v
	}

	C4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C4 = make(map[uint16]uint16)
	for i := 0; i < C4Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.C4[k] = v
	}

	C5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C5 = make(map[uint16]uint32)
	for i := 0; i < C5Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.C5[k] = v
	}

	C6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C6 = make(map[uint16]uint64)
	for i := 0; i < C6Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.C6[k] = v
	}

	C7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C7 = make(map[uint16]int8)
	for i := 0; i < C7Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.C7[k] = v
	}

	C8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C8 = make(map[uint16]int16)
	for i := 0; i < C8Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.C8[k] = v
	}

	C9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C9 = make(map[uint16]int32)
	for i := 0; i < C9Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.C9[k] = v
	}

	C10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C10 = make(map[uint16]int64)
	for i := 0; i < C10Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.C10[k] = v
	}

	C11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C11 = make(map[uint16]float32)
	for i := 0; i < C11Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.C11[k] = v
	}

	C12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.C12 = make(map[uint16]float64)
	for i := 0; i < C12Len; i++ {
		k, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.C12[k] = v
	}

	D1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D1 = make(map[uint32]bool)
	for i := 0; i < D1Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.D1[k] = v
	}

	D2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D2 = make(map[uint32]string)
	for i := 0; i < D2Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.D2[k] = v
	}

	D3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D3 = make(map[uint32]uint8)
	for i := 0; i < D3Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.D3[k] = v
	}

	D4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D4 = make(map[uint32]uint16)
	for i := 0; i < D4Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.D4[k] = v
	}

	D5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D5 = make(map[uint32]uint32)
	for i := 0; i < D5Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.D5[k] = v
	}

	D6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D6 = make(map[uint32]uint64)
	for i := 0; i < D6Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.D6[k] = v
	}

	D7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D7 = make(map[uint32]int8)
	for i := 0; i < D7Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.D7[k] = v
	}

	D8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D8 = make(map[uint32]int16)
	for i := 0; i < D8Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.D8[k] = v
	}

	D9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D9 = make(map[uint32]int32)
	for i := 0; i < D9Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.D9[k] = v
	}

	D10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D10 = make(map[uint32]int64)
	for i := 0; i < D10Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.D10[k] = v
	}

	D11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D11 = make(map[uint32]float32)
	for i := 0; i < D11Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.D11[k] = v
	}

	D12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.D12 = make(map[uint32]float64)
	for i := 0; i < D12Len; i++ {
		k, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.D12[k] = v
	}

	E1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E1 = make(map[uint64]bool)
	for i := 0; i < E1Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.E1[k] = v
	}

	E2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E2 = make(map[uint64]string)
	for i := 0; i < E2Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.E2[k] = v
	}

	E3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E3 = make(map[uint64]uint8)
	for i := 0; i < E3Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.E3[k] = v
	}

	E4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E4 = make(map[uint64]uint16)
	for i := 0; i < E4Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.E4[k] = v
	}

	E5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E5 = make(map[uint64]uint32)
	for i := 0; i < E5Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.E5[k] = v
	}

	E6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E6 = make(map[uint64]uint64)
	for i := 0; i < E6Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.E6[k] = v
	}

	E7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E7 = make(map[uint64]int8)
	for i := 0; i < E7Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.E7[k] = v
	}

	E8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E8 = make(map[uint64]int16)
	for i := 0; i < E8Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.E8[k] = v
	}

	E9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E9 = make(map[uint64]int32)
	for i := 0; i < E9Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.E9[k] = v
	}

	E10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E10 = make(map[uint64]int64)
	for i := 0; i < E10Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.E10[k] = v
	}

	E11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E11 = make(map[uint64]float32)
	for i := 0; i < E11Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.E11[k] = v
	}

	E12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.E12 = make(map[uint64]float64)
	for i := 0; i < E12Len; i++ {
		k, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.E12[k] = v
	}

	F1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F1 = make(map[int8]bool)
	for i := 0; i < F1Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.F1[k] = v
	}

	F2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F2 = make(map[int8]string)
	for i := 0; i < F2Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.F2[k] = v
	}

	F3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F3 = make(map[int8]uint8)
	for i := 0; i < F3Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.F3[k] = v
	}

	F4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F4 = make(map[int8]uint16)
	for i := 0; i < F4Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.F4[k] = v
	}

	F5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F5 = make(map[int8]uint32)
	for i := 0; i < F5Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.F5[k] = v
	}

	F6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F6 = make(map[int8]uint64)
	for i := 0; i < F6Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.F6[k] = v
	}

	F7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F7 = make(map[int8]int8)
	for i := 0; i < F7Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.F7[k] = v
	}

	F8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F8 = make(map[int8]int16)
	for i := 0; i < F8Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.F8[k] = v
	}

	F9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F9 = make(map[int8]int32)
	for i := 0; i < F9Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.F9[k] = v
	}

	F10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F10 = make(map[int8]int64)
	for i := 0; i < F10Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.F10[k] = v
	}

	F11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F11 = make(map[int8]float32)
	for i := 0; i < F11Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.F11[k] = v
	}

	F12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.F12 = make(map[int8]float64)
	for i := 0; i < F12Len; i++ {
		k, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.F12[k] = v
	}

	G1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G1 = make(map[int16]bool)
	for i := 0; i < G1Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.G1[k] = v
	}

	G2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G2 = make(map[int16]string)
	for i := 0; i < G2Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.G2[k] = v
	}

	G3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G3 = make(map[int16]uint8)
	for i := 0; i < G3Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.G3[k] = v
	}

	G4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G4 = make(map[int16]uint16)
	for i := 0; i < G4Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.G4[k] = v
	}

	G5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G5 = make(map[int16]uint32)
	for i := 0; i < G5Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.G5[k] = v
	}

	G6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G6 = make(map[int16]uint64)
	for i := 0; i < G6Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.G6[k] = v
	}

	G7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G7 = make(map[int16]int8)
	for i := 0; i < G7Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.G7[k] = v
	}

	G8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G8 = make(map[int16]int16)
	for i := 0; i < G8Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.G8[k] = v
	}

	G9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G9 = make(map[int16]int32)
	for i := 0; i < G9Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.G9[k] = v
	}

	G10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G10 = make(map[int16]int64)
	for i := 0; i < G10Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.G10[k] = v
	}

	G11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G11 = make(map[int16]float32)
	for i := 0; i < G11Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.G11[k] = v
	}

	G12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.G12 = make(map[int16]float64)
	for i := 0; i < G12Len; i++ {
		k, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.G12[k] = v
	}

	H1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H1 = make(map[int32]bool)
	for i := 0; i < H1Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.H1[k] = v
	}

	H2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H2 = make(map[int32]string)
	for i := 0; i < H2Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.H2[k] = v
	}

	H3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H3 = make(map[int32]uint8)
	for i := 0; i < H3Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.H3[k] = v
	}

	H4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H4 = make(map[int32]uint16)
	for i := 0; i < H4Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.H4[k] = v
	}

	H5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H5 = make(map[int32]uint32)
	for i := 0; i < H5Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.H5[k] = v
	}

	H6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H6 = make(map[int32]uint64)
	for i := 0; i < H6Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.H6[k] = v
	}

	H7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H7 = make(map[int32]int8)
	for i := 0; i < H7Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.H7[k] = v
	}

	H8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H8 = make(map[int32]int16)
	for i := 0; i < H8Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.H8[k] = v
	}

	H9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H9 = make(map[int32]int32)
	for i := 0; i < H9Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.H9[k] = v
	}

	H10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H10 = make(map[int32]int64)
	for i := 0; i < H10Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.H10[k] = v
	}

	H11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H11 = make(map[int32]float32)
	for i := 0; i < H11Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.H11[k] = v
	}

	H12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.H12 = make(map[int32]float64)
	for i := 0; i < H12Len; i++ {
		k, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.H12[k] = v
	}

	I1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I1 = make(map[int64]bool)
	for i := 0; i < I1Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.I1[k] = v
	}

	I2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I2 = make(map[int64]string)
	for i := 0; i < I2Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.I2[k] = v
	}

	I3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I3 = make(map[int64]uint8)
	for i := 0; i < I3Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.I3[k] = v
	}

	I4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I4 = make(map[int64]uint16)
	for i := 0; i < I4Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.I4[k] = v
	}

	I5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I5 = make(map[int64]uint32)
	for i := 0; i < I5Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.I5[k] = v
	}

	I6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I6 = make(map[int64]uint64)
	for i := 0; i < I6Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.I6[k] = v
	}

	I7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I7 = make(map[int64]int8)
	for i := 0; i < I7Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.I7[k] = v
	}

	I8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I8 = make(map[int64]int16)
	for i := 0; i < I8Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.I8[k] = v
	}

	I9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I9 = make(map[int64]int32)
	for i := 0; i < I9Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.I9[k] = v
	}

	I10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I10 = make(map[int64]int64)
	for i := 0; i < I10Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.I10[k] = v
	}

	I11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I11 = make(map[int64]float32)
	for i := 0; i < I11Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.I11[k] = v
	}

	I12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.I12 = make(map[int64]float64)
	for i := 0; i < I12Len; i++ {
		k, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.I12[k] = v
	}

	J1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J1 = make(map[float32]bool)
	for i := 0; i < J1Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.J1[k] = v
	}

	J2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J2 = make(map[float32]string)
	for i := 0; i < J2Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.J2[k] = v
	}

	J3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J3 = make(map[float32]uint8)
	for i := 0; i < J3Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.J3[k] = v
	}

	J4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J4 = make(map[float32]uint16)
	for i := 0; i < J4Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.J4[k] = v
	}

	J5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J5 = make(map[float32]uint32)
	for i := 0; i < J5Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.J5[k] = v
	}

	J6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J6 = make(map[float32]uint64)
	for i := 0; i < J6Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.J6[k] = v
	}

	J7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J7 = make(map[float32]int8)
	for i := 0; i < J7Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.J7[k] = v
	}

	J8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J8 = make(map[float32]int16)
	for i := 0; i < J8Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.J8[k] = v
	}

	J9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J9 = make(map[float32]int32)
	for i := 0; i < J9Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.J9[k] = v
	}

	J10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J10 = make(map[float32]int64)
	for i := 0; i < J10Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.J10[k] = v
	}

	J11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J11 = make(map[float32]float32)
	for i := 0; i < J11Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.J11[k] = v
	}

	J12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.J12 = make(map[float32]float64)
	for i := 0; i < J12Len; i++ {
		k, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.J12[k] = v
	}

	K1Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K1 = make(map[float64]bool)
	for i := 0; i < K1Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeBool()
		if err != nil {
			return err
		}

		u.K1[k] = v
	}

	K2Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K2 = make(map[float64]string)
	for i := 0; i < K2Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		u.K2[k] = v
	}

	K3Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K3 = make(map[float64]uint8)
	for i := 0; i < K3Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint8()
		if err != nil {
			return err
		}

		u.K3[k] = v
	}

	K4Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K4 = make(map[float64]uint16)
	for i := 0; i < K4Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint16()
		if err != nil {
			return err
		}

		u.K4[k] = v
	}

	K5Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K5 = make(map[float64]uint32)
	for i := 0; i < K5Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint32()
		if err != nil {
			return err
		}

		u.K5[k] = v
	}

	K6Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K6 = make(map[float64]uint64)
	for i := 0; i < K6Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.K6[k] = v
	}

	K7Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K7 = make(map[float64]int8)
	for i := 0; i < K7Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.K7[k] = v
	}

	K8Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K8 = make(map[float64]int16)
	for i := 0; i < K8Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt16()
		if err != nil {
			return err
		}

		u.K8[k] = v
	}

	K9Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K9 = make(map[float64]int32)
	for i := 0; i < K9Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt32()
		if err != nil {
			return err
		}

		u.K9[k] = v
	}

	K10Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K10 = make(map[float64]int64)
	for i := 0; i < K10Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeInt64()
		if err != nil {
			return err
		}

		u.K10[k] = v
	}

	K11Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K11 = make(map[float64]float32)
	for i := 0; i < K11Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat32()
		if err != nil {
			return err
		}

		u.K11[k] = v
	}

	K12Len, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.K12 = make(map[float64]float64)
	for i := 0; i < K12Len; i++ {
		k, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeFloat64()
		if err != nil {
			return err
		}

		u.K12[k] = v
	}

	return nil
}

func (u *MapTest) MergeUsing(other *MapTest) {
	u.A1 = other.A1
	u.A2 = other.A2
	u.A3 = other.A3
	u.A4 = other.A4
	u.A5 = other.A5
	u.A6 = other.A6
	u.A7 = other.A7
	u.A8 = other.A8
	u.A9 = other.A9
	u.A10 = other.A10
	u.A11 = other.A11
	u.A12 = other.A12
	u.B1 = other.B1
	u.B2 = other.B2
	u.B3 = other.B3
	u.B4 = other.B4
	u.B5 = other.B5
	u.B6 = other.B6
	u.B7 = other.B7
	u.B8 = other.B8
	u.B9 = other.B9
	u.B10 = other.B10
	u.B11 = other.B11
	u.B12 = other.B12
	u.C1 = other.C1
	u.C2 = other.C2
	u.C3 = other.C3
	u.C4 = other.C4
	u.C5 = other.C5
	u.C6 = other.C6
	u.C7 = other.C7
	u.C8 = other.C8
	u.C9 = other.C9
	u.C10 = other.C10
	u.C11 = other.C11
	u.C12 = other.C12
	u.D1 = other.D1
	u.D2 = other.D2
	u.D3 = other.D3
	u.D4 = other.D4
	u.D5 = other.D5
	u.D6 = other.D6
	u.D7 = other.D7
	u.D8 = other.D8
	u.D9 = other.D9
	u.D10 = other.D10
	u.D11 = other.D11
	u.D12 = other.D12
	u.E1 = other.E1
	u.E2 = other.E2
	u.E3 = other.E3
	u.E4 = other.E4
	u.E5 = other.E5
	u.E6 = other.E6
	u.E7 = other.E7
	u.E8 = other.E8
	u.E9 = other.E9
	u.E10 = other.E10
	u.E11 = other.E11
	u.E12 = other.E12
	u.F1 = other.F1
	u.F2 = other.F2
	u.F3 = other.F3
	u.F4 = other.F4
	u.F5 = other.F5
	u.F6 = other.F6
	u.F7 = other.F7
	u.F8 = other.F8
	u.F9 = other.F9
	u.F10 = other.F10
	u.F11 = other.F11
	u.F12 = other.F12
	u.G1 = other.G1
	u.G2 = other.G2
	u.G3 = other.G3
	u.G4 = other.G4
	u.G5 = other.G5
	u.G6 = other.G6
	u.G7 = other.G7
	u.G8 = other.G8
	u.G9 = other.G9
	u.G10 = other.G10
	u.G11 = other.G11
	u.G12 = other.G12
	u.H1 = other.H1
	u.H2 = other.H2
	u.H3 = other.H3
	u.H4 = other.H4
	u.H5 = other.H5
	u.H6 = other.H6
	u.H7 = other.H7
	u.H8 = other.H8
	u.H9 = other.H9
	u.H10 = other.H10
	u.H11 = other.H11
	u.H12 = other.H12
	u.I1 = other.I1
	u.I2 = other.I2
	u.I3 = other.I3
	u.I4 = other.I4
	u.I5 = other.I5
	u.I6 = other.I6
	u.I7 = other.I7
	u.I8 = other.I8
	u.I9 = other.I9
	u.I10 = other.I10
	u.I11 = other.I11
	u.I12 = other.I12
	u.J1 = other.J1
	u.J2 = other.J2
	u.J3 = other.J3
	u.J4 = other.J4
	u.J5 = other.J5
	u.J6 = other.J6
	u.J7 = other.J7
	u.J8 = other.J8
	u.J9 = other.J9
	u.J10 = other.J10
	u.J11 = other.J11
	u.J12 = other.J12
	u.K1 = other.K1
	u.K2 = other.K2
	u.K3 = other.K3
	u.K4 = other.K4
	u.K5 = other.K5
	u.K6 = other.K6
	u.K7 = other.K7
	u.K8 = other.K8
	u.K9 = other.K9
	u.K10 = other.K10
	u.K11 = other.K11
	u.K12 = other.K12
}
