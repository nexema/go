package generator

import (
	"fmt"

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
