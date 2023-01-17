package generator

import (
	"fmt"
	"text/template"
)

const rawFieldEncoderTemplateString = `
{{- if .ValueType.IsList}}
	{{- if .ValueType.IsNullable}}
	encoder.BeginArray(int64(len(value)))
	for _, element := range value {
		{{- if (index .ValueType.TypeArguments 0).IsNullable}}
		if element.IsNull() {
			encoder.EncodeNull()
		} else {
			{{template "encodePrimitive" (index .ValueType.TypeArguments 0)}}(*element.Value)
		}
		{{- else}}
		{{template "encodePrimitive" (index .ValueType.TypeArguments 0)}}(element)
		{{- end}}
	}
	{{- else}}
	{{- end}}

{{- else if .ValueType.IsEnum}}
	encoder.EncodeUint8(u.{{.FieldName}}.Index())
{{- else if .ValueType.IsMap}}
	{{- if .ValueType.IsNullable}}
	encoder.BeginMap(int64(len(value)))
	for key, value := range value {
	{{- else}}
	encoder.BeginMap(int64(len(u.{{.FieldName}})))
	for key, value := range u.{{.FieldName}} {
	{{- end}}
		{{template "encodePrimitive" (index .ValueType.TypeArguments 0)}}(key)
		{{- if (index .ValueType.TypeArguments 1).IsNullable}}
		if value.IsNull() {
			encoder.EncodeNull()
		} else {
			{{template "encodePrimitive" (index .ValueType.TypeArguments 1)}}(*value.Value)
		}
		{{- else}}
		{{template "encodePrimitive" (index .ValueType.TypeArguments 1)}}(value)
		{{- end}}
	}
{{- else if .ValueType.IsPrimitive}}
	{{- if .ValueType.IsNullable}}
	{{template "encodePrimitive" .ValueType}}(value)
	{{- else}}
	{{template "encodePrimitive" .ValueType}}(u.{{getEncoderVariableName .FieldName .ValueType.ImportTypeName .IsFromUnion}})
	{{- end}}
{{- else if .ValueType.IsType}}
	{{.LowerFieldName}}Bytes, err := u.{{.FieldName}}.Encode()
	if err != nil {
		return nil, err
	}

	encoder.EncodeBinary({{.LowerFieldName}}Bytes)
{{end}}
`

