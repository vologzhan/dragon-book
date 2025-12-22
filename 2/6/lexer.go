package example

import (
	"bytes"
	"io"
	"strings"
)

type Tag int

const (
	TagEof Tag = 0
	TagNum Tag = iota + 256
	TagId
	TagTrue
	TagFalse
)

type Token struct {
	Tag
}

type Num struct {
	Token
	Value int
}

type Word struct {
	Token
	Lexeme string
}

func NewLexer(buf string) *Lexer {
	l := &Lexer{
		reader: strings.NewReader(buf),
		line:   1,
		words:  map[string]Word{},
	}
	_ = l.newWord(TagTrue, "true")
	_ = l.newWord(TagFalse, "false")

	return l
}

type Lexer struct {
	reader io.ByteReader
	line   int
	words  map[string]Word
}

func (l *Lexer) Scan() interface{} {
	var peek byte = ' '
	for ; ; peek, _ = l.reader.ReadByte() {
		if peek == ' ' || peek == '\t' {
			// nothing
		} else if peek == '\n' {
			l.line++
		} else {
			break
		}
	}

	switch {
	case isDigit(peek):
		v := 0
		for ; isDigit(peek); peek, _ = l.reader.ReadByte() {
			v = 10*v + int(peek-'0')
		}

		return Num{
			Token: Token{
				Tag: TagNum,
			},
			Value: v,
		}
	case isLetter(peek):
		buf := bytes.Buffer{}
		for ; isLetter(peek) || isDigit(peek); peek, _ = l.reader.ReadByte() {
			buf.WriteByte(peek)
		}

		lexeme := buf.String()
		if word, ok := l.words[lexeme]; ok {
			return word
		}

		return l.newWord(TagId, lexeme)
	default:
		return Token{
			Tag: Tag(peek),
		}
	}
}

func (l *Lexer) ScanAll() []interface{} {
	eof := Token{TagEof}
	var buf []interface{}
	for {
		t := l.Scan()
		if t == eof {
			return buf
		}
		buf = append(buf, t)
	}
}

func (l *Lexer) newWord(tag Tag, lexeme string) Word {
	w := Word{
		Token: Token{
			Tag: tag,
		},
		Lexeme: lexeme,
	}
	l.words[lexeme] = w

	return w
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_'
}
