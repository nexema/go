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

var decodeTypeMapper map[string]string = map[string]string{
	"string":  "DecodeString",
	"bool":    "DecodeBool",
	"uint8":   "DecodeUint8",
	"uint16":  "DecodeUint16",
	"uint32":  "DecodeUint32",
	"uint64":  "DecodeUvarint",
	"int8":    "DecodeInt8",
	"int16":   "DecodeInt16",
	"int32":   "DecodeInt32",
	"int64":   "DecodeVarint",
	"float32": "DecodeFloat32",
	"float64": "DecodeFloat64",
	"bytes":   "DecodeBinary",
}
