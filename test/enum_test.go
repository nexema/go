package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnumQuality(t *testing.T) {
	enum1 := ColorPicker.Blue()
	enum2 := ColorPicker.Red()
	enum3 := ColorPicker.Blue()

	require.NotEqual(t, enum1, enum2)
	require.Equal(t, enum1, enum3)
}

func TestEnumLookupByIndex(t *testing.T) {
	enum1 := ColorPicker.ByIndex(2)
	require.Equal(t, enum1, ColorPicker.Blue())
	require.Panics(t, func() { ColorPicker.ByIndex(9) })
}

func TestEnumLookupByName(t *testing.T) {
	enum1 := ColorPicker.ByName("orange")
	require.Equal(t, enum1, ColorPicker.Orange())
	require.Panics(t, func() { ColorPicker.ByName("hello") })
}

// red 0
// blue 1
// orange 2
type Color struct {
	index uint8
	name  string
}

type colorPicker struct{}

var ColorPicker colorPicker = colorPicker{}
var colorRed Color = Color{index: 0, name: "red"}
var colorBlue Color = Color{index: 1, name: "blue"}
var colorOrange Color = Color{index: 2, name: "orange"}

func (c Color) Index() uint8 {
	return c.index
}

func (c Color) Name() string {
	return c.name
}

func (colorPicker) Blue() Color {
	return colorBlue
}

func (colorPicker) Red() Color {
	return colorRed
}

func (colorPicker) Orange() Color {
	return colorOrange
}

func (colorPicker) ByIndex(index uint8) Color {
	switch index {
	case 0:
		return colorRed

	case 1:
		return colorBlue

	case 2:
		return colorBlue

	default:
		panic(fmt.Sprintf("Colors with value %v not found", index))
	}
}

func (colorPicker) ByName(name string) Color {
	switch name {
	case "blue":
		return colorBlue

	case "red":
		return colorRed

	case "orange":
		return colorOrange

	default:
		panic(fmt.Sprintf("unknown Color with name %s", name))
	}
}
