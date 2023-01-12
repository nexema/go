package nexemab

import (
	"bufio"
	"io"
)

type Decoder struct {
	r *bufio.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{r: bufio.NewReader(reader)}
}
