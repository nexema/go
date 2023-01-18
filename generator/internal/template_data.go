package generator

type EnumTemplateData struct {
	TypeName      string
	LowerTypeName string
	Fields        []EnumFieldTemplateData
	Documentation []string
	HasDocs       bool
}

type EnumFieldTemplateData struct {
	LowerName string
	Name      string
	Index     int64
}

type StructTemplateData struct {
	TypeName      string
	Documentation []string
	HasDocs       bool
	Fields        []TypeFieldTemplateData
}

type UnionTemplateData struct {
	TypeName      string
	Documentation []string
	HasDocs       bool
	LowerName     string
	Fields        []TypeFieldTemplateData
}

type TypeFieldTemplateData struct {
	FieldName       string // FieldName is field's name but in Go (CamelCase)
	LowerFieldName  string // LowerFieldName is field's name in Go, but lowerCamelCase
	FieldIndex      int64
	IsFromUnion     bool
	ValueType       TypeFieldValueKindTemplate
	DefaultValue    interface{}
	HasDefaultValue bool
}

// TypeFieldValueKindTemplate contains information about the field's value type.
// ex: string, MyStruct, MyEnum, list(string), map(int64, MyEnum), etc
type TypeFieldValueKindTemplate struct {
	IsNullable     bool
	IsList         bool
	IsMap          bool
	IsEnum         bool
	IsPrimitive    bool
	IsType         bool
	ImportTypeName string // Field type's name but with import, if neccessary (ex: models.User)
	IsNumeric      bool
	TypeArguments  []TypeFieldValueKindTemplate
}

type StructFieldTemplateData struct {
	FieldName      string // FieldName is field's name but in Go (CamelCase)
	LowerFieldName string // LowerFieldName is field's name in Go, but lowerCamelCase
	Index          int64
	TypeName       string
	IsNullable     bool
	ImportTypeName string // the same as TypeName but with import e.x: models.User
	TypeId         string
	FieldDef       *NexemaTypeFieldDefinition
}
