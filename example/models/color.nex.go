package models

import (
	"fmt"
)

type Colors struct {
	index uint8
	name  string
}

func (enum Colors) Index() uint8 {
	return enum.index
}

func (enum Colors) Name() string {
	return enum.name
}

type colorsPicker struct{}

var ColorsPicker colorsPicker = colorsPicker{}

var colorsUnknown Colors = Colors{index: 0, name: "unknown"}

func (colorsPicker) Unknown() Colors {
	return colorsUnknown
}

var colorsRed Colors = Colors{index: 1, name: "red"}

func (colorsPicker) Red() Colors {
	return colorsRed
}

var colorsGreen Colors = Colors{index: 2, name: "green"}

func (colorsPicker) Green() Colors {
	return colorsGreen
}

var colorsBlue Colors = Colors{index: 3, name: "blue"}

func (colorsPicker) Blue() Colors {
	return colorsBlue
}

func (colorsPicker) ByIndex(index uint8) Colors {
	switch index {

	case 0:
		return colorsUnknown

	case 1:
		return colorsRed

	case 2:
		return colorsGreen

	case 3:
		return colorsBlue

	default:
		panic(fmt.Sprintf("unknown Colors with value %v", index))
	}
}

func (colorsPicker) ByName(name string) Colors {
	switch name {

	case "unknown":
		return colorsUnknown

	case "red":
		return colorsRed

	case "green":
		return colorsGreen

	case "blue":
		return colorsBlue

	default:
		panic(fmt.Sprintf("unknown Colors with name %s", name))
	}
}
