package test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/messagepack-schema/go/runtime"
	"github.com/stretchr/testify/require"
)

func TestCreateNullable(t *testing.T) {
	nullable := runtime.NewNullable(25)
	require.True(t, nullable.HasValue())
	require.Equal(t, 25, *nullable.Value)
}

func TestClearNullable(t *testing.T) {
	nullable := runtime.NewNullable("a random string")
	require.True(t, nullable.HasValue())
	require.Equal(t, "a random string", *nullable.Value)

	nullable.Clear()
	require.False(t, nullable.HasValue())

	nullable.SetValue("hello world")
	require.Equal(t, "hello world", *nullable.Value)
}

func TestValueOrDefault(t *testing.T) {
	nullable := runtime.NewNullable(32.23)
	nullable.Clear()
	require.Equal(t, 0.0, *nullable.ValueOrDefault())
}

func TestStructNullable(t *testing.T) {
	type user struct {
		FirstName string
		LastName  string
		Age       int
		Date      *time.Time
	}

	nullable := runtime.NewNullable(user{
		FirstName: "a",
		LastName:  "b",
		Age:       5,
		Date:      &time.Time{},
	})

	nullable.Value.Age = 54
	require.Equal(t, 54, nullable.Value.Age)

	nullable.Clear()
	require.Equal(t, user{}, *nullable.ValueOrDefault())
}

func TestJsonMarshal(t *testing.T) {
	nullable := runtime.NewNullable[float32](24.35)
	nullable.Clear()
	var buf []byte
	var err error

	buf, err = json.Marshal(nullable)

	require.Nil(t, err)
	require.Equal(t, "null", string(buf))

	usr := struct {
		FirstName string
		Email     runtime.Nullable[string]
	}{
		FirstName: "hello",
		Email:     runtime.NewNull[string](),
	}

	buf, err = json.Marshal(usr)
	require.Nil(t, err)
	require.Equal(t, `{"FirstName":"hello","Email":null}`, string(buf))

	usr.Email.SetValue("hello@example.com")
	buf, err = json.Marshal(usr)
	require.Nil(t, err)
	require.Equal(t, `{"FirstName":"hello","Email":"hello@example.com"}`, string(buf))
}

func TestJsonUnmarshal(t *testing.T) {
	nullable := runtime.Nullable[string]{}
	var err error

	err = json.Unmarshal([]byte("null"), &nullable)
	require.Nil(t, err)
	require.False(t, nullable.HasValue())

	type aStruct struct {
		FirstName string
		Email     runtime.Nullable[string]
	}
	a := aStruct{}

	err = json.Unmarshal([]byte(`{"FirstName":"hello","Email":null}`), &a)
	require.Nil(t, err)
	require.Exactly(t, aStruct{FirstName: "hello", Email: runtime.NewNull[string]()}, a)

	err = json.Unmarshal([]byte(`{"FirstName":"hello","Email":"hello@example.com"}`), &a)
	require.Nil(t, err)
	require.Exactly(t, aStruct{FirstName: "hello", Email: runtime.NewNullable("hello@example.com")}, a)
}
