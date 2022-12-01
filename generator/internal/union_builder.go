package internal

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (b *Builder) generateUnion(t *SchemaTypeDefinition, pkg string) (sourceCode string, err error) {

	pkgPath := (*b.typeRegistry)[t.Id].ImportPath
	sb := new(strings.Builder)

	// Create fields
	fields := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		fields[i] = fmt.Sprintf("%s %s", field.GoName, b.GetGoType(field.Type, pkgPath))
	}

	builderName := fmt.Sprintf("%sBuilder", t.Name)
	builderTypeName := strcase.ToLowerCamel(builderName)
	whichFieldName := fmt.Sprintf("%sWhichField", t.Name)

	// write union struct
	sb.WriteString(fmt.Sprintf(`type %s struct {
		value interface{}
		fieldIndex int8
	}%s`, t.Name, "\n"))

	// write whichField type and constants
	sb.WriteString(fmt.Sprintf("type %s int8", whichFieldName))
	sb.WriteString(fmt.Sprintf(
		`	
	const (
		%s_NotSet %s = -1
		%s
	)
	`, t.Name, whichFieldName, strings.Join(MapArray(t.Fields, func(v *TypeFieldDefinition) string {
			return fmt.Sprintf(`%s_%s %s = %v`, t.Name, v.GoName, whichFieldName, v.Index)
		}), "\n")))

	// write builder type and var
	sb.WriteString(fmt.Sprintf(`type %s struct{}%s`, builderTypeName, "\n"))
	sb.WriteString(fmt.Sprintf(`var %[1]s *%[2]s = &%[2]s{}%[3]s`, builderName, builderTypeName, "\n"))

	// write builder methods
	for _, field := range t.Fields {
		sb.WriteString(fmt.Sprintf(
			`
			func (*%[1]s) %[2]s(value %[3]s) *%[4]s {
				return &%[4]s{
					value: value,
					fieldIndex: %v,
				}
			}%s
		`, builderTypeName, field.GoName, primitiveMapper[field.Type.Primitive], t.Name, field.Index, "\n"))
	}

	// write IsSet method
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) IsSet() bool {
		return u.fieldIndex != -1
	}
	`, t.Name))

	// write Clear method
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) Clear() {
		u.value = nil
		u.fieldIndex = -1
	}
	`, t.Name))

	// write WhichField method
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) WhichField() %[2]s {
		return %[2]s(u.fieldIndex)
	}
	`, t.Name, whichFieldName))

	// write CurrentValue method
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) CurrentValue() interface{} {
		return u.value
	}
	`, t.Name))

	// write field setters and getters
	for _, field := range t.Fields {
		sb.WriteString(fmt.Sprintf(
			`
		func (u %s) Set%s(value %s) {
			u.value = value
			u.fieldIndex = %v
		}
		`, t.Name, field.GoName, primitiveMapper[field.Type.Primitive], field.Index))

		sb.WriteString(fmt.Sprintf(
			`
		func (u %s) Get%s() %[3]s {
			return u.value.(%[3]s)
		}
		`, t.Name, field.GoName, primitiveMapper[field.Type.Primitive]))
	}

	// writer helper methods to check which field is set
	for _, field := range t.Fields {
		sb.WriteString(fmt.Sprintf(
			`
		func (u %s) Is%s() bool {
			return u.fieldIndex == %v
		}
		`, t.Name, field.GoName, field.Index))
	}

	// write Serialize
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) Serialize() ([]byte, error) {
		buf := new(bytes.Buffer)
		writer := msgpack.NewEncoder(buf)
		var err error

		err = writer.EncodeInt8(u.fieldIndex)
		if err != nil {
			return nil, err
		}

		if u.fieldIndex > -1 {
			switch u.fieldIndex {
				%s
			}
		}

		return buf.Bytes(), nil
	}
	`, t.Name, strings.Join(MapArray(t.Fields, func(v *TypeFieldDefinition) string {
			if v.Type.Primitive == customPrimitive {
				return fmt.Sprintf(
					`
					case %v:
					bytes, err := (u.value.(%s)).Serialize()
					if err != nil {
						return nil, err
					}
					err = writer.%s(bytes)
					if err != nil {
						return nil, err
					}
				`, v.Index, encoderMapper[v.Type.Primitive], primitiveMapper[v.Type.Primitive])
			} else {
				return fmt.Sprintf(
					`
					case %v:
					err = writer.%s(u.value.(%s))
					if err != nil {
						return nil, err
					}
				`, v.Index, encoderMapper[v.Type.Primitive], primitiveMapper[v.Type.Primitive])
			}
		}), "\n")))

	// write MustSerialize
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) MustSerialize() []byte {
		buf, err := u.Serialize()
		if err != nil {
			panic(err)
		}

		return buf
	}
	`, t.Name))

	// write MergeFrom
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) MergeFrom(buffer []byte) error {
		reader := bytes.NewBuffer(buffer)
		decoder := msgpack.NewDecoder(reader)
		var err error

		fieldIndex, err := decoder.DecodeInt8()
		if err != nil {
			return err
		}

		u.fieldIndex = fieldIndex
		switch fieldIndex {
		case -1:
			u.value = nil

			%s
		}

		return nil
	}
	`, t.Name, strings.Join(MapArray(t.Fields, func(v *TypeFieldDefinition) string {
			if v.Type.Primitive == customPrimitive {
				importId := v.Type.ImportId
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

					varName := fmt.Sprintf("%sIdx", v.GoName)
					// %[3]sPicker.ByIndex(%[1]s)
					return fmt.Sprintf(
						`
						case %[1]v:
							%[2]s, err := decoder.DecodeUint8()
							if err != nil {
								return err
							}
							u.value = %[3]sPicker.ByIndex(%[2]s)
					`, v.Index, varName, typeUsage)

				} else {
					typeUsage := ""
					if samePackage {
						typeUsage = td.Definition.Name
					} else {
						b.currentContext.MustImportAs(td.ImportPath, b.packageName)
						typeUsage = fmt.Sprintf("%s.%s", b.packageName, td.Definition.Name)
					}

					return fmt.Sprintf(
						`
						case %v:
							bytes, err := decoder.DecodeBytes()
							if err != nil {
								return err
							}
	
							value := %s{}
							err = value.MergeFrom(bytes)
							if err != nil {
								return err
							}
	
							u.value = value
					`, v.Index, typeUsage)
				}
			} else {
				return fmt.Sprintf(
					`
					case %v:
						u.value, err = decoder.%s()
						if err != nil {
							return err
						}
				`, v.Index, decoderMapper[v.Type.Primitive])
			}
		}), "\n")))

	return sb.String(), nil
}
