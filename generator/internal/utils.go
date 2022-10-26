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

func (b *Builder) WriteEncodeField(field TypeFieldDefinition) []*Statement {
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
	} else if field.Type.Primitive == "custom" {
		importId := field.Type.ImportId
		_, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found", importId))
		}

		varName := fmt.Sprintf("%sBinary", field.Name)
		stmts = append(stmts, List(Id(varName), Id("err")).Op(":=").Id("u").Dot(field.Name).Dot("Serialize").Call())
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeBytes").Params(Id(varName)))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
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

func (b *Builder) WriteDecodeField(field TypeFieldDefinition) []*Statement {
	stmts := []*Statement{}

	if field.Type.Primitive == "list" {
		typeArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		arrName := fmt.Sprintf("%sLen", field.Name)

		stmts = append(stmts, List(Id(arrName), Id("err")).Op(":=").Id("decoder").Dot("DecodeArrayLen").Call())
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, Id("u").Dot(field.Name).Op("=").Make(Index().Custom(Options{}, primitiveMapper[typeArgumentPrimitive]), Id(arrName)))
		stmts = append(stmts, For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<").Id(arrName),
			Id("i").Op("++"),
		).Block(
			List(Id("u").Dot(field.Name).Index(Id("i")), Id("err")).Op("=").Id("decoder").Dot(decoderMapper[typeArgumentPrimitive]).Call(),
			If(Id("err").Op("!=").Nil()).Block(
				Return().Nil(),
			),
		))
	} else if field.Type.Primitive == "map" {
		keyArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		valueArgumentPrimitive := field.Type.TypeArguments[1].Primitive

		mapName := fmt.Sprintf("%sLen", field.Name)

		stmts = append(stmts, List(Id(mapName), Id("err")).Op(":=").Id("decoder").Dot("DecodeMapLen").Call())
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		stmts = append(stmts, Id("u").Dot(field.Name).Op("=").Make(Map(primitiveMapper[keyArgumentPrimitive]).Custom(Options{}, primitiveMapper[valueArgumentPrimitive])))
		stmts = append(stmts, For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<").Id(mapName),
			Id("i").Op("++"),
		).Block(
			List(Id("k"), Id("err")).Op(":=").Id("decoder").Dot(decoderMapper[keyArgumentPrimitive]).Call(),
			If(Id("err").Op("!=").Nil()).Block(
				Return().Nil(),
			),

			List(Id("v"), Id("err")).Op(":=").Id("decoder").Dot(decoderMapper[valueArgumentPrimitive]).Call(),
			If(Id("err").Op("!=").Nil()).Block(
				Return().Nil(),
			),

			Id("u").Dot(field.Name).Index(Id("k")).Op("=").Id("v"),
		))
	} else if field.Type.Primitive == "custom" {
		importId := field.Type.ImportId
		t, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found", importId))
		}

		_ = t
	} else {
		stmts = append(stmts, List(Id("u").Dot(field.Name), Id("err")).Op("=").Custom(Options{}, GetDecodeTypeStatement(field.Type.Primitive)).Call())
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
