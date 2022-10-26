package internal

import (
	"bytes"
	"fmt"
	"io"

	. "github.com/dave/jennifer/jen"
)

type Builder struct {
	input       *GenerateInput
	packageName string
	files       map[string]string
	writer      *io.Writer
}

func NewBuilder(input *GenerateInput) *Builder {
	builder := &Builder{
		input: input,
		files: make(map[string]string),
	}

	builder.packageName = input.Options["packageName"].(string)

	return builder
}

func (b *Builder) SetWriter(w *io.Writer) {
	b.writer = w
}

func (b *Builder) Build() error {

	err := b.generateNode(b.input.RootPackage)
	if err != nil {
		return err
	}

	for name, content := range b.files {
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
			err := b.generateStruct(t, file)
			if err != nil {
				return "", err
			}

		case "union":
			break

		case "enum":
			break

		default:
			return "", fmt.Errorf("unknown type modifier %s", t.Modifier)
		}
	}

	buf := new(bytes.Buffer)
	err = file.Render(buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (b *Builder) generateStruct(t SchemaTypeDefinition, file *File) error {

	// Create fields
	fields := make([]Code, len(t.Fields))
	for i, field := range t.Fields {
		stmt := Id(field.Name)

		WriteField(stmt, field.Type)

		fields[i] = stmt
	}

	// Write fields
	file.Type().Id(t.Name).Struct(fields...)

	// Write Serialize method
	writeSerializeMethod(file, t)

	// Write MustSerialize method
	writeMustSerializeMethod(file, t)

	// Write MergeFrom method
	writeMergeFromMethod(file, t)

	return nil
}

func writeSerializeMethod(file *File, t SchemaTypeDefinition) {
	body := []Code{
		Id("buf").Op(":=").Id("new").Params(Qual("bytes", "Buffer")),
		Id("writer").Op(":=").Qual("github.com/vmihailenco/msgpack/v5", "NewEncoder").Params(Id("buf")),
		Var().Id("err").Id("error"),
	}

	for _, field := range t.Fields {
		stmts := WriteEncodeField(field)
		for _, stmt := range stmts {
			body = append(body, stmt)
		}
	}

	// append return
	body = append(body, Return().List(Id("buf").Dot("Bytes").Call(), Nil()))

	file.Func().Params(
		Id("u").Op("*").Id(t.Name), //receiver
	).Id("Serialize").Params().Params(Index().Byte(), Err()).Block(body...)
}

func writeMustSerializeMethod(file *File, t SchemaTypeDefinition) {
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

func writeMergeFromMethod(file *File, t SchemaTypeDefinition) {
	body := []Code{
		Id("reader").Op(":=").Qual("bytes", "NewBuffer").Call(Id("buffer")),
		Id("decoder").Op(":=").Qual("github.com/vmihailenco/msgpack/v5", "NewDecoder").Params(Id("reader")),
		Var().Id("err").Id("error"),
	}

	for _, field := range t.Fields {
		stmts := WriteEncodeField(field)
		for _, stmt := range stmts {
			body = append(body, stmt)
		}
	}

	// append return
	body = append(body, Return().List(Id("buf").Dot("Bytes").Call(), Nil()))

	file.Func().Params(
		Id("u").Op("*").Id(t.Name), //receiver
	).Id("MergeFrom").Params(Id("buffer").Index().Byte()).Params(Index().Byte(), Err()).Block(body...)
}
