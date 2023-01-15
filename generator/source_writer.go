package generator

import (
	"bytes"
	"fmt"
)

const (
	newline byte = '\n'
	quotes  byte = '"'
	rParen  byte = ')'
)

type sourceWriter struct {
	sb *bytes.Buffer
}

func newSourceWriter() *sourceWriter {
	return &sourceWriter{sb: new(bytes.Buffer)}
}

func (sw *sourceWriter) writeln(text string) {
	sw.sb.WriteString(text)
	sw.sb.WriteByte(newline)
}

func (sw *sourceWriter) writelnf(text string, args ...any) {
	sw.sb.WriteString(fmt.Sprintf(text, args...))
	sw.sb.WriteByte(newline)
}

func (sw *sourceWriter) reset() {
	sw.sb.Reset()
}
