package symbol_table

import (
	"fmt"
	"strings"
)

type Translator struct {
	lexer     *Lexer
	buf       []string
	lookahead string
	env       *Env
}

func NewTranslator(input string) *Translator {
	l := NewLexer(input)
	return &Translator{
		l,
		nil,
		l.Next(),
		nil,
	}
}

func (t *Translator) Translate() string {
	if t.lookahead != "" {
		t.block()
	}

	return strings.Join(t.buf, " ")
}

func (t *Translator) block() {
	t.match("{")

	prev := t.env
	t.env = NewEnv(prev)
	t.buf = append(t.buf, "{")

	t.decls()

	t.match("}")

	t.env = prev
	t.buf = append(t.buf, "}")
}

func (t *Translator) decls() {
	for {
		switch t.lookahead {
		case "{":
			t.block()
		case "}":
			return
		case "int", "char", "bool":
			t.decl()
		default:
			t.factor()
		}
	}
}

func (t *Translator) decl() {
	typ := t.lookahead
	t.matchAny()

	lex := t.lookahead
	t.matchAny()

	t.match(";")

	t.env.Put(lex, typ)
}

func (t *Translator) factor() {
	typ := t.env.Get(t.lookahead)
	t.buf = append(t.buf, fmt.Sprintf("%s:%s;", t.lookahead, typ))
	t.matchAny()
	t.match(";")
}

func (t *Translator) match(terminal string) {
	if t.lookahead != terminal {
		panic("invalid token")
	}
	t.lookahead = t.lexer.Next()
}

func (t *Translator) matchAny() {
	if t.lookahead == "" {
		panic("invalid token")
	}
	t.lookahead = t.lexer.Next()
}
