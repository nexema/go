package internal

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

const runtimeImport = "github.com/messagepack-schema/go/runtime"

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
	"boolean":  "EncodeBool",
	"string":   "EncodeString",
	"uint8":    "EncodeUint8",
	"uint16":   "EncodeUint16",
	"uint32":   "EncodeUint32",
	"uint64":   "EncodeUint64",
	"int8":     "EncodeInt8",
	"int16":    "EncodeInt16",
	"int32":    "EncodeInt32",
	"int64":    "EncodeInt64",
	"float32":  "EncodeFloat32",
	"float64":  "EncodeFloat64",
	"binary":   "EncodeBytes",
	"nullable": "EncodeNullable",
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

func (b *Builder) WriteField(stmt *Statement, t SchemaFieldType, referencePkg string) {
	b.WritePrimitive(stmt, t.Primitive, t.Nullable, t.TypeArguments, t.ImportId, referencePkg)
}

func (b *Builder) WritePrimitive(stmt *Statement, t string, nullable bool, arguments []SchemaFieldType, importId string, referencePkg string) {
	switch t {
	case "string":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(String())
		} else {
			stmt.String()
		}

	case "int8":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Int8())
		} else {
			stmt.Int8()
		}

	case "int16":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Int16())
		} else {
			stmt.Int16()
		}

	case "int32":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Int32())
		} else {
			stmt.Int32()
		}

	case "int64":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Int64())
		} else {
			stmt.Int64()
		}

	case "uint8":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Uint8())
		} else {
			stmt.Uint8()
		}

	case "uint16":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Uint16())
		} else {
			stmt.Uint16()
		}

	case "uint32":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Uint32())
		} else {
			stmt.Uint32()
		}

	case "uint64":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Uint64())
		} else {
			stmt.Uint64()
		}

	case "boolean":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Bool())
		} else {
			stmt.Bool()
		}

	case "float32":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Float32())
		} else {
			stmt.Float32()
		}

	case "float64":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Float64())
		} else {
			stmt.Float64()
		}

	case "binary":
		if nullable {
			stmt.Qual(runtimeImport, "Nullable").Types(Index().Byte())
		} else {
			stmt.Index().Byte()
		}

	case "list":
		arrType := arguments[0].Primitive
		if arguments[0].Nullable {
			b.WritePrimitive(stmt.Index(), arrType, true, nil, importId, referencePkg)
		} else {
			b.WritePrimitive(stmt.Index(), arrType, false, nil, importId, referencePkg)
		}

	case "map":
		keyType := arguments[0].Primitive
		valueType := arguments[1].Primitive
		if arguments[1].Nullable {
			b.WritePrimitive(stmt.Map(primitiveMapper[keyType]), valueType, true, nil, importId, referencePkg)
		} else {
			b.WritePrimitive(stmt.Map(primitiveMapper[keyType]), valueType, false, nil, importId, referencePkg)
		}

	case "custom":
		importType, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found in registry", importId))
		}

		samePkg := importType.ImportPath == referencePkg
		if nullable {

		} else {
			if samePkg {
				stmt.Id(importType.Definition.Name)
			} else {
				stmt.Qual(importType.ImportPath, importType.Definition.Name)
			}
		}
	}
}

