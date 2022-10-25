package internal

import (
	"fmt"

	. "github.com/dave/jennifer/jen"
)

type Builder struct {
	input       *GenerateInput
	packageName string
	files       map[string]string
}

func NewBuilder(input *GenerateInput) *Builder {
	builder := &Builder{
		input: input,
		files: make(map[string]string),
	}

	builder.packageName = input.Options["packageName"].(string)

	return builder
}

func (b *Builder) Build() error {
	return nil
}

func (b *Builder) generateNode(node DeclarationNode) error {
	for _, child := range node.Children {
		if child.IsPackage {
			err := b.generateNode(child)
			if err != nil {
				return err
			}
		} else {
			file := child.Value.(FileDeclaration)
			code, err := b.generateFile(file)
			if err != nil {
				return err
			}

			b.files[file.Path] = code
		}
	}

	return nil
}

func (b *Builder) generateFile(fd FileDeclaration) (code string, err error) {
	file := NewFile(b.packageName)

	for _, t := range fd.Types {
		switch t.Modifier {
		case "struct":
			break

		case "union":
			break

		case "enum":
			break

		default:
			return "", fmt.Errorf("unknown type modifier %s", t.Modifier)
		}
	}

	return "", nil
}

func (b *Builder) generateStruct(t SchemaTypeDefinition, file *File) error {

	return nil
}
