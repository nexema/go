package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	const snapshotPath = "../testdata/snapshot.json"
	buf, _ := os.ReadFile(snapshotPath)
	generator := NewGenerator(&PluginConfig{
		OutputPath: "../example/models",
		ModuleName: "github.com/example",
	})
	result, err := generator.Generate(buf)
	if err != nil {
		println(err.Error())
	}

	require.Nil(t, err)
	require.NotNil(t, result)

	os.RemoveAll("example/models")

	for _, gen := range result {
		path := fmt.Sprintf("%s.go", gen.Path)
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		require.Nil(t, err)

		err = os.WriteFile(path, gen.Content, os.ModePerm)
		require.Nil(t, err)
	}
}