func (b *Builder) WriteEncodeField(field *TypeFieldDefinition) []*Statement {
	stmts := []*Statement{}

	if field.Type.Primitive == "list" {
		typeArgumentPrimitive := field.Type.TypeArguments[0].Primitive

		stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeArrayLen").Params(Id("len").Params(Id("u").Dot(field.GoName))))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		if field.Type.TypeArguments[0].Nullable {
			stmts = append(stmts, For(List(Id("_"), Id("v")).Op(":=").Range().Id("u").Dot(field.GoName)).Block(
				GetEncodeNullable(typeArgumentPrimitive, Id("v")),
				If(Id("err").Op("!=").Nil()).Block(
					Return(List(Nil(), Id("err"))),
				),
			))
		} else {
			stmts = append(stmts, For(List(Id("_"), Id("v")).Op(":=").Range().Id("u").Dot(field.GoName)).Block(
				Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(typeArgumentPrimitive, Id("v"))),
				If(Id("err").Op("!=").Nil()).Block(
					Return(List(Nil(), Id("err"))),
				),
			))
		}
	} else if field.Type.Primitive == "map" {
		keyArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		valueArgumentPrimitive := field.Type.TypeArguments[1].Primitive

		stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeMapLen").Params(Id("len").Params(Id("u").Dot(field.GoName))))
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

		if field.Type.TypeArguments[1].Nullable {
			stmts = append(stmts, For(List(Id("k"), Id("v")).Op(":=").Range().Id("u").Dot(field.GoName)).Block(
				Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(keyArgumentPrimitive, Id("k"))),
				If(Id("err").Op("!=").Nil()).Block(
					Return(List(Nil(), Id("err"))),
				),

				GetEncodeNullable(valueArgumentPrimitive, Id("v")),
				If(Id("err").Op("!=").Nil()).Block(
					Return(List(Nil(), Id("err"))),
				),
			))
		} else {
			stmts = append(stmts, For(List(Id("k"), Id("v")).Op(":=").Range().Id("u").Dot(field.GoName)).Block(
				Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(keyArgumentPrimitive, Id("k"))),
				If(Id("err").Op("!=").Nil()).Block(
					Return(List(Nil(), Id("err"))),
				),

				Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(valueArgumentPrimitive, Id("v"))),
				If(Id("err").Op("!=").Nil()).Block(
					Return(List(Nil(), Id("err"))),
				),
			))
		}
	} else if field.Type.Primitive == "custom" {
		importId := field.Type.ImportId
		td, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found", importId))
		}

		if td.Definition.Modifier == "enum" {
			stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeUint8").Params(Id("u").Dot(field.GoName).Dot("Index").Call()))
			stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			))
		} else {
			varName := fmt.Sprintf("%sBinary", field.GoName)
			stmts = append(stmts, List(Id(varName), Id("err")).Op(":=").Id("u").Dot(field.GoName).Dot("Serialize").Call())
			stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			))

			stmts = append(stmts, Id("err").Op("=").Id("writer").Dot("EncodeBytes").Params(Id(varName)))
			stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
				Return(List(Nil(), Id("err"))),
			))
		}

	} else {
		if field.Type.Nullable {
			stmts = append(stmts, If(
				Id("u").Dot(field.GoName).Dot("HasValue").Call(),
			).Block(
				Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(field.Type.Primitive, Id("u").Dot(field.GoName).Dot("GetValue").Call()))).Else().Block(
				Id("err").Op("=").Custom(Options{}, GetEncodeNil()),
			))
		} else {
			stmts = append(stmts, Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(field.Type.Primitive, Id("u").Dot(field.GoName))))
		}

		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(List(Nil(), Id("err"))),
		))

	}

	return stmts
}

func GetEncodeNullable(fieldPrimitive string, field *Statement) *Statement {
	return If(
		field.Clone().Dot("HasValue").Call(),
	).Block(
		Id("err").Op("=").Custom(Options{}, GetEncodeTypeStatement(fieldPrimitive, field.Clone().Dot("GetValue").Call()))).Else().Block(
		Id("err").Op("=").Custom(Options{}, GetEncodeNil()),
	)
}

func GetEncodeTypeStatement(primitive string, fieldName *Statement) *Statement {
	encoder, ok := encoderMapper[primitive]
	if !ok {
		panic(fmt.Sprintf("unknown encoder for %s", primitive))
	}

	return Id("writer").Dot(encoder).Params(fieldName)
}

func GetEncodeNil() *Statement {
	return Id("writer").Dot("EncodeNil").Call()
}

