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
			// err := b.generateEnum(t, file)
			// if err != nil {
			// 	return "", err
			// }

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

func (b *Builder) generateStruct(t *SchemaTypeDefinition) (sourceCode string, err error) {

	pkgPath := (*b.typeRegistry)[t.Id].ImportPath
	sb := new(strings.Builder)

	// Create fields
	fields := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		fields[i] = fmt.Sprintf("%s %s", field.GoName, b.GetGoType(field.Type, pkgPath))
	}

	// write struct
	sb.WriteString(fmt.Sprintf("type %s struct {%s}", t.Name, strings.Join(fields, "\n")))

	b.writeSerializeMethod(sb, t)
	b.writeMustSerializeMethod(sb, t)
	b.writeMergeFromMethod(sb, t, pkgPath)
	b.writeMergeUsing(sb, t)

	return sb.String(), nil
}

// func (b *Builder) generateEnum(t *SchemaTypeDefinition, file *File) error {

// 	// Create Enum struct
// 	file.Type().Id(t.Name).Struct(
// 		Id("index").Uint8(),
// 		Id("name").String(),
// 	)

// 	lowerCamel := strcase.ToLowerCamel(t.Name)
// 	pickerName := fmt.Sprintf("%sPicker", lowerCamel)

// 	// create enum picker struct
// 	file.Type().Id(pickerName).Struct()

// 	// create enum picker instance
// 	file.Var().Id(fmt.Sprintf("%sPicker", t.Name)).Id(pickerName).Op("=").Id(pickerName).Block()

// 	//write each item in enum
// 	for _, field := range t.Fields {
// 		file.Var().Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name))).Id(t.Name).Op("=").Id(t.Name).Values(Dict{
// 			Id("index"): Lit(field.Index),
// 			Id("name"):  Lit(field.Name),
// 		})
// 	}

// 	// Write Picker method for each field
// 	for _, field := range t.Fields {
// 		file.Func().Params(Id(pickerName)).Id(strcase.ToCamel(field.Name)).Params().Params(Id(t.Name)).Block(
// 			Return(Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name)))),
// 		)
// 	}

// 	// Write Index() method
// 	file.Func().Params(Id("e").Id(t.Name)).Id("Index").Params().Params(Uint8()).Block(
// 		Return(Id("e").Dot("index")),
// 	)

// 	// Write Name() method
// 	file.Func().Params(Id("e").Id(t.Name)).Id("Name").Params().Params(String()).Block(
// 		Return(Id("e").Dot("name")),
// 	)

// 	// Write ByIndex method
// 	byIndexStmts := make([]Code, 0)
// 	for _, field := range t.Fields {
// 		byIndexStmts = append(byIndexStmts, Case(
// 			Lit(field.Index),
// 		).Block(
// 			Return(Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name)))),
// 		))
// 	}

// 	byIndexStmts = append(byIndexStmts, Default().Block(
// 		Panic(Qual("fmt", "Sprintf").Call(Lit(fmt.Sprintf("%s with index %%v not found", t.Name)), Id("index"))),
// 	))

// 	file.Func().Params(Id(pickerName)).Id("ByIndex").Params(Id("index").Uint8()).Params(Id(t.Name)).Block(
// 		Switch(Id("index")).Block(byIndexStmts...),
// 	)

// 	byNameStmts := make([]Code, 0)
// 	for _, field := range t.Fields {
// 		byNameStmts = append(byNameStmts, Case(
// 			Lit(field.Name),
// 		).Block(
// 			Return(Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name)))),
// 		))
// 	}

// 	byNameStmts = append(byNameStmts, Default().Block(
// 		Panic(Qual("fmt", "Sprintf").Call(Lit(fmt.Sprintf("%s with name %%v not found", t.Name)), Id("name"))),
// 	))

// 	file.Func().Params(Id(pickerName)).Id("ByName").Params(Id("name").String()).Params(Id(t.Name)).Block(
// 		Switch(Id("name")).Block(byNameStmts...),
// 	)

// 	return nil
// }

func (b *Builder) writeSerializeMethod(sb *strings.Builder, t *SchemaTypeDefinition) {

	stmts := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		stmts[i] = b.WriteEncodeField(field)
	}

	body := fmt.Sprintf(
		`
func (u *%s) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	var err error

	%s

	return buf.Bytes(), nil
}%s`, t.Name, strings.Join(stmts, "\n"), "\n")

	sb.WriteString(body)

}

func (b *Builder) writeMustSerializeMethod(sb *strings.Builder, t *SchemaTypeDefinition) {
	sb.WriteString(fmt.Sprintf(
		`
	func (u *%s) MustSerialize() []byte {
		buf, err := u.Serialize()
		if err != nil {
			panic(err)
		}

		return buf
	}%s
	`, t.Name, "\n"))
}

func (b *Builder) writeMergeFromMethod(sb *strings.Builder, t *SchemaTypeDefinition, pkgName string) {
	stmts := make([]string, len(t.Fields))
	isAnyNullable := false
	for i, field := range t.Fields {
		stmts[i] = b.WriteDecodeField(field, pkgName)
		if field.Type.Nullable {
			isAnyNullable = true
		}
	}

	var isNextNilStmt string
	if isAnyNullable {
		isNextNilStmt = "var isNextNil bool"
	}

	body := fmt.Sprintf(
		`
		func (u *%s) MergeFrom(buffer []byte) error {
			reader := bytes.NewBuffer(buffer)
			decoder := msgpack.NewDecoder(reader)
			var err error
			%s
			%s

			return nil
		}%s
	`, t.Name, isNextNilStmt, strings.Join(stmts, "\n"), "\n")

	sb.WriteString(body)
}

func (b *Builder) writeMergeUsing(sb *strings.Builder, t *SchemaTypeDefinition) {
	stmts := make([]string, len(t.Fields))
	for i, field := range t.Fields {
		stmts[i] = fmt.Sprintf("u.%[1]s = other.%[1]s", field.GoName)
	}

	body := fmt.Sprintf(
		`
		func (u *%[1]s) MergeUsing(other *%[1]s) {
			%s
		}%s
	`, t.Name, strings.Join(stmts, "\n"), "\n")

	sb.WriteString(body)
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
