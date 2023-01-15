package generator

import "encoding/json"

// NexemaDefinition represents an analyzed and built list of Ast.
// This type is next sent to a plugin to generate source code.
type NexemaDefinition struct {
	Version  int          `json:"version"`  // The Nexema specification's version used to build this definition
	Hashcode uint64       `json:"hashcode"` // Hashcode of the current generation
	Files    []NexemaFile `json:"files"`    // A list of nexema files
}

// NexemaFile represents a .nex file with many NexemaTypeDefinition's
type NexemaFile struct {
	Name  string                 `json:"name"`  // The relative path to the file. Its relative to nexema.yaml
	Types []NexemaTypeDefinition `json:"types"` // The list of types declared in this file
}

// NexemaTypeDefinition contains information about a parsed Nexema type
type NexemaTypeDefinition struct {
	Id            string                      `json:"id"`            // An id generated for this type. It's: sha256(NexemaFilePath-TypeName)
	Name          string                      `json:"name"`          // The name of the type
	Modifier      string                      `json:"modifier"`      // The type's modifier
	Documentation []string                    `json:"documentation"` // The documentation for the type
	Fields        []NexemaTypeFieldDefinition `json:"fields"`        // The list of fields declared in this type
}

// NexemaTypeFieldDefinition contains information about a field declared in a Nexema type
type NexemaTypeFieldDefinition struct {
	Index        int64           `json:"index"`        // The field's index
	Name         string          `json:"name"`         // The field's name
	Metadata     map[string]any  `json:"metadata"`     // The field's metadata
	DefaultValue any             `json:"defaultValue"` // The field's default value
	Type         NexemaValueType `json:"type"`         // The field's value type
}

// BaseNexemaValueType is a base struct for every Nexema's type
type BaseNexemaValueType struct {
	Kind     string `json:"$type"`    // NexemaPrimitiveValueType or NexemaTypeValueType
	Nullable bool   `json:"nullable"` // True if the type is nullable
}

type NexemaValueType interface {
	t() // just to allow NexemaPrimitiveValueType and NexemaTypeValueType be part of this
}

// NexemaPrimitiveValueType represents the value type of a NexemaTypeFieldDefinition
// which has a primitive type.
type NexemaPrimitiveValueType struct {
	BaseNexemaValueType `json:",inline"`
	Primitive           string            `json:"primitive"`     // Value's type primitive
	TypeArguments       []NexemaValueType `json:"typeArguments"` // Any generic type argument
}

// NexemaTypeValueType represents the value type of a NexemaTypeFieldDefinition
// which has another Nexema type as value type.
type NexemaTypeValueType struct {
	BaseNexemaValueType `json:",inline"`
	TypeId              string  `json:"typeId"`      // The imported type's id
	ImportAlias         *string `json:"importAlias"` // the import alias, if specified
}

func (NexemaPrimitiveValueType) t() {}
func (NexemaTypeValueType) t()      {}

func (n *NexemaTypeFieldDefinition) UnmarshalJSON(b []byte) error {
	m := map[string]interface{}{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	n.Index = int64(m["index"].(float64))
	n.Name = m["name"].(string)
	n.Metadata = m["metadata"].(map[string]any)
	n.DefaultValue = m["defaultValue"]

	rawValueType := m["type"]
	if rawValueType != nil {
		valueType := rawValueType.(map[string]any)
		if valueType["$type"] == "NexemaPrimitiveValueType" {
			n.Type = getNexemaPrimitiveValueType(valueType)
		} else if valueType["$type"] == "NexemaTypeValueType" {
			importAlias := valueType["importAlias"]
			nexemaValueType := NexemaTypeValueType{
				BaseNexemaValueType: BaseNexemaValueType{
					Kind:     "NexemaTypeValueType",
					Nullable: valueType["nullable"].(bool),
				},
				TypeId: valueType["typeId"].(string),
			}

			if importAlias != nil {
				nexemaValueType.ImportAlias = stringPtr(importAlias.(string))
			}

			n.Type = nexemaValueType
		}
	}

	return nil
}

func getNexemaPrimitiveValueType(m map[string]any) NexemaPrimitiveValueType {
	primitive := NexemaPrimitiveValueType{
		BaseNexemaValueType: BaseNexemaValueType{
			Kind:     "NexemaPrimitiveValueType",
			Nullable: m["nullable"].(bool),
		},
		Primitive:     m["primitive"].(string),
		TypeArguments: make([]NexemaValueType, 0),
	}

	typeArgumentsArr := m["typeArguments"].([]interface{})
	if len(typeArgumentsArr) > 0 {
		for _, ta := range typeArgumentsArr {
			primitive.TypeArguments = append(primitive.TypeArguments, getNexemaPrimitiveValueType(ta.(map[string]any)))
		}
	}

	return primitive
}
