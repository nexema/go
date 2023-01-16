package generator

import (
	"fmt"
	"text/template"
)

const rawFieldEncoderTemplateString = `
{{if .ValueType.IsList}}
	{{if .ValueType.IsNullable}}
	encoder.BeginArray(int64(len(value)))
	for _, element := range value {
		{{if (index .ValueType.TypeArguments 0).IsNullable}}
		if element.IsNull() {
			encoder.EncodeNull()
		} else {
			{{template "encodePrimitive" (index .ValueType.TypeArguments 0)}}(*element.Value)
		}
		{{else}}
		{{template "encodePrimitive" (index .ValueType.TypeArguments 0)}}(element)
		{{end}}
	}
	{{else}}
	{{end}}

{{else if .ValueType.IsEnum}}
	encoder.EncodeUint8(u.{{.FieldName}}.Index())
{{else if .ValueType.IsMap}}
	encoder.BeginMap(int64(len(u.{{.FieldName}})))
	for key, value := range u.{{.FieldName}} {
		{{template "encodePrimitive" (index .ValueType.TypeArguments 0)}}(key)
		
		{{if (index .ValueType.TypeArguments 1).IsNullable}}
		if value.IsNull() {
			encoder.EncodeNull()
		} else {
			{{template "encodePrimitive" (index .ValueType.TypeArguments 1)}}(*value.Value)
		}
		{{else}}
		{{template "encodePrimitive" (index .ValueType.TypeArguments 1)}}(value)
		{{end}}
	}
{{else if .ValueType.IsPrimitive}}
	{{if .ValueType.IsNullable}}
	{{template "encodePrimitive" .ValueType}}(value)
	{{else}}
	{{template "encodePrimitive" .ValueType}}(u.{{.FieldName}})
	{{end}}
{{else if .ValueType.IsType}}
	{{.LowerFieldName}}Bytes, err := {{.LowerFieldName}}.Encode()
	if err != nil {
		return nil, err
	}

	encoder.EncodeBinary({{.LowerFieldName}}Bytes)
{{end}}
`

const rawFieldDecoderTemplateString = `
{{if .ValueType.IsList}}
	{{.LowerFieldName}}ArrayLen, err := decoder.BeginDecodeArray()
	if err != nil {
		return err
	}

	u.{{.FieldName}} = make({{template "fieldDeclaration" .ValueType}}, {{.LowerFieldName}}ArrayLen)
	for i := int64(0); i < {{.LowerFieldName}}ArrayLen; i++ {
		{{if (index .ValueType.TypeArguments 0).IsNullable}}
		if decoder.IsNextNull() {
			u.{{.FieldName}}[i] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 0).ImportTypeName}}]()
		} else {
			{{.LowerFieldName}},err := {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
			if err != nil {
				return err
			} 

			u.{{.FieldName}}[i] = runtime.NewNullable[{{toGoType (index .ValueType.TypeArguments 0).ImportTypeName}}]({{.LowerFieldName}})
		}
		{{else}}
		u.{{.FieldName}}[i],err = {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
		if err != nil {
			return err
		} 
		{{end}}
	}

{{else if .ValueType.IsEnum}}
	{{.LowerFieldName}}EnumIndex, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.{{.FieldName}} = {{.ValueType.ImportTypeName}}Picker.ByIndex({{.LowerFieldName}}EnumIndex)
{{else if .ValueType.IsMap}}
	{{.LowerFieldName}}MapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	u.{{.FieldName}} = make({{template "fieldDeclaration" .ValueType}}, {{.LowerFieldName}}MapLen)
	for i:=int64(0); i < {{.LowerFieldName}}MapLen; i++ {
		key, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
		if err != nil {
			return err
		}
		
		{{if (index .ValueType.TypeArguments 1).IsNullable}}
		if decoder.IsNextNull() {
			u.{{.FieldName}}[key] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 1).ImportTypeName}}]()
		} else {
			value, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 1)}}()
			if err != nil {
				return err
			}

			u.{{.FieldName}}[key] = runtime.NewNullable[{{toGoType (index .ValueType.TypeArguments 1).ImportTypeName}}](value)
		}
		{{else}}
		value, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 1)}}()
		if err != nil {
			return err
		}

		u.{{.FieldName}}[key] = value
		{{end}}
	}
{{else if .ValueType.IsPrimitive}}
	{{if .ValueType.IsNullable}}
	var value {{.ValueType.ImportTypeName}}
	value, err = {{template "decodePrimitive" .ValueType}}()
	if err != nil {
		return err
	}

	u.{{.FieldName}}.SetValue(value)
	{{else}}
	u.{{.FieldName}}, err = {{template "decodePrimitive" .ValueType}}()
	if err != nil {
		return err
	}
	{{end}}
{{else if .ValueType.IsType}}
	{{.LowerFieldName}}Bytes, err := decoder.DecodeBinary()
	if err != nil {
		return err
	}

	{{if .ValueType.IsNullable}}
	value := {{.ValueType.ImportTypeName}}{}
	err = value.MergeFrom({{.LowerFieldName}}Bytes)
	if err != nil {
		return err
	}

	u.{{.FieldName}}.SetValue(value)
	{{else}}
	err = u.{{.FieldName}}.MergeFrom({{.LowerFieldName}}Bytes)
	if err != nil {
		return err
	}
	{{end}}
{{end}}
`

