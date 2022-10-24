package runtime

import "fmt"

type FieldSet map[string]*FieldInfo

type FieldInfo struct {
	Name          string
	Index         int
	Nullable      bool
	ValueType     FieldValueType
	IsEnum        bool
	TypeArguments []FieldValueType
}

func (f FieldInfo) String() string {
	return fmt.Sprintf("%s (nullable: %v) (index: %v)", f.Name, f.Nullable, f.Index)
}

type FieldValueType int8

const (
	Boolean FieldValueType = 0
	String  FieldValueType = 1
	Uint8   FieldValueType = 2
	Uint16  FieldValueType = 3
	Uint32  FieldValueType = 4
	Uint64  FieldValueType = 5
	Int8    FieldValueType = 6
	Int16   FieldValueType = 7
	Int32   FieldValueType = 8
	Int64   FieldValueType = 9
	Float32 FieldValueType = 10
	Float64 FieldValueType = 11
	Binary  FieldValueType = 12
	List    FieldValueType = 13
	Map     FieldValueType = 14
	Custom  FieldValueType = 15
)
