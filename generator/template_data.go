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
}
