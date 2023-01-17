package main

import (
	"fmt"

	"github.com/example/models"
	"github.com/example/models/identity"
	"github.com/nexema/go/runtime"
)

func main() {
	// printEnum()
	// printUnion()
	printStruct()
}

func printStruct() {
	println("===========================")

	primitives1 := &identity.NexemaPrimitives{
		MyString:  "hello world",
		MyBoolean: true,
		MyFloat64: 98942.242,
		MyInt64:   123131,
		MyBinary:  []byte{0x5, 0x3, 0x2},
	}

	fmt.Println(primitives1)

	primitives2 := identity.NexemaPrimitives{}
	err := primitives2.MergeFrom(primitives1.MustEncode())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	primitives2.MyBinary[1] = 0x2

	fmt.Println("equals:", primitives1.Equals(&primitives2))

	list := identity.NewNexemaList(identity.NexemaListBuilder{
		List1: []string{"hola"},
		List4: []runtime.Nullable[string]{runtime.NewNull[string](), runtime.NewNullable("hola")},
		List5: [][]byte{{5, 3}, {5, 2}},
	})

	fmt.Println(list)

	defaultValues := identity.NewNexemaDefaultType()
	fmt.Println(defaultValues)

	defaultValues.MyNullablelist[0].Clear()
	fmt.Println(defaultValues)
}

func printEnum() {
	println("===========================")
	color := models.ColorsPicker.Blue()
	fmt.Println(color.String())
}

func printUnion() {
	println("===========================")
	union := identity.SingleUnionBuilder.Field1("hello world")
	fmt.Println(union)

	union.Clear()
	fmt.Println(union)

	union.SetField2(true)
	fmt.Println(union)

	union.SetField3([]runtime.Nullable[string]{
		runtime.NewNull[string](),
		runtime.NewNullable("hola"),
		runtime.NewNull[string](),
	})
	fmt.Println(union)

	union2 := union.Clone()
	fmt.Println(union2)

	union2.Clear()
	fmt.Println(union2)
	fmt.Println(union)
}
