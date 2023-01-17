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

	primitives1 := identity.NexemaPrimitives{
		MyString:  "hello world",
		MyBoolean: true,
		MyFloat64: 98942.242,
		MyInt64:   123131,
	}

	primitives2 := identity.NexemaPrimitives{}
	err := primitives2.MergeFrom(primitives1.MustEncode())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("equals:", primitives1 == primitives2)
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
