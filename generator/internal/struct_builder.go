package internal

import (
	"fmt"
	"strings"
)

func (b *Builder) generateStruct(t *SchemaTypeDefinition) (sourceCode string, err error) {

	pkgPath := (*b.typeRegistry)[t.Id].ImportPath
	sb := new(strings.Builder)

	// Create fields
	fields := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		fields[i] = fmt.Sprintf("%s %s", field.GoName, b.GetGoType(field.Type, pkgPath))
	}

	// write struct
	sb.WriteString(fmt.Sprintf("type %s struct {%s}", t.Name, strings.Join(fields, "\n")))

	b.writeSerializeMethod(sb, t)
	b.writeMustSerializeMethod(sb, t)
	b.writeMergeFromMethod(sb, t, pkgPath)
	b.writeMergeUsing(sb, t)

	return sb.String(), nil
}

func (b *Builder) writeSerializeMethod(sb *strings.Builder, t *SchemaTypeDefinition) {

	stmts := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		stmts[i] = b.WriteEncodeField(field)
	}

	body := fmt.Sprintf(
		`
func (u *%s) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	%s

	return buf.Bytes(), nil
}%s`, t.Name, strings.Join(stmts, "\n"), "\n")

	sb.WriteString(body)
}

func (b *Builder) writeMustSerializeMethod(sb *strings.Builder, t *SchemaTypeDefinition) {
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) MustSerialize() []byte {
		buf, err := u.Serialize()
		if err != nil {
			panic(err)
		}

		return buf
	}%s
	`, t.Name, "\n"))
}

func (b *Builder) writeMergeFromMethod(sb *strings.Builder, t *SchemaTypeDefinition, pkgName string) {
	stmts := make([]string, len(t.Fields))
	isAnyNullable := false
	for i, field := range t.Fields {
		stmts[i] = b.WriteDecodeField(field, pkgName)
		if field.Type.Nullable {
			isAnyNullable = true
		}
	}

	var isNextNilStmt string
	if isAnyNullable {
		isNextNilStmt = "var isNextNil bool"
	}

	body := fmt.Sprintf(
		`
		func (u *%s) MergeFrom(buffer []byte) error {
			reader := bytes.NewBuffer(buffer)
			decoder := msgpack.NewDecoder(reader)
			var err error
			%s
			%s

			return nil
		}%s
	`, t.Name, isNextNilStmt, strings.Join(stmts, "\n"), "\n")

	sb.WriteString(body)
}

func (b *Builder) writeMergeUsing(sb *strings.Builder, t *SchemaTypeDefinition) {
	stmts := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		stmts[i] = fmt.Sprintf("u.%[1]s = other.%[1]s", field.GoName)
	}

	body := fmt.Sprintf(
		`
		func (u *%[1]s) MergeUsing(other *%[1]s) {
			%s
		}%s
	`, t.Name, strings.Join(stmts, "\n"), "\n")

	sb.WriteString(body)
}
