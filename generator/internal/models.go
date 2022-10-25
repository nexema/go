package internal

import (
	"encoding/json"
)

type GenerateInput struct {
	RootFolder  string                 `json:"root"`
	Output      string                 `json:"outputPath"`
	Options     map[string]interface{} `json:"options"`
	RootPackage DeclarationNode        `json:"packages"`
}

type DeclarationNode struct {
	Name      string            `json:"name"`
	Value     interface{}       `json:"value"`
	Children  []DeclarationNode `json:"children"`
	IsPackage bool              `json:"-"`
}

func (d *DeclarationNode) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)

	if err != nil {
		return err
	}

	value := m["value"].(map[string]interface{})
	_, ok := value["fileName"]
	if ok {
		fd := struct {
			Name     string            `json:"name"`
			Value    FileDeclaration   `json:"value"`
			Children []DeclarationNode `json:"children"`
		}{}

		err = json.Unmarshal(b, &fd)
		if err != nil {
			return err
		}

		d.Name = fd.Name
		d.Children = fd.Children
		d.Value = fd.Value
	} else {
		pd := struct {
			Name     string             `json:"name"`
			Value    PackageDeclaration `json:"value"`
			Children []DeclarationNode  `json:"children"`
		}{}

		err = json.Unmarshal(b, &pd)
		if err != nil {
			return err
		}

		d.Name = pd.Name
		d.Children = pd.Children
		d.Value = pd.Value
		d.IsPackage = true
	}

	return nil
}

type PackageDeclaration struct {
	PackageName string `json:"packageName"`
	Path        string `json:"path"`
}

type FileDeclaration struct {
	FileName string                 `json:"fileName"`
	Path     string                 `json:"path"`
	Id       string                 `json:"id"`
	Types    []SchemaTypeDefinition `json:"types"`
	Imports  []string               `json:"imports"`
}

type SchemaTypeDefinition struct {
	Id       string                `json:"id"`
	Name     string                `json:"name"`
	Modifier string                `json:"modifier"`
	Fields   []TypeFieldDefinition `json:"fields"`
}

type TypeFieldDefinition struct {
	Name         string                 `json:"name"`
	Index        int                    `json:"index"`
	DefaultValue interface{}            `json:"defaultValue"`
	Metadata     map[string]interface{} `json:"metadata"`
	Type         SchemaFieldType        `json:"type"`
}

type SchemaFieldType struct {
	Primitive     string            `json:"primitive"`
	TypeName      *string           `json:"typeName"`
	Nullable      bool              `json:"nullable"`
	ImportId      string            `json:"importId"`
	TypeArguments []SchemaFieldType `json:"typeArguments"`
}
