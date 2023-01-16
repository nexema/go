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
