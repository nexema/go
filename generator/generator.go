package generator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"go/scanner"
	"path/filepath"

	"github.com/iancoleman/strcase"
)

const (
	modifierEnum   string = "enum"
	modifierUnion  string = "union"
	modifierStruct string = "struct"
)

const (
	runtimeImport string = "github.com/nexema/go/runtime"
	fmtImport     string = "fmt"
	ioImport      string = "io"
)

const (
	importSentence string = "import ("
)

// Generator is the main entry point of the Nexema's Go generator
type Generator struct {
	sw     *sourceWriter
	config PluginConfig

	typeMapping map[string]typeMap
	fileMapping map[string]string

	currentFileImports map[string]any // its a map to avoid duplicated
	currentPkg         string
}

type typeMap struct {
	PackageName string
	TypeDef     *NexemaTypeDefinition
	ImportPath  string
}

var defaultGenerator *Generator

// NewGenerator creates a new Generator instance
func NewGenerator(cfg PluginConfig) *Generator {
	defaultGenerator = &Generator{
		sw:     newSourceWriter(),
		config: cfg,

		typeMapping:        make(map[string]typeMap),
		fileMapping:        make(map[string]string),
		currentFileImports: make(map[string]any, 0),
	}

	return defaultGenerator
}

// Generate takes an input buffer which represents a NexemaDefinition and outputs a list of
// files to be generated
func (g *Generator) Generate(input []byte) ([]*GeneratedFile, error) {
	// deserialize definition
	def := new(NexemaDefinition)
	err := json.Unmarshal(input, def)
	if err != nil {
		return nil, err
	}

	// first, create a mapping of types with their output path and package name
	g.generateMapping(def)

	// start generating
	generatedFiles := make([]*GeneratedFile, len(def.Files))
	for i, file := range def.Files {
		f, err := g.generateFile(&file)
		if err != nil {
			return nil, err
		}

		generatedFiles[i] = f
		println(string(f.Content))
		println("=======================")
		g.sw.reset()
	}

	return generatedFiles, nil
}

// generateFile generates the Go source code that represents f
func (g *Generator) generateFile(f *NexemaFile) (*GeneratedFile, error) {
	// write pkg name
	pkgName := g.fileMapping[f.Name]
	g.currentPkg = pkgName

	// generate each type
	for _, t := range f.Types {
		var err error
		if t.Modifier == modifierEnum {
			g.addImport(fmtImport)
			err = g.generateEnum(&t)
		} else if t.Modifier == modifierStruct {
			g.addImport(runtimeImport)
			g.addImport(ioImport)
			err = g.generateStruct(&t, pkgName)
		} else if t.Modifier == modifierUnion {
			g.addImport(runtimeImport)
			g.addImport(ioImport)
			err = g.generateUnion(&t, pkgName)
		}

		if err != nil {
			return nil, err
		}
	}

	// append imports
	buffer := g.sw.sb.Bytes()
	if len(g.currentFileImports) > 0 {

		importsBuffer := new(bytes.Buffer)
		importsBuffer.WriteString(importSentence)
		for fileImport := range g.currentFileImports {
			importsBuffer.WriteByte(quotes)
			importsBuffer.WriteString(fileImport)
			importsBuffer.WriteByte(quotes)
			importsBuffer.WriteByte(newline)
		}
		importsBuffer.Write([]byte{rParen, newline})

		buffer = prepend(buffer, importsBuffer.Bytes())

		// clear
		g.currentFileImports = make(map[string]any)
	}

	// prepend package keyword
	buffer = prepend(buffer, []byte(fmt.Sprintf("package %s\n", pkgName)))

	// format Go source code
	src, err := format.Source(buffer)
	if err != nil {
		// dump while debugging
		fmt.Println("DUMPING ==============")
		fmt.Println(string(buffer))

		errs := err.(scanner.ErrorList)
		return nil, errors.New(errs.Error())
	}

	return &GeneratedFile{
		Path:    filepath.Join(g.config.OutputPath, f.Name),
		Content: src,
	}, nil
}

func (g *Generator) generateStruct(t *NexemaTypeDefinition, pkgName string) error {
	data := StructTemplateData{
		TypeName: strcase.ToCamel(t.Name),
		Fields: mapArray(t.Fields, func(field NexemaTypeFieldDefinition) TypeFieldTemplateData {
			return TypeFieldTemplateData{
				FieldName:      strcase.ToCamel(field.Name),
				LowerFieldName: strcase.ToLowerCamel(field.Name),
				FieldIndex:     field.Index,
				ValueType:      nexemaTypeFieldDefinitionToTemplateData(field.Type),
			}
		}),
	}

	return structTemplate.ExecuteTemplate(g.sw.sb, "struct", data)
}

func (g *Generator) generateUnion(t *NexemaTypeDefinition, pkgName string) error {
	data := UnionTemplateData{
		TypeName:  strcase.ToCamel(t.Name),
		LowerName: strcase.ToLowerCamel(t.Name),
		Fields: mapArray(t.Fields, func(field NexemaTypeFieldDefinition) TypeFieldTemplateData {
			return TypeFieldTemplateData{
				FieldName:      strcase.ToCamel(field.Name),
				LowerFieldName: strcase.ToLowerCamel(field.Name),
				FieldIndex:     field.Index,
				ValueType:      nexemaTypeFieldDefinitionToTemplateData(field.Type),
			}
		}),
	}

	return unionTemplate.ExecuteTemplate(g.sw.sb, "union", data)
}

func (g *Generator) generateEnum(t *NexemaTypeDefinition) error {
	data := EnumTemplateData{
		TypeName:      t.Name,
		LowerTypeName: strcase.ToLowerCamel(t.Name),
		Fields: mapArray(t.Fields, func(field NexemaTypeFieldDefinition) EnumFieldTemplateData {
			return EnumFieldTemplateData{
				Index:     field.Index,
				Name:      strcase.ToCamel(field.Name),
				LowerName: strcase.ToLowerCamel(field.Name),
			}
		}),
	}

	// build the template
	return enumTemplate.Execute(g.sw.sb, data)
}

func (g *Generator) generateMapping(def *NexemaDefinition) {

	for _, file := range def.Files {
		packageName := filepath.Base(filepath.Dir(filepath.Join(g.config.OutputPath, file.Name)))
		g.fileMapping[file.Name] = packageName

		// importPath for package is: module/output-folder/filepath.Dir(file.Name)
		importPath := filepath.Join(g.config.ModuleName, g.config.OutputPath, filepath.Dir(file.Name))

		for _, typeDef := range file.Types {
			g.typeMapping[typeDef.Id] = typeMap{
				PackageName: packageName,
				ImportPath:  importPath,
				TypeDef:     &typeDef,
			}
		}
	}
}

func (g *Generator) addImport(v string) {
	g.currentFileImports[v] = true
}
