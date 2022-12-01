package internal

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (b *Builder) generateEnum(t *SchemaTypeDefinition) (sourceCode string, err error) {

	b.currentContext.MustImport("fmt")

	lowerCamel := strcase.ToLowerCamel(t.Name)
	pickerName := fmt.Sprintf("%sPicker", lowerCamel)
	sb := new(strings.Builder)

	// Create enum struct
	sb.WriteString(fmt.Sprintf(
		`
		type %s struct {
			index uint8
			name string
		}%s
	`, t.Name, "\n"))

	// create enum picker struct
	sb.WriteString(fmt.Sprintf(`type %s struct {}%s`, pickerName, "\n"))

	// create enum picker instance
	sb.WriteString(fmt.Sprintf(`var %[1]sPicker %[2]s = %[2]s{}%s`, t.Name, pickerName, "\n"))

	// write each field as enum a value
	stmts := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		stmts[i] = fmt.Sprintf(`var %s%s %[3]s = %[3]s{index: %v,name: "%s"}`, lowerCamel, strcase.ToCamel(field.Name), t.Name, field.Index, field.Name)
	}
	sb.WriteString(strings.Join(stmts, "\n"))

	// write picker for each enum value
	for i, field := range t.Fields {
		stmts[i] = fmt.Sprintf(
			`
			func (%s) %[2]s() %s {
				return %s%[2]s
			}
		`, pickerName, strcase.ToCamel(field.Name), t.Name, lowerCamel)
	}
	sb.WriteString(strings.Join(stmts, "\n"))

	// write Index() method
	sb.WriteString(fmt.Sprintf(
		`
	func (e %s) Index() uint8 {
		return e.index
	}
	%s`, t.Name, "\n"))

	// write Name() method
	sb.WriteString(fmt.Sprintf(
		`
	func (e %s) Name() string {
		return e.name
	}
	%s`, t.Name, "\n"))

	// write ByIndex() method
	for i, field := range t.Fields {
		stmts[i] = fmt.Sprintf(
			`
		case %v:
			return %s%s
		`, field.Index, lowerCamel, strcase.ToCamel(field.Name))
	}

	sb.WriteString(fmt.Sprintf(
		`
	func (%s) ByIndex(index uint8) %s {
		switch index {
			%s
		default:
			panic(fmt.Sprintf("%s with index %%v not found", index))
		}
	}
	%s`, pickerName, t.Name, strings.Join(stmts, "\n"), t.Name, "\n"))

	// write ByName method
	for i, field := range t.Fields {
		stmts[i] = fmt.Sprintf(
			`
		case "%s":
			return %s%s
		`, field.Name, lowerCamel, strcase.ToCamel(field.Name))
	}

	sb.WriteString(fmt.Sprintf(
		`
	func (%s) ByName(name string) %s {
		switch name {
			%s
		default:
			panic(fmt.Sprintf("%s with name %%v not found", name))
		}
	}
	%s`, pickerName, t.Name, strings.Join(stmts, "\n"), t.Name, "\n"))

	return sb.String(), nil
}