const rawFieldDecoderTemplateString = `
{{- if .ValueType.IsList}}
	{{- if .ValueType.IsNullable}}
		{{.LowerFieldName}}ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}

		u.{{.FieldName}}.SetValue(make({{template "instantiateList" .ValueType}}, {{.LowerFieldName}}ArrayLen))
		for i := int64(0); i < {{.LowerFieldName}}ArrayLen; i++ {
			{{- if (index .ValueType.TypeArguments 0).IsNullable}}
			if decoder.IsNextNull() {
				{{- if .ValueType.IsNullable}}
				(*u.{{.FieldName}}.Value)[i] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 0).ImportTypeName}}]()
				{{- else}}
				u.{{.FieldName}}[i] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 0).ImportTypeName}}]()
				{{- end}}
			} else {
				{{.LowerFieldName}},err := {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
				if err != nil {
					return err
				} 

				{{- if .ValueType.IsNullable}}
				(*u.{{.FieldName}}.Value)[i] = runtime.NewNullable({{.LowerFieldName}})
				{{- else}}
				u.{{.FieldName}}[i] = runtime.NewNullable({{.LowerFieldName}})
				{{- end}}
			}
			{{- else}}
			{{- if .ValueType.IsNullable}}
			(*u.{{.FieldName}}.Value)[i],err = {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
			if err != nil {
				return err
			} 
			{{- else}}
			u.{{.FieldName}}[i],err = {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
			if err != nil {
				return err
			} 
			{{- end}}
			{{- end}}
		}
	{{- else}}
		{{.LowerFieldName}}ArrayLen, err := decoder.BeginDecodeArray()
		if err != nil {
			return err
		}

		{{- if .IsFromUnion}}
		value := make({{template "fieldDeclaration" .ValueType}}, {{.LowerFieldName}}ArrayLen)
		{{- else}}
		u.{{.FieldName}} = make({{template "fieldDeclaration" .ValueType}}, {{.LowerFieldName}}ArrayLen)
		{{- end}}
		for i := int64(0); i < {{.LowerFieldName}}ArrayLen; i++ {
			{{- if (index .ValueType.TypeArguments 0).IsNullable}}
			if decoder.IsNextNull() {
				{{- if .IsFromUnion}}value{{- else}}u.{{.FieldName}}{{- end}}[i] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 0).ImportTypeName}}]()
			} else {
				{{.LowerFieldName}},err := {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
				if err != nil {
					return err
				} 

				{{if .IsFromUnion}}value{{- else}}u.{{.FieldName}}{{- end}}[i] = runtime.NewNullable({{.LowerFieldName}})
			}
			{{- else}}
			u.{{.FieldName}}[i],err = {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
			if err != nil {
				return err
			} 
			{{- end}}
		}
		{{- if .IsFromUnion}}
		u.value = value
		{{- end}}
	{{- end}}
{{- else if .ValueType.IsEnum}}
	{{.LowerFieldName}}EnumIndex, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.{{.FieldName}} = {{.ValueType.ImportTypeName}}Picker.ByIndex({{.LowerFieldName}}EnumIndex)
{{- else if .ValueType.IsMap}}
	{{.LowerFieldName}}MapLen, err := decoder.BeginDecodeMap()
	if err != nil {
		return err
	}

	{{- if .ValueType.IsNullable}}
	u.{{.FieldName}}.SetValue(make({{template "instantiateMap" .ValueType}}, {{.LowerFieldName}}MapLen))
	for i:=int64(0); i < {{.LowerFieldName}}MapLen; i++ {
		key, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
		if err != nil {
			return err
		}
		
		{{- if (index .ValueType.TypeArguments 1).IsNullable}}
		if decoder.IsNextNull() {
			(*u.{{.FieldName}}.Value)[key] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 1).ImportTypeName}}]()
		} else {
			value, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 1)}}()
			if err != nil {
				return err
			}

			(*u.{{.FieldName}}.Value)[key] = runtime.NewNullable(value)
		}
		{{- else}}
		value, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 1)}}()
		if err != nil {
			return err
		}

		(*u.{{.FieldName}}.Value)[key] = value
		{{- end}}
	}
	{{- else}}
	u.{{.FieldName}} = make({{template "fieldDeclaration" .ValueType}}, {{.LowerFieldName}}MapLen)
	for i:=int64(0); i < {{.LowerFieldName}}MapLen; i++ {
		key, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 0)}}()
		if err != nil {
			return err
		}
		
		{{- if (index .ValueType.TypeArguments 1).IsNullable}}
		if decoder.IsNextNull() {
			{{- if .ValueType.IsNullable}}
			(*u.{{.FieldName}}.Value)[key] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 1).ImportTypeName}}]()
			{{- else}}
			u.{{.FieldName}}[key] = runtime.NewNull[{{toGoType (index .ValueType.TypeArguments 1).ImportTypeName}}]()
			{{- end}}
		} else {
			value, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 1)}}()
			if err != nil {
				return err
			}

			{{- if .ValueType.IsNullable}}
			(*u.{{.FieldName}}.Value)[key] = runtime.NewNullable(value)
			{{- else}}
			u.{{.FieldName}}[key] = runtime.NewNullable(value)
			{{- end}}
		}
		{{- else}}
		value, err := {{template "decodePrimitive" (index .ValueType.TypeArguments 1)}}()
		if err != nil {
			return err
		}

		u.{{.FieldName}}[key] = value
		{{- end}}
	}
	{{- end}}
{{- else if .ValueType.IsPrimitive}}
	{{if .ValueType.IsNullable}}
	var value {{toGoType .ValueType.ImportTypeName}}
	value, err = {{template "decodePrimitive" .ValueType}}()
	if err != nil {
		return err
	}

	u.{{.FieldName}}.SetValue(value)
	{{else}}
	u.{{getDecoderVariableName .FieldName .IsFromUnion}}, err = {{template "decodePrimitive" .ValueType}}()
	if err != nil {
		return err
	}
	{{end}}
{{- else if .ValueType.IsType}}
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
	{{- end}}
{{- end}}
`

