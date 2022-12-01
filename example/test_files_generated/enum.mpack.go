package test_files

import "fmt"

type EnumType struct {
	index uint8
	name  string
}

type enumTypePicker struct{}

var EnumTypePicker enumTypePicker = enumTypePicker{}
var enumTypeUnknown EnumType = EnumType{index: 0, name: "unknown"}
var enumTypeValue EnumType = EnumType{index: 1, name: "value"}
var enumTypeSecond EnumType = EnumType{index: 2, name: "second"}
var enumTypeRandom1 EnumType = EnumType{index: 3, name: "random1"}

func (enumTypePicker) Unknown() EnumType {
	return enumTypeUnknown
}

func (enumTypePicker) Value() EnumType {
	return enumTypeValue
}

func (enumTypePicker) Second() EnumType {
	return enumTypeSecond
}

func (enumTypePicker) Random1() EnumType {
	return enumTypeRandom1
}

func (e EnumType) Index() uint8 {
	return e.index
}

func (e EnumType) Name() string {
	return e.name
}

func (enumTypePicker) ByIndex(index uint8) EnumType {
	switch index {

	case 0:
		return enumTypeUnknown

	case 1:
		return enumTypeValue

	case 2:
		return enumTypeSecond

	case 3:
		return enumTypeRandom1

	default:
		panic(fmt.Sprintf("EnumType with index %v not found", index))
	}
}

func (enumTypePicker) ByName(name string) EnumType {
	switch name {

	case "unknown":
		return enumTypeUnknown

	case "value":
		return enumTypeValue

	case "second":
		return enumTypeSecond

	case "random1":
		return enumTypeRandom1

	default:
		panic(fmt.Sprintf("EnumType with name %v not found", name))
	}
}
