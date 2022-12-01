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

func (b *Builder) GetGoType(field SchemaFieldType, referencePkg string) string {
	primitive := field.Primitive
	if primitive == listPrimitive {
		return fmt.Sprintf("[]%s", b.GetGoType(field.TypeArguments[0], referencePkg))
	} else if primitive == mapPrimitive {
		return fmt.Sprintf("map[%s]%s", b.GetGoType(field.TypeArguments[0], referencePkg), b.GetGoType(field.TypeArguments[1], referencePkg))
	} else if primitive == customPrimitive {
		importType, ok := (*b.typeRegistry)[field.ImportId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found in registry", field.ImportId))
		}

		samePkg := importType.ImportPath == referencePkg
		if field.Nullable {
			b.currentContext.MustImport(runtimeImport)

			if samePkg {
				return fmt.Sprintf("runtime.Nullable[%s]", importType.Definition.Name)
			} else {
				b.currentContext.MustImportAs(importType.ImportPath, b.packageName)
				return fmt.Sprintf("runtime.Nullable[%s.%s]", b.packageName, importType.Definition.Name)
			}
		} else {
			if samePkg {
				return importType.Definition.Name
			} else {
				b.currentContext.MustImportAs(importType.ImportPath, b.packageName)
				return fmt.Sprintf("%s.%s", b.packageName, importType.Definition.Name)
			}
		}
	} else {
		if field.Nullable {
			b.currentContext.MustImport(runtimeImport)
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

// func (b *Builder) WriteField(t SchemaFieldType, referencePkg string) string {
// 	return b.WritePrimitive(t.Primitive, t.Nullable, t.TypeArguments, t.ImportId, referencePkg)
// }

// func (b *Builder) WritePrimitive(t string, nullable bool, arguments []SchemaFieldType, importId string, referencePkg string) string {
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
					if u.%[2]s.HasValue() {
						%[1]s, err := u.%[2]s.Value.Serialize()
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

func (b *Builder) WriteDecodeField(field *TypeFieldDefinition, pkg string) string {
	if field.Type.Primitive == listPrimitive {
		typeArgument := field.Type.TypeArguments[0]
		arrName := fmt.Sprintf("%sLen", field.GoName)

		if typeArgument.Nullable {
			return fmt.Sprintf(
				`
				%[1]s, err := decoder.DecodeArrayLen()
				if err != nil {
					return err
				}
	
				u.%[2]s = make([]runtime.Nullable[%[3]s], %[1]s)
				for i := 0; i < %[1]s; i++ {
					value, err := decoder.%[4]s()
					if err != nil {
						return err
					}

					u.%[2]s[i] = runtime.NewNullable(value)
				}
			`, arrName, field.GoName, primitiveMapper[typeArgument.Primitive], decoderMapper[typeArgument.Primitive])
		} else {
			return fmt.Sprintf(
				` 
				%[1]s, err := decoder.DecodeArrayLen()
				if err != nil {
					return err
				}
	
				u.%[2]s = make([]%[3]s, %[1]s)
				for i := 0; i < %[1]s; i++ {
					u.%[2]s[i], err = decoder.%[4]s()
					if err != nil {
						return err
					}
				}
			`, arrName, field.GoName, primitiveMapper[typeArgument.Primitive], decoderMapper[typeArgument.Primitive])
		}
	} else if field.Type.Primitive == mapPrimitive {
		keyArgument := field.Type.TypeArguments[0]
		valueArgument := field.Type.TypeArguments[1]
		mapName := fmt.Sprintf("%sLen", field.GoName)

		if valueArgument.Nullable {
			return fmt.Sprintf(
				`
				%[1]s, err := decoder.DecodeMapLen()
				if err != nil {
					return err
				}
	
				u.%[2]s = make(map[%[3]s]runtime.Nullable[%[4]s])
				for i := 0; i < %[1]s; i++ {
					k, err := decoder.%[5]s()
					if err != nil {
						return err
					}
	
					v, err := decoder.%[6]s()
					if err != nil {
						return err
					}
	
					u.%[2]s[k] = runtime.NewNullable(v)
				}
			`, mapName, field.GoName, primitiveMapper[keyArgument.Primitive], primitiveMapper[valueArgument.Primitive], decoderMapper[keyArgument.Primitive], decoderMapper[valueArgument.Primitive])
		} else {
			return fmt.Sprintf(
				`
				%[1]s, err := decoder.DecodeMapLen()
				if err != nil {
					return err
				}
	
				u.%[2]s = make(map[%[3]s]%[4]s)
				for i := 0; i < %[1]s; i++ {
					k, err := decoder.%[5]s()
					if err != nil {
						return err
					}
	
					v, err := decoder.%[6]s()
					if err != nil {
						return err
					}
	
					u.%[2]s[k] = v
				}
			`, mapName, field.GoName, primitiveMapper[keyArgument.Primitive], primitiveMapper[valueArgument.Primitive], decoderMapper[keyArgument.Primitive], decoderMapper[valueArgument.Primitive])
		}
	} else if field.Type.Primitive == customPrimitive {
		importId := field.Type.ImportId
		td, ok := (*b.typeRegistry)[importId]
		if !ok {
			panic(fmt.Sprintf("type with id %s not found", importId))
		}

		samePackage := td.ImportPath == pkg

		if td.Definition.Modifier == "enum" {
			typeUsage := ""
			if samePackage {
				typeUsage = td.Definition.Name
			} else {
				b.currentContext.MustImportAs(td.ImportPath, b.packageName)
				typeUsage = fmt.Sprintf("%s.%s", b.packageName, td.Definition.Name)
			}

			varName := fmt.Sprintf("%sIdx", field.GoName)
			return fmt.Sprintf(
				`
				%[1]s, err := decoder.DecodeUint8()
				if err != nil {
					return err
				}

				u.%[2]s = %[3]sPicker.ByIndex(%[1]s)
			`, varName, field.GoName, typeUsage)

		} else {
			varName := fmt.Sprintf("%sBinary", field.GoName)
			typeUsage := ""
			if samePackage {
				typeUsage = td.Definition.Name
			} else {
				b.currentContext.MustImportAs(td.ImportPath, b.packageName)
				typeUsage = fmt.Sprintf("%s.%s", b.packageName, td.Definition.Name)
			}

			if field.Type.Nullable {
				return fmt.Sprintf(
					`
					isNextNil, err = decoder.IsNextNil()
					if err != nil {
						return err
					}

					if isNextNil {
						u.%[1]s = runtime.NewNull[%[3]s]()
					} else {
						%[2]s, err := decoder.DecodeBytes()
						if err != nil {
							return err
						}

						value := %[3]s{}
						err = value.MergeFrom(%[2]s)
						if err != nil {
							return err
						}
						u.%[1]s = runtime.NewNullable(value)
					}
				`, field.GoName, varName, typeUsage)
			} else {
				return fmt.Sprintf(
					`
					%[1]s, err := decoder.DecodeBytes()
					if err != nil {
						return err
					}
	
					err = u.%[2]s.MergeFrom(%[1]s)
					if err != nil {
						return err
					}
				`, varName, field.GoName, typeUsage)
			}
		}
	} else {
		if field.Type.Nullable {
			return fmt.Sprintf(
				`
				isNextNil, err = decoder.IsNextNil()
				if err != nil {
					return err
				}

				if isNextNil {
					u.%[1]s = runtime.NewNull[%[2]s]()
				} else {
					value, err := decoder.%[3]s()
					if err != nil {
						return err
					}

					u.%[1]s = runtime.NewNullable(value)
				}
			`, field.GoName, primitiveMapper[field.Type.Primitive], decoderMapper[field.Type.Primitive])
		} else {
			return fmt.Sprintf(
				`
				u.%s, err = decoder.%s()
				if err != nil {
					return err
				}
			`, field.GoName, decoderMapper[field.Type.Primitive])
		}
	}
}
