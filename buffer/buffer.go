package buffer

type Buffer struct {
	backstore []*backstore // maintain a list of pointers to all the buffers
	bufIdx    int          // the current buffer's index on backstore
	offset    int          // current writing offset
	size      int          // backstore[bufIdx] size
	totalSize int          // the total size of all filled positions in all backstores
}

type BufferConfig struct {
	Capacity          int
	InitialBufferSize int
}

type backstore struct {
	wpos   int    // the last write position
	buffer []byte // the buffer we are writing to
}

func NewBuffer(cfg ...BufferConfig) *Buffer {
	config := BufferConfig{Capacity: 5, InitialBufferSize: 10}
	if len(cfg) > 0 {
		config = cfg[0]
	}

	buffer := &Buffer{
		offset:    0,
		bufIdx:    0,
		backstore: make([]*backstore, 0, config.Capacity),
		size:      config.InitialBufferSize,
		totalSize: 0,
	}

	// create initial backstore buffer
	buffer.backstore = append(buffer.backstore, &backstore{
		buffer: make([]byte, buffer.size),
		wpos:   0,
	})

	return buffer
}

func (b *Buffer) Write(buf []byte) error {
	b.ensure(len(buf))
	b.offset += copy(b.backstore[b.bufIdx].buffer[b.offset:], buf)
	return nil
}

func (b *Buffer) WriteByte(v byte) error {
	b.ensure(1)
	(b.backstore[b.bufIdx].buffer)[b.offset] = v
	b.offset++
	return nil
}

// ensure ensures the current buffer has enough space to store size bytes.
// Returns true if a new buffer was created, otherwise false
func (b *Buffer) ensure(size int) bool {
	if b.size-b.offset < int(size) {
		// update last backstore
		b.backstore[b.bufIdx].wpos = b.offset
		b.totalSize += b.offset

		// create a new buffer
		b.bufIdx++
		b.offset = 0
		b.size = int(size) * 2
		b.backstore = append(b.backstore, &backstore{
			wpos:   0,
			buffer: make([]byte, b.size),
		})
		return true
	}

	return false
}

// Bytes returns a concatenation of all the byte arrays in the underlying backstore
func (b *Buffer) Bytes() []byte {
	// update totalSize with current buffer
	b.totalSize += b.offset
	b.backstore[b.bufIdx].wpos = b.offset

	buffer := make([]byte, b.totalSize)
	offset := 0
	for _, ba := range b.backstore {
		offset += copy(buffer[offset:], ba.buffer[:ba.wpos])
	}

	return buffer
}

// Returns the byte at specified index.
// todo: support go back into other backstores
func (b *Buffer) At(index int) byte {
	if index == -1 {
		index = b.offset - 1
	}

	return b.backstore[b.bufIdx].buffer[index]
}

// Set overwrites the content of the buffer at index
// todo: support go back into other backstores
func (b *Buffer) Set(index int, value byte) {
	if index == -1 {
		index = b.offset - 1
	}

	b.backstore[b.bufIdx].buffer[index] = value
}
