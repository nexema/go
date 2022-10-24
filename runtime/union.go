package runtime

type SchemaUnion[T any, TWhichField any] interface {
	WhichField() TWhichField
	Clear()
	IsSet() bool
	CurrentValue() interface{}
}
