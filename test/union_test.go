package test

import (
	"bytes"
	"testing"

	"github.com/messagepack-schema/go/runtime"
	"github.com/messagepack-schema/go/runtime/msgpack"
	"github.com/stretchr/testify/require"
)

func TestCreateUnion(t *testing.T) {
	myUnion := MyUnionBuilder.Age(26)
	require.Equal(t, myUnion.value, int(26))
	require.Equal(t, myUnion.fieldIndex, 1)

	var u runtime.SchemaUnion[MyUnion, MyUnionWhichField] = MyUnionBuilder.Age(22)
	require.Equal(t, int(22), u.CurrentValue())

	u.Clear()

	require.False(t, u.IsSet())
}

func ClearUnion(u *MyUnion) {
	u.Clear()
}

type MyUnion struct {
	value      interface{}
	fieldIndex int
}

type MyUnionWhichField int8

const (
	MyUnion_NotSet    MyUnionWhichField = -1
	MyUnion_Age       MyUnionWhichField = 1
	MyUnion_FirstName MyUnionWhichField = 2
)

type myUnionBuilder struct{}

var MyUnionBuilder *myUnionBuilder = &myUnionBuilder{}

func (u *MyUnion) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := msgpack.GetEncoder()
	encoder.Reset(buf)

	return buf.Bytes(), nil
}

func (u *MyUnion) MergeFrom(buf []byte) error {
	return nil
}

func (u *MyUnion) MergeUsing(other *MyUnion) error {
	u.fieldIndex = other.fieldIndex
	u.value = other.value
	return nil
}

func (u *MyUnion) Clone() *MyUnion {
	return &MyUnion{
		fieldIndex: u.fieldIndex,
		value:      u.value,
	}
}

func (*myUnionBuilder) Age(value int) *MyUnion {
	return &MyUnion{
		value:      value,
		fieldIndex: 1,
	}
}

func (*myUnionBuilder) FirstName(value string) MyUnion {
	return MyUnion{
		value:      value,
		fieldIndex: 2,
	}
}

func (u MyUnion) IsSet() bool {
	return u.fieldIndex != -1
}

func (u *MyUnion) Clear() {
	u.value = nil
	u.fieldIndex = -1
}

func (u *MyUnion) WhichField() MyUnionWhichField {
	return MyUnionWhichField(u.fieldIndex)
}

func (u *MyUnion) CurrentValue() interface{} {
	return u.value
}

func (u *MyUnion) SetAge(age int) {
	u.value = age
	u.fieldIndex = 1
}

func (u *MyUnion) SetFirstName(name string) {
	u.value = name
	u.fieldIndex = 2
}
