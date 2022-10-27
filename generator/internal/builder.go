package internal

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
)

type Builder struct {
	input        *GenerateInput
	packageName  string
	files        map[string]string
	typeRegistry *TypeRegistry
}

func NewBuilder(input *GenerateInput) *Builder {
	builder := &Builder{
		input:        input,
		files:        make(map[string]string),
		typeRegistry: NewTypeRegistry(),
	}

	builder.packageName = input.Options["packageName"].(string)
	builder.typeRegistry.Fill(input.RootPackage, builder.packageName, input.Output)
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

func (b *Builder) generateFile(fd *FileDeclaration) (code string, err error) {
	file := NewFile(filepath.Base(b.packageName))

	for _, t := range fd.Types {
		switch t.Modifier {
		case "struct":
			err := b.generateStruct(t, file)
			if err != nil {
				return "", err
			}

		case "union":
			break

		case "enum":
			err := b.generateEnum(t, file)
			if err != nil {
				return "", err
			}

		default:
			return "", fmt.Errorf("unknown type modifier %s", t.Modifier)
		}
	}

	buf := new(bytes.Buffer)
	err = file.Render(buf)
	if err != nil {
		fmt.Println("Printing file:::::")

		fmt.Printf("%#v\n", file)

		fmt.Println("========================")

		return "", err
	}

	return buf.String(), nil
}

func (b *Builder) generateStruct(t *SchemaTypeDefinition, file *File) error {

	pkgPath := (*b.typeRegistry)[t.Id].ImportPath

	// Create fields
	fields := make([]Code, len(t.Fields))
	for i, field := range t.Fields {
		stmt := Id(field.GoName)

		b.WriteField(stmt, field.Type, pkgPath)

		fields[i] = stmt
	}

	// Write fields
	file.Type().Id(t.Name).Struct(fields...)

	// Write Serialize method
	b.writeSerializeMethod(file, t)

	// Write MustSerialize method
	writeMustSerializeMethod(file, t)

	// Write MergeFrom method
	b.writeMergeFromMethod(file, t, pkgPath)

	// Write MergeUsing method
	writeMergeUsing(file, t)

	return nil
}

func (b *Builder) generateEnum(t *SchemaTypeDefinition, file *File) error {

	// Create Enum struct
	file.Type().Id(t.Name).Struct(
		Id("index").Uint8(),
		Id("name").String(),
	)

	lowerCamel := strcase.ToLowerCamel(t.Name)
	pickerName := fmt.Sprintf("%sPicker", lowerCamel)

	// create enum picker struct
	file.Type().Id(pickerName).Struct()

	// create enum picker instance
	file.Var().Id(fmt.Sprintf("%sPicker", t.Name)).Id(pickerName).Op("=").Id(pickerName).Block()

	//write each item in enum
	for _, field := range t.Fields {
		file.Var().Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name))).Id(t.Name).Op("=").Id(t.Name).Values(Dict{
			Id("index"): Lit(field.Index),
			Id("name"):  Lit(field.Name),
		})
	}

	// Write Picker method for each field
	for _, field := range t.Fields {
		file.Func().Params(Id(pickerName)).Id(strcase.ToCamel(field.Name)).Params().Params(Id(t.Name)).Block(
			Return(Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name)))),
		)
	}

	// Write Index() method
	file.Func().Params(Id("e").Id(t.Name)).Id("Index").Params().Params(Uint8()).Block(
		Return(Id("e").Dot("index")),
	)

	// Write Name() method
	file.Func().Params(Id("e").Id(t.Name)).Id("Name").Params().Params(String()).Block(
		Return(Id("e").Dot("name")),
	)

	// Write ByIndex method
	byIndexStmts := make([]Code, 0)
	for _, field := range t.Fields {
		byIndexStmts = append(byIndexStmts, Case(
			Lit(field.Index),
		).Block(
			Return(Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name)))),
		))
	}

	byIndexStmts = append(byIndexStmts, Default().Block(
		Panic(Qual("fmt", "Sprintf").Call(Lit(fmt.Sprintf("%s with index %%v not found", t.Name)), Id("index"))),
	))

	file.Func().Params(Id(pickerName)).Id("ByIndex").Params(Id("index").Uint8()).Params(Id(t.Name)).Block(
		Switch(Id("index")).Block(byIndexStmts...),
	)

	byNameStmts := make([]Code, 0)
	for _, field := range t.Fields {
		byNameStmts = append(byNameStmts, Case(
			Lit(field.Name),
		).Block(
			Return(Id(fmt.Sprintf("%s%s", lowerCamel, strcase.ToCamel(field.Name)))),
		))
	}

	byNameStmts = append(byNameStmts, Default().Block(
		Panic(Qual("fmt", "Sprintf").Call(Lit(fmt.Sprintf("%s with name %%v not found", t.Name)), Id("name"))),
	))

	file.Func().Params(Id(pickerName)).Id("ByName").Params(Id("name").String()).Params(Id(t.Name)).Block(
		Switch(Id("name")).Block(byNameStmts...),
	)

	return nil
}

