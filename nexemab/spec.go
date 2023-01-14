package nexemab

const (
	null       byte = 0xc0
	boolTrue   byte = 0x01
	boolFalse  byte = 0x00
	arrayBegin byte = 0xdc
	mapBegin   byte = 0xdf
)

const (
	maxVarintLen = 10
)
