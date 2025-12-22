package lexer

import "strings"

func NewWords(buf string) *Words {
	return &Words{
		buf: buf,
	}
}

type Words struct {
	buf string
	pos int
	end int
}

func (l *Words) Next() string {
	lookahead := strings.Builder{}
	for ; l.end < len(l.buf); l.end++ {
		c := l.buf[l.end]
		if c == ' ' || c == '\t' || c == '\n' {
			if lookahead.Len() > 0 {
				break
			}
			continue
		}

		if lookahead.Len() == 0 {
			l.pos = l.end
		}

		lookahead.WriteByte(c)
	}

	return lookahead.String()
}

func (l *Words) Line() string {
	return l.buf
}

func (l *Words) Pos() int {
	return l.pos
}
