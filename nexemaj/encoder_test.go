package nexemaj

import (
	"encoding/json"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	encoder := NewEncoder()
	b.ResetTimer()

	/*
		replicate the following:
		{"hello":true,"array":["string", 123123.42423]}
	*/
	for i := 0; i < b.N; i++ {
		encoder.WriteObjStart()
		encoder.WriteStringKey("hello")
		encoder.WriteBool(true)
		encoder.WriteComma()

		encoder.WriteStringKey("array")
		encoder.WriteArrayStart()
		encoder.WriteString("string")
		encoder.WriteComma()
		encoder.WriteFloat64(123123.213124)
		encoder.WriteArrayEnd()
		encoder.WriteObjEnd()

		result := encoder.String()
		_ = result
	}
}

func BenchmarkEncodeJson(b *testing.B) {
	/*
		replicate the following:
		{"hello":true,"array":["string", 123123.42423]}
	*/
	m := map[string]interface{}{
		"hello": true,
		"array": []interface{}{"string", 123123.42423},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, _ := json.Marshal(m)
		_ = string(r)
	}
}
