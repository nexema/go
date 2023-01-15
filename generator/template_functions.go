package generator

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

// toNullableEncoder takes a normalEncoder (ex: encoder.EncodeString(u.MyString)) and a flag isNullable.
// If isNullable is true, then it will convert normalEncoder to the following syntax:
//
//	if u.MyString.IsNull() {
//		encoder.EncodeNull()
//	} else {
//		myStringValue := *u.MyString.Value
//		encoder.EncodeString(myStringValue)
//	}
func toNullableEncoder(normalEncoder string, isNullable bool) string {
	return ""
}

func capitalizeFirstFunc(v string) string {
	return strings.ToUpper(string(v[0])) + v[1:]
}

func getEncoderForFieldFunc(field *NexemaTypeFieldDefinition) string {
	fieldName := strcase.ToCamel(field.Name) // Go field's name
	switch t := field.Type.(type) {
	case NexemaPrimitiveValueType:
		switch t.Primitive {
		case "string", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "bool", "binary":
			method := encodeTypeMapper[t.Primitive]
			if t.Nullable {
				return fmt.Sprintf(`if u.%[1]s.IsNull() {
					encoder.EncodeNull()
				} else {
					%[2]s := *u.%[1]s.Value
					encoder.%[3]s(%[2]s)
				}`, fieldName, strcase.ToLowerCamel(field.Name), method)
			} else {
				return fmt.Sprintf(`encoder.%s(u.%s);`, method, fieldName)
			}

		case "list":
			// typeArgument := t.TypeArguments[0]
			return fmt.Sprintf(`
			
			`)

		case "map":
		}

	case NexemaTypeValueType:
		typeDef := defaultGenerator.typeMapping[t.TypeId].TypeDef // typeName here is the typeId
		fieldName := strcase.ToCamel(field.Name)

		// enums encode their index
		if typeDef.Modifier == modifierEnum {
			return fmt.Sprintf("encoder.EncodeUint8(u.%s.Index());", fieldName)
		} else { // struct and union must encode their own first, then encode as binary
			return fmt.Sprintf(`
			%[1]sBytes, err := u.%[2]s.Encode(); 
			if err != nil {
				return nil, err
			}
			encoder.EncodeBinary(%[1]sBytes)
			`, strcase.ToLowerCamel(field.Name), fieldName)
		}
	}

	return ""
}

func getDecoderForFieldFunc(field *NexemaTypeFieldDefinition) string {
	fieldName := strcase.ToCamel(field.Name) // go field's name
	switch t := field.Type.(type) {
	case NexemaPrimitiveValueType:
		switch t.Primitive {
		case "string", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "bool", "binary":
			method := decodeTypeMapper[t.Primitive]
			return fmt.Sprintf(`u.%s, err = decoder.%s(); if err != nil {return err}`, fieldName, method)

		case "list":
		case "map":
		}

	case NexemaTypeValueType:
		typeDef := defaultGenerator.typeMapping[t.TypeId].TypeDef // typeName here is the typeId

		// enums decode their index
		if typeDef.Modifier == modifierEnum {
			return fmt.Sprintf(`%[1]sEnumIndex, err := decoder.DecodeUint8();
			if err != nil {
				return err
			}

			u.%[2]s = %[3]sPicker.ByIndex(%[1]sEnumIndex)
			`, strcase.ToLowerCamel(field.Name), fieldName, typeDef.Name)
		} else { // struct and union must decode itself as binary
			return fmt.Sprintf(`
			%[1]sBytes, err := decoder.DecodeBinary(); 
			if err != nil {
				return nil, err
			}
			err = u.%[2]s.Decode(%[1]sBytes)
			if err != nil {
				return err
			}
			`, strcase.ToLowerCamel(field.Name), fieldName)
		}
	}

	return ""
}

func getTypeNameFunc(field *NexemaTypeFieldDefinition) string {
	return nexemaTypeToGoType(field.Type)
}

func nexemaTypeToGoType(v NexemaValueType) string {
	switch t := v.(type) {
	case NexemaPrimitiveValueType:
		switch t.Primitive {
		case "list":
			goType := nexemaTypeToGoType(t.TypeArguments[0])
			return fmt.Sprintf("[]%s", goType)

		case "map":
			keyArgument := nexemaTypeToGoType(t.TypeArguments[0])
			valueArgument := nexemaTypeToGoType(t.TypeArguments[1])

			return fmt.Sprintf("map[%s]%s", keyArgument, valueArgument)
		default:
			if t.Nullable {
				return fmt.Sprintf("runtime.Nullable[%s]", t.Primitive)
			}

			return t.Primitive
		}

	case NexemaTypeValueType:
		typeDef := defaultGenerator.typeMapping[t.TypeId]
		typeName := ""
		// check if should add import
		if typeDef.PackageName != defaultGenerator.currentPkg {
			defaultGenerator.addImport(typeDef.ImportPath)
			typeName = fmt.Sprintf("%s.%s", filepath.Base(typeDef.ImportPath), typeDef.TypeDef.Name)
		} else {
			typeName = typeDef.TypeDef.Name
		}

		if t.Nullable {
			return fmt.Sprintf("runtime.Nullable[%s]", typeName)
		} else {
			return typeName
		}

	default:
		panic("unknown field type")
	}
}