func (b *Builder) writeSerializeMethod(file *File, t *SchemaTypeDefinition) {
	body := []Code{
		Id("buf").Op(":=").Id("new").Params(Qual("bytes", "Buffer")),
		Id("writer").Op(":=").Qual("github.com/vmihailenco/msgpack/v5", "NewEncoder").Params(Id("buf")),
		Var().Id("err").Id("error"),
	}

	for _, field := range t.Fields {
		stmts := b.WriteEncodeField(field)
		for _, stmt := range stmts {
			body = append(body, stmt)
		}
	}

	// append return
	body = append(body, Return().List(Id("buf").Dot("Bytes").Call(), Nil()))

	file.Func().Params(
		Id("u").Op("*").Id(t.Name), //receiver
	).Id("Serialize").Params().Params(Index().Byte(), Error()).Block(body...)
}

func writeMustSerializeMethod(file *File, t *SchemaTypeDefinition) {
	file.Func().Params(
		Id("u").Op("*").Id(t.Name), //receiver
	).Id("MustSerialize").Params().Index().Byte().Block(
		List(Id("buf"), Id("err")).Op(":=").Id("u").Dot("Serialize").Call(),
		If(Id("err").Op("!=").Nil()).Block(
			Panic(Id("err")),
		),

		Return(Id("buf")),
	)
}

func (b *Builder) writeMergeFromMethod(file *File, t *SchemaTypeDefinition, pkgName string) {
	body := []Code{
		Id("reader").Op(":=").Qual("bytes", "NewBuffer").Call(Id("buffer")),
		Id("decoder").Op(":=").Qual("github.com/vmihailenco/msgpack/v5", "NewDecoder").Params(Id("reader")),
		Var().Id("err").Id("error"),
	}

	for _, field := range t.Fields {
		stmts := b.WriteDecodeField(field, pkgName)
		for _, stmt := range stmts {
			body = append(body, stmt)
		}
	}

	// append return
	body = append(body, Return(Nil()))

	file.Func().Params(
		Id("u").Op("*").Id(t.Name), //receiver
	).Id("MergeFrom").Params(Id("buffer").Index().Byte()).Params(Error()).Block(body...)
}

func writeMergeUsing(file *File, t *SchemaTypeDefinition) {
	body := []Code{}

	for _, field := range t.Fields {
		body = append(body, Id("u").Dot(field.GoName).Op("=").Id("other").Dot(field.GoName))
	}

	// append return
	body = append(body, Return().Nil())

	file.Func().Params(
		Id("u").Op("*").Id(t.Name), //receiver
	).Id("MergeUsing").Params(Id("other").Op("*").Id(t.Name)).Params(Error()).Block(body...)
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