const encodePrimitiveTemplateString = `encoder.Encode{{capitalizeFirst .ImportTypeName}}`
const decodePrimitiveTemplateString = `decoder.Decode{{capitalizeFirst .ImportTypeName}}`

const fieldDeclarationTemplateString = `
{{- if .IsList -}}
	{{- setNullableList (setNullable (toGoType (index .TypeArguments 0).ImportTypeName) (index .TypeArguments 0).IsNullable) .IsNullable}}
{{- else if .IsMap -}}
	{{- setNullableMap (setNullable (toGoType (index .TypeArguments 0).ImportTypeName) (index .TypeArguments 0).IsNullable) (setNullable (toGoType (index .TypeArguments 1).ImportTypeName) (index .TypeArguments 1).IsNullable) .IsNullable}}
{{- else -}}
	{{- setNullable (toGoType .ImportTypeName) .IsNullable}}
{{- end}}`

const instantiateListTemplateString = `{{- if .IsNullable -}}
[]{{template "fieldDeclaration" (index .TypeArguments 0)}}
{{- end -}}`

const instantiateMapTemplateString = `{{- if .IsNullable -}}
map[{{template "fieldDeclaration" (index .TypeArguments 0)}}]{{template "fieldDeclaration" (index .TypeArguments 1)}}
{{- end -}}`

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

func (enum {{.TypeName}}) String() string {
	return fmt.Sprintf("{{.TypeName}}(%s)", enum.name)
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
	{{- range .Fields}}
	{{- if .ValueType.IsNullable}}
		if u.{{.FieldName}}.IsNull() {
			encoder.EncodeNull()
		} else {
			value := *u.{{.FieldName}}.Value
			{{template "rawFieldEncoder" .}}
		}
	{{- else}}
		{{template "rawFieldEncoder" .}}
	{{- end}}
	{{- end}}
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
	{{- range .Fields}}
	{{- if .ValueType.IsNullable}}
		if decoder.IsNextNull() {
			u.{{.FieldName}}.Clear()
		} else {
			{{template "rawFieldDecoder" .}}
		}
	{{- else}}
		{{template "rawFieldDecoder" .}}
	{{- end}}
	{{- end}}
	return nil
}

func (u *{{.TypeName}}) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u *{{.TypeName}}) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u {{.TypeName}}) String() string {
	parts := make([]string, {{len .Fields}})
	{{- range $i, $a := .Fields}}
	parts[{{$i}}] = fmt.Sprintf("{{.FieldName}}: %v", u.{{$a.FieldName}})
	{{- end}}
	return fmt.Sprintf("{{.TypeName}}(%s)", strings.Join(parts, ", "))
}

