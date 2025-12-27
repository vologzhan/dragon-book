package union

import "strings"

const (
	IF   = "if"
	THEN = "then"
	ELSE = "else"

	LT = "<"
	LE = "<="
	EQ = "="
	NE = "<>"
	GT = ">"
	GE = ">="
)

type Relop string
type Word string
type KeyWord string
type Number string
type Float string

type lexer struct {
	buf string
	pos int
}

func parse(buf string) []interface{} {
	var out []interface{}
	l := lexer{buf: buf}
	for {
		next := l.nextToken()
		if next == nil {
			break
		}
		out = append(out, next)
	}

	return out
}

func (l *lexer) nextToken() interface{} {
	buf := strings.Builder{}
	var state int
	for {
		switch state {
		case 0:
			c := l.nextChar()
			switch {
			case c == '<':
				state = 1
			case c == '=':
				state = 5
			case c == '>':
				state = 6
			case isDelim(c):
				state = 0
			case isLetter(c):
				buf.WriteByte(c)
				state = 10
			case isDigit(c):
				buf.WriteByte(c)
				state = 13
			case c == 0:
				return nil
			default:
				panic("invalid token")
			}
		case 1:
			switch l.nextChar() {
			case '=':
				state = 2
			case '>':
				state = 3
			default:
				state = 4
			}
		case 2:
			return Relop(LE)
		case 3:
			return Relop(NE)
		case 4:
			l.retract()
			return Relop(LT)
		case 5:
			return Relop(EQ)
		case 6:
			switch l.nextChar() {
			case '=':
				state = 7
			default:
				state = 8
			}
		case 7:
			return Relop(GE)
		case 8:
			l.retract()
			return Relop(GT)
		case 10:
			c := l.nextChar()
			switch {
			case isLetter(c) || isDigit(c):
				buf.WriteByte(c)
			default:
				l.retract()

				word := buf.String()
				switch word {
				case IF, THEN, ELSE:
					return KeyWord(word)
				default:
					return Word(word)
				}
			}
		case 13:
			c := l.nextChar()
			switch {
			case isDigit(c):
				buf.WriteByte(c)
				continue
			case c == '.':
				buf.WriteByte(c)
				state = 14
			case c == 'E':
				buf.WriteByte(c)
				state = 16
			default:
				l.retract()
				return Number(buf.String())
			}
		case 14:
			c := l.nextChar()
			switch {
			case isDigit(c):
				buf.WriteByte(c)
				state = 15
			default:
				panic("invalid token")
			}
		case 15:
			c := l.nextChar()
			switch {
			case isDigit(c):
				buf.WriteByte(c)
			case c == 'E':
				buf.WriteByte(c)
				state = 16
			default:
				l.retract()
				return Float(buf.String())
			}
		case 16:
			c := l.nextChar()
			switch {
			case c == '+' || c == '-':
				buf.WriteByte(c)
				state = 17
			case isDigit(c):
				buf.WriteByte(c)
				state = 18
			default:
				panic("invalid token")
			}
		case 17:
			c := l.nextChar()
			switch {
			case isDigit(c):
				buf.WriteByte(c)
				state = 18
			default:
				panic("invalid token")
			}
		case 18:
			c := l.nextChar()
			switch {
			case isDigit(c):
				buf.WriteByte(c)
			default:
				l.retract()
				return Float(buf.String())
			}
		case 23:
			c := l.nextChar()
			switch {
			case isDelim(c):
				state = 23
			default:
				l.retract()
				state = 0
			}
		}
	}
}

func (l *lexer) nextChar() byte {
	if l.pos < len(l.buf) {
		c := l.buf[l.pos]
		l.pos++
		return c
	}
	return 0
}

func (l *lexer) retract() {
	if l.pos != len(l.buf) {
		l.pos--
	}
}

func isDelim(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
