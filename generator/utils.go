package generator

import (
	"fmt"
	"path/filepath"
)

const (
	quotes  byte = '"'
	newline byte = '\n'
	rParen  byte = ')'
)

func stringPtr(v string) *string {
	return &v
}

func mapArray[I any, O any](in []I, f func(v I) O) []O {
	out := make([]O, len(in))
	for i, v := range in {
		out[i] = f(v)
	}

	return out
}

func prepend(x []byte, y []byte) []byte {
	return append(y, x...) // todo: maybe there is a better performant way?
}

func nexemaTypeFieldDefinitionToTemplateData(fieldValueType NexemaValueType) TypeFieldValueKindTemplate {
	switch t := fieldValueType.(type) {
	case NexemaPrimitiveValueType:
		valueType := TypeFieldValueKindTemplate{
			IsNullable:     t.Nullable,
			IsList:         t.Primitive == "list",
			IsMap:          t.Primitive == "map",
			IsEnum:         false,
			IsType:         false,
			IsPrimitive:    t.Primitive != "list" && t.Primitive != "map",
			ImportTypeName: t.Primitive,
			TypeArguments:  make([]TypeFieldValueKindTemplate, 0),
		}

		for _, typeArgument := range t.TypeArguments {
			valueType.TypeArguments = append(valueType.TypeArguments, nexemaTypeFieldDefinitionToTemplateData(typeArgument))
		}

		return valueType

	case NexemaTypeValueType:
		typeDef := defaultGenerator.typeMapping[t.TypeId]
		typeName := typeDef.TypeDef.Name
		var importTypeName string

		// check if should add import
		if typeDef.PackageName != defaultGenerator.currentPkg {
			defaultGenerator.addImport(typeDef.ImportPath)
			importTypeName = fmt.Sprintf("%s.%s", filepath.Base(typeDef.ImportPath), typeName)
		} else {
			importTypeName = typeName
		}

		valueType := TypeFieldValueKindTemplate{
			IsNullable:     t.Nullable,
			IsEnum:         typeDef.TypeDef.Modifier == modifierEnum,
			IsType:         typeDef.TypeDef.Modifier != modifierEnum,
			IsPrimitive:    false,
			IsList:         false,
			IsMap:          false,
			ImportTypeName: importTypeName,
		}

		return valueType

	default:
		panic("invalid value type")
	}
}
