package test

import (
	"bytes"
	"testing"

	"github.com/messagepack-schema/go/runtime/msgpack"
	"github.com/stretchr/testify/require"
)

func TestStructEquality(t *testing.T) {
	instance1 := ExampleStruct{
		Name:  "hello world",
		Names: []string{"a", "b", "goodbye world"},
		Config: map[string]uint64{
			"another": 128712895718295789,
			"min":     0,
			"ads":     4454,
		},
		Another: AnotherStruct{
			Name:  "another struct name",
			Color: ColorPicker.Blue(),
		},
	}

	instance2 := ExampleStruct{
		Name:  "hello world",
		Names: []string{"a", "b", "goodbye world"},
		Config: map[string]uint64{
			"another": 128712895718295789,
			"min":     0,
			"ads":     4454,
		},
		Another: AnotherStruct{
			Name:  "another struct name",
			Color: ColorPicker.Blue(),
		},
	}

	instance3 := ExampleStruct{
		Name:  "hello world",
		Names: []string{"a", "b", "goodbye world"},
		Config: map[string]uint64{
			"another": 128712895718295789,
			"min":     0,
			"ads":     4454,
		},
		Another: AnotherStruct{
			Name:  "another struct name",
			Color: ColorPicker.Red(),
		},
	}

	require.Equal(t, instance1, instance2)
	require.NotEqual(t, instance1, instance3)

	instance3.Another.Color = ColorPicker.Blue()
	require.Equal(t, instance1, instance3)

	instance3.Config["new"] = 223
	require.NotEqual(t, instance1, instance3)

	instance2.Names = instance2.Names[1:1]
	require.NotEqual(t, instance1, instance2)
}

func TestSerializeDeserializeStruct(t *testing.T) {
	instance1 := ExampleStruct{
		Name:  "hello world",
		Names: []string{"a", "b", "goodbye world"},
		Config: map[string]uint64{
			"another": 128712895718295789,
			"min":     0,
			"ads":     4454,
		},
		Another: AnotherStruct{
			Name:  "another struct name",
			Color: ColorPicker.Blue(),
		},
	}

	instance2 := ExampleStruct{}
	err := instance2.MergeFrom(instance1.MustSerialize())

	require.Nil(t, err)
	require.Equal(t, instance1, instance2)

	instance1.Names = make([]string, 0)
	err = instance2.MergeFrom(instance1.MustSerialize())

	require.Nil(t, err)
	require.Equal(t, instance1, instance2)
}

func TestMergeOtherInstanceStruct(t *testing.T) {
	instance1 := ExampleStruct{
		Name:  "hello world",
		Names: []string{"a", "b", "goodbye world"},
		Config: map[string]uint64{
			"another": 128712895718295789,
			"min":     0,
			"ads":     4454,
		},
		Another: AnotherStruct{
			Name:  "another struct name",
			Color: ColorPicker.Blue(),
		},
	}

	instance2 := ExampleStruct{
		Name:  "bye world",
		Names: []string{"b", "goodbye world"},
		Config: map[string]uint64{
			"ads": 4454,
		},
		Another: AnotherStruct{
			Name:  "another struct name",
			Color: ColorPicker.Orange(),
		},
	}

	require.NotEqual(t, instance1, instance2)

	instance1.MergeUsing(&instance2)
	require.Equal(t, instance1, instance2)
}

type ExampleStruct struct {
	Name    string
	Names   []string
	Config  map[string]uint64
	Another AnotherStruct
}

type AnotherStruct struct {
	Name  string
	Color Color
}

func (u *ExampleStruct) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	err := writer.EncodeString(u.Name)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeArrayLen(len(u.Names))
	if err != nil {
		return nil, err
	}

	for _, v := range u.Names {
		err = writer.EncodeString(v)
		if err != nil {
			return nil, err
		}
	}

	err = writer.EncodeMapLen(len(u.Config))
	if err != nil {
		return nil, err
	}

	for k, v := range u.Config {
		err = writer.EncodeString(k)
		if err != nil {
			return nil, err
		}

		err = writer.EncodeUint64(v)
		if err != nil {
			return nil, err
		}
	}

	anotherBytes, err := u.Another.Serialize()
	if err != nil {
		return nil, err
	}

	err = writer.EncodeBytes(anotherBytes)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (u *ExampleStruct) MustSerialize() []byte {
	buf, err := u.Serialize()
	if err != nil {
		panic(err)
	}

	return buf
}

func (u *ExampleStruct) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := msgpack.NewDecoder(reader)
	var err error

	u.Name, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	namesLen, err := decoder.DecodeArrayLen()
	if err != nil {
		return err
	}

	u.Names = make([]string, namesLen)
	for i := 0; i < namesLen; i++ {
		u.Names[i], err = decoder.DecodeString()
		if err != nil {
			return err
		}
	}

	configLen, err := decoder.DecodeMapLen()
	if err != nil {
		return err
	}

	u.Config = make(map[string]uint64)
	for i := 0; i < configLen; i++ {
		k, err := decoder.DecodeString()
		if err != nil {
			return err
		}

		v, err := decoder.DecodeUint64()
		if err != nil {
			return err
		}

		u.Config[k] = v
	}

	anotherBytes, err := decoder.DecodeBytes()
	if err != nil {
		return err
	}

	u.Another = AnotherStruct{}
	u.Another.MergeFrom(anotherBytes)

	return nil
}

func (u *ExampleStruct) MergeUsing(other *ExampleStruct) error {
	u.Name = other.Name
	u.Names = other.Names
	u.Config = other.Config
	u.Another = other.Another
	return nil
}

func (u *AnotherStruct) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	writer := msgpack.NewEncoder(buf)
	err := writer.EncodeString(u.Name)
	if err != nil {
		return nil, err
	}

	err = writer.EncodeUint8(u.Color.Index())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (u *AnotherStruct) MergeFrom(buffer []byte) error {
	reader := bytes.NewBuffer(buffer)
	decoder := msgpack.NewDecoder(reader)
	var err error

	u.Name, err = decoder.DecodeString()
	if err != nil {
		return err
	}

	colorIdx, err := decoder.DecodeUint8()
	if err != nil {
		return err
	}

	u.Color = ColorPicker.ByIndex(colorIdx)
	return nil
}

func (u *AnotherStruct) MergeUsing(other *AnotherStruct) error {
	u.Name = other.Name
	u.Color = other.Color
	return nil
}
