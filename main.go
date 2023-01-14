package main

import (
	"time"

	"github.com/nexema/go/nexemaj"
)

func main() {
	now := time.Now()

	encoder := nexemaj.NewEncoder()
	encoder.WriteObjStart()
	encoder.WriteStringKey("hello world")
	// encoder.WriteInt64(123123)
	encoder.WriteFloat64(123123.213124)
	encoder.WriteObjEnd()

	diff := time.Since(now)
	println(encoder.String())
	println("it took ", diff.Milliseconds(), " ns")
}
