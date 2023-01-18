package generator

import (
	"fmt"
	"reflect"
	"strings"
)

func formatDefaultValueFunc(defaultValue interface{}, valueType TypeFieldValueKindTemplate) string {
	if defaultValue == nil {
		return fmt.Sprintf("runtime.NewNull[%s]()", valueType.ImportTypeName)
	}

	switch t := defaultValue.(type) {
	case string:
		return convertToNullable(valueType.ImportTypeName, fmt.Sprintf("%q", t), valueType.IsNullable)

	case int64, float64, bool:
		return convertToNullable(valueType.ImportTypeName, fmt.Sprint(defaultValue), valueType.IsNullable)

	default:
		value := reflect.ValueOf(defaultValue)
		switch value.Kind() {
		case reflect.Slice:
			values := make([]string, value.Len())
			keyValueType := valueType.TypeArguments[0]
			for i := 0; i < value.Len(); i++ {
				values[i] = formatDefaultValueFunc(value.Index(i).Interface(), keyValueType)
			}

			if valueType.IsNullable {
				return fmt.Sprintf("[]runtime.Nullable[%s]{%s}", keyValueType.ImportTypeName, strings.Join(values, ","))
			} else {
				return fmt.Sprintf("[]%s{%s}", keyValueType.ImportTypeName, strings.Join(values, ","))
			}

		case reflect.Map:
			values := make([]string, value.Len())
			keyValueType := valueType.TypeArguments[0]
			valueValueType := valueType.TypeArguments[1]
			for i, k := range value.MapKeys() {
				key := formatDefaultValueFunc(k.Interface(), keyValueType)
				value := formatDefaultValueFunc(value.MapIndex(k).Interface(), valueValueType)

				values[i] = fmt.Sprintf("%s: %s", key, value)
			}

			if valueType.IsNullable {
				return fmt.Sprintf("map[%s]runtime.Nullable[%s]{%s}", keyValueType.ImportTypeName, valueValueType.ImportTypeName, strings.Join(values, ","))
			} else {
				return fmt.Sprintf("map[%s]%s{%s}", keyValueType.ImportTypeName, valueValueType.ImportTypeName, strings.Join(values, ","))
			}
		default:
			panic(fmt.Errorf("unable to handle default type with kind %s", value.Kind()))
		}
	}
}

func convertToNullable(valueType, declaration string, nullable bool) string {
	if nullable {
		return fmt.Sprintf("runtime.NewNullable[%s](%s)", valueType, declaration)
	} else {
		return declaration
	}
}

func setNullableFunc(fieldDeclaration string, isNullable bool) string {
	if isNullable {
		return fmt.Sprintf(`runtime.Nullable[%s]`, fieldDeclaration)
	} else {
		return fieldDeclaration
	}
}

func setNullableListFunc(fieldDeclaration string, isNullable bool) string {
	if isNullable {
		return fmt.Sprintf(`runtime.Nullable[[]%s]`, fieldDeclaration)
	} else {
		return fmt.Sprintf("[]%s", fieldDeclaration)
	}
}

func setNullableMapFunc(keyDecl, valueDecl string, isNullable bool) string {
	if isNullable {
		return fmt.Sprintf(`runtime.Nullable[map[%s]%s]`, keyDecl, valueDecl)
	} else {
		return fmt.Sprintf("map[%s]%s", keyDecl, valueDecl)
	}
}

// getEncoderVariableName will output "value" when calling EncodeX()
func getEncoderVariableNameFunc(fieldName, valueType string, isUnion bool) string {
	if isUnion {
		return fmt.Sprintf("value.(%s)", valueType)
	} else {
		return fieldName
	}
}

// getDecoderVariableName will output "value" when calling EncodeX() or u.value = DecodeX()
func getDecoderVariableNameFunc(fieldName string, isUnion bool) string {
	if isUnion {
		return "value"
	} else {
		return fieldName
	}
}

// stripRuntimeNullableDecl strips from a declaration like runtime.Nullable[X] the runtime.Nullable[] part, leaving X alone
func stripRuntimeNullableDeclFunc(declaration string) string {
	return declaration[17 : len(declaration)-1]
}

func toGoTypeFunc(primitive string) string {
	switch primitive {
	case "binary":
		return "[]byte"

	default:
		return primitive
	}
}

func capitalizeFirstFunc(v string) string {
	return strings.ToUpper(string(v[0])) + v[1:]
}
