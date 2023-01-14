// Nexemab is the binary serialization protocol of Nexema
//
// Specification
// ----------------------------
// boolean uses 1 byte being:
//
//	0x00 for false
//	0x01 for true
//
// ----------------------------
// null uses 1 byte being 0xc0
//
// ----------------------------
// strings up to uint64.MaxValue of length, encoding:
//
//	len of string as uint64
//	string contents
//
// ----------------------------
// integers
//
//	int8-uint8 encodes as 1 byte
//	int16-uint16 encodes as 2 bytes
//	int32-uint32 encodes as 4 bytes
//	int64-uint64 encodes as a varint
//
// ----------------------------
// float
//
//	float32 encodes as 4 byte
//	float64 encodes as 8 bytes
//
// ----------------------------
package nexemab
