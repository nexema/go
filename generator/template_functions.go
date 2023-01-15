package generator

import (
	"fmt"
	"path/filepath"

	"github.com/iancoleman/strcase"
)

func getEncoderForTypeFunc(fieldName, typeName, typeId string) string {
	switch typeName {
	case "string", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "bool", "binary":
		method := encodeTypeMapper[typeName]
		return fmt.Sprintf(`encoder.%s(u.%s);`, method, fieldName)

	case "list":
	case "map":
	default:
		typeDef := defaultGenerator.typeMapping[typeId].TypeDef // typeName here is the typeId

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
			`, strcase.ToLowerCamel(fieldName), fieldName)
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
