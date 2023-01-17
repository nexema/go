package generator

import (
	"fmt"
	"strings"
)

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
