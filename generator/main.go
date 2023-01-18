package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	generator "github.com/nexema/go/generator/internal"
)

const pluginConfigPrefix = "--config="
const newline = '\n'

func main() {
	// create reader and writer for stdin and stdout
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	// parse plugin config
	pluginConfig := &generator.PluginConfig{}
	args := os.Args[1:]
	for _, arg := range args {
		if strings.HasPrefix(arg, pluginConfigPrefix) {
			err := json.Unmarshal([]byte(arg[len(pluginConfigPrefix):]), pluginConfig)
			if err != nil {
				write(writer, "failed to parse plugin config")
				return
			}
		}
	}

	// read buffer
	content, err := reader.ReadString('\n')
	if err != nil {
		write(writer, "could not parse input")
		return
	}

	content = strings.TrimSpace(content)
	gen := generator.NewGenerator(pluginConfig)
	generatedFiles, err := gen.Generate([]byte(content))
	if err != nil {
		write(writer, fmt.Sprintf("failed to generate source files. Error: %s", err))
		return
	}

	// write files to output
	for _, file := range generatedFiles {
		path := fmt.Sprintf("%s.go", file.Path)

		// create directory structure
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil && !errors.Is(err, os.ErrExist) {
			write(writer, fmt.Sprintf("Could not create directory %s for output. Error: %s", filepath.Dir(path), err.Error()))
			return
		}

		err = os.WriteFile(path, file.Content, os.ModePerm)
		if err != nil {
			write(writer, fmt.Sprintf("Could not write file %s. Error: %s", path, err.Error()))
			return
		}
	}

	// write successs
	write(writer, "ok")
}

func write(w *bufio.Writer, content string) {
	w.WriteString(content)
	w.WriteRune(newline)
	w.Flush()
}
