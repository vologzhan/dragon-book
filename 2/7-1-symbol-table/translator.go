package symbol_table

import (
	"comp/parser"
	"fmt"
	"strings"
)

type translator struct {
	*parser.Parser
	buf []string
	env *env
}

func translate(input string) string {
	t := &translator{
		Parser: parser.NewFromString(input),
	}
	t.program()

	return strings.Join(t.buf, " ")
}

func (t *translator) program() {
	if t.Lookahead() != "" {
		t.block()
	}
}

func (t *translator) block() {
	t.Match("{")

	prev := t.env
	t.env = newEnv(prev)
	t.buf = append(t.buf, "{")

	t.decls()

	t.Match("}")

	t.env = prev
	t.buf = append(t.buf, "}")
}

func (t *translator) decls() {
	for {
		switch t.Lookahead() {
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

func (t *translator) decl() {
	typ := t.Lookahead()
	t.MatchAny()

	lex := t.Lookahead()
	t.MatchAny()

	t.Match(";")

	t.env.put(lex, typ)
}

func (t *translator) factor() {
	typ := t.env.get(t.Lookahead())
	t.buf = append(t.buf, fmt.Sprintf("%s:%s;", t.Lookahead(), typ))
	t.MatchAny()
	t.Match(";")
}
