package runtime

import (
	"sync"

	"github.com/nexema/go/buffer"
	"github.com/nexema/go/nexemab"
)

var encoderPool sync.Pool = sync.Pool{
	New: func() any {
		return nexemab.NewEncoder(buffer.NewBytesBuffer())
	},
}

// GetEncoder retrieves a nexemab.Encoder instance
func GetEncoder() *nexemab.Encoder {
	enc := encoderPool.Get().(*nexemab.Encoder)
	enc.Reset()

	return enc
}
