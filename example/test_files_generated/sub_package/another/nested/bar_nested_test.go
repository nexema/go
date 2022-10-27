package test_files

import (
	"testing"

	test_files "github.com/example/test_files/test_files_generated/sub_package/another"
)

func TestSerialize(t *testing.T) {
	instance := MyBarNested{
		First: MyNestedType{
			First: test_files.AnotherType{
				Status: test_files.AnotherEnumPicker.Success(),
				First:  "hello world",
				Second: true,
			},
		},
	}

	buf, _ := instance.Serialize()

	n := MyBarNested{}
	err := n.MergeFrom(buf)
	_ = err
}
