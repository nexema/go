package generator

var encodeTypeMapper map[string]string = map[string]string{
	"string":  "EncodeString",
	"bool":    "EncodeBool",
	"uint8":   "EncodeUint8",
	"uint16":  "EncodeUint16",
	"uint32":  "EncodeUint32",
	"uint64":  "EncodeUvarint",
	"int8":    "EncodeInt8",
	"int16":   "EncodeInt16",
	"int32":   "EncodeInt32",
	"int64":   "EncodeVarint",
	"float32": "EncodeFloat32",
	"float64": "EncodeFloat64",
	"bytes":   "EncodeBinary",
}
