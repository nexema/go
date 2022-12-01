package internal

import "fmt"

const runtimeImport = "github.com/messagepack-schema/go/runtime"

const listPrimitive = "list"
const mapPrimitive = "map"
const customPrimitive = "custom"
const enumModifier = "enum"
const structModifier = "struct"

var primitiveMapper map[string]string = map[string]string{
	"boolean": "bool",
	"string":  "string",
	"uint8":   "uint8",
	"uint16":  "uint16",
	"uint32":  "uint32",
	"uint64":  "uint64",
	"int8":    "int8",
	"int16":   "int16",
	"int32":   "int32",
	"int64":   "int64",
	"float32": "float32",
	"float64": "float64",
	"binary":  "[]byte",
}

func GetGoType(field SchemaFieldType) string {
	primitive := field.Primitive
	if primitive == listPrimitive {
		return fmt.Sprintf("[]%s", GetGoType(field.TypeArguments[0]))
	} else if primitive == mapPrimitive {
		return fmt.Sprintf("map[%s]%s", GetGoType(field.TypeArguments[0]), GetGoType(field.TypeArguments[1]))
	} else if primitive == customPrimitive {
		return "custom"
	} else {
		if field.Nullable {
			return fmt.Sprintf("runtime.Nullable[%s]", primitiveMapper[field.Primitive])
		}

		return primitiveMapper[field.Primitive]
	}
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

// var decoderMapper map[string]string = map[string]string{
// 	"boolean": "DecodeBool",
// 	"string":  "DecodeString",
// 	"uint8":   "DecodeUint8",
// 	"uint16":  "DecodeUint16",
// 	"uint32":  "DecodeUint32",
// 	"uint64":  "DecodeUint64",
// 	"int8":    "DecodeInt8",
// 	"int16":   "DecodeInt16",
// 	"int32":   "DecodeInt32",
// 	"int64":   "DecodeInt64",
// 	"float32": "DecodeFloat32",
// 	"float64": "DecodeFloat64",
// 	"binary":  "DecodeBytes",
// }

// func (b *Builder) WriteField(stmt *Statement, t SchemaFieldType, referencePkg string) {
// 	b.WritePrimitive(stmt, t.Primitive, t.Nullable, t.TypeArguments, t.ImportId, referencePkg)
// }

// func (b *Builder) WritePrimitive(stmt *Statement, t string, nullable bool, arguments []SchemaFieldType, importId string, referencePkg string) {
// 	switch t {
// 	case "string":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(String())
// 		} else {
// 			stmt.String()
// 		}

// 	case "int8":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Int8())
// 		} else {
// 			stmt.Int8()
// 		}

// 	case "int16":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Int16())
// 		} else {
// 			stmt.Int16()
// 		}

// 	case "int32":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Int32())
// 		} else {
// 			stmt.Int32()
// 		}

// 	case "int64":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Int64())
// 		} else {
// 			stmt.Int64()
// 		}

// 	case "uint8":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Uint8())
// 		} else {
// 			stmt.Uint8()
// 		}

// 	case "uint16":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Uint16())
// 		} else {
// 			stmt.Uint16()
// 		}

// 	case "uint32":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Uint32())
// 		} else {
// 			stmt.Uint32()
// 		}

// 	case "uint64":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Uint64())
// 		} else {
// 			stmt.Uint64()
// 		}

// 	case "boolean":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Bool())
// 		} else {
// 			stmt.Bool()
// 		}

// 	case "float32":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Float32())
// 		} else {
// 			stmt.Float32()
// 		}

// 	case "float64":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Float64())
// 		} else {
// 			stmt.Float64()
// 		}

// 	case "binary":
// 		if nullable {
// 			stmt.Qual(runtimeImport, "Nullable").Types(Index().Byte())
// 		} else {
// 			stmt.Index().Byte()
// 		}

// 	case "list":
// 		arrType := arguments[0].Primitive
// 		if arguments[0].Nullable {
// 			b.WritePrimitive(stmt.Index(), arrType, true, nil, importId, referencePkg)
// 		} else {
// 			b.WritePrimitive(stmt.Index(), arrType, false, nil, importId, referencePkg)
// 		}

// 	case "map":
// 		keyType := arguments[0].Primitive
// 		valueType := arguments[1].Primitive
// 		if arguments[1].Nullable {
// 			b.WritePrimitive(stmt.Map(primitiveMapper[keyType]), valueType, true, nil, importId, referencePkg)
// 		} else {
// 			b.WritePrimitive(stmt.Map(primitiveMapper[keyType]), valueType, false, nil, importId, referencePkg)
// 		}

// 	case "custom":
// 		importType, ok := (*b.typeRegistry)[importId]
// 		if !ok {
// 			panic(fmt.Sprintf("type with id %s not found in registry", importId))
// 		}

// 		samePkg := importType.ImportPath == referencePkg
// 		if nullable {

// 		} else {
// 			if samePkg {
// 				stmt.Id(importType.Definition.Name)
// 			} else {
// 				stmt.Qual(importType.ImportPath, importType.Definition.Name)
// 			}
// 		}
// 	}
// }

func (b *Builder) WriteEncodeField(field *TypeFieldDefinition) string {
	if field.Type.Primitive == listPrimitive {
		typeArgument := field.Type.TypeArguments[0]
		if typeArgument.Nullable {
			return fmt.Sprintf(
				`
				err = writer.EncodeArrayLen(len(u.%[1]s))
				if err != nil {
					return nil, err
				}

				for _, v := range u.%[1]s {
					if v.HasValue() {
						err = writer.%s(v.GetValue())
					} else {
						err = writer.EncodeNil()
					}
				}
			`, field.GoName, encoderMapper[typeArgument.Primitive])
		} else {
			return fmt.Sprintf(
				`
				err = writer.EncodeArrayLen(len(u.%[1]s))
				if err != nil {
					return nil, err
				}
	
				for _, v := range u.%[1]s {
					err = writer.%s(v)
					if err != nil {
						return nil, err
					}
				}
			`, field.GoName, encoderMapper[typeArgument.Primitive])
		}
	} else if field.Type.Primitive == mapPrimitive {
		keyArgument := field.Type.TypeArguments[0]
		valueArgument := field.Type.TypeArguments[1]

		if valueArgument.Nullable {
			return fmt.Sprintf(
				`
				err = writer.EncodeMapLen(len(u.%[1]s))
				if err != nil {
					return nil, err
				}

				for k, v := range u.%[1]s {
					err = writer.%s(k)
					if err != nil {
						return nil, err
					}

					if v.HasValue() {
						err = writer.%s(v.GetValue())
					} else {
						err = writer.EncodeNil()
					}

					if err != nil {
						return nil, err
					}
				}
			`, field.GoName, encoderMapper[keyArgument.Primitive], encoderMapper[valueArgument.Primitive])
		} else {
			return fmt.Sprintf(
				`
				err = writer.EncodeMapLen(len(u.%[1]s))
				if err != nil {
					return nil, err
				}

				for k, v := range u.%[1]s {
					err = writer.%s(k)
					if err != nil {
						return nil, err
					}

					err = writer.%s(v)
					if err != nil {
						return nil, err
					}
				}
			`, field.GoName, encoderMapper[keyArgument.Primitive], encoderMapper[valueArgument.Primitive])
		}
	} else if field.Type.Primitive == customPrimitive {
		importId := field.Type.ImportId
		td, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found", importId))
		}

		if td.Definition.Modifier == enumModifier {
			return fmt.Sprintf(
				`
				err = writer.EncodeUint8(u.%s.Index())
				if err != nil {
					return nil, err
				}
			`, field.GoName)
		} else {
			varName := fmt.Sprintf("%sBinary", field.GoName)
			if field.Type.Nullable {
				return fmt.Sprintf(
					`
					if u.HasValue() {
						%[1]s, err := u.%s.GetValue().Serialize()
						if err != nil {
							return nil, err
						}
		
						err = writer.EncodeBytes(%[1]s)
						if err != nil {
							return nil, err
						}
					} else {
						err = writer.EncodeNil()
						if err != nil {
							return nil, err
						}
					}
					
				`, varName, field.GoName)
			} else {
				return fmt.Sprintf(
					`
					%[1]s, err := u.%s.Serialize()
					if err != nil {
						return nil, err
					}
	
					err = writer.EncodeBytes(%[1]s)
					if err != nil {
						return nil, err
					}
				`, varName, field.GoName)
			}
		}
	} else {
		if field.Type.Nullable {
			return fmt.Sprintf(
				`
				if u.%[1]s.HasValue() {
					err = writer.%s(u.%[1]s.GetValue())
				} else {
					err = writer.EncodeNil()
				}

				if err != nil {
					return nil, err
				}
			`, field.GoName, encoderMapper[field.Type.Primitive])
		} else {
			return fmt.Sprintf(
				`
				err = writer.%s(u.%s)
				if err != nil {
					return nil, err
				}
			`, encoderMapper[field.Type.Primitive], field.GoName)
		}
	}
}

