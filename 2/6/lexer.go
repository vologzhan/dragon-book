package lexer

import (
	"bytes"
	"io"
	"strings"
)

type Lexer struct {
	reader io.ByteScanner
	peek   byte
	line   int
	words  map[string]Word
}

func NewLexer(buf string) *Lexer {
	l := &Lexer{
		strings.NewReader(buf),
		' ',
		1,
		map[string]Word{},
	}
	_ = l.reserve(TagTrue, "true")
	_ = l.reserve(TagFalse, "false")

	return l
}

func (l *Lexer) Scan() interface{} {
	for ; ; l.peek, _ = l.reader.ReadByte() {
		if l.peek == ' ' || l.peek == '\t' {
			// nothing
		} else if l.peek == '\n' {
			l.line++
		} else if l.peek == '/' {
			prev := l.peek
			l.peek, _ = l.reader.ReadByte()

			if l.peek == '/' {
				l.skipComment()
			} else if l.peek == '*' {
				l.skipMultilineComment()
			} else {
				_ = l.reader.UnreadByte()
				l.peek = prev
				break
			}
		} else {
			break
		}
	}

	switch {
	case isDigit(l.peek):
		v := 0
		for ; isDigit(l.peek); l.peek, _ = l.reader.ReadByte() {
			v = 10*v + int(l.peek-'0')
		}

		return newNum(v)
	case isLetter(l.peek):
		buf := bytes.Buffer{}
		for ; isLetter(l.peek) || isDigit(l.peek); l.peek, _ = l.reader.ReadByte() {
			buf.WriteByte(l.peek)
		}

		lexeme := buf.String()
		if word, ok := l.words[lexeme]; ok {
			return word
		}

		return l.reserve(TagId, lexeme)
	default:
		t := newToken(Tag(l.peek))
		l.peek = ' '

		return t
	}
}

func (l *Lexer) ScanAll() []interface{} {
	eof := newToken(TagEof)
	var buf []interface{}
	for {
		t := l.Scan()
		if t == eof {
			return buf
		}
		buf = append(buf, t)
	}
}

func (l *Lexer) reserve(tag Tag, lexeme string) Word {
	w := newWord(tag, lexeme)
	l.words[lexeme] = w

	return w
}

func (l *Lexer) skipComment() {
	for ; ; l.peek, _ = l.reader.ReadByte() {
		switch l.peek {
		case '\n':
			l.line++
			return
		case 0:
			return
		}
	}
}

func (l *Lexer) skipMultilineComment() {
	for ; ; l.peek, _ = l.reader.ReadByte() {
		switch l.peek {
		case '\n':
			l.line++
		case '*':
			l.peek, _ = l.reader.ReadByte()
			if l.peek == '/' {
				return
			}
			if l.peek == '\n' {
				l.line++
			}
		case 0:
			return
		}
	}
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}
