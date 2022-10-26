package internal

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

var primitiveMapper map[string]Code = map[string]Code{
	"boolean": Bool(),
	"string":  String(),
	"uint8":   Uint8(),
	"uint16":  Uint16(),
	"uint32":  Uint32(),
	"uint64":  Uint64(),
	"int8":    Int8(),
	"int16":   Int16(),
	"int32":   Int32(),
	"int64":   Int64(),
	"float32": Float32(),
	"float64": Float64(),
	"binary":  Index().Byte(),
}

var encoderMapper map[string]string = map[string]string{
	"boolean": "EncodeBool",
	"string":  "EncodeString",
	"uint8":   "EncodeUint8",
	"uint16":  "EncodeUint16",
	"uint32":  "EncodeUint32",
	"uint64":  "EncodeUint64",
	"int8":    "EncodeInt8",
	"int16":   "EncodeInt16",
	"int32":   "EncodeInt32",
	"int64":   "EncodeInt64",
	"float32": "EncodeFloat32",
	"float64": "EncodeFloat64",
	"binary":  "EncodeBytes",
}

var decoderMapper map[string]string = map[string]string{
	"boolean": "DecodeBool",
	"string":  "DecodeString",
	"uint8":   "DecodeUint8",
	"uint16":  "DecodeUint16",
	"uint32":  "DecodeUint32",
	"uint64":  "DecodeUint64",
	"int8":    "DecodeInt8",
	"int16":   "DecodeInt16",
	"int32":   "DecodeInt32",
	"int64":   "DecodeInt64",
	"float32": "DecodeFloat32",
	"float64": "DecodeFloat64",
	"binary":  "DecodeBytes",
}

func WriteField(stmt *Statement, t SchemaFieldType) {
	WritePrimitive(stmt, t.Primitive, t.TypeArguments)
}

func WritePrimitive(stmt *Statement, t string, arguments []SchemaFieldType) {
	switch t {
	case "string":
		stmt.String()

	case "int8":
		stmt.Int8()

	case "int16":
		stmt.Int16()

	case "int32":
		stmt.Int32()

	case "int64":
		stmt.Int64()

	case "uint8":
		stmt.Uint8()

	case "uint16":
		stmt.Uint16()

	case "uint32":
		stmt.Uint32()

	case "uint64":
		stmt.Uint64()

	case "boolean":
		stmt.Bool()

	case "float32":
		stmt.Float32()

	case "float64":
		stmt.Float64()

	case "binary":
		stmt.Index().Byte()

	case "list":
		arrType := arguments[0].Primitive
		WritePrimitive(stmt.Index(), arrType, nil)

	case "map":
		keyType := arguments[0].Primitive
		valueType := arguments[1].Primitive
		WritePrimitive(stmt.Map(primitiveMapper[keyType]), valueType, nil)

	case "custom":

	}
}

func WriteEncodeField(field TypeFieldDefinition) []*Statement {
	stmts := []*Statement{}

	if field.Type.Primitive == "list" {
		typeArgumentPrimitive := field.Type.TypeArguments[0].Primitive

		stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeArrayLen").Params(Id("len").Params(Id("u").Dot(field.Name))))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, For(List(Id("_"), Id("v")).Op(":=").Range().Id("u").Dot(field.Name)).Block(
			Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(typeArgumentPrimitive, Id("v"))),
			If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			),
		))
	} else if field.Type.Primitive == "map" {
		keyArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		valueArgumentPrimitive := field.Type.TypeArguments[1].Primitive

		stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeMapLen").Params(Id("len").Params(Id("u").Dot(field.Name))))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, For(List(Id("k"), Id("v")).Op(":=").Range().Id("u").Dot(field.Name)).Block(
			Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(keyArgumentPrimitive, Id("k"))),
			If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			),

			Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(valueArgumentPrimitive, Id("v"))),
			If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			),
		))
	} else {
		stmts = append(stmts, Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(field.Type.Primitive, Id("u").Dot(field.Name))))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))
	}

	return stmts
}

func GetEncodeTypeStatement(primitive string, fieldName *Statement) *Statement {
	encoder, ok := encoderMapper[primitive]
	if !ok {
		panic(fmt.Sprintf("unknown encoder for %s", primitive))
	}

	return Id("writer").Dot(encoder).Params(fieldName)
}

func WriteDecodeField(field TypeFieldDefinition) []*Statement {
	stmts := []*Statement{}

	if field.Type.Primitive == "list" {
		typeArgumentPrimitive := field.Type.TypeArguments[0].Primitive

		stmts = append(stmts, List(Id(fmt.Sprintf("%sLen", field.Name)), Id("err")).Op(":=").Id("decoder").Dot("DecodeArrayLen"))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, Id("u").Dot(field.Name).Op("=").Make(Index().Custom(Options{}, primitiveMapper[typeArgumentPrimitive]), Id(fmt.Sprintf("%sLen", field.Name))))
		stmts = append(stmts, For(Id("i").Op(":=").Lit(0))) // TODO: continue loop
	} else if field.Type.Primitive == "map" {
		keyArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		valueArgumentPrimitive := field.Type.TypeArguments[1].Primitive

		stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeMapLen").Params(Id("len").Params(Id("u").Dot(field.Name))))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, For(List(Id("k"), Id("v")).Op(":=").Range().Id("u").Dot(field.Name)).Block(
			Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(keyArgumentPrimitive, Id("k"))),
			If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			),

			Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(valueArgumentPrimitive, Id("v"))),
			If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			),
		))
	} else {
		stmts = append(stmts, List(Id("u").Dot(field.Name), Id("err")).Op("=").Custom(Options{}, GetDecodeTypeStatement(field.Type.Primitive)))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))
	}

	return stmts
}

func GetDecodeTypeStatement(primitive string) *Statement {
	decoder, ok := decoderMapper[primitive]
	if !ok {
		panic(fmt.Sprintf("unknown decoder for %s", primitive))
	}

	return Id("decoder").Dot(decoder)
}