func (u *{{.TypeName}}) Equals(other *{{.TypeName}}) bool {
	{{- range .Fields}}
	{{- if .ValueType.IsPrimitive}}
		{{- if .ValueType.IsNullable}}
			{{- if eq .ValueType.ImportTypeName "binary"}}
			if (u.{{.FieldName}}.IsNull() != other.{{.FieldName}}.IsNull()) || (!bytes.Equal(*u.{{.FieldName}}.Value, *other.{{.FieldName}}.Value)) {
				return false
			}
			{{- else}}
			if (u.{{.FieldName}}.IsNull() != other.{{.FieldName}}.IsNull()) || (*u.{{.FieldName}}.Value != *other.{{.FieldName}}.Value){
				return false
			}
			{{- end}}
		{{- else}}
			{{- if eq .ValueType.ImportTypeName "binary"}}
			if !bytes.Equal(u.{{.FieldName}}, other.{{.FieldName}}) {
				return false
			}
			{{- else}}
			if u.{{.FieldName}} != other.{{.FieldName}} {
				return false
			}
			{{- end}}
		{{- end}}
	{{- else if .ValueType.IsList}}
		{{- if .ValueType.IsNullable}}
		if u.{{.FieldName}}.IsNull() != other.{{.FieldName}}.IsNull() {
			return false
		}
		
		{{.LowerFieldName}}LenThis := len(*u.{{.FieldName}}.Value)
		{{.LowerFieldName}}LenOther := len(*other.{{.FieldName}}.Value)
		if {{.LowerFieldName}}LenThis != {{.LowerFieldName}}LenOther {
			return false
		}

		for i := 0; i < {{.LowerFieldName}}LenThis; i++ {
			a := (*u.{{.FieldName}}.Value)[i]
			b := (*other.{{.FieldName}}.Value)[i]

			{{$typeArgument := (index .ValueType.TypeArguments 0)}}
			{{- if $typeArgument.IsNullable}}
				{{- if eq $typeArgument.ImportTypeName "binary"}}
				if (a.IsNull() != b.IsNull()) || (!bytes.Equal(*a.Value, *b.Value)) {
					return false
				} 
				{{- else}}
				if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
					return false
				} 
				{{- end}}
			{{- else}}
				{{- if eq $typeArgument.ImportTypeName "binary"}}
				if (!bytes.Equal(a, b)) {
					return false
				}
				{{- else}}
				if a != b {
					return false
				}
				{{- end}}
			{{- end}}
		} 
		{{- else}}
		{{.LowerFieldName}}LenThis := len(u.{{.FieldName}})
		{{.LowerFieldName}}LenOther := len(other.{{.FieldName}})
		if {{.LowerFieldName}}LenThis != {{.LowerFieldName}}LenOther {
			return false
		}

		for i := 0; i < {{.LowerFieldName}}LenThis; i++ {
			a := u.{{.FieldName}}[i]
			b := other.{{.FieldName}}[i]

			{{$typeArgument := (index .ValueType.TypeArguments 0)}}
			{{- if $typeArgument.IsNullable}}
				{{- if eq $typeArgument.ImportTypeName "binary"}}
				if (a.IsNull() != b.IsNull()) || (!bytes.Equal(*a.Value, *b.Value)) {
					return false
				} 
				{{- else}}
				if (a.IsNull() != b.IsNull()) || (*a.Value != *b.Value) {
					return false
				} 
				{{- end}}
			{{- else}}
				{{- if eq $typeArgument.ImportTypeName "binary"}}
				if (!bytes.Equal(a, b)) {
					return false
				}
				{{- else}}
				if a != b {
					return false
				}
				{{- end}}
			{{- end}}
		}
		{{- end}}
	{{- else if .ValueType.IsMap}}
		{{- if .ValueType.IsNullable}}
		{{- else}}
		{{- end}}
	{{- else if .ValueType.IsEnum}}
		{{- if .ValueType.IsNullable}}
		if (u.{{.FieldName}}.IsNull() != && other.{{.FieldName}}.IsNull()) || (u.{{.FieldName}}.Value.Index() != other.{{.FieldName}}.Value.Index()) {
			return false
		}
		{{- else}}
		if u.{{.FieldName}}.Index() != other.{{.FieldName}}.Index() {
			return false
		}
		{{- end}}
	{{- else if .ValueType.IsType}}
		{{- if .ValueType.IsNullable}}
		if (u.{{.FieldName}}.IsNull() != other.{{.FieldName}}.IsNull()) || (!u.{{.FieldName}}.Equals(other.{{.FieldName}})) {
			return false
		}
		{{- else}}
		if !u.{{.FieldName}}.Equals(other.{{.FieldName}}) {
			return false
		}
		{{- end}}
	{{- end}}
	{{- end}}
	return true
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