const encodePrimitiveTemplateString = `encoder.Encode{{capitalizeFirst .ImportTypeName}}`
const decodePrimitiveTemplateString = `decoder.Decode{{capitalizeFirst .ImportTypeName}}`

const fieldDeclarationTemplateString = `
{{- if .IsList -}}
	{{setNullableList (setNullable (toGoType (index .TypeArguments 0).ImportTypeName) (index .TypeArguments 0).IsNullable) .IsNullable}}
{{- else if .IsMap -}}
	map[{{setNullable (toGoType (index .TypeArguments 0).ImportTypeName) (index .TypeArguments 0).IsNullable}}]{{setNullable (toGoType (index .TypeArguments 1).ImportTypeName) (index .TypeArguments 1).IsNullable}}
{{- else -}}
	{{setNullable (toGoType .ImportTypeName) .IsNullable}}
{{- end}}`

const enumTemplateString = `type {{.TypeName}} struct {
	index uint8
	name  string
}

func (enum {{.TypeName}}) Index() uint8 {
	return enum.index
}

func (enum {{.TypeName}}) Name() string {
	return enum.name
}

type {{.LowerTypeName}}Picker struct{}

var {{.TypeName}}Picker {{.LowerTypeName}}Picker = {{.LowerTypeName}}Picker{}
{{range .Fields}}
var {{$.LowerTypeName}}{{.Name}} {{$.TypeName}} = {{$.TypeName}}{index: {{.Index}}, name: "{{.LowerName}}"}

func ({{$.LowerTypeName}}Picker) {{.Name}}() {{$.TypeName}} {
	return {{$.LowerTypeName}}{{.Name}}
}
{{end}}

func ({{.LowerTypeName}}Picker) ByIndex(index uint8) {{.TypeName}} {
	switch index {
	{{range .Fields}}
	case {{.Index}}:
		return {{$.LowerTypeName}}{{.Name}}
	{{end}}

	default:
		panic(fmt.Sprintf("unknown {{.TypeName}} with value %v", index))
	}
}

func ({{.LowerTypeName}}Picker) ByName(name string) {{.TypeName}} {
	switch name {
	{{range .Fields}}
	case "{{.LowerName}}":
		return {{$.LowerTypeName}}{{.Name}}
	{{end}}

	default:
		panic(fmt.Sprintf("unknown {{.TypeName}} with name %s", name))
	}
}
`

const structTemplateString = `type {{.TypeName}} struct {
{{range .Fields}}{{.FieldName}} {{template "fieldDeclaration" .ValueType}}
{{end}}
}

func (u *{{.TypeName}}) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	{{range .Fields}}
	{{if .ValueType.IsNullable}}
		if u.{{.FieldName}}.IsNull() {
			encoder.EncodeNull()
		} else {
			value := *u.{{.FieldName}}.Value
			{{template "rawFieldEncoder" .}}
		}
	{{else}}
		{{template "rawFieldEncoder" .}}
	{{end}}
	{{end}}

	return encoder.TakeBytes(), nil
}

func (u *{{.TypeName}}) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u {{.TypeName}}) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	{{range .Fields}}
	{{if .ValueType.IsNullable}}
		if decoder.IsNextNull() {
			u.{{.FieldName}}.Clear()
		} else {
			{{template "rawFieldDecoder" .}}
		}
	{{else}}
		{{template "rawFieldDecoder" .}}
	{{end}}
	{{end}}
	return nil
}

func (u {{.TypeName}}) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u {{.TypeName}}) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}
`