// func (b *Builder) WriteDecodeField(field *TypeFieldDefinition, pkg string) []*Statement {
// 	stmts := []*Statement{}

// 	if field.Type.Primitive == "list" {
// 		typeArgumentPrimitive := field.Type.TypeArguments[0].Primitive
// 		arrName := fmt.Sprintf("%sLen", field.GoName)

// 		stmts = append(stmts, List(Id(arrName), Id("err")).Op(":=").Id("decoder").Dot("DecodeArrayLen").Call())
// 		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
// 			Return().Id("err"),
// 		))

// 		stmts = append(stmts, Id("u").Dot(field.GoName).Op("=").Make(Index().Custom(Options{}, primitiveMapper[typeArgumentPrimitive]), Id(arrName)))
// 		stmts = append(stmts, For(
// 			Id("i").Op(":=").Lit(0),
// 			Id("i").Op("<").Id(arrName),
// 			Id("i").Op("++"),
// 		).Block(
// 			List(Id("u").Dot(field.GoName).Index(Id("i")), Id("err")).Op("=").Id("decoder").Dot(decoderMapper[typeArgumentPrimitive]).Call(),
// 			If(Id("err").Op("!=").Nil()).Block(
// 				Return().Id("err"),
// 			),
// 		))
// 	} else if field.Type.Primitive == "map" {
// 		keyArgumentPrimitive := field.Type.TypeArguments[0].Primitive
// 		valueArgumentPrimitive := field.Type.TypeArguments[1].Primitive

