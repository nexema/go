package internal

import (
	"fmt"
	"path/filepath"
	"strings"
)

type TypeRegistry map[string]*SchemaTypeReference
type SchemaTypeReference struct {
	Definition *SchemaTypeDefinition
	ImportPath string
}

func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{}
}

func (tr *TypeRegistry) Fill(root DeclarationNode, packageName, relativePath string) {
	if !root.IsPackage {
		panic("cannot fill TypeRegistry with non-package root DeclarationNode")
	}

	for _, entry := range root.Children {
		if entry.IsPackage {
			tr.Fill(entry, packageName, relativePath)
		} else {
			file := entry.Value.(*FileDeclaration)
			for _, t := range file.Types {
				f := strings.ReplaceAll(filepath.Dir(file.Path), "\\", "/")
				importPath := filepath.Clean(fmt.Sprintf("%s/%s/%s", packageName, filepath.Base(relativePath), f))
				(*tr)[t.Id] = &SchemaTypeReference{
					Definition: t,
					ImportPath: importPath,
				}
			}
		}
	}
}
