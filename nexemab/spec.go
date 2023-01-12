package nexemab

const (
	Null      byte = 0xc0
	BoolTrue  byte = 0x01
	BoolFalse byte = 0x00
)

const (
	MSB    uint64 = 0x80
	REST   uint64 = 0x7F
	MSBALL uint64 = ^REST
	INT    uint64 = 2147483648
)
