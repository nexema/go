package internal

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateStruct(t *testing.T) {
	t.Skip()
	input := GenerateInput{
		Options: map[string]interface{}{
			"packageName": "test",
		},
		RootPackage: DeclarationNode{
			Name:      "root",
			IsPackage: true,
			Value: PackageDeclaration{
				PackageName: "root",
				Path:        "root",
			},
			Children: []DeclarationNode{
				{
					Name:      "a",
					IsPackage: false,
					Children:  nil,
					Value: FileDeclaration{
						FileName: "a.mpack",
						Path:     "root/a.mpack",
						Id:       "123456789",
						Types: []SchemaTypeDefinition{
							{
								Id:       "454545",
								Name:     "MyType",
								Modifier: "struct",
								Fields: []TypeFieldDefinition{
									{
										Name:  "a",
										Index: 1,
										Type: SchemaFieldType{
											Primitive: "map",
											TypeArguments: []SchemaFieldType{
												{Primitive: "string"},
												{Primitive: "int8"},
											},
											Nullable: false,
										},
									},
									{
										Name:  "b",
										Index: 2,
										Type: SchemaFieldType{
											Primitive: "list",
											TypeArguments: []SchemaFieldType{
												{Primitive: "uint64"},
											},
											Nullable: false,
										},
									},
									{
										Name:  "c",
										Index: 3,
										Type: SchemaFieldType{
											Primitive: "binary",
											Nullable:  false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	builder := NewBuilder(&input)
	err := builder.Build()
	require.Nil(t, err)
}

func TestGenerateEnum(t *testing.T) {
	t.Skip()
	input := GenerateInput{
		Options: map[string]interface{}{
			"packageName": "test",
		},
		RootPackage: DeclarationNode{
			Name:      "root",
			IsPackage: true,
			Value: PackageDeclaration{
				PackageName: "root",
				Path:        "root",
			},
			Children: []DeclarationNode{
				{
					Name:      "a",
					IsPackage: false,
					Children:  nil,
					Value: FileDeclaration{
						FileName: "a.mpack",
						Path:     "root/a.mpack",
						Id:       "123456789",
						Types: []SchemaTypeDefinition{
							{
								Id:       "454545",
								Name:     "MyType",
								Modifier: "struct",
								Fields: []TypeFieldDefinition{
									{
										Name:  "a",
										Index: 1,
										Type: SchemaFieldType{
											Primitive: "map",
											TypeArguments: []SchemaFieldType{
												{Primitive: "string"},
												{Primitive: "int8"},
											},
											Nullable: false,
										},
									},
									{
										Name:  "b",
										Index: 2,
										Type: SchemaFieldType{
											Primitive: "list",
											TypeArguments: []SchemaFieldType{
												{Primitive: "uint64"},
											},
											Nullable: false,
										},
									},
									{
										Name:  "c",
										Index: 3,
										Type: SchemaFieldType{
											Primitive: "binary",
											Nullable:  false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	builder := NewBuilder(&input)
	err := builder.Build()
	require.Nil(t, err)
}

func TestGenerateRaw(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString(raw)
	generateInput := GenerateInput{}
	json.Unmarshal(data, &generateInput)

	generateInput.Options["packageName"] = "github.com/example/test_files"
	generateInput.Output = "../../example/test_files_generated"

	builder := NewBuilder(&generateInput)
	err := builder.Build()
	require.Nil(t, err)
}
