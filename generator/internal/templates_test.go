package generator

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFieldDeclarationTemplate(t *testing.T) {
	tests := []struct {
		description string
		input       TypeFieldValueKindTemplate
		want        string
	}{
		{
			description: "list",
			want:        "[]string",
			input: TypeFieldValueKindTemplate{
				IsList:         true,
				IsNullable:     false,
				ImportTypeName: "list",
				TypeArguments: []TypeFieldValueKindTemplate{
					{
						IsPrimitive:    true,
						ImportTypeName: "string",
					},
				},
			},
		},
		{
			description: "list with nullable argument",
			want:        "[]runtime.Nullable[string]",
			input: TypeFieldValueKindTemplate{
				IsList:         true,
				IsNullable:     false,
				ImportTypeName: "list",
				TypeArguments: []TypeFieldValueKindTemplate{
					{
						IsPrimitive:    true,
						IsNullable:     true,
						ImportTypeName: "string",
					},
				},
			},
		},
		{
			description: "nullable list with nullable argument",
			want:        "runtime.Nullable[[]runtime.Nullable[string]]",
			input: TypeFieldValueKindTemplate{
				IsList:         true,
				IsNullable:     true,
				ImportTypeName: "list",
				TypeArguments: []TypeFieldValueKindTemplate{
					{
						IsPrimitive:    true,
						IsNullable:     true,
						ImportTypeName: "string",
					},
				},
			},
		},
		{
			description: "nullable list",
			want:        "runtime.Nullable[[]string]",
			input: TypeFieldValueKindTemplate{
				IsList:         true,
				IsNullable:     true,
				ImportTypeName: "list",
				TypeArguments: []TypeFieldValueKindTemplate{
					{
						IsPrimitive:    true,
						ImportTypeName: "string",
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			result := new(bytes.Buffer)
			err := fieldDeclarationTemplate.Execute(result, tc.input)
			require.Nil(t, err)
			require.Equal(t, tc.want, result.String())
		})
	}
}
