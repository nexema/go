package generator

type EnumTemplateData struct {
	TypeName      string
	LowerTypeName string
	Fields        []EnumFieldTemplateData
}

type EnumFieldTemplateData struct {
	LowerName string
	Name      string
	Index     int64
}

type StructTemplateData struct {
	TypeName string
	Fields   []StructFieldTemplateData
}

type StructFieldTemplateData struct {
	Name           string
	Index          int64
	TypeName       string
	ImportTypeName string // the same as TypeName but with import e.x: models.User
	TypeId         string
	FieldDef       *NexemaTypeFieldDefinition
}

type UnionTemplateData struct {
	TypeName  string
	LowerName string
	Fields    []UnionFieldTemplateData
}

type UnionFieldTemplateData struct {
	FieldName      string
	FieldIndex     int64
	ImportTypeName string // the same as TypeName but with import e.x: models.User
	TypeId         string
	FieldDef       *NexemaTypeFieldDefinition
}
