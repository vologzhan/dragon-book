package lexer

func NewSymbols(buf string) *Symbols {
	return &Symbols{
		buf: buf,
		pos: -1,
	}
}

type Symbols struct {
	buf string
	pos int
}

func (l *Symbols) Next() string {
	l.pos++
	if l.pos < len(l.buf) {
		return string(l.buf[l.pos])
	}

	return ""
}

func (l *Symbols) Line() string {
	return l.buf
}

func (l *Symbols) Pos() int {
	return l.pos
}
