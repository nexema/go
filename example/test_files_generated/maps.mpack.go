package test_files

import "github.com/messagepack-schema/go/runtime"
import "github.com/messagepack-schema/go/runtime/msgpack"
import "bytes"

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