const unionTemplateString = `type {{.TypeName}} struct {
	value interface{}
	fieldIndex int64
}

type {{.TypeName}}WhichField int8
const (
	{{.TypeName}}_NotSet {{.TypeName}}WhichField = -1
	{{- range .Fields}}
	{{$.TypeName}}_{{.FieldName}} {{$.TypeName}}WhichField = {{.FieldIndex}}
	{{- end}}
)

type {{.LowerName}}Builder struct{}
var {{.TypeName}}Builder *{{.LowerName}}Builder = &{{.LowerName}}Builder{}

func (u *{{.TypeName}}) IsSet() bool {
	return u.fieldIndex != -1
}

func (u *{{.TypeName}}) Clear() {
	u.value = nil
	u.fieldIndex = -1
}

func (u *{{.TypeName}}) WhichField() {{.TypeName}}WhichField {
	return {{.TypeName}}WhichField(u.fieldIndex)
}

func (u *{{.TypeName}}) CurrentValue() interface{} {
	return u.value
}

{{range .Fields}}
func (*{{$.LowerName}}Builder) {{.FieldName}}(value {{template "fieldDeclaration" .ValueType}}) *{{$.TypeName}} {
	return &{{$.TypeName}}{
		value:      value,
		fieldIndex: {{.FieldIndex}},
	}
}

func (u *{{$.TypeName}}) Set{{.FieldName}}(value {{template "fieldDeclaration" .ValueType}}) {
	u.value = value
	u.fieldIndex = {{.FieldIndex}}
}
{{end}}

func (u *{{.TypeName}}) MergeUsing(other *{{.TypeName}}) error {
	u.fieldIndex = other.fieldIndex
	u.value = other.value
	return nil
}

func (u *{{.TypeName}}) Clone() *{{.TypeName}} {
	return &{{.TypeName}}{
		fieldIndex: u.fieldIndex,
		value:      u.value,
	}
}

func (u *{{.TypeName}}) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	encoder.EncodeVarint(u.fieldIndex)
	switch u.fieldIndex {
	{{range .Fields}}
	case {{.FieldIndex}}:
		{{template "rawFieldEncoder" .}}
	{{end}}
	}

	return encoder.TakeBytes(), nil
}

func (u *{{.TypeName}}) MustEncode() []byte {
	bytes, err := u.Encode()
	if err != nil {
		panic(err)
	}

	return bytes
}

func (u {{.TypeName}}) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	u.fieldIndex, err = decoder.DecodeVarint()
	if err != nil {
		return err
	}

	switch u.fieldIndex {
	case -1:
		u.value = nil

	{{range .Fields}}
	case {{.FieldIndex}}:
		{{template "rawFieldDecoder" .}}
	{{end}}
	}
	
	return nil
}

func (u {{.TypeName}}) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}
`

var enumTemplate, structTemplate, unionTemplate, rawFieldEncoder, rawFieldDecoder, primitiveEncoder, primitiveDecoder, fieldDeclarationTemplate *template.Template

func init() {
	funcMap := template.FuncMap{
		"capitalizeFirst": capitalizeFirstFunc,
		"setNullable":     setNullableFunc,
		"setNullableList": setNullableListFunc,
		"toGoType":        toGoTypeFunc,
	}
	rawFieldEncoder = parseTemplateWithFunc("rawFieldEncoder", rawFieldEncoderTemplateString, funcMap)
	primitiveEncoder = parseTemplateWithFunc("encodePrimitive", encodePrimitiveTemplateString, funcMap)
	fieldDeclarationTemplate = parseTemplateWithFunc("fieldDeclaration", fieldDeclarationTemplateString, funcMap)

	rawFieldDecoder = parseTemplateWithFunc("rawFieldDecoder", rawFieldDecoderTemplateString, funcMap)
	primitiveDecoder = parseTemplateWithFunc("decodePrimitive", decodePrimitiveTemplateString, funcMap)

	enumTemplate = parseTemplate("enum", enumTemplateString)
	structTemplate = parseTemplateWithFunc("struct", structTemplateString, funcMap)
	unionTemplate = parseTemplateWithFunc("union", unionTemplateString, funcMap)

	structTemplate, _ = structTemplate.AddParseTree("fieldDeclaration", fieldDeclarationTemplate.Tree)
	structTemplate, _ = structTemplate.AddParseTree("rawFieldEncoder", rawFieldEncoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("rawFieldDecoder", rawFieldDecoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("encodePrimitive", primitiveEncoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("decodePrimitive", primitiveDecoder.Tree)

	unionTemplate, _ = unionTemplate.AddParseTree("fieldDeclaration", fieldDeclarationTemplate.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("rawFieldEncoder", rawFieldEncoder.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("rawFieldDecoder", rawFieldDecoder.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("encodePrimitive", primitiveEncoder.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("decodePrimitive", primitiveDecoder.Tree)
}

func parseTemplate(name, content string) *template.Template {
	temp, err := template.New(name).Parse(content)
	if err != nil {
		panic(err)
	}

	return temp
}

func parseTemplateWithFunc(name, content string, funcs template.FuncMap) *template.Template {
	temp, err := template.New(name).Funcs(funcs).Parse(content)
	if err != nil {
		panic(fmt.Errorf("unable to parse template %q, error: %q", name, err.Error()))
	}

	return temp
}