// 		mapName := fmt.Sprintf("%sLen", field.GoName)

// 		stmts = append(stmts, List(Id(mapName), Id("err")).Op(":=").Id("decoder").Dot("DecodeMapLen").Call())
// 		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
// 			Return().Id("err"),
// 		))

// 		stmts = append(stmts, Id("u").Dot(field.GoName).Op("=").Make(Map(primitiveMapper[keyArgumentPrimitive]).Custom(Options{}, primitiveMapper[valueArgumentPrimitive])))
// 		stmts = append(stmts, For(
// 			Id("i").Op(":=").Lit(0),
// 			Id("i").Op("<").Id(mapName),
// 			Id("i").Op("++"),
// 		).Block(
// 			List(Id("k"), Id("err")).Op(":=").Id("decoder").Dot(decoderMapper[keyArgumentPrimitive]).Call(),
// 			If(Id("err").Op("!=").Nil()).Block(
// 				Return().Id("err"),
// 			),

// 			List(Id("v"), Id("err")).Op(":=").Id("decoder").Dot(decoderMapper[valueArgumentPrimitive]).Call(),
// 			If(Id("err").Op("!=").Nil()).Block(
// 				Return().Id("err"),
// 			),

// 			Id("u").Dot(field.GoName).Index(Id("k")).Op("=").Id("v"),
// 		))
// 	} else if field.Type.Primitive == "custom" {
// 		importId := field.Type.ImportId
// 		td, ok := (*b.typeRegistry)[importId]
// 		if !ok {
// 			panic(fmt.Sprintf("type with id %s not found", importId))
// 		}

// 		samePackage := td.ImportPath == pkg

// 		if td.Definition.Modifier == "enum" {
// 			varName := fmt.Sprintf("%sIdx", field.GoName)
// 			stmts = append(stmts, List(Id(varName), Err()).Op(":=").Id("decoder").Dot("DecodeUint8").Call())
// 			stmts = append(stmts, If(Err().Op("!=").Nil()).Block(
// 				Return(Err()),
// 			))

// 			op := Id("u").Dot(field.GoName).Op("=")
// 			if samePackage {
// 				op.Id(fmt.Sprintf("%sPicker", td.Definition.Name))
// 			} else {
// 				op.Qual(td.ImportPath, fmt.Sprintf("%sPicker", td.Definition.Name))
// 			}

// 			stmts = append(stmts, op.Dot("ByIndex").Call(Id(varName)))
// 		} else {
// 			varName := fmt.Sprintf("%sBinary", field.GoName)

// 			stmts = append(stmts, List(Id(varName), Err()).Op(":=").Id("decoder").Dot("DecodeBytes").Call())
// 			stmts = append(stmts, If(Err().Op("!=").Nil()).Block(
// 				Return(Err()),
// 			))

// 			op := Id("u").Dot(field.GoName).Op("=")
// 			if samePackage {
// 				op.Id(td.Definition.Name)
// 			} else {
// 				op.Qual(td.ImportPath, td.Definition.Name)
// 			}

// 			stmts = append(stmts, op.Values())
// 			stmts = append(stmts, Id("u").Dot(field.GoName).Dot("MergeFrom").Call(Id(varName)))
// 		}

// 	} else {
// 		stmts = append(stmts, List(Id("u").Dot(field.GoName), Id("err")).Op("=").Custom(Options{}, GetDecodeTypeStatement(field.Type.Primitive)).Call())
// 		stmts = append(stmts, If(Id("err").Op("!=").Nil()).Block(
// 			Return(Id("err")),
// 		))
// 	}

// 	return stmts
// }

// func GetDecodeTypeStatement(primitive string) *Statement {
// 	decoder, ok := decoderMapper[primitive]
// 	if !ok {
// 		panic(fmt.Sprintf("unknown decoder for %s", primitive))
// 	}

// 	return Id("decoder").Dot(decoder)
// }
