package generator

import (
	"html/template"
)

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
{{range .Fields}}
{{.Name}} {{getTypeName .FieldDef}}
{{end}}
}

func (u *{{.TypeName}}) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	var err error
	{{range .Fields}}
	{{getEncoder .FieldDef}}
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

func (u *{{.TypeName}}) Decode(buffer []byte) error {
	return nil
}

func (u *{{.TypeName}}) MustDecode(buffer []byte) {
	err := u.Decode(buffer)
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
	{{range .Fields}}
	{{$.TypeName}}_{{.FieldName}} {{$.TypeName}}WhichField = {{.FieldIndex}}
	{{end}}
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
func (*{{$.LowerName}}Builder) {{.FieldName}}(value {{getTypeName .FieldDef}}) *{{$.TypeName}} {
	return &{{$.TypeName}}{
		value:      value,
		fieldIndex: {{.FieldIndex}},
	}
}

func (u *{{$.TypeName}}) Set{{.FieldName}}(value {{getTypeName .FieldDef}}) {
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
		{{getEncoder .FieldDef}}
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

func (u *{{.TypeName}}) Decode(reader io.Reader) error {
	decoder := runtime.GetDecoder(reader)
	var err error
	u.fieldIndex, err = decoder.DecodeVarint()
	if err != nil {
		return err
	}

	switch fieldIndex {
	case -1:
		u.value = nil

	{{range .Fields}}
	case {{.FieldIndex}}:
		{{getDecoder .FieldDef}}
	{{end}}
	}
	
	return nil
}

func (u *{{.TypeName}}) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}
`

var enumTemplate, structTemplate, unionTemplate *template.Template

func init() {
	funcMap := template.FuncMap{
		"getEncoder":  getEncoderForFieldFunc,
		"getDecoder":  getDecoderForFieldFunc,
		"getTypeName": getTypeNameFunc,
	}

	enumTemplate = parseTemplate("enum", enumTemplateString)
	structTemplate = parseTemplateWithFunc("struct", structTemplateString, funcMap)

	unionTemplate = parseTemplateWithFunc("union", unionTemplateString, funcMap)
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
		panic(err)
	}

	return temp
}
