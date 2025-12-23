package lexer

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
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
	for ; ; l.next() {
		if l.match(' ') || l.match('\t') {
			// nothing
		} else if l.match('\n') {
			l.line++
		} else if l.match('/') {
			prev := l.peek
			l.next()

			if l.match('/') {
				l.scanComment()
			} else if l.match('*') {
				l.scanMultilineComment()
			} else {
				return newToken(Tag(prev))
			}
		} else if l.match('.') {
			return l.scanMantissa(0)
		} else if l.match('=') || l.match('!') || l.match('<') || l.match('>') {
			prev := l.peek
			l.next()

			if !l.match('=') {
				return newToken(Tag(prev))
			}

			var tok Token
			switch prev {
			case '=':
				tok = newToken(TagEqual)
			case '!':
				tok = newToken(TagNotEqual)
			case '<':
				tok = newToken(TagLessOrEqual)
			case '>':
				tok = newToken(TagGreaterOrEqual)
			}

			l.peek = ' '

			return tok
		} else {
			break
		}
	}

	switch {
	case l.matchDigit():
		v := 0
		for ; l.matchDigit(); l.next() {
			v = addDigit(v, l.peek)
		}

		if l.match('.') {
			return l.scanMantissa(v)
		}

		return newNum(v)
	case l.matchLetter():
		buf := bytes.Buffer{}
		for ; l.matchLetter() || l.matchDigit(); l.next() {
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

func (l *Lexer) scanComment() {
	for ; ; l.next() {
		switch l.peek {
		case '\n':
			l.line++
			return
		case 0:
			return
		}
	}
}

func (l *Lexer) scanMultilineComment() {
	for ; ; l.next() {
		switch l.peek {
		case '\n':
			l.line++
		case '*':
			l.next()
			if l.match('/') {
				return
			}
			if l.match('\n') {
				l.line++
			}
		case 0:
			return
		}
	}
}

func (l *Lexer) scanMantissa(intPart int) Float {
	mantissa := 0
	for l.next(); l.matchDigit(); l.next() {
		mantissa = addDigit(mantissa, l.peek)
	}

	str := fmt.Sprintf("%d.%d", intPart, mantissa)
	f, _ := strconv.ParseFloat(str, 64)

	return newFloat(f)
}

func (l *Lexer) next() {
	l.peek, _ = l.reader.ReadByte()
}

func (l *Lexer) prev() {
	_ = l.reader.UnreadByte()
}

func (l *Lexer) reserve(tag Tag, lexeme string) Word {
	w := newWord(tag, lexeme)
	l.words[lexeme] = w

	return w
}

func (l *Lexer) match(b byte) bool {
	return l.peek == b
}

func (l *Lexer) matchDigit() bool {
	return l.peek >= '0' && l.peek <= '9'
}

func (l *Lexer) matchLetter() bool {
	return (l.peek >= 'a' && l.peek <= 'z') || (l.peek >= 'A' && l.peek <= 'Z') || l.match('_')
}

func addDigit(v int, b byte) int {
	return 10*v + int(b-'0')
}
