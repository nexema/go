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
{{.Name}} {{.ImportTypeName}}
{{end}}
}

func (u *{{.TypeName}}) Encode() ([]byte, error) {
	encoder := runtime.GetEncoder()
	var err error
	{{range .Fields}}
	{{getEncoder .Name .TypeName .TypeId}}
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

var enumTemplate, structTemplate *template.Template

func init() {
	enumTemplate = parseTemplate("enum", enumTemplateString)
	structTemplate = parseTemplateWithFunc("struct", structTemplateString, template.FuncMap{
		"getEncoder": getEncoderForTypeFunc,
	})
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