func (b *Builder) WriteDecodeField(field *TypeFieldDefinition, pkg string) []*Statement {
	stmts := []*Statement{}

	if field.Type.Primitive == "list" {
		typeArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		arrName := fmt.Sprintf("%sLen", field.GoName)

		stmts = append(stmts, List(Id(arrName), Id("err")).Op(":=").Id("decoder").Dot("DecodeArrayLen").Call())
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return().Id("err"),
		))

		stmts = append(stmts, Id("u").Dot(field.GoName).Op("=").Make(Index().Custom(Options{}, primitiveMapper[typeArgumentPrimitive]), Id(arrName)))
		stmts = append(stmts, For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<").Id(arrName),
			Id("i").Op("++"),
		).Block(
			List(Id("u").Dot(field.GoName).Index(Id("i")), Id("err")).Op("=").Id("decoder").Dot(decoderMapper[typeArgumentPrimitive]).Call(),
			If(Id("err").Op("!=").Nil()).Block(
				Return().Id("err"),
			),
		))
	} else if field.Type.Primitive == "map" {
		keyArgumentPrimitive := field.Type.TypeArguments[0].Primitive
		valueArgumentPrimitive := field.Type.TypeArguments[1].Primitive

		mapName := fmt.Sprintf("%sLen", field.GoName)

		stmts = append(stmts, List(Id(mapName), Id("err")).Op(":=").Id("decoder").Dot("DecodeMapLen").Call())
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return().Id("err"),
		))

		stmts = append(stmts, Id("u").Dot(field.GoName).Op("=").Make(Map(primitiveMapper[keyArgumentPrimitive]).Custom(Options{}, primitiveMapper[valueArgumentPrimitive])))
		stmts = append(stmts, For(
			Id("i").Op(":=").Lit(0),
			Id("i").Op("<").Id(mapName),
			Id("i").Op("++"),
		).Block(
			List(Id("k"), Id("err")).Op(":=").Id("decoder").Dot(decoderMapper[keyArgumentPrimitive]).Call(),
			If(Id("err").Op("!=").Nil()).Block(
				Return().Id("err"),
			),

			List(Id("v"), Id("err")).Op(":=").Id("decoder").Dot(decoderMapper[valueArgumentPrimitive]).Call(),
			If(Id("err").Op("!=").Nil()).Block(
				Return().Id("err"),
			),

			Id("u").Dot(field.GoName).Index(Id("k")).Op("=").Id("v"),
		))
	} else if field.Type.Primitive == "custom" {
		importId := field.Type.ImportId
		td, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found", importId))
		}

		samePackage := td.ImportPath == pkg

		if td.Definition.Modifier == "enum" {
			varName := fmt.Sprintf("%sIdx", field.GoName)
			stmts = append(stmts, List(Id(varName), Err()).Op(":=").Id("decoder").Dot("DecodeUint8").Call())
			stmts = append(stmts, If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			))

			op := Id("u").Dot(field.GoName).Op("=")
			if samePackage {
				op.Id(fmt.Sprintf("%sPicker", td.Definition.Name))
			} else {
				op.Qual(td.ImportPath, fmt.Sprintf("%sPicker", td.Definition.Name))
			}

			stmts = append(stmts, op.Dot("ByIndex").Call(Id(varName)))
		} else {
			varName := fmt.Sprintf("%sBinary", field.GoName)

			stmts = append(stmts, List(Id(varName), Err()).Op(":=").Id("decoder").Dot("DecodeBytes").Call())
			stmts = append(stmts, If(Err().Op("!=").Nil()).Block(
				Return(Err()),
			))

			op := Id("u").Dot(field.GoName).Op("=")
			if samePackage {
				op.Id(td.Definition.Name)
			} else {
				op.Qual(td.ImportPath, td.Definition.Name)
			}

			stmts = append(stmts, op.Values())
			stmts = append(stmts, Id("u").Dot(field.GoName).Dot("MergeFrom").Call(Id(varName)))
		}

	} else {
		stmts = append(stmts, List(Id("u").Dot(field.GoName), Id("err")).Op("=").Custom(Options{}, GetDecodeTypeStatement(field.Type.Primitive)).Call())
		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
			Return(Id("err")),
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
