package symbol_table

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

type Lexer struct {
	buf  io.ByteReader
	peek byte
}

func NewLexer(buf string) *Lexer {
	return &Lexer{
		strings.NewReader(buf),
		' ',
	}
}

func (l *Lexer) Next() string {
	for ; ; l.peek, _ = l.buf.ReadByte() {
		if l.peek == ' ' || l.peek == '\t' || l.peek == '\n' {
			continue
		}
		break
	}

	switch {
	case l.isDigit():
		v := 0
		for ; l.isDigit(); l.peek, _ = l.buf.ReadByte() {
			v = 10*v + int(l.peek-'0')
		}

		return strconv.Itoa(v)
	case l.isLetter():
		var buf bytes.Buffer
		for ; l.isLetter(); l.peek, _ = l.buf.ReadByte() {
			buf.WriteByte(l.peek)
		}

		return buf.String()
	case l.peek == 0:
		return ""
	default:
		t := string(l.peek)
		l.peek = ' '

		return t
	}
}

func (l *Lexer) isLetter() bool {
	return (l.peek >= 'a' && l.peek <= 'z') || (l.peek >= 'A' && l.peek <= 'Z')
}

func (l *Lexer) isDigit() bool {
	return l.peek >= '0' && l.peek <= '9'
}
