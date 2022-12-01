package internal

import (
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
)

type Builder struct {
	input          *GenerateInput
	packagePath    string
	packageName    string
	files          map[string]string
	typeRegistry   *TypeRegistry
	currentContext *generationContext
}

type generationContext struct {
	fileId  string
	imports map[string]string
}

func NewBuilder(input *GenerateInput) *Builder {
	builder := &Builder{
		input:        input,
		files:        make(map[string]string),
		typeRegistry: NewTypeRegistry(),
	}

	builder.packagePath = input.Options["packageName"].(string)
	builder.packageName = path.Base(builder.packagePath)
	builder.typeRegistry.Fill(input.RootPackage, builder.packagePath, input.Output)
	builder.sanitize()

	return builder
}

func (b *Builder) Build() error {

	err := b.generateNode(b.input.RootPackage)
	if err != nil {
		return err
	}

	for name, content := range b.files {
		err := os.MkdirAll(filepath.Dir(name), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(fmt.Sprintf("%s.go", name), []byte(content), 0644)
		if err != nil {
			return err
		}

		fmt.Printf("File: %v\n", name)
		fmt.Println(content)
		fmt.Println("=================================")
	}

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
			file := child.Value.(*FileDeclaration)
			code, err := b.generateFile(file)
			if err != nil {
				return err
			}
			b.files[filepath.Join(b.input.Output, file.Path)] = code
		}
	}

	return nil
}

func (b *Builder) generateFile(fd *FileDeclaration) (sourceCode string, err error) {

	b.newGenerationContext(fd.Id)

	pkgDeclaration := filepath.Base(b.packagePath)
	stringBuilder := new(strings.Builder)
	stringBuilder.WriteString(fmt.Sprintf("package %s\n", pkgDeclaration))

	// generate types
	typesSourceCode := make([]string, len(fd.Types))
	for i, t := range fd.Types {
		switch t.Modifier {
		case "struct":
			sourceCode, err := b.generateStruct(t)
			if err != nil {
				return "", err
			}

			typesSourceCode[i] = sourceCode
			b.currentContext.MustImport("github.com/messagepack-schema/go/runtime/msgpack")
			b.currentContext.MustImport("bytes")

		case "union":
			break

		case "enum":
			sourceCode, err := b.generateEnum(t)
			if err != nil {
				return "", err
			}

			typesSourceCode[i] = sourceCode

		default:
			return "", fmt.Errorf("unknown type modifier %s", t.Modifier)
		}
	}

	// write imports
	for importName, importAlias := range b.currentContext.imports {
		if importAlias != "" {
			stringBuilder.WriteString(fmt.Sprintf("import %s \"%s\"\n", importAlias, importName))
		} else {
			stringBuilder.WriteString(fmt.Sprintf("import \"%s\"\n", importName))
		}
	}

	// write types
	stringBuilder.WriteString(strings.Join(typesSourceCode, "\n"))

	// "compile" stringBuilder
	raw := stringBuilder.String()
	fmt.Printf("raw: %v\n", raw)

	// format source code
	buf, err := format.Source([]byte(raw))
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (b *Builder) sanitize() {
	b.sanitizePkg(b.input.RootPackage)
}

func (b *Builder) sanitizePkg(t DeclarationNode) {
	for _, n := range t.Children {
		if n.IsPackage {
			b.sanitizePkg(n)
		} else {
			fd := n.Value.(*FileDeclaration)
			for _, td := range fd.Types {
				for _, field := range td.Fields {
					field.GoName = strcase.ToCamel(field.Name)
				}
			}
		}
	}
}

func (b *Builder) newGenerationContext(fileId string) {
	b.currentContext = &generationContext{
		fileId:  fileId,
		imports: make(map[string]string),
	}
}

func (b *generationContext) MustImport(importName string) {
	_, ok := b.imports[importName]
	if !ok {
		b.imports[importName] = ""
	}
}

func (b *generationContext) MustImportAs(importName string, alias string) {
	_, ok := b.imports[importName]
	if !ok {
		b.imports[importName] = alias
	}
}
