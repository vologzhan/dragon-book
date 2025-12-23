package parser_homework

func newLexer(buf string) *lexer {
	return &lexer{
		buf: buf,
		pos: -1,
	}
}

type lexer struct {
	buf string
	pos int
}

func (l *lexer) Next() string {
	l.pos++
	if l.pos < len(l.buf) {
		return string(l.buf[l.pos])
	}

	return ""
}