func (u *{{.TypeName}}) Decode(reader io.Reader) error {
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

func (u *{{.TypeName}}) MustDecode(reader io.Reader) {
	err := u.Decode(reader)
	if err != nil {
		panic(err)
	}
}

func (u *{{.TypeName}}) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	return u.Decode(reader)
}

func (u {{.TypeName}}) String() string {
	value := "not-set"
	if u.fieldIndex != -1 {
		switch u.fieldIndex {
		{{range .Fields}}
		case {{.FieldIndex}}:
			value = fmt.Sprintf("%d: %v", u.fieldIndex, u.value)
		{{end}}
		}
	}

	return fmt.Sprintf("{{.TypeName}}(%s)", value)
}
`

var enumTemplate, structTemplate, unionTemplate, rawFieldEncoder, rawFieldDecoder, primitiveEncoder, primitiveDecoder, fieldDeclarationTemplate, instantiateList, instantiateMap *template.Template

func init() {
	funcMap := template.FuncMap{
		"capitalizeFirst":          capitalizeFirstFunc,
		"setNullable":              setNullableFunc,
		"setNullableList":          setNullableListFunc,
		"setNullableMap":           setNullableMapFunc,
		"toGoType":                 toGoTypeFunc,
		"stripRuntimeNullableDecl": stripRuntimeNullableDeclFunc,
		"getEncoderVariableName":   getEncoderVariableNameFunc,
		"getDecoderVariableName":   getDecoderVariableNameFunc,
	}
	rawFieldEncoder = parseTemplateWithFunc("rawFieldEncoder", rawFieldEncoderTemplateString, funcMap)
	primitiveEncoder = parseTemplateWithFunc("encodePrimitive", encodePrimitiveTemplateString, funcMap)
	fieldDeclarationTemplate = parseTemplateWithFunc("fieldDeclaration", fieldDeclarationTemplateString, funcMap)
	instantiateList = parseTemplate("instantiateList", instantiateListTemplateString)
	instantiateMap = parseTemplate("instantiateMap", instantiateMapTemplateString)

	rawFieldDecoder = parseTemplateWithFunc("rawFieldDecoder", rawFieldDecoderTemplateString, funcMap)
	primitiveDecoder = parseTemplateWithFunc("decodePrimitive", decodePrimitiveTemplateString, funcMap)

	enumTemplate = parseTemplate("enum", enumTemplateString)
	structTemplate = parseTemplateWithFunc("struct", structTemplateString, funcMap)
	unionTemplate = parseTemplateWithFunc("union", unionTemplateString, funcMap)

	structTemplate, _ = structTemplate.AddParseTree("fieldDeclaration", fieldDeclarationTemplate.Tree)
	structTemplate, _ = structTemplate.AddParseTree("rawFieldEncoder", rawFieldEncoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("rawFieldDecoder", rawFieldDecoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("instantiateList", instantiateList.Tree)
	structTemplate, _ = structTemplate.AddParseTree("instantiateMap", instantiateMap.Tree)
	structTemplate, _ = structTemplate.AddParseTree("encodePrimitive", primitiveEncoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("decodePrimitive", primitiveDecoder.Tree)

	unionTemplate, _ = unionTemplate.AddParseTree("fieldDeclaration", fieldDeclarationTemplate.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("rawFieldEncoder", rawFieldEncoder.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("rawFieldDecoder", rawFieldDecoder.Tree)
	structTemplate, _ = structTemplate.AddParseTree("instantiateMap", instantiateMap.Tree)
	unionTemplate, _ = unionTemplate.AddParseTree("instantiateList", instantiateList.Tree)
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
